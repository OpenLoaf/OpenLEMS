package c_device

type IHumitureBasic interface {
	GetTemperature() (float64, error) // 获取温度
	GetHumidity() (float64, error)    // 湿度
}

type IHumiture interface {
	IInfo
	IHumitureBasic
}
