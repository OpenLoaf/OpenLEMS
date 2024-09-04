package station

import (
	"application/internal/model/entity"
	common "ems-plan"
	"ems-plan/c_base"
	"ems-plan/util"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sStation) GetEnergyStoreStatus() *entity.EnergyStoreStatus {
	ess := common.GetStationEnergyStore()

	if ess == nil {
		return nil
	}

	lastUpdateTimeObj := ess.GetLastUpdateTime()
	var lastUpdateTimeStr string
	if lastUpdateTimeObj != nil {
		lastUpdateTimeStr = lastUpdateTimeObj.Format("2006-01-02 15:04:05.000")
	}
	status, err := ess.GetStatus()
	if err != nil {
		status = c_base.EPcsStatusUnknown
	}
	essStatus := &entity.EnergyStoreStatus{
		DeviceId:            ess.GetDeviceConfig().Id,
		I18nName:            ess.GetDeviceConfig().Name,
		LastUpdateTime:      lastUpdateTimeStr,
		Status:              status,
		TargetPower:         gconv.String(ess.GetTargetPower()),
		TargetReactivePower: gconv.String(ess.GetTargetReactivePower()),
		TargetPowerFactor:   util.Float32ToString(ess.GetTargetPowerFactor(), 2),
	}
	if value, err := ess.GetPower(); err == nil {
		essStatus.Power = util.Float64ToString(value, 2)
	}

	if value, err := ess.GetApparentPower(); err == nil {
		essStatus.ApparentPower = util.Float64ToString(value, 2)
	}

	if value, err := ess.GetReactivePower(); err == nil {
		essStatus.ReactivePower = util.Float64ToString(value, 2)
	}
	if value, err := ess.GetSoc(); err == nil {
		essStatus.Soc = util.Float32ToString(value, 2)
	}

	if value, err := ess.GetTodayIncomingQuantity(); err == nil {
		essStatus.TodayCharge = util.Float64ToString(value, 4)
	}

	if value, err := ess.GetTodayOutgoingQuantity(); err == nil {
		essStatus.TodayDischarge = util.Float64ToString(value, 4)
	}

	if value, err := ess.GetHistoryIncomingQuantity(); err == nil {
		essStatus.HistoryCharge = util.Float64ToString(value, 4)
	}
	if value, err := ess.GetHistoryOutgoingQuantity(); err == nil {
		essStatus.HistoryDischarge = util.Float64ToString(value, 4)
	}
	return essStatus
}

func (s *sStation) SetEnergyStorePower(power int32) error {
	ess := common.GetStationEnergyStore()
	if ess == nil {
		return gerror.New("设备不存在")
	}
	return ess.SetPower(power)
}

func (s *sStation) SetEnergyStoreStatus(status c_base.EEnergyStoreStatus) error {
	ess := common.GetStationEnergyStore()
	if ess == nil {
		return gerror.New("设备不存在")
	}
	return ess.SetStatus(status)
}
