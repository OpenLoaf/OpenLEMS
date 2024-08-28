package pylonCheckwatt_v1

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"fmt"
)

type PylonCheckwattEss struct {
	alarmChan chan *c_base.EAlarmLevel

	//*c_base.SConfigImpl // 配置信息
	ctx    context.Context
	unitId uint8 // modbus转发的id

	ammeter c_device.IAmmeter

	//bms      *common_cabinet.CabinetBms      // 电池
	//pcs      *common_cabinet.CabinetPcs      // 逆变器
	//fire     *common_cabinet.CabinetFire     // 消防
	//cooling  *common_cabinet.CabinetCooling  // 制冷
	//humidity *common_cabinet.CabinetHumidity // 温湿度
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
	}
*/

func (p *PylonCheckwattEss) GetAlarmChan() chan<- *c_base.EAlarmLevel {
	return p.alarmChan
}

func (p *PylonCheckwattEss) HandlerAlarm(alarm *c_base.SAlarmDetail) {
	fmt.Printf("handler alarm: %v\n", alarm)
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

func (p *PylonCheckwattEss) SetPower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetReactivePower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetTargetPower() float64 {
	//TODO implement me
	panic("implement me")
}

func (p *PylonCheckwattEss) GetTargetReactivePower() float64 {
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
