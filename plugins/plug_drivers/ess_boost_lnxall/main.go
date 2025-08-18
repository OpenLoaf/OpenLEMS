package main

import (
	"common/c_base"
	"context"
	"ess_lnxall/ess_boost_lnxall_v1"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(ctx context.Context) c_base.IDevice {
	plugin := ess_boost_lnxall_v1.NewPlugin(ctx)

	plugin.GetDriverDescription().BuildTime = buildTime
	plugin.GetDriverDescription().CommitHash = commitHash
	return plugin
}
