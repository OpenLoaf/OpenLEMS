package internal

// ModbusRtuProtocolConfig rtu的协议配置
type ModbusRtuProtocolConfig struct {
	baudRate int    // 波特率
	dataBits int    // DataBits sets the number of bits per serial character (rtu only)
	parity   string // // Parity: N - None, E - Even, O - Odd (default E)
	stopBits int    // StopBits sets the number of serial stop bits (rtu only)
}

func (m *ModbusRtuProtocolConfig) GetBaudRate() int {
	if m.baudRate == 0 {
		m.baudRate = 9600 // 默认9600
	}
	return m.baudRate
}

func (m *ModbusRtuProtocolConfig) GetDataBits() int {
	if m.dataBits == 0 {
		m.dataBits = 8
	}
	return m.dataBits
}

func (m *ModbusRtuProtocolConfig) GetParity() string {
	if m.parity == "" {
		m.parity = "N"
	}
	return m.parity
}

func (m *ModbusRtuProtocolConfig) GetStopBits() int {
	if m.stopBits == 0 {
		m.stopBits = 1
	}
	return m.stopBits
}
