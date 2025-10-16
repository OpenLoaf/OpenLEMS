package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetBizLogReq struct {
	g.Meta   `path:"/log/biz" method:"get" tags:"日志" summary:"读取业务日志" role:"user"`
	Type     string `json:"type"   v:"in:ems,device,protocol,policy,remote,all" dc:"业务类型(可空/为all返回全部)"`
	Id       string `json:"id"     dc:"相关ID(ems可空)"`
	Date     string `json:"date"   dc:"日期，格式：2006-01-02，默认为今天"`
	Page     int    `json:"page"   d:"1" dc:"页码，从1开始"`
	PageSize int    `json:"pageSize" d:"100" dc:"每页条数(最大1000)"`
	Level    string `json:"level"  v:"in:DEBUG,INFO,WARN,ERROR,ALL" dc:"日志等级(可空/为ALL返回全部；仅支持: DEBUG/INFO/WARN/ERROR)"`
}

// LogLine 结构化日志行
type LogLine struct {
	CreatedAt *time.Time `json:"createdAt" dc:"时间"`
	Id        string     `json:"id"`
	Type      string     `json:"type" dc:"日志类型：ems、device、protocol、policy、remote"`
	Level     string     `json:"level"     dc:"日志等级"`
	Content   string     `json:"content"   dc:"日志内容"`
}

type GetBizLogRes struct {
	Total int       `json:"total" dc:"总行数"`
	Lines []LogLine `json:"lines" dc:"日志行列表(倒序)"`
}
