package c_base

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

type STelemetry struct {
	Name                string `json:"name,omitempty" yaml:"name"`                     // 遥测名称
	NationalizationName string `json:"nationalizationName" yaml:"nationalizationName"` // 遥测名称的国际化覆盖
	Unit                string `json:"unit,omitempty" yaml:"unit"`                     // 单位
	Remark              string `json:"remark,omitempty" yaml:"remark"`                 // 备注
}

func (s *STelemetry) String() string {
	i18nKey := "-"
	if s.NationalizationName != "" {
		i18nKey = s.NationalizationName
	}
	return fmt.Sprintf("%s\t%s\t%s\t%s\t", s.Name, i18nKey, s.Unit, s.Remark)
}

func (s *STelemetry) ToI18n(ctx context.Context) *STelemetry {
	key := g.I18n().T(ctx, s.NationalizationName)
	if key == "" {
		key = s.Remark
	}
	return &STelemetry{
		Name:                s.Name,
		NationalizationName: key,
		Unit:                s.Unit,
		Remark:              s.Remark,
	}
}

func TelemetryListI18n(ctx context.Context, list []*STelemetry) []*STelemetry {
	var result []*STelemetry
	for _, item := range list {
		result = append(result, item.ToI18n(ctx))
	}
	return result

}
