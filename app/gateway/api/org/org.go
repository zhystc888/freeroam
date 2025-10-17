// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package org

import (
	"context"

	"freeroam/app/gateway/api/org/v1"
)

type IOrgV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
}
