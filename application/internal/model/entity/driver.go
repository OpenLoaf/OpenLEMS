package entity

import "common/c_base"

type SDriver struct {
	DriverName         string `json:"driverName" dc:"驱动名称"`
	DriverVersion      string `json:"driverVersion" dc:"版本号"`
	DriverType         string `json:"driverType" dc:"设备类型"`
	DriverDescription  string `json:"driverDescription" dc:"驱动描述"`
	DriverStatus       bool   `json:"driverStatus" dc:"运行状态"`
	DriverIsEmbedded   bool   `json:"driverIsEmbedded" dc:"运行状态"`
	DriverLastUpdate   string `json:"driverLastUpdate" dc:"更新时间	"`
	DriverPath         string `json:"driverPath" dc:"驱动路径"`
	DriverHash         string `json:"driverHash" dc:"哈希值"`
	DriverFileSizeByte int64  `json:"driverFileSizeByte" dc:"文件大小(byte)"`
	DriverImage        string `json:"driverImage" dc:"base64图片"`

	ProtocolType string `json:"protocolType" dc:"协议类型"`

	// 扩展字段（参考 common/c_base/driver_description_s_f.go 与 driver_info_s.go）
	Brand      string                     `json:"brand" dc:"品牌"`
	Model      string                     `json:"model" dc:"型号"`
	BuildTime  string                     `json:"buildTime" dc:"编译时间"`
	CommitHash string                     `json:"commitHash" dc:"提交哈希"`
	Remark     string                     `json:"remark" dc:"备注"`
	Author     string                     `json:"author" dc:"作者"`
	Telemetry  []*c_base.SFieldDefinition `json:"telemetry" dc:"遥测描述列表"`
}

// type SDriverDetail struct {
// 	DriverName        string `json:"driverName" dc:"驱动名称"`
// 	DriverVersion     string `json:"driverVersion" dc:"版本号"`
// 	DriverType        string `json:"driverType" dc:"设备类型"`
// 	DriverDescription string `json:"driverDescription" dc:"驱动描述"`
// 	DriverStatus      bool   `json:"driverStatus" dc:"运行状态"`
// 	DriverLastUpdate  string `json:"driverLastUpdate" dc:"更新时间	"`
// 	ProtocolType      string `json:"protocolType" dc:"协议类型"`

// }
