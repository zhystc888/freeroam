// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/org/api/role/v1"
)

type (
	IRole interface {
		// CreateRole 创建角色
		CreateRole(ctx context.Context, in *v1.CreateRoleReq) (*v1.CreateRoleRes, error)
		// UpdateRole 更新角色
		UpdateRole(ctx context.Context, in *v1.UpdateRoleReq) (*v1.UpdateRoleRes, error)
		// DeleteRole 删除角色
		DeleteRole(ctx context.Context, in *v1.DeleteRoleReq) (*v1.DeleteRoleRes, error)
		// GetRole 获取角色详情
		GetRole(ctx context.Context, in *v1.GetRoleReq) (*v1.GetRoleRes, error)
		// ListRole 获取角色列表
		ListRole(ctx context.Context, in *v1.ListRoleReq) (*v1.ListRoleRes, error)
		// GetRolePositionList 查询角色绑定的职务列表
		GetRolePositionList(ctx context.Context, in *v1.GetRolePositionListReq) (*v1.GetRolePositionListRes, error)
		// BatchAssignRolePosition 批量绑定职务到角色（覆盖式）
		BatchAssignRolePosition(ctx context.Context, in *v1.BatchAssignRolePositionReq) (*v1.BatchAssignRolePositionRes, error)
		// GetRolePositionIds 查询角色当前绑定职务ID集合（表单回显）
		GetRolePositionIds(ctx context.Context, in *v1.GetRolePositionIdsReq) (*v1.GetRolePositionIdsRes, error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
