package permissions

import (
	"context"
	"freeroam/app/gateway/api/permissions/v1"
	oPermissions "freeroam/app/org/api/permissions/v1"
)

func (c *ControllerV1) GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (res *v1.GetRolePermissionsRes, err error) {
	rpcReq := &oPermissions.GetRolePermissionsReq{
		RoleId: req.RoleId,
	}

	rpcRes, err := c.PermissionsRpcService.GetRolePermissions(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	permissionsIds := rpcRes.PermissionsIds
	if len(permissionsIds) == 0 {
		permissionsIds = make([]int64, 0)
	}

	return &v1.GetRolePermissionsRes{
		PermissionsIds: permissionsIds,
	}, nil
}
