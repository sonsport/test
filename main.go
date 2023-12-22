package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"fuya-ark/internal/cmd"
	_ "fuya-ark/internal/logic"
	_ "fuya-ark/internal/mq"
	_ "fuya-ark/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
