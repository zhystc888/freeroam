package org

import (
	"context"

	v1 "freeroam/app/gateway/api/org/v1"
	oOrg "freeroam/app/org/api/org/v1"
)

func (c *ControllerV1) ListOrg(ctx context.Context, req *v1.ListOrgReq) (res *v1.ListOrgRes, err error) {
	rpcReq := &oOrg.ListOrgReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Code:     req.Code,
		Category: req.Category,
		Status:   req.Status,
	}

	result, err := c.OrgRpcService.ListOrg(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	list := make([]*v1.OrgListItem, 0, len(result.List))
	for _, item := range result.List {
		list = append(list, &v1.OrgListItem{
			Id:       item.Id,
			Name:     item.Name,
			FullName: item.FullName,
			Code:     item.Code,
			Status:   item.Status,
			Category: item.Category,
			CreateAt: item.CreateAt,
		})
	}

	res = &v1.ListOrgRes{
		List:     list,
		Total:    result.Total,
		Page:     result.Page,
		PageSize: result.PageSize,
	}
	return res, nil
}
