package main

import (
	_ "gf_star/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_star/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
