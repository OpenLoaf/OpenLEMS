package c_base

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type STelemetry struct {
	Name    string // 遥测名称
	I18nKey string // 遥测名称的国际化覆盖
	Unit    string // 单位
	Remark  string // 备注
}

func (s *STelemetry) ToI18n(ctx context.Context) *STelemetry {
	key := g.I18n().T(ctx, s.I18nKey)
	if key == "" {
		key = s.Remark
	}
	return &STelemetry{
		Name:    s.Name,
		I18nKey: key,
		Unit:    s.Unit,
		Remark:  s.Remark,
	}
}

func TelemetryListI18n(ctx context.Context, list []*STelemetry) []*STelemetry {
	var result []*STelemetry
	for _, item := range list {
		result = append(result, item.ToI18n(ctx))
	}
	return result

}
