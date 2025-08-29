package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 当前告警分页查询 ====================

type GetCurrentAlarmsReq struct {
	g.Meta   `path:"/alarm/current" method:"get" tags:"告警" summary:"获取当前告警分页列表"`
	DeviceId string `json:"deviceId" dc:"可选，指定设备ID；为空则查询全部"`
	Level    string `json:"level" v:"in:NONE,WARN,ALARM,ERROR,ALL" dc:"告警级别过滤(可空/为ALL返回全部)"`
	Point    string `json:"point" dc:"可选，告警点位过滤"`
	Page     int    `json:"page" d:"1" dc:"页码，从1开始"`
	PageSize int    `json:"pageSize" d:"20" dc:"每页条数(最大100)"`
}

type CurrentAlarmItem struct {
	SourceDeviceId   string     `json:"sourceDeviceId" dc:"源告警设备ID"`
	SourceDeviceName string     `json:"sourceDeviceName" dc:"源设备名称"`
	DeviceId         string     `json:"deviceId" dc:"设备ID"`
	DeviceName       string     `json:"deviceName" dc:"设备名称"`
	Point            string     `json:"point" dc:"告警点位"`
	PointName        string     `json:"pointName" dc:"告警点位名称"`
	Level            string     `json:"level" dc:"告警级别"`
	Detail           string     `json:"detail" dc:"告警详情"`
	CreatedAt        *time.Time `json:"createdAt" dc:"创建时间"`
}

type GetCurrentAlarmsRes struct {
	Total        int                 `json:"total" dc:"总记录数"`
	IgnoreTotal  int                 `json:"ignoreTotal" dc:"过滤总数"`
	HistoryTotal int                 `json:"historyTotal" dc:"历史总数"`
	Items        []*CurrentAlarmItem `json:"items" dc:"当前告警列表"`
}

// ==================== 历史告警分页查询 ====================

type GetHistoryAlarmsReq struct {
	g.Meta    `path:"/alarm/history" method:"get" tags:"告警" summary:"获取历史告警分页列表"`
	DeviceId  string `json:"deviceId" dc:"可选，指定设备ID；为空则查询全部"`
	Level     string `json:"level" v:"in:NONE,WARN,ALARM,ERROR,ALL" dc:"告警级别过滤(可空/为ALL返回全部)"`
	Point     string `json:"point" dc:"可选，告警点位名称过滤"`
	PointName string `json:"pointName" dc:"可选，告警标题模糊搜索"`
	Date      string `json:"date" dc:"可选，日期过滤，格式：2006-01-02"`
	Page      int    `json:"page" d:"1" dc:"页码，从1开始"`
	PageSize  int    `json:"pageSize" d:"20" dc:"每页条数(最大100)"`
}

type HistoryAlarmItem struct {
	Id               int        `json:"id" dc:"告警ID"`
	DeviceId         string     `json:"deviceId" dc:"设备ID"`
	DeviceName       string     `json:"deviceName" dc:"设备名称"`
	SourceDeviceId   string     `json:"sourceDeviceId" dc:"源告警设备ID"`
	SourceDeviceName string     `json:"sourceDeviceName" dc:"源设备名称"`
	Point            string     `json:"point" dc:"告警点位名称"`
	PointName        string     `json:"pointName" dc:"告警标题"`
	Level            string     `json:"level" dc:"告警级别"`
	Detail           string     `json:"detail" dc:"告警详情"`
	TriggerAt        *time.Time `json:"triggerAt" dc:"触发时间"`
	ClearAt          *time.Time `json:"clearAt" dc:"清除时间"`
}

type GetHistoryAlarmsRes struct {
	Total int                `json:"total" dc:"总记录数"`
	Items []HistoryAlarmItem `json:"items" dc:"历史告警列表"`
}
