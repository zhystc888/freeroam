package main

import (
	_ "freeroam/app/gateway/boot"
	_ "freeroam/app/gateway/internal/packed"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/os/gctx"

	"freeroam/app/gateway/internal/cmd"
)

func main() {
	// 设置进程全局时区
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
