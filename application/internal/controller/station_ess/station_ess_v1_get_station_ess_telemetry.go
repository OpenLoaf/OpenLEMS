package station_ess

import (
	"application/internal/model/entity"
	"context"
	common "ems-plan"
	"ems-plan/util"
	"github.com/gogf/gf/v2/util/gconv"

	"application/api/station_ess/v1"
)

func (c *ControllerV1) GetStationEssTelemetry(ctx context.Context, req *v1.GetStationEssTelemetryReq) (res *v1.GetStationEssTelemetryRes, err error) {
	ess := common.DeviceInstance.GetStationEnergyStore()

	if ess == nil {
		return nil, nil
	}

	lastUpdateTimeObj := ess.GetLastUpdateTime()
	var lastUpdateTimeStr string
	if lastUpdateTimeObj != nil {
		lastUpdateTimeStr = lastUpdateTimeObj.Format("2006-01-02 15:04:05.000")
	}

	e := &entity.EssStatus{
		DeviceId:       ess.GetDeviceConfig().Id,
		I18nName:       ess.GetDeviceConfig().Name,
		LastUpdateTime: lastUpdateTimeStr,

		TargetPower:         gconv.String(ess.GetTargetPower()),
		TargetReactivePower: gconv.String(ess.GetTargetReactivePower()),
		TargetPowerFactor:   util.Float32ToString(ess.GetTargetPowerFactor(), 2),
	}
	if value, err := ess.GetPower(); err == nil {
		e.Power = util.Float64ToString(value, 4)
	}

	if value, err := ess.GetApparentPower(); err == nil {
		e.ApparentPower = util.Float64ToString(value, 4)
	}

	if value, err := ess.GetReactivePower(); err == nil {
		e.ReactivePower = util.Float64ToString(value, 4)
	}
	if value, err := ess.GetSoc(); err == nil {
		e.Soc = util.Float32ToString(value, 2)
	}

	if value, err := ess.GetTodayIncomingQuantity(); err == nil {
		e.TodayCharge = util.Float64ToString(value, 4)
	}

	if value, err := ess.GetTodayOutgoingQuantity(); err == nil {
		e.TodayDischarge = util.Float64ToString(value, 4)
	}

	if value, err := ess.GetHistoryIncomingQuantity(); err == nil {
		e.HistoryCharge = util.Float64ToString(value, 4)
	}
	if value, err := ess.GetHistoryOutgoingQuantity(); err == nil {
		e.HistoryDischarge = util.Float64ToString(value, 4)
	}

	return &v1.GetStationEssTelemetryRes{EssStatus: e}, nil
}
