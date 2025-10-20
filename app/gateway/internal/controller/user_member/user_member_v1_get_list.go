package user_member

import (
	"context"
	sUserMember "freeroam/app/org/api/user_member/v1"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"freeroam/app/gateway/api/user_member/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	dto := sUserMember.GetListReq{
		Username: req.Username,
		Name:     req.Name,
		Mobile:   req.Mobile,
		Gender:   nil,
		Status:   nil,
		Page:     req.Page,
		Limit:    req.Limit,
	}

	if req.Gender != nil {
		dto.Gender = wrapperspb.Int64(*req.Gender)
	}

	if req.Status != nil {
		dto.Status = wrapperspb.Int64(*req.Status)
	}

	result, err := c.UserMemberRpcService.GetList(ctx, &dto)
	if err != nil {
		return
	}

	list := make([]*v1.GetListItem, len(result.List))

	for i, item := range result.List {
		list[i] = &v1.GetListItem{
			UserId:   item.UserId,
			Username: item.Username,
			Name:     item.Name,
			Mobile:   item.Mobile,
			Gender:   item.Gender,
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
