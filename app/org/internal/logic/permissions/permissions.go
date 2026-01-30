package permissions

import (
	"context"
	v1 "freeroam/app/org/api/permissions/v1"
	"freeroam/app/org/internal/dao"
	"freeroam/app/org/internal/model/entity"
	"freeroam/app/org/internal/service"
	"freeroam/common/berror"
	"freeroam/common/tools/authsession"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sPermissions struct{}

func init() {
	service.RegisterPermissions(&sPermissions{})
}

// AssignRolePermissions 角色权限分配
func (s *sPermissions) AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (*v1.AssignRolePermissionsRes, error) {
	// 检验角色是否存在
	m := dao.Role
	exist, err := m.Ctx(ctx).
		Where(m.Columns().Id, req.RoleId).
		Where(m.Columns().IsDeleted, false).
		Exist()
	if !exist {
		return nil, gerror.NewCodef(berror.DataNotExist, "角色 %d 不存在", req.RoleId)
	}

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

	// 角色权限分配后，角色对应的用户版本增加
	memberIds, err := s.getRoleMember(ctx, req.RoleId)
	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询角色成员失败")
	}

	for _, id := range memberIds {
		_, err = authsession.IncrMemberVersion(ctx, id)
		if err != nil {
			return nil, gerror.WrapCodef(berror.DBErr, err, "更新用户 %d 版本失败", id)
		}
	}

	return &v1.AssignRolePermissionsRes{Success: true}, nil
}

// GetRolePermissions 查询角色权限
func (s *sPermissions) GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (*v1.GetRolePermissionsRes, error) {
	m := dao.RolePermissions

	// 查询角色关联的权限
	permIds, err := m.Ctx(ctx).
		Where(m.Columns().RoleId, req.RoleId).
		Where(m.Columns().IsDeleted, false).
		Fields(m.Columns().PermissionsId).
		Distinct().
		Array()

	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询角色关联的权限失败")
	}

	return &v1.GetRolePermissionsRes{PermissionsIds: permIds.Int64s()}, nil
}

// GetPermissionsTree 查询权限资源树（授权用）
func (s *sPermissions) GetPermissionsTree(ctx context.Context, req *v1.GetPermissionsTreeReq) (*v1.GetPermissionsTreeRes, error) {
	m := dao.Permissions

	// 查询所有有效权限
	var permList []*entity.Permissions
	err := m.Ctx(ctx).
		Where(m.Columns().IsDeleted, false).
		Where(m.Columns().IsEnabled, true).
		OrderAsc(m.Columns().Sort).
		Scan(&permList)

	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询所有有效权限")
	}

	// 构建树结构
	tree := s.buildPermissionsTree(permList, 0)

	return &v1.GetPermissionsTreeRes{Tree: tree}, nil
}

// GetMemberPermissions 获取用户权限
func (s *sPermissions) GetMemberPermissions(ctx context.Context, req *v1.GetMemberPermissionsReq) (*v1.GetMemberPermissionsRes, error) {
	// 1. 获取成员的所有职务ID
	mop := dao.MemberOrgPosition
	subPositions, err := mop.Ctx(ctx).
		Where(mop.Columns().MemberId, req.MemberId).
		Where(mop.Columns().IsDeleted, false).
		Fields(mop.Columns().PositionId).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询成员所有职务 ID 失败")
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
		return nil, gerror.WrapCode(berror.DBErr, err, "查询职务对应的所有角色 ID 失败")
	}

	// 3. 获取角色对应的所有权限ID
	rp := dao.RolePermissions
	subPermIds, err := rp.Ctx(ctx).
		WhereIn(rp.Columns().RoleId, subRoles).
		Where(rp.Columns().IsDeleted, false).
		Fields(rp.Columns().PermissionsId).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询角色对应的所有权限 ID 失败")
	}

	// 4. 查询最终权限列表
	p := dao.Permissions
	permCodes, err := p.Ctx(ctx).
		WhereIn(p.Columns().Id, subPermIds).
		WhereIn(p.Columns().PermType, req.PermType).
		Where(p.Columns().IsDeleted, false).
		Where(p.Columns().IsEnabled, true).
		Fields(p.Columns().PermCode).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询最终权限列表失败")
	}

	return &v1.GetMemberPermissionsRes{Permissions: permCodes.Strings()}, nil
}

// 校验权限有效性
func (s *sPermissions) validatePermCodes(ctx context.Context, permCodes []string) ([]*entity.Permissions, error) {
	m := dao.Permissions

	var permList []*entity.Permissions
	err := m.Ctx(ctx).
		Where(m.Columns().IsDeleted, false).
		Where(m.Columns().IsEnabled, true).
		WhereIn(m.Columns().PermCode, permCodes).
		Scan(&permList)

	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询权限列表失败")
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
		return nil, gerror.NewCodef(berror.InvalidPermissions, "无效的权限标识: %v", invalidCodes)
	}

	return permList, nil
}

// 覆盖写入角色权限
func (s *sPermissions) saveRolePermissions(ctx context.Context, roleId int64, permList []*entity.Permissions) error {
	m := dao.RolePermissions

	return m.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除该角色已有的权限关联
		_, err := m.Ctx(ctx).
			Where(m.Columns().RoleId, roleId).
			Unscoped().Delete()

		if err != nil {
			return gerror.WrapCode(berror.DBErr, err, "删除该角色已有的权限关联失败")
		}

		// 批量插入新的权限关联
		if len(permList) > 0 {
			insertData := make([]map[string]interface{}, 0, len(permList))
			for _, perm := range permList {
				insertData = append(insertData, map[string]interface{}{
					m.Columns().RoleId:        roleId,
					m.Columns().PermissionsId: perm.Id,
					m.Columns().IsDeleted:     false,
				})
			}

			_, err = m.Ctx(ctx).Data(insertData).Insert()
			if err != nil {
				return gerror.WrapCode(berror.DBErr, err, "批量插入新的权限关联失败")
			}
		}

		return nil
	})
}

// 构建权限树
func (s *sPermissions) buildPermissionsTree(permList []*entity.Permissions, parentId uint64) []*v1.PermissionsTreeNode {
	var nodes []*v1.PermissionsTreeNode

	for _, p := range permList {
		if p.ParentId == parentId {
			node := &v1.PermissionsTreeNode{
				Id:       int64(p.Id),
				PermCode: p.PermCode,
				Name:     p.Name,
				PermType: p.PermType,
				Children: s.buildPermissionsTree(permList, p.Id),
			}
			nodes = append(nodes, node)
		}
	}

	return nodes
}

// 获取角色下的所有用户
func (s *sPermissions) getRoleMember(ctx context.Context, roleId int64) ([]uint64, error) {
	// 1. 获取角色对应的所有职务ID
	pr := dao.PositionRole
	subPositionIds, err := pr.Ctx(ctx).
		Where(pr.Columns().RoleId, roleId).
		Where(pr.Columns().IsDeleted, false).
		Fields(pr.Columns().PositionId).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询角色对应的所有职务 ID 失败")
	}

	// 2. 获取职务的所有成员ID
	mop := dao.MemberOrgPosition
	subMemberIds, err := mop.Ctx(ctx).
		WhereIn(mop.Columns().PositionId, subPositionIds).
		Where(mop.Columns().IsDeleted, false).
		Fields(mop.Columns().MemberId).
		Distinct().
		Array()
	if err != nil {
		return nil, gerror.WrapCode(berror.DBErr, err, "查询职务所有成员 ID 失败")
	}

	return subMemberIds.Uint64s(), nil
}
