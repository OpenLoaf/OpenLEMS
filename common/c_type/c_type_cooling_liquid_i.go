package c_type

import (
	"common/c_base"
	"common/c_enum"
)

// ICoolingLiquid 液冷
type ICoolingLiquid interface {
	c_base.IDriver

	GetLiquidCoolingStatus() (*c_enum.ECoolingStatus, error) // 获取液冷状态
	GetInputWaterPressure() (*float32, error)                // 回水压力
	GetInputWaterTemperature() (*float32, error)             // 回水温度
	GetOutputWaterPressure() (*float32, error)               // 出水压力
	GetOutputWaterTemperature() (*float32, error)            // 出水温度
}
