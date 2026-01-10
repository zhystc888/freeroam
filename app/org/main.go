package main

import (
	_ "freeroam/app/org/internal/packed"

	_ "freeroam/app/org/internal/logic"

	"freeroam/app/org/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func main() {
	// 设置进程全局时区
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
