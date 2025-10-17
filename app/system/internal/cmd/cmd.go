package cmd

import (
	"context"
	"freeroam/app/system/internal/controller/enum"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := grpcx.Server.New()
			enum.Register(s)
			s.Run()
			return nil
		},
	}
)
