package permission

import (
	"context"
	"fmt"
	"strings"

	v1 "freeroam/app/system/api/permission/v1"
	"freeroam/app/system/internal/dao"
	"freeroam/app/system/internal/model/entity"
	"freeroam/app/system/internal/service"
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
	// 1. 校验权限点有效性
	permList, err := s.validatePermCodes(ctx, req.PermCodes)
	if err != nil {
		return nil, err
	}

	// 2. 写时扩散计算：补齐父链 + 包含子节点
	expandedPermIds, err := s.expandPermissions(ctx, permList)
	if err != nil {
		return nil, err
	}

	// 3. 覆盖写入角色权限
	err = s.saveRolePermissions(ctx, req.RoleId, expandedPermIds)
	if err != nil {
		return nil, err
	}

	return &v1.AssignRolePermissionsRes{Success: true}, nil
}

// GetRolePermissions 查询角色已分配权限
func (s *sPermission) GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (*v1.GetRolePermissionsRes, error) {
	m := dao.RolePermissions
	p := dao.Permissions

	// 查询角色关联的权限（仅页面和组件，不返回接口权限）
	var permCodes []string
	err := m.Ctx(ctx).
		Fields(p.Table()+"."+p.Columns().PermCode).
		LeftJoin(p.Table(), fmt.Sprintf("%s.%s = %s.%s",
			m.Table(), m.Columns().PermissionId,
			p.Table(), p.Columns().Id)).
		Where(m.Columns().RoleId, req.RoleId).
		Where(m.Columns().IsDeleted, 0).
		Where(p.Table()+"."+p.Columns().IsDeleted, 0).
		Where(p.Table()+"."+p.Columns().IsEnabled, 1).
		WhereIn(p.Table()+"."+p.Columns().PermType, []int{1, 2}). // 仅页面和组件
		Scan(&permCodes)

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.GetRolePermissionsRes{PermCodes: permCodes}, nil
}

// GetPermissionTree 查询权限资源树（授权用）
func (s *sPermission) GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (*v1.GetPermissionTreeRes, error) {
	m := dao.Permissions

	// 查询所有有效权限（仅页面和组件，不包含接口权限）
	var permList []*entity.Permissions
	err := m.Ctx(ctx).
		Where(m.Columns().IsDeleted, 0).
		Where(m.Columns().IsEnabled, 1).
		WhereIn(m.Columns().PermType, []int{1, 2}). // 仅页面和组件
		OrderAsc(m.Columns().Sort).
		Scan(&permList)

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	// 构建树结构
	tree := s.buildPermissionTree(permList, 0)

	return &v1.GetPermissionTreeRes{Tree: tree}, nil
}

// GetFrontPermissions 获取前端权限集合（页面+组件）
func (s *sPermission) GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (*v1.GetFrontPermissionsRes, error) {
	// 通过成员→职务→角色链路聚合权限
	// SQL: member_positions -> position_roles -> role_permissions -> permissions
	p := dao.Permissions
	rp := dao.RolePermissions

	// 查询成员拥有的前端权限（页面+组件）
	var permCodes []string
	err := p.Ctx(ctx).
		Distinct().
		Fields(p.Columns().PermCode).
		InnerJoin(rp.Table(), fmt.Sprintf("%s.%s = %s.%s",
			rp.Table(), rp.Columns().PermissionId,
			p.Table(), p.Columns().Id)).
		InnerJoin("free_position_role pr", fmt.Sprintf("pr.role_id = %s.%s AND pr.is_deleted = 0",
			rp.Table(), rp.Columns().RoleId)).
		InnerJoin("free_member_position mp", "mp.position_id = pr.position_id AND mp.is_deleted = 0").
		Where("mp.member_id", req.MemberId).
		Where(rp.Table()+"."+rp.Columns().IsDeleted, 0).
		Where(p.Columns().IsDeleted, 0).
		Where(p.Columns().IsEnabled, 1).
		WhereIn(p.Columns().PermType, []int{1, 2}). // 仅页面和组件
		Scan(&permCodes)

	if err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	return &v1.GetFrontPermissionsRes{Permissions: permCodes}, nil
}

// validatePermCodes 校验权限点有效性
func (s *sPermission) validatePermCodes(ctx context.Context, permCodes []string) ([]*entity.Permissions, error) {
	m := dao.Permissions

	var permList []*entity.Permissions
	err := m.Ctx(ctx).
		Where(m.Columns().IsDeleted, 0).
		Where(m.Columns().IsEnabled, 1).
		WhereIn(m.Columns().PermType, []int{1, 2}). // 仅允许页面和组件
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

// expandPermissions 写时扩散计算：补齐父链 + 包含子节点
func (s *sPermission) expandPermissions(ctx context.Context, permList []*entity.Permissions) ([]uint64, error) {
	m := dao.Permissions
	permIdSet := make(map[uint64]bool)

	// 收集所有输入权限的ID
	for _, p := range permList {
		permIdSet[p.Id] = true
	}

	// 1. 向上补齐父链
	for _, p := range permList {
		if p.IdPath != "" {
			// 解析 id_path，格式如 /1/12/88/
			ids := strings.Split(strings.Trim(p.IdPath, "/"), "/")
			for _, idStr := range ids {
				if idStr == "" {
					continue
				}
				var id uint64
				fmt.Sscanf(idStr, "%d", &id)
				if id > 0 {
					permIdSet[id] = true
				}
			}
		}
	}

	// 2. 向下包含子节点
	for _, p := range permList {
		// 查询所有以当前权限 id_path 为前缀的子权限
		var children []*entity.Permissions
		err := m.Ctx(ctx).
			Where(m.Columns().IsDeleted, 0).
			Where(m.Columns().IsEnabled, 1).
			WhereLike(m.Columns().IdPath, fmt.Sprintf("%%/%d/%%", p.Id)).
			Scan(&children)

		if err != nil {
			return nil, gerror.NewCode(berror.DBErr, err.Error())
		}

		for _, child := range children {
			permIdSet[child.Id] = true
		}
	}

	// 转换为切片
	permIds := make([]uint64, 0, len(permIdSet))
	for id := range permIdSet {
		permIds = append(permIds, id)
	}

	return permIds, nil
}

// saveRolePermissions 覆盖写入角色权限
func (s *sPermission) saveRolePermissions(ctx context.Context, roleId int64, permIds []uint64) error {
	m := dao.RolePermissions

	return m.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 软删除该角色已有的权限关联
		_, err := m.Ctx(ctx).
			Data(map[string]interface{}{
				m.Columns().IsDeleted: 1,
			}).
			Where(m.Columns().RoleId, roleId).
			Where(m.Columns().IsDeleted, 0).
			Update()

		if err != nil {
			return gerror.NewCode(berror.DBErr, err.Error())
		}

		// 2. 批量插入新的权限关联
		if len(permIds) > 0 {
			insertData := make([]map[string]interface{}, 0, len(permIds))
			for _, permId := range permIds {
				insertData = append(insertData, map[string]interface{}{
					m.Columns().RoleId:       roleId,
					m.Columns().PermissionId: permId,
					m.Columns().IsDeleted:    0,
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
				PermType: int32(p.PermType),
				Children: s.buildPermissionTree(permList, p.Id),
			}
			nodes = append(nodes, node)
		}
	}

	return nodes
}
