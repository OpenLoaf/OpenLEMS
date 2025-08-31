package internal

type ITest interface {
	GetProviderPower() (float64, error)          // 获取发电功率
	GetConsumerPower() (float64, error)          // 获取用电功率
	GetFixedSafeCapacityPower() (float64, error) // 获取固定的安全功率(容量功率）

}
