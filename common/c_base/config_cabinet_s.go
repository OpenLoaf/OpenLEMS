package c_base

type SCabinetConfig struct {
	CabinetId    uint8             // 柜子ID, 不同柜子ID对应不同对柜子
	Name         string            // 设备名称
	Driver       string            // 驱动名称，不需要带版本号
	Enable       bool              // 是否启用
	ModbusEnable bool              // 是否启用modbus
	Params       map[string]string // 额外参数
}
