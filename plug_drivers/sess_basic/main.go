//go:generate ../build.sh
package main

import (
	"context"
	_ "embed"
	"ems-plan/c_base"
	"station_energy_store/sess_basic_v1"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(ctx context.Context) c_base.IDriver {
	plugin := sess_basic_v1.NewPlugin(ctx)
	plugin.GetDescription().BuildTime = buildTime
	plugin.GetDescription().CommitHash = commitHash
	return plugin
}

func main() {
	command := c_base.PluginDriverCommand(func() c_base.IDriver {
		return NewPlugin(context.Background())
	})

	// 此处可添加自定义命令
	//command.AddCommand()
	command.Run(context.Background())
}
