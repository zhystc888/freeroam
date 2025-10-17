package org

import (
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"bbk/app/gateway/api/org/v1"
	sorg "bbk/app/org/api/org/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	dto := sorg.GetListReq{
		Name:     req.Name,
		Type:     req.Type,
		Code:     req.Code,
		Page:     req.Page,
		Limit:    req.Limit,
		ParentId: nil,
	}

	if req.ParentId != nil {
		dto.ParentId = wrapperspb.Int64(*req.ParentId)
	}

	result, err := c.OrgRpcService.GetList(ctx, &dto)
	if err != nil {
		return
	}

	list := make([]*v1.GetListItem, len(result.List))

	for i, item := range result.List {
		list[i] = &v1.GetListItem{
			Id:       item.Id,
			Name:     item.Name,
			Code:     item.Code,
			Type:     item.Type,
			Status:   item.Status,
			CreateAt: item.CreateAt,
		}
	}

	res = &v1.GetListRes{
		List:  list,
		Total: result.Total,
	}

	return
}
