package main

import (
	_ "gtoolboxs/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gtoolboxs/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
