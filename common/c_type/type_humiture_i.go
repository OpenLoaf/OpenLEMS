package c_type

import "common/c_base"

type IHumitureBasic interface {
	GetTemperature() (float64, error) // 获取温度
	GetHumidity() (float64, error)    // 湿度
}

type IHumiture interface {
	c_base.IDriver
	IHumitureBasic
}
