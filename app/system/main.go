package main

import (
	_ "freeroam/app/system/boot"

	_ "freeroam/app/system/internal/logic"

	"freeroam/app/system/internal/cmd"
	_ "freeroam/app/system/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func main() {
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
