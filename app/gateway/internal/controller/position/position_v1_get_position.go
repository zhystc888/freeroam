package position

import (
	"context"

	v1 "freeroam/app/gateway/api/position/v1"
	oPosition "freeroam/app/org/api/position/v1"
)

func (c *ControllerV1) GetPosition(ctx context.Context, req *v1.GetPositionReq) (res *v1.GetPositionRes, err error) {
	rpcReq := &oPosition.GetPositionReq{
		Id: req.Id,
	}

	result, err := c.PositionRpcService.GetPosition(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetPositionRes{
		Id:        result.Id,
		Name:      result.Name,
		Status:    result.Status,
		DataScope: result.DataScope,
		OrgIds:    result.OrgIds,
		RoleIds:   result.RoleIds,
		CreateAt:  result.CreateAt,
		UpdateAt:  result.UpdateAt,
	}
	return res, nil
}
