package c_enum

type EProtocolType string // 协议
const (
	EModbusTcp EProtocolType = "modbus_tcp"
	EModbusRtu EProtocolType = "modbus_rtu"
	ECanbus    EProtocolType = "canbus"
	ECanbusUdp EProtocolType = "canbus_udp"
	EGpiod     EProtocolType = "gpiod"
	EGpioIn    EProtocolType = "gpio_in"
	EGpioOut   EProtocolType = "gpio_out"
	EGpioSfs   EProtocolType = "gpiosfs"
)
