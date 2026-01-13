package enum

import (
	"context"
	v1 "freeroam/app/system/api/enum/v1"
	"freeroam/app/system/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedEnumServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterEnumServer(s.Server, &Controller{})
}

func (*Controller) GetByTypeAndCode(ctx context.Context, req *v1.GetByTypeAndCodeReq) (res *v1.GetByTypeAndCodeRes, err error) {
	return service.Enum().GetByTypeAndCode(ctx, req)
}

func (*Controller) GetByType(ctx context.Context, req *v1.GetByTypeReq) (res *v1.GetByTypeRes, err error) {
	return service.Enum().GetByType(ctx, req)
}
