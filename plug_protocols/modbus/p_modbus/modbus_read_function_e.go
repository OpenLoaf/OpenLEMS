package p_modbus

// TODO 改名字加E
type ModbusReadFunction = uint

// 查询的方法
const (
	MqReadCoils ModbusReadFunction = iota
	MqDiscreteInputs
	MqHoldingRegisters
	MqInputRegisters
)
