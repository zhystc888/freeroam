package config

import (
	"context"
	v1 "freeroam/app/system/api/config/v1"
	"freeroam/app/system/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedConfigServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterConfigServer(s.Server, &Controller{})
}

func (*Controller) GetByCode(ctx context.Context, req *v1.GetByCodeReq) (res *v1.GetByCodeRes, err error) {
	return service.Config().GetByCode(ctx, req)
}
