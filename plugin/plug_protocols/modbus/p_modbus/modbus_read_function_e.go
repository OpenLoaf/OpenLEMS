package p_modbus

// TODO 改名字加E
type ModbusReadFunction = uint

// 查询的方法
const (
	MqReadCoils        ModbusReadFunction = iota // 01 线圈（0x)
	MqDiscreteInputs                             // 02 离散输入（1x)
	MqHoldingRegisters                           // 03 保持寄存器（4x)
	MqInputRegisters                             // 04 输入寄存器（3x)
)
