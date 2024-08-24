package util

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"plugin"
)

func OpenPlugin(ctx context.Context, path string) (plugin.Symbol, error) {
	return openPlugin(ctx, path, "NewPlugin")
}

func openPlugin(ctx context.Context, path string, symName string) (plugin.Symbol, error) {
	g.Log().Infof(ctx, "加载插件：%s", path)
	// 打开插件
	p, err := plugin.Open(path)
	if err != nil {
		g.Log().Errorf(ctx, "打开插件失败：%v", err)
		panic(err)
	}

	// 查找并使用结构体和函数
	return p.Lookup(symName)
}
