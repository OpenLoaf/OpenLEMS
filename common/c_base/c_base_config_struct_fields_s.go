package c_base

import (
	"common/c_enum"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
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

// UnmarshalYAML 自定义 YAML 反序列化方法，支持向后兼容
func (s *SFieldDefinition) UnmarshalYAML(value *yaml.Node) error {
	// 创建一个临时结构体来处理 YAML 解析
	type TempFieldDefinition struct {
		Key                string                            `yaml:"key"`
		Name               string                            `yaml:"name"`
		Group              string                            `yaml:"group"`
		ValueType          c_enum.EConfigFieldsValueType     `yaml:"valueType"`
		ComponentType      c_enum.EConfigFieldsComponentType `yaml:"componentType"`
		Step               *float32                          `yaml:"step"`
		Required           bool                              `yaml:"required"`
		Unit               *string                           `yaml:"unit"`
		Min                *int64                            `yaml:"min"`
		Max                *int64                            `yaml:"max"`
		Default            *string                           `yaml:"default"`
		ValueExplain       interface{}                       `yaml:"valueExplain"` // 使用 interface{} 来处理不同的格式
		ParamExplain       interface{}                       `yaml:"paramExplain"` // 使用 interface{} 来处理不同的格式
		Regex              *string                           `yaml:"regex"`
		RegexFailedMessage *string                           `yaml:"regexFailedMessage"`
		Description        string                            `yaml:"description"`
	}

	var temp TempFieldDefinition
	if err := value.Decode(&temp); err != nil {
		return err
	}

	// 复制基本字段
	s.Key = temp.Key
	s.Name = temp.Name
	s.Group = temp.Group
	s.ValueType = temp.ValueType
	s.ComponentType = temp.ComponentType
	s.Step = temp.Step
	s.Required = temp.Required
	s.Unit = temp.Unit
	s.Min = temp.Min
	s.Max = temp.Max
	s.Default = temp.Default
	s.Regex = temp.Regex
	s.RegexFailedMessage = temp.RegexFailedMessage
	s.Description = temp.Description

	// 处理 ValueExplain 字段
	if temp.ValueExplain != nil {
		s.ValueExplain = convertToFieldExplainArray(temp.ValueExplain)
	}

	// 处理 ParamExplain 字段
	if temp.ParamExplain != nil {
		s.ParamExplain = convertToFieldExplainArray(temp.ParamExplain)
	}

	return nil
}

// convertToFieldExplainArray 将不同格式的 explain 数据转换为 []*SFieldExplain
func convertToFieldExplainArray(data interface{}) []*SFieldExplain {
	if data == nil {
		return nil
	}

	// 如果是字符串，使用 ParseExplainString 解析
	if str, ok := data.(string); ok {
		return ParseExplainString(str)
	}

	// 如果是 map，转换为新的格式
	if mapData, ok := data.(map[string]interface{}); ok {
		var explains []*SFieldExplain
		for key, value := range mapData {
			if valueStr, ok := value.(string); ok {
				// 检查是否包含颜色信息
				var explainValue, color string
				if strings.Contains(valueStr, "|") {
					parts := strings.SplitN(valueStr, "|", 2)
					if len(parts) == 2 {
						explainValue = strings.TrimSpace(parts[0])
						color = strings.TrimSpace(parts[1])
					}
				} else {
					explainValue = valueStr
				}

				explains = append(explains, &SFieldExplain{
					Key:   key,
					Value: explainValue,
					Color: color,
				})
			}
		}
		return explains
	}

	// 如果已经是 []*SFieldExplain 格式，直接返回
	if explains, ok := data.([]*SFieldExplain); ok {
		return explains
	}

	// 如果是 []interface{}，尝试转换
	if slice, ok := data.([]interface{}); ok {
		var explains []*SFieldExplain
		for _, item := range slice {
			if explain, ok := item.(*SFieldExplain); ok {
				explains = append(explains, explain)
			} else if mapItem, ok := item.(map[string]interface{}); ok {
				explain := &SFieldExplain{}
				if key, ok := mapItem["key"].(string); ok {
					explain.Key = key
				}
				if value, ok := mapItem["value"].(string); ok {
					explain.Value = value
				}
				if color, ok := mapItem["color"].(string); ok {
					explain.Color = color
				}
				explains = append(explains, explain)
			}
		}
		return explains
	}

	return nil
}
