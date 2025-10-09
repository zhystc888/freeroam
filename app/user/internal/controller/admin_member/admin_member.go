package admin_member

import (
	v1 "bbk/app/user/api/admin_member/v1"
	"bbk/app/user/internal/model"
	"bbk/app/user/internal/service"
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedAdminMemberServer
	AdminMemberServer service.IAdminMember
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAdminMemberServer(s.Server, newAdminMemberController())
}

func newAdminMemberController() *Controller {
	return &Controller{
		AdminMemberServer: service.AdminMember(),
	}
}

func (c *Controller) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	member, err := c.AdminMemberServer.GetMember(ctx, req.UserId)
	if member != nil {
		res = &v1.GetRes{}
		res.Name = member.Name
		res.UserId = member.UserID
		res.Status = member.Status
		res.Username = member.Username
	}
	return
}

func (c *Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	params := &model.AdminGetMemberListDto{
		UserID: req.UserId,
		Name:   req.Name,
		Status: req.Status,
		PageReq: model.PageReq{
			Page:  req.Page,
			Limit: req.Limit,
		},
	}
	list, total, err := c.AdminMemberServer.GetMemberList(ctx, params)
	if err != nil {
		return res, err
	}

	res = &v1.GetListRes{
		List:  []*v1.GetListItem{},
		Total: int64(total),
	}

	for _, v := range list {
		item := &v1.GetListItem{}
		item.UserId = v.UserID
		item.Name = v.Name
		item.Status = v.Status
		item.CreateAt = v.CreateAt.String()
		item.LastIp = v.User.LastIp
		item.LastTime = v.User.LastTime.String()
		res.List = append(res.List, item)
	}

	return res, nil
}
