package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetModbusStatusReq struct {
	g.Meta `path:"/remote/modbus/status" method:"get" tags:"远程管理" summary:"获取Modbus服务状态"`
}

type GetModbusStatusRes struct {
	IsRunning       bool                  `json:"isRunning" dc:"Modbus服务是否正在运行"`
	ListenPort      int                   `json:"listenPort" dc:"监听端口"`
	DeviceCount     int                   `json:"deviceCount" dc:"设备数量"`
	ConnectionCount int                   `json:"connectionCount" dc:"连接数量"`
	DeviceStatus    []*ModbusDeviceStatus `json:"deviceStatus" dc:"设备状态列表"`
}

// ModbusDeviceStatus Modbus设备状态结构体
type ModbusDeviceStatus struct {
	DeviceId       string                `json:"deviceId" dc:"设备ID"`
	ModbusId       uint8                 `json:"modbusId" dc:"Modbus从站ID"`
	StartAddr      uint16                `json:"startAddr" dc:"起始地址"`
	IsOnline       bool                  `json:"isOnline" dc:"设备是否在线"`
	LastUpdateTime time.Time             `json:"lastUpdateTime" dc:"最后更新时间"`
	Error          string                `json:"error" dc:"错误信息（如果有）"`
	TotalRegisters uint16                `json:"totalRegisters" dc:"总寄存器数量"`
	PointMappings  []*ModbusPointMapping `json:"pointMappings" dc:"点位映射列表"`
}

// ModbusPointMapping Modbus点位映射结构体
type ModbusPointMapping struct {
	PointKey      string `json:"pointKey" dc:"点位键名"`
	PointName     string `json:"pointName" dc:"点位名称"`
	ValueType     string `json:"valueType" dc:"数据类型"`
	StartOffset   uint16 `json:"startOffset" dc:"相对起始地址的偏移"`
	RegisterCount uint16 `json:"registerCount" dc:"占用寄存器数量"`
	IsSystemPoint bool   `json:"isSystemPoint" dc:"是否为系统固定点位"`
	Description   string `json:"description" dc:"点位描述"`
}
