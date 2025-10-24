//go:generate stringer -type=EDeviceRunningStatus -trimprefix=EDeviceRunningStatus -output=c_enum_device_running_status_e_string.go
package c_enum

import "strings"

// EDeviceRunningStatus 设备运行状态枚举
type EDeviceRunningStatus int

const (
	EDeviceRunningStatusRunning EDeviceRunningStatus = iota + 1 // 运行中
	EDeviceRunningStatusStopped                                 // 已停止
)

// ParseDeviceRunningStatus 解析设备运行状态字符串
func ParseDeviceRunningStatus(status string) EDeviceRunningStatus {
	status = strings.ToLower(strings.TrimSpace(status))
	switch status {
	case "running", "运行中", "1":
		return EDeviceRunningStatusRunning
	case "stopped", "已停止", "停止", "2":
		return EDeviceRunningStatusStopped
	default:
		return EDeviceRunningStatusRunning // 默认运行中
	}
}

// MarshalJSON 自定义JSON序列化
func (s EDeviceRunningStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON 自定义JSON反序列化
func (s *EDeviceRunningStatus) UnmarshalJSON(data []byte) error {
	status := string(data)
	// 去除引号
	if len(status) > 2 && status[0] == '"' && status[len(status)-1] == '"' {
		status = status[1 : len(status)-1]
	}
	*s = ParseDeviceRunningStatus(status)
	return nil
}
