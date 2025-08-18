package pcs_lnxall_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_log"
	"common/c_modbus"
	"context"
)

type sPcsLnxallPcs struct {
	c_modbus.IModbusProtocol
	ctx          context.Context
	deviceConfig *c_base.SDeviceConfig
	*c_base.SDriverDescription
}

var _ c_device.IPcs = (*sPcsLnxallPcs)(nil)

func (s *sPcsLnxallPcs) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.IModbusProtocol = protocol.(c_modbus.IModbusProtocol)
	s.deviceConfig = deviceConfig

	// 注册轮询
	s.RegisterRead(s.ctx,
		GroupAcInfo, GroupOtherInfo,
	)
}

func (s *sPcsLnxallPcs) Shutdown() {
	c_log.Noticef(s.ctx, "sPcsLnxallPcs Shutdown")
}

func (s *sPcsLnxallPcs) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (s *sPcsLnxallPcs) SetReset() error {
	c_log.Warningf(s.ctx, "SetReset() not support!")
	return nil
}

func (s *sPcsLnxallPcs) SetStatus(status c_base.EEnergyStoreStatus) error {
	c_log.Warningf(s.ctx, "SetStatus() not support!")
	return nil
}

func (s *sPcsLnxallPcs) SetGridMode(mode c_base.EGridMode) error {
	c_log.Warningf(s.ctx, "SetGridMode() not support!")
	return nil
}

func (s *sPcsLnxallPcs) GetStatus() (c_base.EEnergyStoreStatus, error) {
	value, err := s.GetUintValue(WorkState)
	if err != nil {
		return c_base.EPcsStatusUnknown, err
	}
	//c_log.Noticef(s.ctx, "Pcs状态获取%v", value)
	switch value {
	case 0:
		return c_base.EPcsStatusOff, nil
	case 1:
		return c_base.EPcsStatusSync, nil
	case 5:
		return c_base.EPcsStatusStandby, nil
	case 32:
		return c_base.EPcsStatusFault, nil
	case 257:
		apPower, err := s.GetApparentPower()
		if err != nil {
			return c_base.EPcsStatusUnknown, err
		}
		if apPower > 0 {
			return c_base.EPcsStatusCharge, nil
		}
		if apPower == 0 {
			return c_base.EPcsStatusStandby, nil
		}
		if apPower < 0 {
			return c_base.EPcsStatusDischarge, nil
		}

	default:
		return c_base.EPcsStatusUnknown, nil
	}

	return c_base.EPcsStatusUnknown, nil

}

func (s *sPcsLnxallPcs) GetGridMode() (c_base.EGridMode, error) {
	c_log.Warningf(s.ctx, "GetGridMode() not support!")
	return c_base.EGridUnknown, nil
}

func (s *sPcsLnxallPcs) SetPower(power int32) error {
	return s.WriteSingleRegister(PTotal, power)
}

func (s *sPcsLnxallPcs) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsLnxallPcs) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsLnxallPcs) GetTargetPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsLnxallPcs) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsLnxallPcs) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsLnxallPcs) GetPower() (float64, error) {
	return s.GetFloat64Value(PTotal)
}

func (s *sPcsLnxallPcs) GetApparentPower() (float64, error) {
	return s.GetFloat64Value(STotal)
}

func (s *sPcsLnxallPcs) GetReactivePower() (float64, error) {
	return s.GetFloat64Value(QTotal)
}

func (s *sPcsLnxallPcs) GetRatedPower() int32 {
	return -1
}

func (s *sPcsLnxallPcs) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsLnxallPcs) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsLnxallPcs) GetTodayIncomingQuantity() (float64, error) {
	return s.GetFloat64Value(AcDayCharge)
}

func (s *sPcsLnxallPcs) GetHistoryIncomingQuantity() (float64, error) {
	return s.GetFloat64Value(AcHistoryCharge)
}

func (s *sPcsLnxallPcs) GetTodayOutgoingQuantity() (float64, error) {
	return s.GetFloat64Value(AcDayDischarge)
}

func (s *sPcsLnxallPcs) GetHistoryOutgoingQuantity() (float64, error) {
	return s.GetFloat64Value(AcHistoryDischarge)
}
