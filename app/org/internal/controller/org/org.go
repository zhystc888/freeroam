package org

import (
	"context"
	v1 "freeroam/app/org/api/org/v1"
	"freeroam/app/org/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedOrgServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterOrgServer(s.Server, &Controller{})
}

func (*Controller) GetOrgTree(ctx context.Context, req *v1.GetOrgTreeReq) (res *v1.GetOrgTreeRes, err error) {
	return service.Org().GetOrgTree(ctx, req)
}

func (*Controller) ListOrg(ctx context.Context, req *v1.ListOrgReq) (res *v1.ListOrgRes, err error) {
	return service.Org().ListOrg(ctx, req)
}

func (*Controller) GetOrg(ctx context.Context, req *v1.GetOrgReq) (res *v1.GetOrgRes, err error) {
	return service.Org().GetOrg(ctx, req)
}

func (*Controller) CreateOrg(ctx context.Context, req *v1.CreateOrgReq) (res *v1.CreateOrgRes, err error) {
	return service.Org().CreateOrg(ctx, req)
}

func (*Controller) UpdateOrg(ctx context.Context, req *v1.UpdateOrgReq) (res *v1.UpdateOrgRes, err error) {
	return service.Org().UpdateOrg(ctx, req)
}

func (*Controller) DeleteOrg(ctx context.Context, req *v1.DeleteOrgReq) (res *v1.DeleteOrgRes, err error) {
	return service.Org().DeleteOrg(ctx, req)
}

func (*Controller) MoveOrg(ctx context.Context, req *v1.MoveOrgReq) (res *v1.MoveOrgRes, err error) {
	return service.Org().MoveOrg(ctx, req)
}
