package org

import (
	"context"

	v1 "freeroam/app/gateway/api/org/v1"
	oOrg "freeroam/app/org/api/org/v1"
)

func (c *ControllerV1) GetOrg(ctx context.Context, req *v1.GetOrgReq) (res *v1.GetOrgRes, err error) {
	rpcReq := &oOrg.GetOrgReq{
		Id: req.Id,
	}

	result, err := c.OrgRpcService.GetOrg(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetOrgRes{
		Id:       result.Id,
		ParentId: result.ParentId,
		Name:     result.Name,
		FullName: result.FullName,
		Code:     result.Code,
		Category: result.Category,
		Status:   result.Status,
		Sort:     result.Sort,
		Path:     result.Path,
		CreateAt: result.CreateAt,
		UpdateAt: result.UpdateAt,
	}
	return res, nil
}
