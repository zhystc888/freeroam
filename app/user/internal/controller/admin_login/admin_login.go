package admin_login

import (
	v1 "bbk/app/user/api/admin_login/v1"
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedAdminLoginServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAdminLoginServer(s.Server, &Controller{})
}

func (*Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
