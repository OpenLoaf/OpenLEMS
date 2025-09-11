package main

import (
	"bms_pylon_tech_us108/bms_pylon_tech_us108_v1"
	"common/c_base"
	"fmt"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(device c_base.IDevice) c_base.IDevice {
	return bms_pylon_tech_us108_v1.NewPlugin(device)
}

func GetDriverInfo() *c_base.SDriverInfo {
	info := bms_pylon_tech_us108_v1.GetDriverInfo()
	info.BuildTime = buildTime
	info.CommitHash = commitHash
	return info
}

func main() {
	fmt.Println(GetDriverInfo())
}

/*
func main() {
	command := c_base.PluginDriverCommand(func() c_base.IDevice {
		return NewPlugin(context.Background())
	})

	// 此处可添加自定义命令
	//command.AddCommand()
	command.Run(context.Background())
}
*/
