package permissions

import (
	"context"
	"freeroam/app/gateway/api/permissions/v1"
	oPermissions "freeroam/app/org/api/permissions/v1"
)

func (c *ControllerV1) AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (res *v1.AssignRolePermissionsRes, err error) {
	rpcReq := &oPermissions.AssignRolePermissionsReq{
		RoleId:    req.RoleId,
		PermCodes: req.PermCodes,
	}

	rpcRes, err := c.PermissionsRpcService.AssignRolePermissions(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.AssignRolePermissionsRes{
		Success: rpcRes.Success,
	}, nil
}
