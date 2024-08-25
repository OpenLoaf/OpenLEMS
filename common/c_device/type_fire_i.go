package c_device

import "ems-plan/c_telemetry"

type IFireBasic interface {
	GetFireEnvTemperature() (float64, error)          // 获取消防环境温度
	GetCarbonMonoxideConcentration() (float64, error) // 一氧化碳浓度
	HasSmoke() (bool, error)                          // 是否有烟雾报警
}

type IFire interface {
	IInfo
	IFireBasic
	c_telemetry.IAlarmHandler
}
