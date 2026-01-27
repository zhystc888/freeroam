package cmd

import (
	"context"
	"freeroam/app/system/internal/controller/config"
	"freeroam/app/system/internal/controller/enum"
	"freeroam/common/interceptor/cgrpcx"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
					cgrpcx.UnaryServerInterceptorInjectJwtClaims,
				)}...,
			)

			s := grpcx.Server.New(c)

			config.Register(s)
			enum.Register(s)

			// grpc 反射，通过 url 直接获取 grpc 接口信息
			reflection.Register(s.Server)
			s.Run()
			return nil
		},
	}
)
