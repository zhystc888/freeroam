// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/org/api/auth/v1"
)

type (
	IAuth interface {
		// ValidateMemberCredential 校验成员账号密码
		ValidateMemberCredential(ctx context.Context, req *v1.ValidateMemberCredentialReq) (res *v1.ValidateMemberCredentialRes, err error)
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
