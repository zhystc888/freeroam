package role

import (
	"context"
	v1 "freeroam/app/org/api/role/v1"
	"freeroam/app/org/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedRoleServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterRoleServer(s.Server, &Controller{})
}

func (*Controller) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error) {
	return service.Role().CreateRole(ctx, req)
}

func (*Controller) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (res *v1.UpdateRoleRes, err error) {
	return service.Role().UpdateRole(ctx, req)
}

func (*Controller) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {
	return service.Role().DeleteRole(ctx, req)
}

func (*Controller) GetRole(ctx context.Context, req *v1.GetRoleReq) (res *v1.GetRoleRes, err error) {
	return service.Role().GetRole(ctx, req)
}

func (*Controller) ListRole(ctx context.Context, req *v1.ListRoleReq) (res *v1.ListRoleRes, err error) {
	return service.Role().ListRole(ctx, req)
}

func (*Controller) GetRolePositionList(ctx context.Context, req *v1.GetRolePositionListReq) (res *v1.GetRolePositionListRes, err error) {
	return service.Role().GetRolePositionList(ctx, req)
}

func (*Controller) BatchAssignRolePosition(ctx context.Context, req *v1.BatchAssignRolePositionReq) (res *v1.BatchAssignRolePositionRes, err error) {
	return service.Role().BatchAssignRolePosition(ctx, req)
}

func (*Controller) GetRolePositionIds(ctx context.Context, req *v1.GetRolePositionIdsReq) (res *v1.GetRolePositionIdsRes, err error) {
	return service.Role().GetRolePositionIds(ctx, req)
}
