package c_base

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
)

type SConfigFields struct {
	Name               string                     `json:"name" required:"true"`
	Description        string                     `json:"description" required:"true"`
	Code               string                     `json:"code" required:"true"`
	ValueType          string                     `json:"valueType" dc:"string字符串、int整数、float浮点数、bool布尔值" required:"true"`
	ComponentType      EConfigFieldsComponentType `json:"componentType" dc:"组件类型" required:"true"`
	Min                uint8                      `json:"min"`
	Max                uint8                      `json:"max"`
	Default            string                     `json:"default"`
	SelectOptions      map[string]string          `json:"selectOptions"`
	Regex              string                     `json:"regex" dc:"正则表达式"`
	RegexFailedMessage string                     `json:"regexFailedMessage" dc:"正则表达式失败提醒"`
}

func (s *SConfigFields) String() string {
	if s == nil {
		return "SConfigFields(nil)"
	}

	return fmt.Sprintf("SConfigFields{Name:%s, Code:%s, ValueType:%s, ComponentType:%s, Min:%d, Max:%d, Default:%s, Regex:%s}",
		s.Name, s.Code, s.ValueType, s.ComponentType, s.Min, s.Max, s.Default, s.Regex)
}

func (s *SConfigFields) Check() error {
	if s == nil {
		return errors.New("SConfigFields is nil")
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
	if s.Min > s.Max {
		return errors.New("Min value cannot be greater than Max value")
	}

	// 检查正则表达式
	if s.Regex != "" {
		if _, err := regexp.Compile(s.Regex); err != nil {
			return errors.Errorf("Invalid regex pattern: %v", err)
		}
	}

	return nil
}
