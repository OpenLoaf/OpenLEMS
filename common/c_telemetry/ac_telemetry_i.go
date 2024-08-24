package c_telemetry

type IAcTelemetry interface {
	GetPower() (float64, error)         // 有功功率
	GetApparentPower() (float64, error) // 视在功率
	GetReactivePower() (float64, error) // 无功功率
}

type IAcStatisticsQuantity interface {
	IAcTelemetry
	IStatisticsQuantity
}
