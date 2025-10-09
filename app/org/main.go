package main

import (
	_ "bbk/app/org/internal/packed"
	"github.com/gogf/gf/v2/os/gtime"

	_ "bbk/app/org/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	_ "bbk/app/org/boot"

	"bbk/app/org/internal/cmd"
)

func main() {
	// 设置进程全局时区
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
