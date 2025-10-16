package v1

import (
	"common/c_base"

	"github.com/gogf/gf/v2/frame/g"
)

type GetRealDeviceListReq struct {
	g.Meta   `path:"/device/real/list" method:"get" tags:"设备相关" summary:"获取加载的设备列表" role:"user"`
	ShowType uint8 `json:"showType,omitempty" v:"in:0,1,2#请选择显示类型" dc:"显示类型:0全部1物理设备2虚拟设备"`
}

type GetRealDeviceListRes struct {
	Devices []*c_base.SDeviceConfig `json:"devices" dc:"设备详情"`
}
