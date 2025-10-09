package admin_member

import (
	"context"

	"bbk/app/gateway/api/admin_member/v1"
	uam "bbk/app/user/api/admin_member/v1"
)

func (c *ControllerV1) GetMemberList(ctx context.Context, req *v1.GetMemberListReq) (res *v1.GetMemberListRes, err error) {
	result, err := c.UserRpcService.GetList(ctx, &uam.GetListReq{
		UserId: req.UserID,
		Name:   req.Name,
		Status: req.Status,
		Page:   req.Page,
		Limit:  req.Limit,
	})

	if err != nil {
		return
	}

	list := make([]v1.GetMemberListItem, len(result.List))
	for i, admin := range result.List {
		item := v1.GetMemberListItem{}
		item.UserID = admin.UserId
		item.Name = admin.Name
		item.Status = admin.Status
		item.LastIp = admin.LastIp
		item.LastTime = admin.LastTime
		item.CreateAt = admin.CreateAt
		list[i] = item
	}

	res = &v1.GetMemberListRes{
		List:  list,
		Total: result.Total,
	}

	return
}
