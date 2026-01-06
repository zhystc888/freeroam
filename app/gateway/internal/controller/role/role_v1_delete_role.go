package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
	oRole "freeroam/app/org/api/role/v1"
)

func (c *ControllerV1) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {
	result, err := c.RoleRpcService.DeleteRole(ctx, &oRole.DeleteRoleReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &v1.DeleteRoleRes{
		Success: result.Success,
	}, nil
}
