//go:generate stringer -type=EPriceType -trimprefix=EPriceType -output=c_enum_price_type_e_string.go
package c_enum

import "strings"

// EPriceType 电价类型枚举
type EPriceType int

const (
	EPriceTypeValley     EPriceType = iota + 1 // 谷电
	EPriceTypePeak                             // 峰电
	EPriceTypeFlat                             // 平电
	EPriceTypeSharp                            // 尖峰
	EPriceTypeDeepValley                       // 深谷
)

// ParsePriceType 解析电价类型字符串
func ParsePriceType(value string) EPriceType {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case "valley", "谷电", "1":
		return EPriceTypeValley
	case "peak", "峰电", "2":
		return EPriceTypePeak
	case "flat", "平电", "3":
		return EPriceTypeFlat
	case "sharp", "尖峰", "4":
		return EPriceTypeSharp
	case "deep_valley", "deepvalley", "深谷", "5":
		return EPriceTypeDeepValley
	default:
		return EPriceTypeFlat // 默认平电
	}
}

// MarshalJSON 自定义JSON序列化
func (e EPriceType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.String() + `"`), nil
}

// UnmarshalJSON 自定义JSON反序列化
func (e *EPriceType) UnmarshalJSON(data []byte) error {
	value := string(data)
	// 去除引号
	if len(value) > 2 && value[0] == '"' && value[len(value)-1] == '"' {
		value = value[1 : len(value)-1]
	}
	*e = ParsePriceType(value)
	return nil
}
