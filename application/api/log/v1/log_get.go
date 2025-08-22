package v1

import "github.com/gogf/gf/v2/frame/g"

type GetBizLogReq struct {
	g.Meta   `path:"/log/biz" method:"get" tags:"日志" summary:"读取业务日志"`
	Type     string `json:"type"   v:"in:ems,device,protocol,policy,all" dc:"业务类型(可空/为all返回全部)"`
	Id       string `json:"id"     dc:"相关ID(ems可空)"`
	Date     string `json:"date"   dc:"日期，格式：20060102，默认为今天"`
	Page     int    `json:"page"   d:"1" dc:"页码，从1开始"`
	PageSize int    `json:"pageSize" d:"100" dc:"每页条数(最大1000)"`
}

// LogLine 结构化日志行
type LogLine struct {
	Timestamp string `json:"timestamp" dc:"时间戳"`
	Level     string `json:"level"     dc:"日志等级"`
	Content   string `json:"content"   dc:"日志内容"`
}

type GetBizLogRes struct {
	Total int       `json:"total" dc:"总行数"`
	Lines []LogLine `json:"lines" dc:"日志行列表(倒序)"`
}
