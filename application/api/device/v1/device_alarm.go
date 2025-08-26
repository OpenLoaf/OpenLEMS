package v1

import "github.com/gogf/gf/v2/frame/g"

type GetCurrentAlarmsReq struct {
	g.Meta   `path:"/device/alarms" method:"get" tags:"设备" summary:"获取当前告警列表"`
	DeviceId string `json:"deviceId" dc:"可选，指定设备ID；为空则查询全部"`
}

type AlarmItem struct {
	DeviceId string `json:"deviceId" dc:"设备ID"`
	Level    string `json:"level" dc:"告警级别"`
	Message  string `json:"message" dc:"告警信息"`
	Time     string `json:"time" dc:"发送时间(YYYY-MM-DD HH:mm:ss)"`
}

type GetCurrentAlarmsRes struct {
	CurrentTotal int         `json:"currentTotal" dc:"当前告警总数"`
	ShieldTotal  int         `json:"shieldTotal" dc:"屏蔽告警总数"`
	HistoryTotal int         `json:"historyTotal" dc:"历史告警"`
	Items        []AlarmItem `json:"items" dc:"告警列表"`
}
