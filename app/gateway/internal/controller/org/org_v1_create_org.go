package org

import (
	"context"

	v1 "freeroam/app/gateway/api/org/v1"
	oOrg "freeroam/app/org/api/org/v1"
)

func (c *ControllerV1) CreateOrg(ctx context.Context, req *v1.CreateOrgReq) (res *v1.CreateOrgRes, err error) {
	rpcReq := &oOrg.CreateOrgReq{
		ParentId: req.ParentId,
		Name:     req.Name,
		Code:     req.Code,
		Category: req.Category,
		Status:   req.Status,
		Sort:     req.Sort,
	}

	result, err := c.OrgRpcService.CreateOrg(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CreateOrgRes{
		Id: result.Id,
	}
	return res, nil
}
