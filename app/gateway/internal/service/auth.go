// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/gateway/api/auth/v1"
)

type (
	IAuth interface {
		Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
		Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
