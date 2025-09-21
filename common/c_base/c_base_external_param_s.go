package c_base

// SExternalParam 设备对外参数结构体，用于存储设备的对外接口配置
type SExternalParam struct {
	ModbusRegisterAddr uint  `code:"modbusRegisterAddr" name:"Modbus寄存器起始地址" ct:"number" vt:"int" min:"0" max:"65535" def:"40000" desc:"对外提供的Modbus寄存器起始地址"`
	ModbusId           uint8 `code:"modbusId" name:"Modbus设备ID" ct:"number" vt:"int" min:"1" max:"255" def:"1" desc:"对外提供的Modbus设备ID"`
	ModbusAllowControl bool  `code:"modbusAllowControl" name:"允许Modbus控制" ct:"switch" vt:"bool" def:"false" desc:"是否允许通过Modbus进行设备控制"`
}
