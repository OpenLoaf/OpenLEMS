//go:generate ../build.sh
package main

import (
	"common/c_base"
	"context"
	_ "embed"
	"starCharge100E_v1/pcs_star_charge_100E_v1"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(ctx context.Context) c_base.IDevice {
	plugin := pcs_star_charge_100E_v1.NewPlugin(ctx)

	plugin.GetDriverDescription().BuildTime = buildTime
	plugin.GetDriverDescription().CommitHash = commitHash
	return plugin
}

func main() {
	ctx := context.Background()
	command := c_base.PluginDriverCommand(func() c_base.IDevice {
		return NewPlugin(ctx)
	})

	// 此处可添加自定义命令
	//command.AddCommand()
	command.Run(ctx)
}
