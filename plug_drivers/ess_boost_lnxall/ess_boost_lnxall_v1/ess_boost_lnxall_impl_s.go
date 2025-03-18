package ess_boost_lnxall_v1

import (
	"common/c_base"
	"common/c_device"
	"context"
	"time"
)

type sEssBoostLnxallEss struct {
	*c_base.SAlarmHandler
	*c_base.SDescription
	deviceConfig *c_base.SDriverConfig
	ctx          context.Context

	ammeter c_device.IAmmeter // 电表
	bms     c_device.IBms     // 电池
	pcs     c_device.IPcs     // 逆变器

	buttonScram     c_device.IGpio
	buttonDischarge c_device.IGpio
	buttonCharge    c_device.IGpio
	ledRunning      c_device.IGpio
	ledFault        c_device.IGpio
}

func (s *sEssBoostLnxallEss) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	s.deviceConfig = deviceConfig
	// 从配置中获取电表、PCS、BMS的配置

	panic("implement me")
}

func (s *sEssBoostLnxallEss) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetDriverType() c_base.EDeviceType {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetLastUpdateTime() *time.Time {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetMetaValueList() []*c_base.MetaValueWrapper {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetDeviceConfig() *c_base.SDriverConfig {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCellMinTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCellMaxTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCellAvgTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCellMinVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCellMaxVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCellAvgVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetSoc() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetSoh() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCapacity() (uint32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCycleCount() (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) SetPower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetTargetPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetRatedPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetFireEnvTemperature() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetCarbonMonoxideConcentration() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) HasSmoke() (bool, error) {
	//TODO implement me
	panic("implement me")
}
