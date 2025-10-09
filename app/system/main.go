package main

import (
	_ "bbk/app/system/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"bbk/app/system/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
