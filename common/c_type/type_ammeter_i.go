package c_type

import "common/c_base"

type IAmmeterBasic interface {
	GetUa() (float32, error)     // A相电压
	GetUb() (float32, error)     // B相电压
	GetUc() (float32, error)     // C相电压
	GetIa() (float32, error)     // A相电流
	GetIb() (float32, error)     // B相电流
	GetIc() (float32, error)     // C相电流
	GetPa() (float32, error)     // A相有功功率
	GetPb() (float32, error)     // B相有功功率
	GetPc() (float32, error)     // C相有功功率
	GetPTotal() (float32, error) // 总有功功率

	GetQa() (float32, error)     // A相无功功率
	GetQb() (float32, error)     // B相无功功率
	GetQc() (float32, error)     // C相无功功率
	GetQTotal() (float32, error) // 总无功功率

	GetSa() (float32, error)     // A相视在功率
	GetSb() (float32, error)     // B相视在功率
	GetSc() (float32, error)     // C相视在功率
	GetSTotal() (float32, error) // 总视在功率

	GetPfa() (float32, error)     // A相功率因数
	GetPfb() (float32, error)     // B相功率因数
	GetPfc() (float32, error)     // C相功率因数
	GetPfTotal() (float32, error) // 总功率因数

	GetPtCt() (float32, float32, error) // PT CT
	GetFrequency() (float32, error)     // 频率

	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量

}

type IAmmeter interface {
	c_base.IDevice
	IAmmeterBasic
}
