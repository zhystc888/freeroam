// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
)

type IRoleV1 interface {
	CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error)
	UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (res *v1.UpdateRoleRes, err error)
	DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error)
	GetRole(ctx context.Context, req *v1.GetRoleReq) (res *v1.GetRoleRes, err error)
	ListRole(ctx context.Context, req *v1.ListRoleReq) (res *v1.ListRoleRes, err error)
	GetRolePositionList(ctx context.Context, req *v1.GetRolePositionListReq) (res *v1.GetRolePositionListRes, err error)
	BatchAssignRolePosition(ctx context.Context, req *v1.BatchAssignRolePositionReq) (res *v1.BatchAssignRolePositionRes, err error)
	GetRolePositionIds(ctx context.Context, req *v1.GetRolePositionIdsReq) (res *v1.GetRolePositionIdsRes, err error)
}
