package c_device

import "ems-plan/c_base"

// ILiquidCooling 液冷
type ILiquidCooling interface {
	c_base.IDriver

	GetLiquidCoolingStatus() (ECoolingStatus, error) // 获取液冷状态
	GetInputWaterPressure() (float32, error)         // 回水压力
	GetInputWaterTemperature() (float32, error)      // 回水温度
	GetOutputWaterPressure() (float32, error)        // 出水压力
	GetOutputWaterTemperature() (float32, error)     // 出水温度
}
