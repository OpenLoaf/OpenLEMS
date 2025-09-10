package c_type

import "common/c_base"

type IFireBasic interface {
	GetFireEnvTemperature() (float64, error)          // 获取消防环境温度
	GetCarbonMonoxideConcentration() (float64, error) // 一氧化碳浓度
	HasSmoke() (bool, error)                          // 是否有烟雾报警
}

type IFire interface {
	c_base.IDriver
	IFireBasic
}
