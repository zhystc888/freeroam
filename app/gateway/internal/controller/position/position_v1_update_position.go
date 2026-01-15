package position

import (
	"context"

	v1 "freeroam/app/gateway/api/position/v1"
	oPosition "freeroam/app/org/api/position/v1"
)

func (c *ControllerV1) UpdatePosition(ctx context.Context, req *v1.UpdatePositionReq) (res *v1.UpdatePositionRes, err error) {
	rpcReq := &oPosition.UpdatePositionReq{
		Id:        req.Id,
		Name:      req.Name,
		Status:    req.Status,
		DataScope: req.DataScope,
		OrgIds:    req.OrgIds,
		RoleIds:   req.RoleIds,
	}

	result, err := c.PositionRpcService.UpdatePosition(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.UpdatePositionRes{
		Success: result.Success,
	}
	return res, nil
}
