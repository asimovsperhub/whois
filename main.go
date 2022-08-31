package main

import (
	_ "dewhois/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"dewhois/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
