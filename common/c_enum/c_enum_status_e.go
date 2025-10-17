//go:generate stringer -type=EStrategyStatus -trimprefix=EStrategyStatus -output=c_enum_status_e_string.go
package c_enum

import "strings"

// EStrategyStatus 状态枚举
type EStrategyStatus int

const (
	EStatusEnable  EStrategyStatus = iota // 启用
	EStatusDisable                        // 禁用
	EStatusDeleted                        // 已删除
)

// ParseEnergyStorageStrategyStatus 解析状态字符串
func ParseEnergyStorageStrategyStatus(status string) EStrategyStatus {
	status = strings.ToLower(strings.TrimSpace(status))
	switch status {
	case "enable", "active", "enabled":
		return EStatusEnable
	case "disable", "inactive", "disabled":
		return EStatusDisable
	case "deleted", "delete":
		return EStatusDeleted
	default:
		return EStatusDisable // 默认禁用
	}
}

// MarshalJSON 自定义JSON序列化
func (s EStrategyStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON 自定义JSON反序列化
func (s *EStrategyStatus) UnmarshalJSON(data []byte) error {
	status := string(data)
	// 去除引号
	if len(status) > 2 && status[0] == '"' && status[len(status)-1] == '"' {
		status = status[1 : len(status)-1]
	}
	*s = ParseEnergyStorageStrategyStatus(status)
	return nil
}
