package entity

type SDriver struct {
	DriverName         string `json:"driverName" dc:"驱动名称"`
	DriverVersion      string `json:"driverVersion" dc:"版本号"`
	DriverProtocolType string `json:"driverProtocolType" dc:"协议类型"`
	DriverDescription  string `json:"driverDescription" dc:"驱动描述"`
	DriverStatus       bool   `json:"driverStatus" dc:"状态"`
	DriverLastUpdate   string `json:"driverLastUpdate" dc:"更新时间	"`
}
