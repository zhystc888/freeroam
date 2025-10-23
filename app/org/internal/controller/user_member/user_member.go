package user_member

import (
	"context"
	v1 "freeroam/app/org/api/user_member/v1"
	"freeroam/app/org/internal/model"
	"freeroam/app/org/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserMemberServer
	userMember service.IUserMember
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserMemberServer(s.Server, &Controller{
		userMember: service.UserMember(),
	})
}

func (c *Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	params := &model.UserMemberListDto{
		Username: req.Username,
		Name:     req.Name,
		Mobile:   req.Mobile,
		PageReq: &model.PageReq{
			Page:  req.Page,
			Limit: req.Limit,
		},
	}

	params.Gender = nil
	if req.Gender != nil {
		value := req.Gender.GetValue()
		params.Gender = &value
	}

	params.Status = nil
	if req.Status != nil {
		value := req.Status.GetValue()
		params.Status = &value
	}

	result, err := c.userMember.GetList(ctx, params)
	if err != nil {
		return res, err
	}

	list := make([]*v1.GetListItem, len(result.List))

	for i, v := range result.List {
		list[i] = &v1.GetListItem{
			UserId:   v.UserId,
			Username: v.Username,
			Name:     v.Name,
			Mobile:   v.Mobile,
			Gender:   v.Gender,
			Status:   v.Status,
			CreateAt: v.CreateAt.Format("2006-01-02 15:04:05"),
		}
	}

	res = &v1.GetListRes{
		List:  list,
		Total: int64(result.Total),
	}

	return
}

func (c *Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	org, err := c.userMember.GetOne(ctx, req.GetUserId())
	if org != nil {
		res = &v1.GetOneRes{}
		res.UserId = org.UserId
		res.Username = org.Username
		res.Name = org.Name
		res.Mobile = org.Mobile
		res.Gender = org.Gender
		res.Status = org.Status
		res.Super = org.Status
	}
	return
}

func (c *Controller) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.UpdatePasswordRes, err error) {
	return c.userMember.UpdatePassword(ctx, req)
}
