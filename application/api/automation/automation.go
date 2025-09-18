// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package automation

import (
	"context"

	v1 "application/api/automation/v1"
)

type IAutomationV1 interface {
	// 获取所有的自动化分页接口（包括enabled = false）
	GetAutomationPage(ctx context.Context, req *v1.GetAutomationPageReq) (res *v1.GetAutomationPageRes, err error)

	// 获取某个设备的所有自动化列表（包括enabled = false）
	GetAutomationsByDevice(ctx context.Context, req *v1.GetAutomationsByDeviceReq) (res *v1.GetAutomationsByDeviceRes, err error)

	// 创建自动化任务
	CreateAutomation(ctx context.Context, req *v1.CreateAutomationReq) (res *v1.CreateAutomationRes, err error)

	// 更新自动化任务
	UpdateAutomation(ctx context.Context, req *v1.UpdateAutomationReq) (res *v1.UpdateAutomationRes, err error)

	// 删除自动化任务
	DeleteAutomation(ctx context.Context, req *v1.DeleteAutomationReq) (res *v1.DeleteAutomationRes, err error)

	// 开启/停用自动化任务
	ToggleAutomation(ctx context.Context, req *v1.ToggleAutomationReq) (res *v1.ToggleAutomationRes, err error)
}
