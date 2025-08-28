package log

import (
	"common/c_base"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

// BizJsonHandler 输出统一的 JSON 格式，并在每条日志中附带分类字段(type/id)
var BizJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	typ := "ems"
	id := ""
	if ctx != nil {
		if v := ctx.Value(c_base.ConstCtxKeyDeviceId); v != nil {
			if s, ok := v.(string); ok && s != "" {
				typ = "device"
				id = s
			}
		}
		if typ == "ems" { // 仅当未匹配设备时再判断协议
			if v := ctx.Value(c_base.ConstCtxKeyProtocolId); v != nil {
				if s, ok := v.(string); ok && s != "" {
					typ = "protocol"
					id = s
				}
			}
		}
		if typ == "ems" { // 兜底策略ID（约定键名）
			if v := ctx.Value("PolicyId"); v != nil {
				if s, ok := v.(string); ok && s != "" {
					typ = "policy"
					id = s
				}
			}
		}
	}
	content := in.Content
	if gstr.Trim(content) == "" {
		// 回退使用 values 内容
		content = fmt.Sprint(in.Values...)
	}
	content = gstr.Trim(content)
	payload := struct {
		Time    string `json:"time"`
		Level   string `json:"level"`
		Content string `json:"content"`
		Type    string `json:"type"`
		Id      string `json:"id"`
	}{
		Time:    in.TimeFormat,
		Level:   gstr.Trim(in.LevelFormat, "[]"),
		Content: content,
		Type:    typ,
		Id:      id,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		in.Buffer.WriteString(err.Error())
		in.Next(ctx)
		return
	}
	in.Buffer.Write(b)
	in.Buffer.WriteString("\n")
	in.Next(ctx)
}

// BizEMS 获取 EMS 业务日志（固定单文件，按天切分）
func BizEMS() *glog.Logger {
	l := g.Log("biz_ems")
	// 统一业务日志：JSON格式 + 异步输出
	l.SetHandlers(BizJsonHandler)
	l.SetAsync(true)
	return l
}
