package main

import (
	_ "bbk/app/system/internal/logic"
	_ "bbk/app/system/internal/packed"

	"github.com/gogf/gf/v2/os/gtime"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	_ "bbk/app/system/boot"
	"bbk/app/system/internal/cmd"
)

func main() {
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
