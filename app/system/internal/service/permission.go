// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/system/api/permission/v1"
)

type (
	IPermission interface {
		// AssignRolePermissions 角色权限分配（写时扩散）
		AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (*v1.AssignRolePermissionsRes, error)
		// GetRolePermissions 查询角色已分配权限
		GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (*v1.GetRolePermissionsRes, error)
		// GetPermissionTree 查询权限资源树（授权用）
		GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (*v1.GetPermissionTreeRes, error)
		// GetFrontPermissions 获取前端权限集合（页面+组件）
		GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (*v1.GetFrontPermissionsRes, error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
