// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"freeroam/app/org/internal/model"
)

type (
	IUserMember interface {
		GetList(ctx context.Context, params *model.UserMemberListDto) (res *model.ListReq[model.UserMemberListVo], err error)
		GetOne(ctx context.Context, userId int64) (res *model.UserMemberVo, err error)
	}
)

var (
	localUserMember IUserMember
)

func UserMember() IUserMember {
	if localUserMember == nil {
		panic("implement not found for interface IUserMember, forgot register?")
	}
	return localUserMember
}

func RegisterUserMember(i IUserMember) {
	localUserMember = i
}
