package c_base

import (
	"common/c_enum"
)

// SConfigPoint 配置点位，用于配置结构体字段描述
type SConfigPoint struct {
	*SPoint // 嵌套基础点位信息

	// 配置特定的字段（不重复SPoint中已有的字段）
	Required           bool                              `json:"required" dc:"是否必填"`                         // 是否必填
	Default            *string                           `json:"default,omitempty" dc:"默认值"`                 // 默认值
	Regex              *string                           `json:"regex,omitempty" dc:"正则表达式验证"`               // 正则表达式验证
	RegexFailedMessage *string                           `json:"regexFailedMessage,omitempty" dc:"正则验证失败提示"` // 正则验证失败提示
	Step               *float32                          `json:"step,omitempty" dc:"步长（用于数字输入）"`             // 步长（用于数字输入）
	ComponentType      c_enum.EConfigFieldsComponentType `json:"componentType" dc:"组件类型"`                    // 组件类型
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

func (s *SConfigPoint) GetComponentType() c_enum.EConfigFieldsComponentType {
	return s.ComponentType
}

// GetValueExplainWithParams 获取Value解释，支持动态参数
func (s *SConfigPoint) GetValueExplainWithParams(value any, params map[string]any) (string, string, error) {
	if s.SPoint == nil {
		return "", "", nil
	}

	// 使用统一的公共函数进行值解释
	return ExplainValueWithColor(value, s.SPoint.ValueExplain, s.SPoint.Precise)
}

// ToFieldDefinition 将SConfigPoint转换为SFieldDefinition对象
func (s *SConfigPoint) ToFieldDefinition() *SFieldDefinition {
	if s == nil || s.SPoint == nil {
		return nil
	}

	// 转换值类型
	valueType := s.convertEValueTypeToConfigFieldsValueType(s.SPoint.ValueType)

	// 使用保存的组件类型，如果没有则根据值类型推断
	componentType := s.ComponentType
	if componentType == "" {
		componentType = s.inferComponentType(valueType)
	}

	// 处理指针类型字段
	var unit *string
	if s.SPoint.Unit != "" {
		unit = &s.SPoint.Unit
	}

	var min, max *int64
	if s.SPoint.Min != 0 {
		min = &s.SPoint.Min
	}
	if s.SPoint.Max != 0 {
		max = &s.SPoint.Max
	}

	// 创建SFieldDefinition
	// 组信息
	var group *SPointGroup
	if s.SPoint.Group != nil {
		group = &SPointGroup{
			GroupKey:  s.SPoint.Group.GroupKey,
			GroupName: s.SPoint.Group.GroupName,
			GroupSort: s.SPoint.Group.GroupSort,
			Disable:   s.SPoint.Group.Disable,
		}
	}
	fieldDef := &SFieldDefinition{
		Key:                s.SPoint.Key,
		Name:               s.SPoint.Name,
		Group:              group,
		ValueType:          valueType,
		ComponentType:      componentType,
		Step:               s.Step,
		Required:           s.Required,
		Unit:               unit,
		Min:                min,
		Max:                max,
		Default:            s.Default,
		ValueExplain:       s.SPoint.ValueExplain,
		Regex:              s.Regex,
		RegexFailedMessage: s.RegexFailedMessage,
		Description:        s.SPoint.Desc,
	}

	return fieldDef
}

// convertEValueTypeToConfigFieldsValueType 将EValueType转换为EConfigFieldsValueType
func (s *SConfigPoint) convertEValueTypeToConfigFieldsValueType(valueType c_enum.EValueType) c_enum.EConfigFieldsValueType {
	switch valueType {
	case c_enum.EBool:
		return c_enum.EConfigFieldsValueTypeBoolean
	case c_enum.EInt8, c_enum.EUint8, c_enum.EInt16, c_enum.EUint16, c_enum.EInt32, c_enum.EUint32, c_enum.EInt64, c_enum.EUint64:
		return c_enum.EConfigFieldsValueTypeInt
	case c_enum.EFloat32, c_enum.EFloat64:
		return c_enum.EConfigFieldsValueTypeFloat
	case c_enum.EString:
		return c_enum.EConfigFieldsValueTypeString
	default:
		return c_enum.EConfigFieldsValueTypeString
	}
}

// inferComponentType 根据值类型推断组件类型
// 注意：label组件类型需要通过ct标签显式指定，不会自动推断
func (s *SConfigPoint) inferComponentType(valueType c_enum.EConfigFieldsValueType) c_enum.EConfigFieldsComponentType {
	switch valueType {
	case c_enum.EConfigFieldsValueTypeBoolean:
		return c_enum.EConfigFieldsComponentTypeSwitch
	case c_enum.EConfigFieldsValueTypeInt, c_enum.EConfigFieldsValueTypeFloat:
		return c_enum.EConfigFieldsComponentTypeNumber
	default:
		return c_enum.EConfigFieldsComponentTypeText
	}
}

// getGroupName 已废弃：Group 现在为 *SPointGroup
