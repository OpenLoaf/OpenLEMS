package c_base

import "time"

type IPolicy interface {
	SetPolicyMode(EPolicyMode)         // 设置策略模式
	GetPolicyMode() EPolicyMode        // 获取策略模式
	RegisterMonitor(*SPolicyMonitor)   // 注册监听器，定时触发策略
	RegisterActiveManualAction(func()) // 注册重置指令，当切换到手动模式时，会先触发重置
}

type SPolicyMonitor struct {
	Name        string
	Duration    *time.Duration // 时间周期
	Modes       []EPolicyMode  // 哪些模式下触发
	TriggerFunc func() bool    // 触发条件
	HandleFunc  func()         // 执行方法
}

type EPolicyMode int

const (
	EPolicyModeAuto     = iota // // 全自动模式
	EPolicyModeSemiAuto        // 半自动模式
	EPolicyModeManual          // 纯手动模式
)
