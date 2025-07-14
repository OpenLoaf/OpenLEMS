package c_base

type EProtocolType string // 协议
const (
	EModbusTcp EProtocolType = "modbus_tcp"
	EModbusRtu EProtocolType = "modbus_rtu"
	ECanbus    EProtocolType = "canbus"
	ECanbusTcp EProtocolType = "canbus_tcp"
	EGpioSysfs EProtocolType = "gpio"
	EVirtual   EProtocolType = "virtual"
)
