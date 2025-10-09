// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bbk/app/user/internal/model"
	"context"
)

type (
	IAdminMember interface {
		GetMemberList(ctx context.Context, params *model.AdminGetMemberListDto) (res []model.AdminGetMemberListVo, total int, err error)
		GetMember(ctx context.Context, userId int64) (res *model.AdminGetMemberVo, err error)
	}
)

var (
	localAdminMember IAdminMember
)

func AdminMember() IAdminMember {
	if localAdminMember == nil {
		panic("implement not found for interface IAdminMember, forgot register?")
	}
	return localAdminMember
}

func RegisterAdminMember(i IAdminMember) {
	localAdminMember = i
}
