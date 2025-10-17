package org

import (
	v1 "bbk/app/org/api/org/v1"
	"bbk/app/org/internal/model"
	"bbk/app/org/internal/service"
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedOrgServer
	OrgServer service.IOrg
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterOrgServer(s.Server, &Controller{
		OrgServer: service.Org(),
	})
}

func (c *Controller) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	org, err := c.OrgServer.Get(ctx, req.GetId())
	if org != nil {
		res = &v1.GetRes{}
		res.Id = org.Id
		res.ParentId = org.ParentId
		res.Name = org.Name
		res.Type = org.Type
		res.Code = org.Code
		res.Status = org.Status
		res.Supervisors = make([]*v1.Supervisor, len(org.Supervisors))
		for k, v := range org.Supervisors {
			res.Supervisors[k] = &v1.Supervisor{
				Id:   v.User.UserId,
				Name: v.User.Name,
			}
		}
	}
	return
}

func (c *Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	params := &model.OrgListDto{
		Name: req.Name,
		Code: req.Code,
		Type: req.Type,
		PageReq: &model.PageReq{
			Page:  req.Page,
			Limit: req.Limit,
		},
	}

	params.ParentId = nil

	if req.ParentId != nil {
		parentIdValue := req.ParentId.GetValue()
		params.ParentId = &parentIdValue
	}

	result, err := c.OrgServer.GetList(ctx, params)
	if err != nil {
		return res, err
	}

	list := make([]*v1.GetListItem, len(result.List))

	for i, v := range result.List {
		list[i] = &v1.GetListItem{
			Id:       v.Id,
			Name:     v.Name,
			Code:     v.Code,
			Type:     v.Type,
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
