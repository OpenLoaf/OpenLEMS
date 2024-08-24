package pylon_checkwatt_v1

import (
	"common_cabinet"
	"context"
	"ems-plan/c_device"
	"ems-plan/c_group"
	"ems-plan/c_telemetry"
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

type PylonCheckwattEss struct {
	*c_device.SConfig // 配置信息
	ctx               context.Context
	unitId            uint8 // modbus转发的id

	ammeter c_device.IAmmeter

	bms      *common_cabinet.CabinetBms      // 电池
	pcs      *common_cabinet.CabinetPcs      // 逆变器
	fire     *common_cabinet.CabinetFire     // 消防
	cooling  *common_cabinet.CabinetCooling  // 制冷
	humidity *common_cabinet.CabinetHumidity // 温湿度
}

func CreateEss(ctx context.Context, cabinetId uint8, params map[string]string,
	ammeter c_device.IAmmeter,
	pcs *common_cabinet.CabinetPcs,
	bms *common_cabinet.CabinetBms,
	fire *common_cabinet.CabinetFire,
	cooling *common_cabinet.CabinetCooling,
	humidity *common_cabinet.CabinetHumidity) c_device.IEnergyStore {
	_ess := &PylonCheckwattEss{
		ctx:      ctx,
		unitId:   cabinetId,
		ammeter:  ammeter,
		bms:      bms,
		pcs:      pcs,
		fire:     fire,
		cooling:  cooling,
		humidity: humidity,
		SConfig:  c_device.NewConfig(cabinetId, c_device.EEnergyStore, c_group.EGroupEnergyStore, false, params),
	}

	if _ess.bms == nil || _ess.pcs == nil {
		panic(fmt.Sprintf("一个柜子需要电池和PCS组成，但是现在有一项不存在，请检查配置！ cabinetId:%d", cabinetId))
	}
	return _ess
}

func (p *PylonCheckwattEss) GetCache() *gcache.Cache {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetLastUpdateTime() *time.Time {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetDescription() c_device.SDescription {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetRatedPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetMaxInputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetMaxOutputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetDcVoltage() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetDcCurrent() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetBmsStatus(status c_device.EBmsStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetBmsStatus() (c_device.EBmsStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetSoc() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetSoh() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetCellTemp() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetCellVoltage() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetCapacity() (uint16, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetCycleCount() (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetStatus(status c_device.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetGridMode(mode c_device.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetStatus() (c_device.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetGridMode() (c_device.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetFireEnvTemperature() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetCarbonMonoxideConcentration() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) HasSmoke() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetDcStatisticsQuantity() c_telemetry.IStatisticsQuantity {
	//TODO implement me
	panic("implement me")
}
