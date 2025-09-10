package v1

import (
	"common/c_base"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceNameListReq struct {
	g.Meta      `path:"/device/name-list" method:"get" tags:"设备相关" summary:"获取所有设备的名称列表"`
	DeviceTypes []c_enum.EDeviceType `json:"deviceTypes,omitempty" dc:"设备类型过滤，可设置多个类型。可选值：ammeter(电表),ca(制冷空调),cl(液冷机组),bms(电池管理系统),fire(消防),hum(温湿度),pcs(电池逆变器),load(负载),pv(光伏),ess(储能柜),cp(充电桩),gen(发电机),gpio(DIY),sess(总站储能柜)"`
}

type GetDeviceNameListRes struct {
	DeviceNames map[string]string `json:"deviceNames" dc:"设备名称列表，格式为{设备ID:设备名称}"`
}
