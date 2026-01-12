package auth

import (
	"context"

	"freeroam/app/gateway/api/auth/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return c.AuthService.Logout(ctx, req)
}
