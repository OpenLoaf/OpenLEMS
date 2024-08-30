package pylonCheckwatt_v1

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

type PylonCheckwattEss struct {
	c_base.IDriverConfig
	*c_base.SAlarmHandler
	cabinetId uint8 // 属于哪个柜子

	//*c_base.SConfigImpl // 配置信息
	ctx     context.Context
	unitId  uint8             // modbus转发的id
	ammeter c_device.IAmmeter // 电表
	bms     c_device.IBms     // 电池
	pcs     c_device.IPcs     // 逆变器
}

/*
func CreateEss(ctx context.Context, cabinetId uint8, params map[string]string,

	ammeter c_base.IAmmeter,
	pcs *common_cabinet.CabinetPcs,
	bms *common_cabinet.CabinetBms,
	fire *common_cabinet.CabinetFire,
	cooling *common_cabinet.CabinetCooling,
	humidity *common_cabinet.CabinetHumidity) c_base.IEnergyStore {
	_ess := &PylonCheckwattEss{
		ctx:         ctx,
		unitId:      cabinetId,
		ammeter:     ammeter,
		bms:         bms,
		pcs:         pcs,
		fire:        fire,
		cooling:     cooling,
		humidity:    humidity,
		SConfigImpl: c_base.NewConfig(cabinetId, c_base.EDeviceEnergyStore, c_group.EGroupEnergyStore, false, params),
	}

	if _ess.bms == nil || _ess.pcs == nil {
		panic(fmt.Sprintf("一个柜子需要电池和PCS组成，但是现在有一项不存在，请检查配置！ cabinetId:%d", cabinetId))
	}
	return _ess
}*/

func (p *PylonCheckwattEss) Init(ctx context.Context, client c_base.IProtocol, cfg any) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetId() string {
	return fmt.Sprintf("pylonCheckwattEss_%d", p.cabinetId)
}

func (p *PylonCheckwattEss) GetType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
}

func (p *PylonCheckwattEss) GetDescription() c_base.SDescription {
	return c_base.SDescription{
		Brand:  "Plyon",
		Model:  "Checkwatt",
		Type:   c_base.EDeviceEnergyStore,
		Remark: "虚拟派能柜，整合PCS与BMS",
	}
}

func (p *PylonCheckwattEss) GetCache() *gcache.Cache {
	p.bms.GetCache()

	return nil
}

func (p *PylonCheckwattEss) GetLastUpdateTime() *time.Time {

	return nil
}

func (p *PylonCheckwattEss) IsActivate() bool {
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

func (p *PylonCheckwattEss) GetRatedPower() (int32, error) {
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

func (p *PylonCheckwattEss) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetPower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetTargetPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetTargetPowerFactor() float32 {
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
