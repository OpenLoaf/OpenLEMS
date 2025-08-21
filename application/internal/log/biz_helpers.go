package log

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// BizEMS 获取 EMS 业务日志（固定单文件，按天切分）
func BizEMS() *glog.Logger {
	l := g.Log("biz_ems")
	l.SetAsync(true)
	return l
}

// BizDevice 获取设备业务日志（同类型不同ID分目录/文件）
func BizDevice(deviceId string) *glog.Logger {
	l := glog.New()
	l.SetConfigWithMap(g.Map{
		"path":   fmt.Sprintf("out/logs/biz/device/%s", deviceId),
		"file":   "{Ymd}.log",
		"level":  "all",
		"stdout": false,
		"header": true,
		"async":  true,
	})
	return l
}

// BizProtocol 获取协议业务日志（同类型不同ID分目录/文件）
func BizProtocol(protocolId string) *glog.Logger {
	l := glog.New()
	l.SetConfigWithMap(g.Map{
		"path":   fmt.Sprintf("out/logs/biz/protocol/%s", protocolId),
		"file":   "{Ymd}.log",
		"level":  "all",
		"stdout": false,
		"header": true,
		"async":  true,
	})
	return l
}

// BizPolicy 获取策略业务日志（同类型不同ID分目录/文件）
func BizPolicy(policyId string) *glog.Logger {
	l := glog.New()
	l.SetConfigWithMap(g.Map{
		"path":   fmt.Sprintf("out/logs/biz/policy/%s", policyId),
		"file":   "{Ymd}.log",
		"level":  "all",
		"stdout": false,
		"header": true,
		"async":  true,
	})
	return l
}
