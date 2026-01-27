// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package permission

import (
	"context"

	"freeroam/app/gateway/api/permission/v1"
)

type IPermissionV1 interface {
	AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (res *v1.AssignRolePermissionsRes, err error)
	GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (res *v1.GetRolePermissionsRes, err error)
	GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (res *v1.GetPermissionTreeRes, err error)
	GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (res *v1.GetFrontPermissionsRes, err error)
}
