package cmd

import (
	"context"
	"freeroam/app/gateway/internal/controller/enum"
	"freeroam/app/gateway/internal/controller/org"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					org.NewV1(),
					enum.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
