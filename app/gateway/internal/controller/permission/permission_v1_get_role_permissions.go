package permission

import (
	"context"

	v1 "freeroam/app/gateway/api/permission/v1"
	oPermission "freeroam/app/org/api/permission/v1"
)

func (c *ControllerV1) GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (res *v1.GetRolePermissionsRes, err error) {
	rpcReq := &oPermission.GetRolePermissionsReq{
		RoleId: req.RoleId,
	}

	rpcRes, err := c.PermissionRpcService.GetRolePermissions(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.GetRolePermissionsRes{
		PermCodes: rpcRes.PermCodes,
	}, nil
}
