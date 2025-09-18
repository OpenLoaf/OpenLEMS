package convert

import (
	"common/c_base"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func SConfigStructFieldsToI18n(ctx context.Context, s *c_base.SConfigStructFields) *c_base.SConfigStructFields {
	key := g.I18n().T(ctx, s.Name)
	if key == "" {
		key = s.Description
	}

	// 处理单位信息
	var unit *string
	if s.Unit != nil {
		unit = s.Unit
	}

	return &c_base.SConfigStructFields{
		Key:                s.Key,
		Name:               key,
		Group:              s.Group,
		ValueType:          s.ValueType,
		ComponentType:      s.ComponentType,
		Step:               s.Step,
		Required:           s.Required,
		Unit:               unit,
		Min:                s.Min,
		Max:                s.Max,
		Default:            s.Default,
		ValueExplain:       s.ValueExplain,
		ParamExplain:       s.ParamExplain,
		Regex:              s.Regex,
		RegexFailedMessage: s.RegexFailedMessage,
		Description:        s.Description,
	}
}

func ConfigStructFieldsListI18n(ctx context.Context, list []*c_base.SConfigStructFields) []*c_base.SConfigStructFields {
	var result []*c_base.SConfigStructFields
	for _, item := range list {
		result = append(result, SConfigStructFieldsToI18n(ctx, item))
	}
	return result
}
