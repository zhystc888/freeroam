package permissions

import (
	"context"
	"freeroam/app/gateway/api/permissions/v1"
	"freeroam/app/gateway/internal/service"
)

func (c *ControllerV1) GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (res *v1.GetFrontPermissionsRes, err error) {
	return service.Permissions().GetFrontPermissions(ctx, req)
}
