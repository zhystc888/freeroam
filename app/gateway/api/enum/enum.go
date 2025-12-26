// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package enum

import (
	"context"

	"freeroam/app/gateway/api/enum/v1"
)

type IEnumV1 interface {
	GetEnumList(ctx context.Context, req *v1.GetEnumListReq) (res *v1.GetEnumListRes, err error)
}
