package cmd

import (
	"bbk/app/user/internal/controller/admin_member"
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
)

type JsonRes struct {
	Code    int         `json:"code" x-apifox-mock:"0"`     // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message" x-apifox-mock:"OK"` // 提示信息
	Data    interface{} `json:"data"`                       // 返回数据(业务接口定义具体数据结构)
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := grpcx.Server.New()
			admin_member.Register(s)
			s.Run()
			return nil
		},
	}
)
