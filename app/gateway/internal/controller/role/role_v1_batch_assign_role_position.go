package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
	oRole "freeroam/app/org/api/role/v1"
)

func (c *ControllerV1) BatchAssignRolePosition(ctx context.Context, req *v1.BatchAssignRolePositionReq) (res *v1.BatchAssignRolePositionRes, err error) {
	result, err := c.RoleRpcService.BatchAssignRolePosition(ctx, &oRole.BatchAssignRolePositionReq{
		RoleId:      req.RoleId,
		PositionIds: req.PositionIds,
	})
	if err != nil {
		return nil, err
	}

	return &v1.BatchAssignRolePositionRes{
		Success: result.Success,
	}, nil
}
