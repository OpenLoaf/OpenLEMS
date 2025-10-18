//go:generate stringer -type=EStatus -trimprefix=EStatus -output=c_enum_status_e_string.go
package c_enum

import "strings"

// EStatus 状态枚举
type EStatus int

const (
	EStatusDisable EStatus = iota // 禁用
	EStatusEnable                 // 启用
	EStatusDeleted                // 已删除
	EStatusExpire                 // 已过期
)

// ParseStatus 解析状态字符串
func ParseStatus(status string) EStatus {
	status = strings.ToLower(strings.TrimSpace(status))
	switch status {
	case "enable", "enabled", "1":
		return EStatusEnable
	case "disable", "disabled", "2":
		return EStatusDisable
	case "deleted", "delete", "3":
		return EStatusDeleted
	default:
		return EStatusDisable // 默认禁用
	}
}

// MarshalJSON 自定义JSON序列化
func (s EStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON 自定义JSON反序列化
func (s *EStatus) UnmarshalJSON(data []byte) error {
	status := string(data)
	// 去除引号
	if len(status) > 2 && status[0] == '"' && status[len(status)-1] == '"' {
		status = status[1 : len(status)-1]
	}
	*s = ParseStatus(status)
	return nil
}
