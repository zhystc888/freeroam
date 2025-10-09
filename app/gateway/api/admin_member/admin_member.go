// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin_member

import (
	"context"

	"bbk/app/gateway/api/admin_member/v1"
)

type IAdminMemberV1 interface {
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
	GetMemberList(ctx context.Context, req *v1.GetMemberListReq) (res *v1.GetMemberListRes, err error)
}
