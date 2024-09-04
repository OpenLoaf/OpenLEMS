package telemetry

import (
	"application/api/telemetry/v1"
	"application/internal/model/entity"
	"context"
	common "ems-plan"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetTelemetryDescription(ctx context.Context, req *v1.GetTelemetryDescriptionReq) (res *v1.GetTelemetryDescriptionRes, err error) {

	return &v1.GetTelemetryDescriptionRes{
		Ess: makeResponse(ctx, string(c_base.EStationEnergyStore)),
	}, nil
}

func makeResponse(ctx context.Context, name string) *v1.TelemetryDescriptionObj {
	stationEnergyStore := common.GetDeviceById(string(c_base.EStationEnergyStore))
	var stationTelemetryList []*entity.DeviceTelemetry
	stationTelemetryList = append(stationTelemetryList, &entity.DeviceTelemetry{
		DeviceId:      stationEnergyStore.GetDeviceConfig().Id,
		I8nName:       stationEnergyStore.GetDeviceConfig().Name,
		TelemetryKeys: c_base.TelemetryListI18n(ctx, stationEnergyStore.GetDescription().Telemetry),
	})
	children := stationEnergyStore.(c_device.IStationEnergyStore).GetChildren()
	for _, child := range children {
		stationTelemetryList = append(stationTelemetryList, &entity.DeviceTelemetry{
			DeviceId:      child.GetDeviceConfig().Id,
			I8nName:       child.GetDeviceConfig().Name,
			TelemetryKeys: c_base.TelemetryListI18n(ctx, child.GetDescription().Telemetry),
		})
	}

	return &v1.TelemetryDescriptionObj{
		Name:     g.I18n().T(ctx, stationEnergyStore.GetDeviceConfig().Name),
		Children: stationTelemetryList,
	}
}
