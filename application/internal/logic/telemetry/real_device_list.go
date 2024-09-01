package telemetry

import (
	"context"
	"example/ems/common/device"
	"example/ems/common/plugin"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTelemetry) RealDeviceList() (devices []plugin.DriverInfo) {

	return copyDriver(device.ListAllRealDevice())
}

func (s *sTelemetry) AllDeviceList() (devices []plugin.DriverInfo) {
	return copyDriver(device.ListAllDevice())
}

func copyDriver(drivers []*plugin.DriverInfo) []plugin.DriverInfo {
	ctx := context.Background()
	result := make([]plugin.DriverInfo, 0)
	for _, i := range drivers {
		result = append(result, plugin.DriverInfo{
			Group:            i.Group,
			Id:               i.Id,
			I18nName:         g.I18n().Tf(ctx, i.I18nName, i.I18nArgs...),
			Brand:            i.Brand,
			Model:            i.Model,
			IsMaster:         i.IsMaster,
			IsStation:        i.IsStation,
			IsVirtual:        i.IsVirtual,
			TelemetryI18nMap: i.TelemetryI18nMap,
			Description:      i.Description,
			Type:             i.Type,
		})
	}
	return result
}
