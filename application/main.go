package main

import (
	_ "application/internal/logic"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gbuild"

	"application/internal/cmd"
	"context"
)

func main() {
	g.Log().Infof(context.Background(), "当前BuildInfo:%+v", gbuild.Info())
	cmd.Main.Run(context.Background())
}
