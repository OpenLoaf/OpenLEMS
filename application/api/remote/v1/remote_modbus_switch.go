package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ModbusSwitchReq struct {
	g.Meta  `path:"/remote/modbus/enabled/{enabled}" method:"put" tags:"远程管理" summary:"Modbus服务开关" role:"admin"`
	Enabled bool `json:"enabled" dc:"是否启用Modbus服务" v:"required#启用状态不能为空"`
}

type ModbusSwitchRes struct {
}
