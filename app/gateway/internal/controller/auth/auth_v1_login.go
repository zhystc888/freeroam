package auth

import (
	"context"
	"freeroam/app/gateway/api/auth/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return c.AuthService.Login(ctx, req)
}
