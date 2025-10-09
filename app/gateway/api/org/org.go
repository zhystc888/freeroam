// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package org

import (
	"context"

	"bbk/app/gateway/api/org/v1"
)

type IOrgV1 interface {
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
}
