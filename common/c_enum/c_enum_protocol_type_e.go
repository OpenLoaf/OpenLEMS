package c_enum

type EProtocolType string // 协议
const (
	EModbusTcp EProtocolType = "modbus_tcp"
	EModbusRtu EProtocolType = "modbus_rtu"
	ECanbus    EProtocolType = "canbus"
	ECanbusUdp EProtocolType = "canbus_udp"
	EGpioIn    EProtocolType = "gpio_in"
	EGpioOut   EProtocolType = "gpio_out"
	EIec101    EProtocolType = "iec101"
	EIec102    EProtocolType = "iec102"
	EIec104    EProtocolType = "iec104"
)
