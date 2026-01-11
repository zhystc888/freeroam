package org

import (
	"context"

	v1 "freeroam/app/gateway/api/org/v1"
	oOrg "freeroam/app/org/api/org/v1"
)

func (c *ControllerV1) DeleteOrg(ctx context.Context, req *v1.DeleteOrgReq) (res *v1.DeleteOrgRes, err error) {
	rpcReq := &oOrg.DeleteOrgReq{
		Id: req.Id,
	}

	result, err := c.OrgRpcService.DeleteOrg(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.DeleteOrgRes{
		Success: result.Success,
	}
	return res, nil
}
