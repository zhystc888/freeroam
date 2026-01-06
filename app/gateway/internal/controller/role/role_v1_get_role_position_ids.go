package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
	oRole "freeroam/app/org/api/role/v1"
)

func (c *ControllerV1) GetRolePositionIds(ctx context.Context, req *v1.GetRolePositionIdsReq) (res *v1.GetRolePositionIdsRes, err error) {
	result, err := c.RoleRpcService.GetRolePositionIds(ctx, &oRole.GetRolePositionIdsReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetRolePositionIdsRes{
		PositionIds: result.PositionIds,
	}, nil
}
