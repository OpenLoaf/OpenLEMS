//go:build linux && arm

package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// startGui 启动GUI界面（不支持平台）
func startGui(ctx context.Context) {
	g.Log().Warningf(ctx, "GUI界面在此平台上不受支持")
}

// startGuiFullscreen 启动全屏GUI界面（不支持平台）
func startGuiFullscreen(ctx context.Context) {
	g.Log().Warningf(ctx, "GUI界面在此平台上不受支持")
}
