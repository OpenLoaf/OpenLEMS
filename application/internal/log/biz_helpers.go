package log

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"sync"

	"common/c_base"

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

var (
	deviceLoggerCache   sync.Map // key: deviceId -> *glog.Logger
	protocolLoggerCache sync.Map // key: protocolId -> *glog.Logger
	policyLoggerCache   sync.Map // key: policyId -> *glog.Logger
)

func mergeConfig(baseCfg map[string]interface{}, kv map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(baseCfg)+len(kv))
	for k, v := range baseCfg {
		out[k] = v
	}
	for k, v := range kv {
		out[k] = v
	}
	return out
}

func getBaseLoggerConfig(ctx context.Context, key string) map[string]interface{} {
	// 合并全局 logger 与命名实例 logger.<key>，命名实例覆盖全局
	globalVar := g.Cfg().MustGet(ctx, "logger")
	globalCfg := map[string]interface{}{}
	if !globalVar.IsNil() {
		if m := globalVar.Map(); m != nil {
			globalCfg = m
		}
	}
	namedVar := g.Cfg().MustGet(ctx, fmt.Sprintf("logger.%s", key))
	namedCfg := map[string]interface{}{}
	if !namedVar.IsNil() {
		if m := namedVar.Map(); m != nil {
			namedCfg = m
		}
	}
	return mergeConfig(globalCfg, namedCfg)
}

func buildSubLogger(ctx context.Context, baseKey, id string) *glog.Logger {
	baseCfg := getBaseLoggerConfig(ctx, baseKey)
	basePath := ""
	if p, ok := baseCfg["path"].(string); ok && p != "" {
		basePath = p
	} else {
		basePath = filepath.Join("logs", "biz", baseKey)
	}
	subCfg := mergeConfig(baseCfg, map[string]interface{}{
		"path":  filepath.Join(basePath, id),
		"async": true,
	})
	// 删除CtxKeys，不需要保存设备的这些信息了
	delete(subCfg, "CtxKeys")
	l := glog.New()
	_ = l.SetConfigWithMap(subCfg)
	// 业务日志统一为JSON格式（同样写入分类字段）
	l.SetHandlers(BizJsonHandler)
	return l
}

// BizDevice 获取设备业务日志（同类型不同ID分目录/文件）
func BizDevice(ctx context.Context, deviceId string) *glog.Logger {
	if v, ok := deviceLoggerCache.Load(deviceId); ok {
		return v.(*glog.Logger)
	}
	l := buildSubLogger(ctx, "biz_device", deviceId)
	deviceLoggerCache.Store(deviceId, l)
	return l
}

// BizProtocol 获取协议业务日志（同类型不同ID分目录/文件）
func BizProtocol(ctx context.Context, protocolId string) *glog.Logger {
	if v, ok := protocolLoggerCache.Load(protocolId); ok {
		return v.(*glog.Logger)
	}
	l := buildSubLogger(ctx, "biz_protocol", protocolId)
	protocolLoggerCache.Store(protocolId, l)
	return l
}

// BizPolicy 获取策略业务日志（同类型不同ID分目录/文件）
func BizPolicy(ctx context.Context, policyId string) *glog.Logger {
	if v, ok := policyLoggerCache.Load(policyId); ok {
		return v.(*glog.Logger)
	}
	l := buildSubLogger(ctx, "biz_policy", policyId)
	policyLoggerCache.Store(policyId, l)
	return l
}
