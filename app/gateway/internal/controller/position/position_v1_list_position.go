package position

import (
	"context"

	v1 "freeroam/app/gateway/api/position/v1"
	oPosition "freeroam/app/org/api/position/v1"
)

func (c *ControllerV1) ListPosition(ctx context.Context, req *v1.ListPositionReq) (res *v1.ListPositionRes, err error) {
	rpcReq := &oPosition.ListPositionReq{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Keyword:  req.Keyword,
		Status:   req.Status,
	}

	result, err := c.PositionRpcService.ListPosition(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	list := make([]*v1.PositionListItem, 0, len(result.List))
	for _, item := range result.List {
		list = append(list, &v1.PositionListItem{
			Id:        item.Id,
			Name:      item.Name,
			Status:    item.Status,
			DataScope: item.DataScope,
			OrgIds:    item.OrgIds,
			RoleIds:   item.RoleIds,
			CreateAt:  item.CreateAt,
		})
	}

	res = &v1.ListPositionRes{
		List:     list,
		Total:    result.Total,
		Page:     result.Page,
		PageSize: result.PageSize,
	}
	return res, nil
}
