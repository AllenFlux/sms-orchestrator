package main

import (
	_ "sms_backend/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"sms_backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
