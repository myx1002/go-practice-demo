package main

import (
	_ "proxima/app/user/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima/app/user/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
