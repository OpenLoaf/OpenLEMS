package c_base

import (
	"common/c_enum"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type SFieldDefinition struct {
	Key                string                            `json:"key" yaml:"key" short:"key"`
	Name               string                            `json:"name" yaml:"name" short:"name" required:"true"`
	Group              string                            `json:"group" yaml:"group" short:"group"`
	ValueType          c_enum.EConfigFieldsValueType     `json:"valueType" yaml:"value_type" short:"vt" dc:"string字符串、int整数、float浮点数、bool布尔值" required:"true"`
	ComponentType      c_enum.EConfigFieldsComponentType `json:"componentType" yaml:"component_type" short:"ct" dc:"组件类型" required:"true"`
	Step               *float32                          `json:"step" yaml:"step" short:"step" default:"1" dc:"步长（小步长）"`
	Required           bool                              `json:"required" yaml:"required" short:"req" required:"true" dc:"是否必填"`
	Unit               *string                           `json:"unit" yaml:"unit" short:"unit" dc:"单位"`
	Min                *int64                            `json:"min" yaml:"min" short:"min"`
	Max                *int64                            `json:"max" yaml:"max" short:"max"`
	Default            *string                           `json:"default" yaml:"default" short:"def"`
	ValueExplain       []*SFieldExplain                  `json:"valueExplain,omitempty" yaml:"valueExplain"` // 值解释
	ParamExplain       []*SFieldExplain                  `json:"paramExplain,omitempty" yaml:"paramExplain"` // 从参数值中读取解释
	Regex              *string                           `json:"regex" yaml:"regex" short:"regex" dc:"正则表达式"`
	RegexFailedMessage *string                           `json:"regexFailedMessage" yaml:"regex_failed_message" short:"rfm" dc:"正则表达式失败提醒"`
	Description        string                            `json:"description" yaml:"description" short:"desc" required:"true"`
}

type SFieldExplain struct {
	Key   string `json:"key" yaml:"key" short:"key" required:"true"`
	Value string `json:"value" yaml:"value" short:"value" required:"true"`
	Color string `json:"color" yaml:"color" short:"color" required:"true"`
}

func (s *SFieldDefinition) String() string {
	if s == nil {
		return "SFieldDefinition(nil)"
	}

	// 处理指针字段
	var minVal, maxVal int64
	if s.Min != nil {
		minVal = *s.Min
	}
	if s.Max != nil {
		maxVal = *s.Max
	}

	var defaultVal, regexVal string
	if s.Default != nil {
		defaultVal = *s.Default
	}
	if s.Regex != nil {
		regexVal = *s.Regex
	}

	return fmt.Sprintf("SFieldDefinition{Key:%s, Name:%s, ValueType:%s, ComponentType:%s, Min:%d, Max:%d, Default:%s, Regex:%s}",
		s.Key, s.Name, s.ValueType, s.ComponentType, minVal, maxVal, defaultVal, regexVal)
}

func (s *SFieldDefinition) Check() error {
	if s == nil {
		return errors.New("SFieldDefinition is nil")
	}

	// 检查必填字段
	if s.Name == "" {
		return errors.New("Name is required")
	}
	if s.Description == "" {
		return errors.New("Description is required")
	}
	if s.Key == "" {
		return errors.New("Key is required")
	}
	if s.ValueType == "" {
		return errors.New("ValueType is required")
	}

	// 检查数值范围
	if s.Min != nil && s.Max != nil && *s.Min > *s.Max {
		return errors.New("Min value cannot be greater than Max value")
	}

	// 检查正则表达式
	if s.Regex != nil && *s.Regex != "" {
		if _, err := regexp.Compile(*s.Regex); err != nil {
			return errors.Errorf("Invalid regex pattern: %v", err)
		}
	}

	return nil
}

// ToPoint 将配置字段转换为点位信息
func (s *SFieldDefinition) ToPoint(valueType c_enum.EValueType, params map[string]any) IPoint {
	valueExplain := make(map[string]string)
	if len(s.ValueExplain) > 0 {
		valueType = c_enum.EString
		for _, explain := range s.ValueExplain {
			valueExplain[explain.Key] = explain.Value
		}
	}
	if len(s.ParamExplain) > 0 {
		valueType = c_enum.EString
		for _, explain := range s.ParamExplain {
			if pv, ok := params[explain.Value]; ok && pv != nil {
				valueExplain[explain.Key] = pv.(string)
			}
		}
	}

	// 处理单位信息
	var unit string
	if s.Unit != nil {
		unit = *s.Unit
	}

	return &SPoint{
		Key:          s.Key,
		Name:         s.Name,
		Group:        &SPointGroup{GroupName: s.Group}, // 将字符串转换为 SPointGroup 指针
		Precise:      0,                                // SFieldDefinition 没有精度字段，使用默认值
		Desc:         s.Description,
		Unit:         unit,
		ValueType:    valueType,
		ValueExplain: valueExplain,
	}
}

// ParseExplainString 解析 Explain 字符串格式为 []*SFieldExplain 对象数组
// 支持两种格式：
// 1. 简单格式：N:无校验,E:偶校验,O:奇校验
// 2. 带颜色格式：N:无校验|#52c41a,E:偶校验|#1890ff,O:奇校验|#f5222d
func ParseExplainString(explainStr string) []*SFieldExplain {
	if explainStr == "" {
		return nil
	}

	var explains []*SFieldExplain
	pairs := strings.Split(explainStr, ",")

	for _, pair := range pairs {
		trimmedPair := strings.TrimSpace(pair)
		if trimmedPair == "" {
			continue
		}

		// 检查是否包含颜色信息（| 分隔符）
		var key, value, color string
		if strings.Contains(trimmedPair, "|") {
			// 带颜色格式：key:value|#color
			parts := strings.SplitN(trimmedPair, "|", 2)
			if len(parts) == 2 {
				color = strings.TrimSpace(parts[1])
				keyValuePart := strings.TrimSpace(parts[0])
				keyValue := strings.SplitN(keyValuePart, ":", 2)
				if len(keyValue) == 2 {
					key = strings.TrimSpace(keyValue[0])
					value = strings.TrimSpace(keyValue[1])
				}
			}
		} else {
			// 简单格式：key:value
			keyValue := strings.SplitN(trimmedPair, ":", 2)
			if len(keyValue) == 2 {
				key = strings.TrimSpace(keyValue[0])
				value = strings.TrimSpace(keyValue[1])
			}
		}

		if key != "" && value != "" {
			explains = append(explains, &SFieldExplain{
				Key:   key,
				Value: value,
				Color: color,
			})
		}
	}

	return explains
}
