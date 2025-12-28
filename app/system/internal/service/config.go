// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/system/api/config/v1"
)

type (
	IConfig interface {
		Boot(ctx context.Context)
		// GetByCode 根据 code 获取一个 配置信息
		GetByCode(ctx context.Context, in *v1.GetByCodeReq) (*v1.GetByCodeRes, error)
	}
)

var (
	localConfig IConfig
)

func Config() IConfig {
	if localConfig == nil {
		panic("implement not found for interface IConfig, forgot register?")
	}
	return localConfig
}

func RegisterConfig(i IConfig) {
	localConfig = i
}
