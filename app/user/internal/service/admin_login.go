// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "bbk/app/user/api/admin_login/v1"
	"bbk/app/user/internal/model"
)

type (
	IAdminLogin interface {
		AdminLogin(dto *model.LoginDto) (res *v1.LoginRes, err error)
	}
)

var (
	localAdminLogin IAdminLogin
)

func AdminLogin() IAdminLogin {
	if localAdminLogin == nil {
		panic("implement not found for interface IAdminLogin, forgot register?")
	}
	return localAdminLogin
}

func RegisterAdminLogin(i IAdminLogin) {
	localAdminLogin = i
}
