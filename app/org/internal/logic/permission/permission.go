package permission

import (
	"context"
	"fmt"
	v1 "freeroam/app/org/api/permission/v1"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model/entity"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(&sPermission{})
}

// AssignRolePermissions 角色权限分配
func (s *sPermission) AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (*v1.AssignRolePermissionsRes, error) {
	// 校验权限有效性
	permList, err := s.validatePermCodes(ctx, req.PermCodes)
	if err != nil {
		return nil, err
	}

	// 覆盖写入角色权限
	err = s.saveRolePermissions(ctx, req.RoleId, permList)
	if err != nil {
		return nil, err
	}

	return &v1.AssignRolePermissionsRes{Success: true}, nil
}

// GetRolePermissions 查询角色权限
func (s *sPermission) GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (*v1.GetRolePermissionsRes, error) {
	m := dao.RolePermissions

	// 查询角色关联的权限
	permIds, err := m.Ctx(ctx).
		Where(m.Columns().RoleId, req.RoleId).
		Where(m.Columns().IsDeleted, false).
		Fields(m.Columns().PermissionId).
		Distinct().
		Array()

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.GetRolePermissionsRes{PermissionIds: permIds.Int64s()}, nil
}

// GetPermissionTree 查询权限资源树（授权用）
func (s *sPermission) GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (*v1.GetPermissionTreeRes, error) {
	m := dao.Permissions

	// 查询所有有效权限
	var permList []*entity.Permissions
	err := m.Ctx(ctx).
		Where(m.Columns().IsDeleted, false).
		Where(m.Columns().IsEnabled, true).
		OrderAsc(m.Columns().Sort).
		Scan(&permList)

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 构建树结构
	tree := s.buildPermissionTree(permList, 0)

	return &v1.GetPermissionTreeRes{Tree: tree}, nil
}

// GetMemberPermissions 获取用户权限
func (s *sPermission) GetMemberPermissions(ctx context.Context, req *v1.GetMemberPermissionsReq) (*v1.GetMemberPermissionsRes, error) {
	// 1. 获取成员的所有职务ID
	mop := dao.MemberOrgPosition
	subPositions, err := mop.Ctx(ctx).
		Where(mop.Columns().MemberId, req.MemberId).
		Where(mop.Columns().IsDeleted, false).
		Fields(mop.Columns().PositionId).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 2. 获取职务对应的所有角色ID
	pr := dao.PositionRole
	subRoles, err := pr.Ctx(ctx).
		WhereIn(pr.Columns().PositionId, subPositions).
		Where(pr.Columns().IsDeleted, false).
		Fields(pr.Columns().RoleId).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 3. 获取角色对应的所有权限ID
	rp := dao.RolePermissions
	subPermIds, err := rp.Ctx(ctx).
		WhereIn(rp.Columns().RoleId, subRoles).
		Where(rp.Columns().IsDeleted, false).
		Fields(rp.Columns().PermissionId).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 4. 查询最终权限列表
	p := dao.Permissions
	query := p.Ctx(ctx).
		WhereIn(p.Columns().Id, subPermIds).
		Where(p.Columns().IsDeleted, false).
		Where(p.Columns().IsEnabled, true)

	if len(req.PermTypes) > 0 {
		query = query.WhereIn(p.Columns().PermType, req.PermTypes)
	}

	permCodes, err := query.Fields(p.Columns().PermCode).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.GetMemberPermissionsRes{Permissions: permCodes.Strings()}, nil
}

// validatePermCodes 校验权限有效性
func (s *sPermission) validatePermCodes(ctx context.Context, permCodes []string) ([]*entity.Permissions, error) {
	m := dao.Permissions

	var permList []*entity.Permissions
	err := m.Ctx(ctx).
		Where(m.Columns().IsDeleted, false).
		Where(m.Columns().IsEnabled, true).
		WhereIn(m.Columns().PermCode, permCodes).
		Scan(&permList)

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 检查是否所有权限点都有效
	if len(permList) != len(permCodes) {
		foundCodes := make(map[string]bool)
		for _, p := range permList {
			foundCodes[p.PermCode] = true
		}
		var invalidCodes []string
		for _, code := range permCodes {
			if !foundCodes[code] {
				invalidCodes = append(invalidCodes, code)
			}
		}
		return nil, gerror.NewCode(berror.InvalidPermissions, fmt.Sprintf("无效的权限标识: %v", invalidCodes))
	}

	return permList, nil
}

// saveRolePermissions 覆盖写入角色权限
func (s *sPermission) saveRolePermissions(ctx context.Context, roleId int64, permList []*entity.Permissions) error {
	m := dao.RolePermissions

	return m.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除该角色已有的权限关联
		_, err := m.Ctx(ctx).
			Where(m.Columns().RoleId, roleId).
			Unscoped().Delete()

		if err != nil {
			return gerror.NewCode(berror.DBErr, err.Error())
		}

		// 批量插入新的权限关联
		if len(permList) > 0 {
			insertData := make([]map[string]interface{}, 0, len(permList))
			for _, perm := range permList {
				insertData = append(insertData, map[string]interface{}{
					m.Columns().RoleId:       roleId,
					m.Columns().PermissionId: perm.Id,
					m.Columns().IsDeleted:    false,
				})
			}

			_, err = m.Ctx(ctx).Data(insertData).Insert()
			if err != nil {
				return gerror.NewCode(berror.DBErr, err.Error())
			}
		}

		return nil
	})
}

// buildPermissionTree 构建权限树
func (s *sPermission) buildPermissionTree(permList []*entity.Permissions, parentId uint64) []*v1.PermissionTreeNode {
	var nodes []*v1.PermissionTreeNode

	for _, p := range permList {
		if p.ParentId == parentId {
			node := &v1.PermissionTreeNode{
				Id:       int64(p.Id),
				PermCode: p.PermCode,
				Name:     p.Name,
				PermType: p.PermType,
				Children: s.buildPermissionTree(permList, p.Id),
			}
			nodes = append(nodes, node)
		}
	}

	return nodes
}
