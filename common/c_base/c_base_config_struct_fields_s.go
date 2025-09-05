package c_base

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

type SConfigStructFields struct {
	Name               string                     `json:"name" required:"true"`
	Code               string                     `json:"code" required:"true"`
	Group              string                     `json:"group"`
	ValueType          string                     `json:"valueType" dc:"string字符串、int整数、float浮点数、bool布尔值" required:"true"`
	ComponentType      EConfigFieldsComponentType `json:"componentType" dc:"组件类型" required:"true"`
	Step               *float32                   `json:"setup" required:"true" default:"1" dc:"步长（小步长）"`
	Required           bool                       `json:"required" required:"true" dc:"是否必填"`
	Unit               *string                    `json:"unit" dc:"单位"`
	Min                *int64                     `json:"min"`
	Max                *int64                     `json:"max"`
	Default            *string                    `json:"default"`
	SelectOptions      map[string]string          `json:"selectOptions"`
	Regex              *string                    `json:"regex" dc:"正则表达式"`
	RegexFailedMessage *string                    `json:"regexFailedMessage" dc:"正则表达式失败提醒"`
	Description        string                     `json:"description" required:"true"`
}

func (s *SConfigStructFields) String() string {
	if s == nil {
		return "SConfigStructFields(nil)"
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

	return fmt.Sprintf("SConfigStructFields{Name:%s, Code:%s, ValueType:%s, ComponentType:%s, Min:%d, Max:%d, Default:%s, Regex:%s}",
		s.Name, s.Code, s.ValueType, s.ComponentType, minVal, maxVal, defaultVal, regexVal)
}

func (s *SConfigStructFields) Check() error {
	if s == nil {
		return errors.New("SConfigStructFields is nil")
	}

	// 检查必填字段
	if s.Name == "" {
		return errors.New("Name is required")
	}
	if s.Description == "" {
		return errors.New("Description is required")
	}
	if s.Code == "" {
		return errors.New("Code is required")
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
