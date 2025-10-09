package c_base

import (
	"strconv"

	"github.com/shockerli/cvt"
)

// SConfigPoint 配置点位，用于配置结构体字段描述
type SConfigPoint struct {
	*SPoint // 嵌套基础点位信息

	// 配置特定的字段（不重复SPoint中已有的字段）
	Required           bool     `json:"required" dc:"是否必填"`                         // 是否必填
	Default            *string  `json:"default,omitempty" dc:"默认值"`                 // 默认值
	Regex              *string  `json:"regex,omitempty" dc:"正则表达式验证"`               // 正则表达式验证
	RegexFailedMessage *string  `json:"regexFailedMessage,omitempty" dc:"正则验证失败提示"` // 正则验证失败提示
	Step               *float32 `json:"step,omitempty" dc:"步长（用于数字输入）"`             // 步长（用于数字输入）
}

// 注意：不需要重复实现IPoint接口方法
// 通过结构体嵌套自动继承SPoint的方法实现
// SPoint字段将在启动时验证是否设置

// 配置相关方法
func (s *SConfigPoint) GetRequired() bool {
	return s.Required
}

func (s *SConfigPoint) GetDefault() *string {
	return s.Default
}

func (s *SConfigPoint) GetRegex() *string {
	return s.Regex
}

func (s *SConfigPoint) GetRegexFailedMessage() *string {
	return s.RegexFailedMessage
}

func (s *SConfigPoint) GetStep() *float32 {
	return s.Step
}

// GetValueExplainWithParams 获取Value解释，支持动态参数
func (s *SConfigPoint) GetValueExplainWithParams(value any, params map[string]any) (string, string, error) {
	if s.SPoint == nil {
		return "", "", nil
	}

	// 1. 将value转换为字符串
	var valueStr string
	var err error

	// 检查值是否为数值类型（整数或浮点数）
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, bool, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *bool:
		// 数值类型直接转换为字符串
		valueStr, err = cvt.StringE(value)
		if err != nil {
			return "", "", err
		}
	default:
		// 非数值类型（如枚举）先转为int再转为字符串
		intVal, err := cvt.IntE(value)
		if err != nil {
			return "", "", err
		}
		valueStr, err = cvt.StringE(intVal)
		if err != nil {
			return "", "", err
		}
	}

	// 2. 从ValueExplain中查找匹配的解释
	if len(s.SPoint.ValueExplain) > 0 {
		for _, explain := range s.SPoint.ValueExplain {
			if explain.Key == valueStr {
				// 如果FromParam为true，从参数中获取值
				if explain.FromParam && params != nil {
					if paramValue, ok := params[explain.Value]; ok && paramValue != nil {
						return cvt.String(paramValue), explain.Color, nil
					}
				}
				// 否则直接返回Value
				return explain.Value, explain.Color, nil
			}
		}
	}

	// 3. 浮点数据进行格式化输出
	if floatVal, err := cvt.Float64E(value); err == nil {
		formatted := strconv.FormatFloat(floatVal, 'f', int(s.SPoint.Precise), 64)
		return formatted, "", nil
	}

	// 如果无法转换为浮点数，返回转换后的字符串
	return valueStr, "", nil
}
