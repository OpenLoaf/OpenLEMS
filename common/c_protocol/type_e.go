package c_protocol

type EType string // 协议
const (
	EModbusTcp EType = "modbus_tcp"
	EModbusRtu EType = "modbus_rtu"
	ECanbus    EType = "canbus"
	ECanbusTcp EType = "canbus_tcp"
	EGpioSysfs EType = "gpio"
)
