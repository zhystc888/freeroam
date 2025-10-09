package org

import (
	v1 "bbk/app/org/api/org/v1"
	"bbk/app/org/internal/service"
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
