//go:generate stringer -type=EAlarmLevel -trimprefix=EAlarmLevel -output=c_enum_alarm_level_e_string.go
package c_enum

import (
	"encoding/json"
	"fmt"
	"strings"
)

type EAlarmLevel int

const (
	EAlarmLevelNone  EAlarmLevel = iota // 默认非告警
	EAlarmLevelWarn                     // 警告，不影响系统
	EAlarmLevelAlert                    // 警报，系统会限制功能
	EAlarmLevelError                    // 故障，系统会使得设备停机
)

// UnmarshalJSON 实现自定义 JSON 反序列化
func (e *EAlarmLevel) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	switch v := value.(type) {
	case float64:
		// 处理数字类型
		*e = EAlarmLevel(int(v))
	case int:
		*e = EAlarmLevel(v)
	case string:
		// 处理字符串类型，移除前缀并忽略大小写
		normalized := strings.ToUpper(strings.TrimSpace(v))
		switch normalized {
		case "NONE", "0":
			*e = EAlarmLevelNone
		case "WARN", "1":
			*e = EAlarmLevelWarn
		case "ALERT", "2":
			*e = EAlarmLevelAlert
		case "ERROR", "3":
			*e = EAlarmLevelError
		default:
			return fmt.Errorf("invalid EAlarmLevel value: %s", v)
		}
	default:
		return fmt.Errorf("invalid type for EAlarmLevel: %T", v)
	}

	return nil
}

// MarshalJSON 实现自定义 JSON 序列化（可选）
func (e EAlarmLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String()) // 使用 stringer 生成的字符串方法
}
