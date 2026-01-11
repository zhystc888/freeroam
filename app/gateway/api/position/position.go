// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package position

import (
	"context"

	"freeroam/app/gateway/api/position/v1"
)

type IPositionV1 interface {
	ListPosition(ctx context.Context, req *v1.ListPositionReq) (res *v1.ListPositionRes, err error)
	GetPosition(ctx context.Context, req *v1.GetPositionReq) (res *v1.GetPositionRes, err error)
	CreatePosition(ctx context.Context, req *v1.CreatePositionReq) (res *v1.CreatePositionRes, err error)
	UpdatePosition(ctx context.Context, req *v1.UpdatePositionReq) (res *v1.UpdatePositionRes, err error)
	DeletePosition(ctx context.Context, req *v1.DeletePositionReq) (res *v1.DeletePositionRes, err error)
	GetPositionOptions(ctx context.Context, req *v1.GetPositionOptionsReq) (res *v1.GetPositionOptionsRes, err error)
}
