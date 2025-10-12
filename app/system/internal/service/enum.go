// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bbk/app/system/internal/model"
	"context"
)

type (
	IEnum interface {
		GetEnumList(ctx context.Context, dto *model.GetEnumListDto) (res *model.GetEnumListRes, err error)
	}
)

var (
	localEnum IEnum
)

func Enum() IEnum {
	if localEnum == nil {
		panic("implement not found for interface IEnum, forgot register?")
	}
	return localEnum
}

func RegisterEnum(i IEnum) {
	localEnum = i
}
