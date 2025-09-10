package c_type

import (
	"common/c_base"
	"common/c_enum"
)

type IBmsBasic interface {
	SetReset() error // 复位

	SetBmsStatus(status c_enum.EBmsStatus) error // 设置BMS状态

	GetCellMinTemp() (float32, error)    // 电芯最低温度
	GetCellMaxTemp() (float32, error)    // 电芯最高温度
	GetCellAvgTemp() (float32, error)    // 电芯平均温度
	GetCellMinVoltage() (float32, error) // 电芯最低电压
	GetCellMaxVoltage() (float32, error) // 电芯最高电压
	GetCellAvgVoltage() (float32, error) // 电芯平均电压

	GetBmsStatus() (c_enum.EBmsStatus, error) // BMS状态
	GetSoc() (float32, error)                 // 电池当前容量 %
	GetSoh() (float32, error)                 // 电池健康 %
	GetCapacity() (uint32, error)             // 电池容量kWh
	GetCycleCount() (uint, error)             // 循环次数

	GetRatedPower() (uint32, error)      // 额定功率
	GetMaxInputPower() (float32, error)  // 最大充电功率、最大输入功率限制
	GetMaxOutputPower() (float32, error) // 最大放电功率、最大输出功率限制

	GetDcPower() (float64, error)   // 直流功率
	GetDcVoltage() (float64, error) // 直流电压
	GetDcCurrent() (float64, error) // 直流电流

	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量

}

type IBms interface {
	c_base.IDriver
	IBmsBasic
}

//type Bms
