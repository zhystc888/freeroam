package position

import (
	"context"
	v1 "freeroam/app/org/api/position/v1"
	"freeroam/app/org/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedPositionServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPositionServer(s.Server, &Controller{})
}

func (*Controller) ListPosition(ctx context.Context, req *v1.ListPositionReq) (res *v1.ListPositionRes, err error) {
	return service.Position().ListPosition(ctx, req)
}

func (*Controller) GetPosition(ctx context.Context, req *v1.GetPositionReq) (res *v1.GetPositionRes, err error) {
	return service.Position().GetPosition(ctx, req)
}

func (*Controller) CreatePosition(ctx context.Context, req *v1.CreatePositionReq) (res *v1.CreatePositionRes, err error) {
	return service.Position().CreatePosition(ctx, req)
}

func (*Controller) UpdatePosition(ctx context.Context, req *v1.UpdatePositionReq) (res *v1.UpdatePositionRes, err error) {
	return service.Position().UpdatePosition(ctx, req)
}

func (*Controller) DeletePosition(ctx context.Context, req *v1.DeletePositionReq) (res *v1.DeletePositionRes, err error) {
	return service.Position().DeletePosition(ctx, req)
}

func (*Controller) GetPositionOptions(ctx context.Context, req *v1.GetPositionOptionsReq) (res *v1.GetPositionOptionsRes, err error) {
	return service.Position().GetPositionOptions(ctx, req)
}
