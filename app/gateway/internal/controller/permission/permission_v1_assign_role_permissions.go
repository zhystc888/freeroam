package permission

import (
	"context"

	"freeroam/app/gateway/api/permission/v1"
	sPermission "freeroam/app/system/api/permission/v1"
)

func (c *ControllerV1) AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (res *v1.AssignRolePermissionsRes, err error) {
	rpcReq := &sPermission.AssignRolePermissionsReq{
		RoleId:    req.RoleId,
		PermCodes: req.PermCodes,
	}

	rpcRes, err := c.PermissionRpcService.AssignRolePermissions(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.AssignRolePermissionsRes{
		Success: rpcRes.Success,
	}, nil
}
