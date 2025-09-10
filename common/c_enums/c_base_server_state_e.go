//go:generate stringer -type=EServerState -trimprefix=EState -output=c_base_server_state_e_string.go
package c_base

// EServerState 定义所有 Manager 通用状态
type EServerState int

const (
	EStateInit    EServerState = iota // 初始化
	EStateRunning                     // 正在运行
	EStateStopped                     // 已停止
	EStateError                       // 异常
)
