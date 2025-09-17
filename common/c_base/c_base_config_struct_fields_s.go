package c_base

import (
	"common/c_enum"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type SConfigStructFields struct {
	Name               string                            `json:"name" yaml:"name" short:"name" required:"true"`
	Code               string                            `json:"code" yaml:"code" short:"code" `
	Group              string                            `json:"group" yaml:"group" short:"group"`
	ValueType          c_enum.EConfigFieldsValueType     `json:"valueType" yaml:"value_type" short:"vt" dc:"string字符串、int整数、float浮点数、bool布尔值" required:"true"`
	ComponentType      c_enum.EConfigFieldsComponentType `json:"componentType" yaml:"component_type" short:"ct" dc:"组件类型" required:"true"`
	Step               *float32                          `json:"step" yaml:"step" short:"step" default:"1" dc:"步长（小步长）"`
	Required           bool                              `json:"required" yaml:"required" short:"req" required:"true" dc:"是否必填"`
	Unit               *string                           `json:"unit" yaml:"unit" short:"unit" dc:"单位"`
	Min                *int64                            `json:"min" yaml:"min" short:"min"`
	Max                *int64                            `json:"max" yaml:"max" short:"max"`
	Default            *string                           `json:"default" yaml:"default" short:"def"`
	SelectOptions      []string                          `json:"selectOptions" yaml:"select_options" short:"opts"`
	Regex              *string                           `json:"regex" yaml:"regex" short:"regex" dc:"正则表达式"`
	RegexFailedMessage *string                           `json:"regexFailedMessage" yaml:"regex_failed_message" short:"rfm" dc:"正则表达式失败提醒"`
	Description        string                            `json:"description" yaml:"description" short:"desc" required:"true"`
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

	// 处理 SelectOptions 字段
	var selectOptionsStr string
	if len(s.SelectOptions) > 0 {
		selectOptionsStr = fmt.Sprintf("[%s]", strings.Join(s.SelectOptions, ", "))
	} else {
		selectOptionsStr = "[]"
	}

	return fmt.Sprintf("SConfigStructFields{Key:%s, Code:%s, ValueType:%s, ComponentType:%s, After:%d, Before:%d, Default:%s, Regex:%s, SelectOptions:%s}",
		s.Name, s.Code, s.ValueType, s.ComponentType, minVal, maxVal, defaultVal, regexVal, selectOptionsStr)
}

func (s *SConfigStructFields) Check() error {
	if s == nil {
		return errors.New("SConfigStructFields is nil")
	}

	// 检查必填字段
	if s.Name == "" {
		return errors.New("Key is required")
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
		return errors.New("After value cannot be greater than Before value")
	}

	// 检查正则表达式
	if s.Regex != nil && *s.Regex != "" {
		if _, err := regexp.Compile(*s.Regex); err != nil {
			return errors.Errorf("Invalid regex pattern: %v", err)
		}
	}

	return nil
}
