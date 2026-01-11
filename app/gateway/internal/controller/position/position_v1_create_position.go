package position

import (
	"context"

	v1 "freeroam/app/gateway/api/position/v1"
	oPosition "freeroam/app/org/api/position/v1"
)

func (c *ControllerV1) CreatePosition(ctx context.Context, req *v1.CreatePositionReq) (res *v1.CreatePositionRes, err error) {
	rpcReq := &oPosition.CreatePositionReq{
		Name:      req.Name,
		Status:    req.Status,
		DataScope: req.DataScope,
		OrgIds:    req.OrgIds,
		RoleIds:   req.RoleIds,
	}

	result, err := c.PositionRpcService.CreatePosition(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CreatePositionRes{
		Id: result.Id,
	}
	return res, nil
}
