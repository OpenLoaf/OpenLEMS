//go:generate stringer -type=EModbusReadFunction -trimprefix=EMq -output=c_modbus_read_function_e_string.go
package c_modbus

type EModbusReadFunction uint8

// 查询的方法
const (
	EMqReadCoils        EModbusReadFunction = iota // 01 线圈（0x)
	EMqDiscreteInputs                              // 02 离散输入（1x)
	EMqHoldingRegisters                            // 03 保持寄存器（4x)
	EMqInputRegisters                              // 04 输入寄存器（3x)
)
