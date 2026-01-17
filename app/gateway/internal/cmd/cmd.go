package cmd

import (
	"context"
	"freeroam/app/gateway/internal/controller/auth"
	"freeroam/app/gateway/internal/controller/enum"
	"freeroam/app/gateway/internal/controller/org"
	"freeroam/app/gateway/internal/controller/position"
	"freeroam/app/gateway/internal/controller/role"
	"freeroam/app/gateway/internal/middleware"
	cMiddleware "freeroam/common/middleware"

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
				group.Middleware(ghttp.MiddlewareHandlerResponse).Middleware(cMiddleware.ErrorHandler)
				nonAuthGroup := group.Clone()
				nonAuthGroup.Bind(
					auth.NewV1(),
				)

				AuthGroup := group.Clone()
				AuthGroup.Middleware(middleware.AuthSign).Bind(
					enum.NewV1(),
					role.NewV1(),
					org.NewV1(),
					position.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
