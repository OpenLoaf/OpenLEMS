package c_device

import "ems-plan/c_telemetry"

type EInputType = string

// 定义基础的输入设备类型
const (
	EScramButton     EInputType = "ScramButton"     // 急停按钮
	EChargeButton    EInputType = "ChargeButton"    // 充电按钮
	EDischargeButton EInputType = "DischargeButton" // 放电按钮
)

type IInput interface {
	IInfo
	c_telemetry.IAlarmHandler

	IsUp() (bool, error)
	IsDown() (bool, error)
}
