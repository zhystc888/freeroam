// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/gateway/api/permissions/v1"
)

type (
	IPermissions interface {
		// GetFrontPermissions 获取当前登录用户前端权限
		GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (*v1.GetFrontPermissionsRes, error)
		// VeryApiPermissions 验证当前登录用户是否有指定权限
		VeryApiPermissions(ctx context.Context, permissions string) (bool, error)
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
