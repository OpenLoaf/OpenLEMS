package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 创建忽略告警 ====================

type CreateAlarmIgnoreReq struct {
	g.Meta         `path:"/alarm/ignore" method:"post" tags:"告警" summary:"创建忽略告警"`
	DeviceId       string `json:"deviceId" v:"required" dc:"设备ID"`
	SourceDeviceId string `json:"sourceDeviceId" dc:"源设备ID"`
	Point          string `json:"point" v:"required" dc:"告警点位名称"`
}

type CreateAlarmIgnoreRes struct{}

// ==================== 删除忽略告警 ====================

type DeleteAlarmIgnoreReq struct {
	g.Meta   `path:"/alarm/ignore" method:"delete" tags:"告警" summary:"删除忽略告警"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
	Point    string `json:"point" v:"required" dc:"告警点位名称"`
}

type DeleteAlarmIgnoreRes struct{}

// ==================== 忽略告警分页查询 ====================

type GetAlarmIgnoreReq struct {
	g.Meta   `path:"/alarm/ignore" method:"get" tags:"告警" summary:"获取忽略告警分页列表"`
	DeviceId string `json:"deviceId" dc:"可选，指定设备ID；为空则查询全部"`
	Point    string `json:"point" dc:"可选，告警点位名称过滤"`
	Date     string `json:"date" dc:"可选，日期过滤，格式：2006-01-02"`
	Page     int    `json:"page" d:"1" dc:"页码，从1开始"`
	PageSize int    `json:"pageSize" d:"20" dc:"每页条数(最大100)"`
}

type AlarmIgnoreItem struct {
	Id               int        `json:"id" dc:"忽略记录ID"`
	DeviceId         string     `json:"deviceId" dc:"设备ID"`
	DeviceName       string     `json:"deviceName" dc:"设备名称"`
	SourceDeviceId   string     `json:"sourceDeviceId" dc:"源告警设备ID"`
	SourceDeviceName string     `json:"sourceDeviceName" dc:"源设备名称"`
	Point            string     `json:"point" dc:"告警点位名称"`
	CreatedAt        *time.Time `json:"createdAt" dc:"创建时间"`
}

type GetAlarmIgnoreRes struct {
	Total int `json:"total" dc:"总记录数"`

	Items []AlarmIgnoreItem `json:"items" dc:"忽略告警列表"`
}
