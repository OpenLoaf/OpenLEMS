package c_proto

// SModbusProtocolRtuConfig rtu的协议配置
type SModbusProtocolRtuConfig struct {
	BaudRate int    `json:"baudRate" name:"波特率"  ct:"number" vt:"int" min:"1200" max:"115200" step:"100" default:"9600" unit:"bps" dc:"串口通信波特率，常用值：9600, 19200, 38400, 57600, 115200"`
	DataBits int    `json:"dataBits" name:"数据位" ct:"number" vt:"int" min:"5" max:"8" step:"1" default:"8" dc:"每个串口字符的数据位数，通常为8位"`
	Parity   string `json:"parity" name:"校验位" ct:"singleSelect" vt:"string" selectOptions:"N:无校验,E:偶校验,O:奇校验" default:"N" dc:"串口校验位：N-无校验，E-偶校验，O-奇校验"`
	StopBits int    `json:"stopBits" name:"停止位" ct:"number" vt:"int" min:"1" max:"2" step:"1" default:"1" dc:"串口停止位数，通常为1位或2位"`
}

// SModbusDeviceConfig modbus的基本配置
type SModbusDeviceConfig struct {
	UnitId uint8 `json:"unitId" name:"ModbusID" required:"true" min:"1" max:"255" default:"1" required:"true"  step:"1"` // 单元ID
}
