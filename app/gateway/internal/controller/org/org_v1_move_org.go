package org

import (
	"context"

	v1 "freeroam/app/gateway/api/org/v1"
	oOrg "freeroam/app/org/api/org/v1"
)

func (c *ControllerV1) MoveOrg(ctx context.Context, req *v1.MoveOrgReq) (res *v1.MoveOrgRes, err error) {
	rpcReq := &oOrg.MoveOrgReq{
		Id:          req.Id,
		NewParentId: req.NewParentId,
		NewSort:     req.NewSort,
	}

	result, err := c.OrgRpcService.MoveOrg(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.MoveOrgRes{
		Success: result.Success,
	}
	return res, nil
}
