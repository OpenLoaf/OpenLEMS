package c_base

type EProtocolType string // 协议
const (
	EModbusTcp EProtocolType = "modbus_tcp"
	EModbusRtu EProtocolType = "modbus_rtu"
	ECanbus    EProtocolType = "canbus"
	ECanbusUdp EProtocolType = "canbus_udp"
	EGpioSysfs EProtocolType = "gpio"
)
