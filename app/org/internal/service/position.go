// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "freeroam/app/org/api/position/v1"
)

type (
	IPosition interface {
		// ListPosition 职务列表
		ListPosition(ctx context.Context, in *v1.ListPositionReq) (*v1.ListPositionRes, error)
		// GetPosition 获取职务详情
		GetPosition(ctx context.Context, in *v1.GetPositionReq) (*v1.GetPositionRes, error)
		// CreatePosition 新建职务
		CreatePosition(ctx context.Context, in *v1.CreatePositionReq) (*v1.CreatePositionRes, error)
		// UpdatePosition 编辑职务
		UpdatePosition(ctx context.Context, in *v1.UpdatePositionReq) (*v1.UpdatePositionRes, error)
		// DeletePosition 删除职务
		DeletePosition(ctx context.Context, in *v1.DeletePositionReq) (*v1.DeletePositionRes, error)
		// GetPositionOptions 按组织获取可选职务集合
		GetPositionOptions(ctx context.Context, in *v1.GetPositionOptionsReq) (*v1.GetPositionOptionsRes, error)
	}
)

var (
	localPosition IPosition
)

func Position() IPosition {
	if localPosition == nil {
		panic("implement not found for interface IPosition, forgot register?")
	}
	return localPosition
}

func RegisterPosition(i IPosition) {
	localPosition = i
}
