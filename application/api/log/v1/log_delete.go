package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteBizLogReq struct {
	g.Meta `path:"/log/biz" method:"delete" tags:"日志" summary:"删除业务日志"`
	Type   string `json:"type"   v:"in:ems,device,protocol,policy,all" dc:"业务类型(全部) 可空"`
	Level  string `json:"level"  v:"in:DEBUG,INFO,WARN,ERROR,ALL" dc:"日志等级(全部) 可空；仅支持: DEBUG/INFO/WARN/ERROR/ALL"`
}

type DeleteBizLogRes struct {
	Total int `json:"total" dc:"删除数量"`
}
