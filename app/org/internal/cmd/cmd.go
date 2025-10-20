package cmd

import (
	"context"
	"freeroam/app/org/internal/controller/org"
	"freeroam/app/org/internal/controller/user_member"
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
			org.Register(s)
			user_member.Register(s)
			s.Run()
			return nil
		},
	}
)
