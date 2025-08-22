package c_none

import (
	"common/c_base"
)

type sNoneEnergyStore struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneEnergyStore) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneEnergyStore) Shutdown() {

}

func (s *sNoneEnergyStore) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
}

func (s *sNoneEnergyStore) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneEnergyStore) SetReset() error {
	return NoneErr
}

func (s *sNoneEnergyStore) SetStatus(status c_base.EEnergyStoreStatus) error {
	return NoneErr
}

func (s *sNoneEnergyStore) SetGridMode(mode c_base.EGridMode) error {
	return NoneErr
}

func (s *sNoneEnergyStore) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return c_base.EPcsStatusUnknown, NoneErr
}

func (s *sNoneEnergyStore) GetGridMode() (c_base.EGridMode, error) {
	return c_base.EGridUnknown, NoneErr
}

func (s *sNoneEnergyStore) SetPower(power int32) error {
	return NoneErr
}

func (s *sNoneEnergyStore) SetReactivePower(power int32) error {
	return NoneErr
}

func (s *sNoneEnergyStore) SetPowerFactor(factor float32) error {
	return NoneErr
}

func (s *sNoneEnergyStore) GetTargetPower() int32 {
	return 0
}

func (s *sNoneEnergyStore) GetTargetReactivePower() int32 {
	return 0
}

func (s *sNoneEnergyStore) GetTargetPowerFactor() float32 {
	return 1
}

func (s *sNoneEnergyStore) GetPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetApparentPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetReactivePower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetRatedPower() int32 {
	return 0
}

func (s *sNoneEnergyStore) GetMaxInputPower() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetMaxOutputPower() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetCellMinTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetCellMaxTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetCellAvgTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetCellMinVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetCellMaxVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetCellAvgVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetSoc() (float32, error) {
	return 0.0, NoneErr
}

func (s *sNoneEnergyStore) GetSoh() (float32, error) {
	return 0.0, NoneErr
}

func (s *sNoneEnergyStore) GetCapacity() (uint32, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetCycleCount() (uint, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetDcPower() (float64, error) {
	return 0.0, NoneErr
}

func (s *sNoneEnergyStore) GetFireEnvTemperature() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) GetCarbonMonoxideConcentration() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneEnergyStore) HasSmoke() (bool, error) {
	return false, NoneErr
}
