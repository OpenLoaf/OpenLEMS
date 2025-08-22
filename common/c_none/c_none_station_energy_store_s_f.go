package c_none

import (
	"common/c_base"
)

type sNoneStationEnergyStore struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
	deviceConfig *c_base.SDeviceConfig
	protocol     c_base.IProtocol
	childDevice  []c_base.IDevice
}

func (s *sNoneStationEnergyStore) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.protocol = protocol
	s.childDevice = childDevice
}

func (s *sNoneStationEnergyStore) Shutdown() {

}

func (s *sNoneStationEnergyStore) GetDriverType() c_base.EDeviceType {
	return c_base.EStationEnergyStore
}

func (s *sNoneStationEnergyStore) GetDriverDescription() *c_base.SDriverDescription {
	return nil
}

func (s *sNoneStationEnergyStore) SetReset() error {
	return NoneErr
}

func (s *sNoneStationEnergyStore) SetStatus(status c_base.EEnergyStoreStatus) error {
	return NoneErr
}

func (s *sNoneStationEnergyStore) SetGridMode(mode c_base.EGridMode) error {
	return NoneErr
}

func (s *sNoneStationEnergyStore) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return c_base.EPcsStatusUnknown, NoneErr
}

func (s *sNoneStationEnergyStore) GetGridMode() (c_base.EGridMode, error) {
	return c_base.EGridUnknown, NoneErr
}

func (s *sNoneStationEnergyStore) SetPower(power int32) error {
	return NoneErr
}

func (s *sNoneStationEnergyStore) SetReactivePower(power int32) error {
	return NoneErr
}

func (s *sNoneStationEnergyStore) SetPowerFactor(factor float32) error {
	return NoneErr
}

func (s *sNoneStationEnergyStore) GetTargetPower() int32 {
	return 0
}

func (s *sNoneStationEnergyStore) GetTargetReactivePower() int32 {
	return 0
}

func (s *sNoneStationEnergyStore) GetTargetPowerFactor() float32 {
	return 1
}

func (s *sNoneStationEnergyStore) GetPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetApparentPower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetReactivePower() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetRatedPower() int32 {
	return 0
}

func (s *sNoneStationEnergyStore) GetMaxInputPower() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetMaxOutputPower() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetTodayIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetHistoryIncomingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetTodayOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetHistoryOutgoingQuantity() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCellMinTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCellMaxTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCellAvgTemp() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCellMinVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCellMaxVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCellAvgVoltage() (float32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetSoc() (float32, error) {
	return 0.0, NoneErr
}

func (s *sNoneStationEnergyStore) GetSoh() (float32, error) {
	return 0.0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCapacity() (uint32, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCycleCount() (uint, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetDcPower() (float64, error) {
	return 0.0, NoneErr
}

func (s *sNoneStationEnergyStore) GetFireEnvTemperature() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) GetCarbonMonoxideConcentration() (float64, error) {
	return 0, NoneErr
}

func (s *sNoneStationEnergyStore) HasSmoke() (bool, error) {
	return false, NoneErr
}

func (s *sNoneStationEnergyStore) GetAllowControl() bool {
	return false
}

func (s *sNoneStationEnergyStore) GetChildren() []c_base.IDevice {
	return nil
}
