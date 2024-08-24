package c_telemetry

type IStatisticsIncomingQuantity interface {
	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
}

type IStatisticsOutgoingQuantity interface {
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量
}

type IStatisticsQuantity interface {
	IStatisticsIncomingQuantity
	IStatisticsOutgoingQuantity
}
