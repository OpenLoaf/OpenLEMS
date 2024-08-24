package c_telemetry

type IDcTelemetry interface {
	GetDcPower() (float64, error)   // 直流功率
	GetDcVoltage() (float64, error) // 直流电压
	GetDcCurrent() (float64, error) // 直流电流
}

type IDcStatisticsQuantity interface {
	IDcTelemetry
	IStatisticsQuantity
}
