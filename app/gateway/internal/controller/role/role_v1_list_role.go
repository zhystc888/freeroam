package role

import (
	"context"

	"freeroam/app/gateway/api/role/v1"
	oRole "freeroam/app/org/api/role/v1"
)

func (c *ControllerV1) ListRole(ctx context.Context, req *v1.ListRoleReq) (res *v1.ListRoleRes, err error) {
	result, err := c.RoleRpcService.ListRole(ctx, &oRole.ListRoleReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Status:   req.Status,
		IsSystem: req.IsSystem,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*v1.GetRoleRes, 0, len(result.List))
	for _, item := range result.List {
		list = append(list, &v1.GetRoleRes{
			Id:       item.Id,
			Code:     item.Code,
			Name:     item.Name,
			Status:   item.Status,
			IsSystem: item.IsSystem,
			Remark:   item.Remark,
			CreateAt: item.CreateAt,
			UpdateAt: item.UpdateAt,
		})
	}

	return &v1.ListRoleRes{
		List:     list,
		Total:    result.Total,
		Page:     result.Page,
		PageSize: result.PageSize,
	}, nil
}
