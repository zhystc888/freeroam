// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/system/api/enum/v1"
)

type (
	IEnum interface {
		// GetByTypeAndCode 根据枚举类型和code 获取一个枚举
		GetByTypeAndCode(ctx context.Context, in *v1.GetByTypeAndCodeReq) (*v1.GetByTypeAndCodeRes, error)
		// GetByType 根据枚举类型数组 获取多个枚举
		GetByType(ctx context.Context, in *v1.GetByTypeReq) (*v1.GetByTypeRes, error)
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
