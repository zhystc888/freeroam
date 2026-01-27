package permission

import (
	"context"
	v1 "freeroam/app/system/api/permission/v1"
	"freeroam/app/system/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedPermissionServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPermissionServer(s.Server, &Controller{})
}

func (*Controller) AssignRolePermissions(ctx context.Context, req *v1.AssignRolePermissionsReq) (res *v1.AssignRolePermissionsRes, err error) {
	return service.Permission().AssignRolePermissions(ctx, req)
}

func (*Controller) GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsReq) (res *v1.GetRolePermissionsRes, err error) {
	return service.Permission().GetRolePermissions(ctx, req)
}

func (*Controller) GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (res *v1.GetPermissionTreeRes, err error) {
	return service.Permission().GetPermissionTree(ctx, req)
}

func (*Controller) GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (res *v1.GetFrontPermissionsRes, err error) {
	return service.Permission().GetFrontPermissions(ctx, req)
}
