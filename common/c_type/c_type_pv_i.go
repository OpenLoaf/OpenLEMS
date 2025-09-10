package c_type

import (
	"common/c_base"
	"common/c_enum"
)

type IPvBase interface {
	GetStatus() (c_enum.EPvStatus, error)

	SetPower(power float64) error             // 设置有功功率
	SetReactivePower(power float64) error     // 设置无功功率
	SetPowerFactor(factor float32) error      // 设置功率因数
	GetTargetPower() (float64, error)         // 获取目标有功功率
	GetTargetReactivePower() (float64, error) // 获取目标无功功率
	GetTargetPowerFactor() (float32, error)   // 获取目标功率因数

	GetPower() (float64, error)         // 有功功率
	GetApparentPower() (float64, error) // 视在功率
	GetReactivePower() (float64, error) // 无功功率

	GetDcPower() (float64, error)   // 直流功率
	GetDcVoltage() (float64, error) // 直流电压
	GetDcCurrent() (float64, error) // 直流电流

	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量
	GetCapacity() (uint32, error)                 // 获取装机容量

	GetTemperature() (uint32, error) // 获取温度
	GetIrradiance() (uint32, error)  // 获取辐射值
}

type IPv interface {
	c_base.IDriver
	IPvBase
}
