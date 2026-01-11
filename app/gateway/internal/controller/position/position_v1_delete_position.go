package position

import (
	"context"

	v1 "freeroam/app/gateway/api/position/v1"
	oPosition "freeroam/app/org/api/position/v1"
)

func (c *ControllerV1) DeletePosition(ctx context.Context, req *v1.DeletePositionReq) (res *v1.DeletePositionRes, err error) {
	rpcReq := &oPosition.DeletePositionReq{
		Id: req.Id,
	}

	result, err := c.PositionRpcService.DeletePosition(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DeletePositionRes{
		Success: result.Success,
	}
	return res, nil
}
