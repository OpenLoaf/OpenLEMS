package main

import (
	"common/c_base"
	driver "ess_demo/ess_demo_v1" // 修改此处
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(device c_base.IDevice) c_base.IDevice {
	return driver.NewPlugin(device)
}

func GetDriverInfo() *c_base.SDriverInfo {
	info := driver.GetDriverInfo()
	info.BuildTime = buildTime
	info.CommitHash = commitHash
	return info
}
