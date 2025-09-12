package convert

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func STelemetryToI18n(ctx context.Context, s *c_base.STelemetry) *c_base.STelemetry {
	key := g.I18n().T(ctx, s.Name)
	if key == "" {
		key = s.Desc
	}
	return &c_base.STelemetry{
		Key:  s.Key,
		Name: key,
		Unit: s.Unit,
		Desc: s.Desc,
	}
}

func TelemetryListI18n(ctx context.Context, list []*c_base.STelemetry) []*c_base.STelemetry {
	var result []*c_base.STelemetry
	for _, item := range list {
		result = append(result, STelemetryToI18n(ctx, item))
	}
	return result

}
