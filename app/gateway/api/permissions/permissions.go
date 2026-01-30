// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package permissions

import (
	"context"

	"freeroam/app/gateway/api/permissions/v1"
)

type IPermissionsV1 interface {
	AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (res *v1.AssignRolePermissionsRes, err error)
	GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (res *v1.GetRolePermissionsRes, err error)
	GetPermissionsTree(ctx context.Context, req *v1.GetPermissionsTreeReq) (res *v1.GetPermissionsTreeRes, err error)
	GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (res *v1.GetFrontPermissionsRes, err error)
}
