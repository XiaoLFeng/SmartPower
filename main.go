package main

import (
	_ "SmartPower/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"SmartPower/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
