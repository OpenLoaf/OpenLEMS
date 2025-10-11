package c_base

import (
	"common/c_enum"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

// SFieldDefinition 可以是配置的定义，也可以是点位telemetry的定义
type SFieldDefinition struct {
	Key                string                            `json:"key" yaml:"key" short:"code"`
	Name               string                            `json:"name,omitempty" yaml:"name" short:"name" required:"true"`
	Group              string                            `json:"group,omitempty" yaml:"group" short:"group"`
	ValueType          c_enum.EConfigFieldsValueType     `json:"valueType" yaml:"value_type" short:"vt" dc:"string字符串、int整数、float浮点数、bool布尔值" required:"true"`
	ComponentType      c_enum.EConfigFieldsComponentType `json:"componentType" yaml:"component_type" short:"ct" dc:"组件类型" required:"true"`
	Step               *float32                          `json:"step,omitempty" yaml:"step" short:"step" default:"1" dc:"步长（小步长）"`
	Required           bool                              `json:"required" yaml:"required" short:"req" required:"true" dc:"是否必填"`
	Unit               *string                           `json:"unit,omitempty" yaml:"unit" short:"unit" dc:"单位"`
	Min                *int64                            `json:"min,omitempty" yaml:"min" short:"min"`
	Max                *int64                            `json:"max,omitempty" yaml:"max" short:"max"`
	Default            *string                           `json:"default,omitempty" yaml:"default" short:"def"`
	ValueExplain       []*SFieldExplain                  `json:"valueExplain,omitempty" yaml:"valueExplain" short:"ve"` // 值解释
	ParamExplain       []*SFieldExplain                  `json:"paramExplain,omitempty" yaml:"paramExplain" short:"pe"` // 从参数值中读取解释
	Regex              *string                           `json:"regex,omitempty" yaml:"regex" short:"regex" dc:"正则表达式"`
	RegexFailedMessage *string                           `json:"regexFailedMessage,omitempty" yaml:"regex_failed_message" short:"rfm" dc:"正则表达式失败提醒"`
	Description        string                            `json:"description,omitempty" yaml:"description" short:"desc" required:"true"`
}

type SFieldExplain struct {
	Key       string `json:"key" yaml:"key" short:"key" required:"true"`
	Value     string `json:"value" yaml:"value" short:"value" required:"true"`
	FromParam bool   `json:"fromParam" yaml:"fromParam" short:"fromParam" required:"true"`
	Color     string `json:"color" yaml:"color" short:"color" required:"true"`
}

func (s *SFieldDefinition) GetValueExplain() []*SFieldExplain {
	return s.ValueExplain
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

// 实现 IPoint 接口的所有方法

// GetKey 获取点位Key
func (s *SFieldDefinition) GetKey() string {
	return s.Key
}

// GetName 获取名称
func (s *SFieldDefinition) GetName() string {
	return s.Name
}

// GetGroup 获取分组
func (s *SFieldDefinition) GetGroup() *SPointGroup {
	if s.Group == "" {
		return nil
	}
	return &SPointGroup{
		GroupKey:  s.Group,
		GroupName: s.Group,
		GroupSort: 0,
		Disable:   false,
	}
}

// GetUnit 获取单位
func (s *SFieldDefinition) GetUnit() string {
	if s.Unit != nil {
		return *s.Unit
	}
	return ""
}

// GetDesc 获取备注
func (s *SFieldDefinition) GetDesc() string {
	return s.Description
}

// GetSort 获取排序
func (s *SFieldDefinition) GetSort() int {
	// SFieldDefinition 没有排序字段，返回默认值
	return 0
}

// GetMin 获取点位理论最小值
func (s *SFieldDefinition) GetMin() int64 {
	if s.Min != nil {
		return *s.Min
	}
	return 0
}

// GetMax 获取点位理论最大值
func (s *SFieldDefinition) GetMax() int64 {
	if s.Max != nil {
		return *s.Max
	}
	return 0
}

// GetPrecise 获取小数点精度
func (s *SFieldDefinition) GetPrecise() uint8 {
	// SFieldDefinition 没有精度字段，返回默认值
	return 0
}

// GetValueType 获取值类型
func (s *SFieldDefinition) GetValueType() c_enum.EValueType {
	return s.convertValueType(s.ValueType)
}

// GetValueExplain 获取Value解释
func (s *SFieldDefinition) GetValueExplainByValue(value any) (string, error) {
	return s.GetValueExplainWithParams(value, nil)
}

// GetValueExplainWithParams 获取Value解释，支持动态参数
func (s *SFieldDefinition) GetValueExplainWithParams(value any, params map[string]any) (string, error) {
	// 1. 将value转换为字符串
	var valueStr string
	var err error

	// 检查值是否为数值类型（整数或浮点数）
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, bool, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *bool:
		// 数值类型直接转换为字符串
		valueStr, err = cvt.StringE(value)
		if err != nil {
			return "", err
		}
	default:
		// 非数值类型（如枚举）先转为int再转为字符串
		intVal, err := cvt.IntE(value)
		if err != nil {
			return "", err
		}
		valueStr, err = cvt.StringE(intVal)
		if err != nil {
			return "", err
		}
	}

	// 2. 从ValueExplain中查找匹配的解释
	if len(s.ValueExplain) > 0 {
		for _, explain := range s.ValueExplain {
			if explain.Key == valueStr {
				return explain.Value, nil
			}
		}
	}

	// 3. 从ParamExplain中查找匹配的解释（支持动态参数）
	if len(s.ParamExplain) > 0 && params != nil {
		for _, explain := range s.ParamExplain {
			if explain.Key == valueStr {
				if pv, ok := params[explain.Value]; ok && pv != nil {
					return cvt.String(pv), nil
				}
			}
		}
	}

	// 4. 浮点数据进行格式化输出
	if floatVal, err := cvt.Float64E(value); err == nil {
		// 使用 strconv.FormatFloat 进行精确格式化
		formatted := strconv.FormatFloat(floatVal, 'f', int(s.GetPrecise()), 64)
		return formatted, nil
	}

	// 如果无法转换为浮点数，返回转换后的字符串
	return valueStr, nil
}

// IsHidden 是否隐藏不显示
func (s *SFieldDefinition) IsHidden() bool {
	// SFieldDefinition 没有隐藏字段，返回默认值
	return false
}

// IsAlarmPoint 是否是告警点位
func (s *SFieldDefinition) IsAlarmPoint() bool {
	// SFieldDefinition 没有告警触发函数，返回默认值
	return false
}

// TriggerAlarm 判断触发或者消除告警
func (s *SFieldDefinition) TriggerAlarm(value any) (trigger bool, level c_enum.EAlarmLevel, err error) {
	// SFieldDefinition 没有告警触发函数，返回默认值
	return false, c_enum.EAlarmLevelNone, nil
}

// convertValueType 将 EConfigFieldsValueType 转换为 EValueType
func (s *SFieldDefinition) convertValueType(valueType c_enum.EConfigFieldsValueType) c_enum.EValueType {
	switch valueType {
	case c_enum.EConfigFieldsValueTypeString:
		return c_enum.EString
	case c_enum.EConfigFieldsValueTypeInt:
		return c_enum.EInt32
	case c_enum.EConfigFieldsValueTypeFloat:
		return c_enum.EFloat32
	case c_enum.EConfigFieldsValueTypeBoolean:
		return c_enum.EBool
	default:
		return c_enum.EAuto
	}
}

// ToPointWithValueType 将配置字段转换为点位信息，支持外部ValueType参数
// 这个方法提供了与原有ToPoint方法相同的功能，用于向后兼容
func (s *SFieldDefinition) ToPointWithValueType(valueType c_enum.EValueType, params map[string]any) IPoint {
	// 如果有ValueExplain或ParamExplain，强制设置为EString（与原有ToPoint方法逻辑一致）
	if len(s.ValueExplain) > 0 || len(s.ParamExplain) > 0 {
		valueType = c_enum.EString
	}

	// 创建一个临时的SFieldDefinition副本，设置正确的ValueType
	tempDef := *s
	tempDef.ValueType = s.convertValueTypeFromEValueType(valueType)

	return &tempDef
}

// convertValueTypeFromEValueType 将 EValueType 转换为 EConfigFieldsValueType
func (s *SFieldDefinition) convertValueTypeFromEValueType(valueType c_enum.EValueType) c_enum.EConfigFieldsValueType {
	switch valueType {
	case c_enum.EString:
		return c_enum.EConfigFieldsValueTypeString
	case c_enum.EInt8, c_enum.EUint8, c_enum.EInt16, c_enum.EUint16, c_enum.EInt32, c_enum.EUint32, c_enum.EInt64, c_enum.EUint64:
		return c_enum.EConfigFieldsValueTypeInt
	case c_enum.EFloat32, c_enum.EFloat64:
		return c_enum.EConfigFieldsValueTypeFloat
	case c_enum.EBool:
		return c_enum.EConfigFieldsValueTypeBoolean
	default:
		return c_enum.EConfigFieldsValueTypeString
	}
}
