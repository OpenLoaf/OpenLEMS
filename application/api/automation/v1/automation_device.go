package v1

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// GetAutomationsByDeviceReq 获取设备自动化列表请求
type GetAutomationsByDeviceReq struct {
	g.Meta   `path:"/automation/device/{deviceId}" method:"get" tags:"自动化相关" summary:"获取设备自动化列表"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

// GetAutomationsByDeviceRes 获取设备自动化列表响应
type GetAutomationsByDeviceRes struct {
	Automations []*entity.SAutomation `json:"automations" dc:"自动化任务列表"`
	Count       int                   `json:"count" dc:"任务数量"`
}
