package v1

import (
	"common/c_base"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceDetailedListReq struct {
	g.Meta      `path:"/device/detailed-list" method:"get" tags:"设备相关" summary:"获取设备详细列表"`
	DeviceTypes []c_enum.EDeviceType `json:"deviceTypes,omitempty" dc:"设备类型过滤，可设置多个类型。可选值：ammeter(电表),ca(制冷空调),cl(液冷机组),bms(电池管理系统),fire(消防),hum(温湿度),pcs(电池逆变器),load(负载),pv(光伏),ess(储能柜),cp(充电桩),gen(发电机),gpio(DIY),sess(总站储能柜)"`
	Enabled     *bool                `json:"enabled,omitempty" dc:"设备启用状态过滤，true-只显示启用的设备，false-只显示禁用的设备，不传-显示所有设备"`
}

type DeviceDetailedInfo struct {
	Id                 string         `json:"id" dc:"设备ID"`
	Pid                string         `json:"pid" dc:"父设备ID"`
	Name               string         `json:"name" dc:"设备名称"`
	ProtocolId         string         `json:"protocolId" dc:"协议配置ID"`
	Driver             string         `json:"driver" dc:"驱动名称"`
	LogLevel           string         `json:"logLevel" dc:"日志等级"`
	Strategy           string         `json:"strategy" dc:"策略名称"`
	StorageEnable      bool           `json:"storageEnable" dc:"是否存储"`
	StorageIntervalSec int32          `json:"storageIntervalSec" dc:"存储间隔(秒)"`
	Sort               int            `json:"sort" dc:"排序"`
	Enabled            bool           `json:"enabled" dc:"是否启用"`
	Params             map[string]any `json:"params" dc:"设备参数"`
	CreatedAt          *time.Time     `json:"createdAt" dc:"创建时间"`
	UpdatedAt          *time.Time     `json:"updatedAt" dc:"更新时间"`

	// 设备类型相关信息
	DeviceType      c_enum.EDeviceType `json:"deviceType" dc:"设备类型"`
	IsVirtualDevice bool               `json:"isVirtualDevice" dc:"是否是虚拟设备"`
	DriverType      string             `json:"driverType" dc:"驱动类型"`
	DriverBrand     string             `json:"driverBrand" dc:"驱动品牌"`
	DriverModel     string             `json:"driverModel" dc:"驱动型号"`
	DriverVersion   string             `json:"driverVersion" dc:"驱动版本"`

	// 协议相关信息
	ProtocolName    string `json:"protocolName" dc:"协议名称"`
	ProtocolType    string `json:"protocolType" dc:"协议类型"`
	ProtocolAddress string `json:"protocolAddress" dc:"协议地址"`

	// 状态信息
	FailedMessage string `json:"failedMessage,omitempty" dc:"初始化失败原因"`
}

type GetDeviceDetailedListRes struct {
	Devices []*DeviceDetailedInfo `json:"devices" dc:"设备详细列表"`
	Total   int                   `json:"total" dc:"设备总数"`
}
