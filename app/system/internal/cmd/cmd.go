package cmd

import (
	"context"
	"freeroam/app/system/internal/controller/config"
	"freeroam/app/system/internal/controller/enum"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc/reflection"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := grpcx.Server.New()

			config.Register(s)
			enum.Register(s)

			// grpc 反射，通过 url 直接获取 grpc 接口信息
			reflection.Register(s.Server)
			s.Run()
			return nil
		},
	}
)
