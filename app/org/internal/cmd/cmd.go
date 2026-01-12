package cmd

import (
	"context"
	"freeroam/app/org/internal/controller/auth"
	"freeroam/common/interceptor/cgrpcx"

	"freeroam/app/org/internal/controller/role"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					cgrpcx.ErrorLogInterceptor,
				)}...,
			)

			s := grpcx.Server.New(c)

			// 注册 controller
			role.Register(s)
			auth.Register(s)
			s.Run()
			return nil
		},
	}
)
