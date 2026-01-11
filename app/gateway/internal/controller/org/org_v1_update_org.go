package org

import (
	"context"

	v1 "freeroam/app/gateway/api/org/v1"
	oOrg "freeroam/app/org/api/org/v1"
)

func (c *ControllerV1) UpdateOrg(ctx context.Context, req *v1.UpdateOrgReq) (res *v1.UpdateOrgRes, err error) {
	rpcReq := &oOrg.UpdateOrgReq{
		Id:       req.Id,
		ParentId: req.ParentId,
		Name:     req.Name,
		FullName: req.FullName,
		Code:     req.Code,
		Category: req.Category,
		Status:   req.Status,
		Sort:     req.Sort,
	}

	result, err := c.OrgRpcService.UpdateOrg(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.UpdateOrgRes{
		Success: result.Success,
	}
	return res, nil
}
