// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/org/api/permissions/v1"
)

type (
	IPermissions interface {
		// AssignRolePermissions 角色权限分配
		AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (*v1.AssignRolePermissionsRes, error)
		// GetRolePermissions 查询角色权限
		GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (*v1.GetRolePermissionsRes, error)
		// GetPermissionsTree 查询权限资源树（授权用）
		GetPermissionsTree(ctx context.Context, req *v1.GetPermissionsTreeReq) (*v1.GetPermissionsTreeRes, error)
		// GetMemberPermissions 获取用户权限
		GetMemberPermissions(ctx context.Context, req *v1.GetMemberPermissionsReq) (*v1.GetMemberPermissionsRes, error)
	}
)

var (
	localPermissions IPermissions
)

func Permissions() IPermissions {
	if localPermissions == nil {
		panic("implement not found for interface IPermissions, forgot register?")
	}
	return localPermissions
}

func RegisterPermissions(i IPermissions) {
	localPermissions = i
}
