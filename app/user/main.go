package main

import (
	_ "bbk/app/user/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "bbk/app/user/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"

	_ "bbk/app/user/boot"
	"bbk/app/user/internal/cmd"
)

func main() {
	// 设置进程全局时区
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
