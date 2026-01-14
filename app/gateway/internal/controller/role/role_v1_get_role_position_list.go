package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
	oRole "freeroam/app/org/api/role/v1"
)

func (c *ControllerV1) GetRolePositionList(ctx context.Context, req *v1.GetRolePositionListReq) (res *v1.GetRolePositionListRes, err error) {
	result, err := c.RoleRpcService.GetRolePositionList(ctx, &oRole.GetRolePositionListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		RoleId:   req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*v1.PositionItem, 0, len(result.List))
	for _, item := range result.List {
		list = append(list, &v1.PositionItem{
			PositionId:     item.PositionId,
			PositionName:   item.PositionName,
			PositionStatus: item.PositionStatus,
			CreateAt:       item.CreateAt,
		})
	}

	return &v1.GetRolePositionListRes{
		List:     list,
		Total:    result.Total,
		Page:     result.Page,
		PageSize: result.PageSize,
	}, nil
}
