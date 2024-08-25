package c_device

import "ems-plan/c_telemetry"

type EOutputType = string

// 定义基础的输出设备类型
const (
	ERunningOutput EOutputType = "RunningOutput" // 运行输出
	EWarningOutput EOutputType = "WarningOutput" // 告警输出
)

type IOutput interface {
	IInfo
	c_telemetry.IAlarmHandler
	SetUp() (bool, error)
	SetDown() (bool, error)
}
