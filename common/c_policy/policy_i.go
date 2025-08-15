package c_policy

import "common/c_base"

type IPolicy interface {
	GetDevices(showAll bool) []c_base.IDriver // 获取策略的控制设备
	Check() error                             // 验证策略
	Active() error                            // 激活策略
	Destroy() error                           // 销毁策略
}
