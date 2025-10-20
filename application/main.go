package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gbuild"

	"application/internal/cmd"
	"hexlib"
)

func main() {
	ctx := context.Background()
	g.Log().Infof(ctx, "当前BuildInfo:%+v", gbuild.Info())
	g.Log().Infof(ctx, "HexVersion: %s", hexlib.HexVersion())
	fmt.Println("HexVersion:", hexlib.HexVersion())
	cmd.Main.Run(ctx)
}
