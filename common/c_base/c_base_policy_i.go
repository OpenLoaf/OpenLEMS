package c_base

import (
	"common/c_enum"
	"time"
)

type IPolicy[T IDriver] interface {
	SetPolicyMode(c_enum.EPolicyMode)  // 设置策略模式
	GetPolicyMode() c_enum.EPolicyMode // 获取策略模式
	RegisterMonitor(*SPolicyMonitor)   // 注册监听器，定时触发策略
	RegisterActiveManualAction(func())
	ExecuteDriverFunc(func(driver T)) // 注册重置指令，当切换到手动模式时，会先触发重置
}

type SPolicyMonitor struct {
	Name        string
	Duration    *time.Duration       // 时间周期
	Modes       []c_enum.EPolicyMode // 哪些模式下触发
	TriggerFunc func() bool          // 触发条件
	HandleFunc  func()               // 执行方法
}
