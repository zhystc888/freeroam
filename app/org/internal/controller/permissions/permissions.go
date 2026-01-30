package permissions

import (
	"context"
	v1 "freeroam/app/org/api/permissions/v1"
	"freeroam/app/org/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedPermissionsServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPermissionsServer(s.Server, &Controller{})
}

func (*Controller) AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (res *v1.AssignRolePermissionsRes, err error) {
	return service.Permissions().AssignRolePermissions(ctx, req)
}

func (*Controller) GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (res *v1.GetRolePermissionsRes, err error) {
	return service.Permissions().GetRolePermissions(ctx, req)
}

func (*Controller) GetPermissionsTree(ctx context.Context, req *v1.GetPermissionsTreeReq) (res *v1.GetPermissionsTreeRes, err error) {
	return service.Permissions().GetPermissionsTree(ctx, req)
}

func (*Controller) GetMemberPermissions(ctx context.Context, req *v1.GetMemberPermissionsReq) (res *v1.GetMemberPermissionsRes, err error) {
	return service.Permissions().GetMemberPermissions(ctx, req)
}
