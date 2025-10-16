package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type QueryBizLogInfoReq struct {
	g.Meta `path:"/log/biz/info" method:"get" tags:"日志" summary:"查询日志信息" role:"user"`
}

type QueryBizLogInfoRes struct {
	Total         int    `json:"total" dc:"日志总数"`
	FirstLogTime  string `json:"firstLogTime" dc:"第一条日志时间"`
	LatestLogTime string `json:"latestLogTime" dc:"最新日志时间"`
}
