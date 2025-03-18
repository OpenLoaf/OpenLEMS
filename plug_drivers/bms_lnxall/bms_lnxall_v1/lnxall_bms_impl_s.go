package bms_lnxall_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_error"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"modbus/p_modbus"
)

type sBmsLnxallBms struct {
	p_modbus.IModbusProtocol
	ctx          context.Context
	log          *glog.Logger
	deviceConfig *c_base.SDriverConfig
	*c_base.SDescription
}

func (s *sBmsLnxallBms) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	s.IModbusProtocol = protocol.(p_modbus.IModbusProtocol)

	// 注册
	s.IModbusProtocol.RegisterRead(s.ctx,
		GroupBasic,
		//GroupStatistic,
	)

}

func (s *sBmsLnxallBms) Destroy() {

}

func (s *sBmsLnxallBms) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceBms
}

func (s *sBmsLnxallBms) SetReset() error {
	return c_error.NonSupportError
}

func (s *sBmsLnxallBms) SetBmsStatus(status c_device.EBmsStatus) error {
	return c_error.NonSupportError
}

func (s *sBmsLnxallBms) GetCellMinTemp() (float32, error) {
	return s.GetFloat32Value(TEMP_MIN)
}

func (s *sBmsLnxallBms) GetCellMaxTemp() (float32, error) {
	return s.GetFloat32Value(TEMP_MAX)
}

func (s *sBmsLnxallBms) GetCellAvgTemp() (float32, error) {
	return s.GetFloat32Value(TEMP_AVG)
}

func (s *sBmsLnxallBms) GetCellMinVoltage() (float32, error) {
	return s.GetFloat32Value(BATT_MIN)
}

func (s *sBmsLnxallBms) GetCellMaxVoltage() (float32, error) {
	return s.GetFloat32Value(BATT_MAX)
}

func (s *sBmsLnxallBms) GetCellAvgVoltage() (float32, error) {
	return s.GetFloat32Value(BATT_AVG)
}

func (s *sBmsLnxallBms) GetBmsStatus() (c_device.EBmsStatus, error) {
	value, err := s.GetUintValue(RACK_STATUS)
	if err != nil {
		return c_device.EBmsStatusUnknown, err
	}
	switch value {
	case 1:
		return c_device.EBmsStatusCharge, nil
	case 2:
		return c_device.EBmsStatusDischarge, nil
	case 3:
		return c_device.EBmsStatusFault, nil
	default:
		return c_device.EBmsStatusUnknown, nil
	}

}

func (s *sBmsLnxallBms) GetSoc() (float32, error) {
	return s.GetFloat32Value(RACK_SOC)
}

func (s *sBmsLnxallBms) GetSoh() (float32, error) {
	return s.GetFloat32Value(RACK_SOH)
}

func (s *sBmsLnxallBms) GetCapacity() (uint32, error) {
	return 0, c_error.NonSupportError
}

func (s *sBmsLnxallBms) GetCycleCount() (uint, error) {
	// 获取 总放电量 / 总用电量= 总循环次数

	provider, err := s.GetHistoryOutgoingQuantity()
	if err != nil {
		return 0, err
	}
	use, err := s.GetHistoryIncomingQuantity()
	if err != nil {
		return 0, err
	}

	return uint(provider / use), nil
}

func (s *sBmsLnxallBms) GetRatedPower() int32 {
	return -1
}

func (s *sBmsLnxallBms) GetMaxInputPower() (float32, error) {
	return 0, c_error.NonSupportError
}

func (s *sBmsLnxallBms) GetMaxOutputPower() (float32, error) {
	return 0, c_error.NonSupportError
}

func (s *sBmsLnxallBms) GetDcPower() (float64, error) {
	return 0, c_error.NonSupportError
}

func (s *sBmsLnxallBms) GetDcVoltage() (float64, error) {
	return s.GetFloat64Value(BATT_VOLTAGE)
}

func (s *sBmsLnxallBms) GetDcCurrent() (float64, error) {
	return s.GetFloat64Value(BATT_CURRENT)
}

func (s *sBmsLnxallBms) GetTodayIncomingQuantity() (float64, error) {
	return 0, c_error.NonSupportError
}

func (s *sBmsLnxallBms) GetHistoryIncomingQuantity() (float64, error) {
	return s.GetFloat64Value(TOTAL_CHARGE_ENERGY)
}

func (s *sBmsLnxallBms) GetTodayOutgoingQuantity() (float64, error) {
	return 0, c_error.NonSupportError
}

func (s *sBmsLnxallBms) GetHistoryOutgoingQuantity() (float64, error) {
	return s.GetFloat64Value(TOTAL_DISCHARGE_ENERGY)
}
