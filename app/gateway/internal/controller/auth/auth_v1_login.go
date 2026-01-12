package auth

import (
	"context"
	"fmt"

	"freeroam/app/gateway/api/auth/v1"

	"golang.org/x/crypto/bcrypt"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	return c.AuthService.Login(ctx, req)
}
