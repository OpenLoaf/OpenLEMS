package v1

import (
	"application/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetRealDeviceListReq struct {
	g.Meta `path:"/device/real/list" method:"get" tags:"设备相关" summary:"获取物理设备列表"`
}

type GetRealDeviceListRes struct {
	Devices []*entity.SDevice `json:"devices" dc:"设备列表"`
}
