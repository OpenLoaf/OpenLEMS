package common

import (
	"common/c_base"
	"context"
)

type IManager interface {
	Start(ctx context.Context)   // 启动服务
	Shutdown()                   // 停止管理器（释放资源、退出 goroutine）
	Cleanup() error              // 清理过期/无效资源（定时调用）
	Status() c_base.EServerState // 返回运行状态
}

/*
type IDeviceManager interface {
	IManager
	Init() error

	GetDeviceConfigTree() []*c_base.SDeviceConfig        // 获取设备配置树
	GetDeviceConfigById(id string) *c_base.SDeviceConfig // 通过id获取设备配置

	GetDeviceById(id string) c_base.IDevice                // 通过Id获取设备
	GetDeviceByType(t c_base.EDeviceType) []c_base.IDevice //通过 类型获取设备类别

	StartDevice(deviceType c_base.EDeviceType) error // 启动设备
	StopDeviceById(id string) error
}

type StrategyManager interface {
	IManager
}

type AlarmManager interface {
	IManager
}

var deviceManager IDeviceManager

func RegisterDeviceManager(d IDeviceManager) {

}

func GetDeviceManager() {

}
*/
