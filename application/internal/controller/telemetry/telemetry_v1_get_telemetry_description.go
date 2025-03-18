package telemetry

import (
	"application/api/telemetry/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"common/c_device"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetTelemetryDescription(ctx context.Context, req *v1.GetTelemetryDescriptionReq) (res *v1.GetTelemetryDescriptionRes, err error) {
	stationEnergyStore := common.GetStationEnergyStore()
	if stationEnergyStore == nil {
		return nil, gerror.New("场站储能不存在！")
	}
	return &v1.GetTelemetryDescriptionRes{
		Ess: makeResponse(ctx, stationEnergyStore),
	}, nil
}

func makeResponse(ctx context.Context, stationEnergyStore c_device.IStationEnergyStore) *v1.TelemetryDescriptionObj {

	var stationTelemetryList []*entity.DeviceTelemetry
	//var stationTelemetryList
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

	// TODO 添加排序

	//garray.NewSortedArray(func(v1, v2 interface{}) int {
	//	if v1.(*entity.DeviceTelemetry).I8nName > v2.(*entity.DeviceTelemetry).I8nName {
	//		return 1
	//	}
	//})

	return &v1.TelemetryDescriptionObj{
		Name:     g.I18n().T(ctx, stationEnergyStore.GetDeviceConfig().Name),
		Children: stationTelemetryList,
	}
}
