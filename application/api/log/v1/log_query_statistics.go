package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type QueryBizLogStatisticsReq struct {
	g.Meta `path:"/log/biz/statistics" method:"get" tags:"日志" summary:"查询日志统计信息"`
}

type QueryBizLogStatisticsRes struct {
	Type struct {
		Ems      int `json:"ems" dc:"ems类型日志数量"`
		Device   int `json:"device" dc:"device类型日志数量"`
		Protocol int `json:"protocol" dc:"protocol类型日志数量"`
		Policy   int `json:"policy" dc:"policy类型日志数量"`
	} `json:"type" dc:"按类型统计"`
	Level struct {
		DEBUG int `json:"DEBUG" dc:"DEBUG等级日志数量"`
		INFO  int `json:"INFO" dc:"INFO等级日志数量"`
		WARN  int `json:"WARN" dc:"WARN等级日志数量"`
		ERROR int `json:"ERROR" dc:"ERROR等级日志数量"`
	} `json:"level" dc:"按等级统计"`
}
