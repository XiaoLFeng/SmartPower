package main

import (
	_ "SmartPower/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mssql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"SmartPower/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
