package c_device

import "ems-plan/c_telemetry"

type ECoolingStatus int

const (
	ECoolingStatusStop    ECoolingStatus = iota // 0:停机
	ECoolingStatusCool                          // 1:制冷
	ECoolingStatusHeat                          // 2:制热
	ECoolingStatusRecycle                       // 3:自循环
)

type ICoolingBasic interface {
	c_telemetry.IAlarmHandler

	SetTemperature(temperature float32) error // 设置温度

	GetTemperature() (float32, error)       // 当前温度
	GetTargetTemperature() (float32, error) // 目标温度
}
