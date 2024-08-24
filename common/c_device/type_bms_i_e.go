package c_device

import "ems-plan/c_telemetry"

type EBmsStatus int

const (
	EBmsStatusUnknown EBmsStatus = iota
	EBmsStatusOff
	EBmsStatusStandby
	EBmsStatusCharge
	EBmsStatusDischarge
	EBmsStatusFault
)

type IBmsBasic interface {
	c_telemetry.IPowerLimit
	c_telemetry.IDcStatisticsQuantity

	SetReset() error                      // 复位
	SetBmsStatus(status EBmsStatus) error // 设置BMS状态

	GetBmsStatus() (EBmsStatus, error)                  // BMS状态
	GetSoc() (float32, error)                           // 电池当前容量 %
	GetSoh() (float32, error)                           // 电池健康 %
	GetCellTemp() (float32, float32, float32, error)    // 电芯最低温度, 电芯最高温度, 电芯平均温度
	GetCellVoltage() (float32, float32, float32, error) // 电芯最低电压, 电芯最高电压, 电芯平均电压
	GetCapacity() (uint16, error)                       // 电池容量kWh
	GetCycleCount() (uint, error)                       // 循环次数
}

type IBms interface {
	IInfo
	IBmsBasic
}
