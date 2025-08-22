package ess_boost_gold_v1

import (
	"common/c_base"
	"common/c_proto"
	"common/c_type"
	"context"
	"fmt"
)

const (
	IdButtonDischarge = "button-discharge" // 放电按钮
	IdButtonCharge    = "button-charge"    // 充电按钮
	IdButtonScram     = "button-scram"     // 急停
	IdLedRunning      = "led-running"      // 运行指示灯
	IdLedFault        = "led-fault"        // 故障指示灯
)

type sEssBoostGoldEss struct {
	c_proto.IModbusProtocol
	*c_base.SDriverDescription
	deviceConfig *c_base.SDeviceConfig
	essConfig    *EssBoostGoldConfig
	ctx          context.Context

	targetPower         int32 // 目标有功功率
	targetReactivePower int32 // 目标无功功率

	pcs c_type.IPcs // 逆变器

}

func (s *sEssBoostGoldEss) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.deviceConfig = deviceConfig
	s.IModbusProtocol = protocol.(c_proto.IModbusProtocol)

	s.essConfig = &EssBoostGoldConfig{}
	err := deviceConfig.ScanParams(s.essConfig)
	if err != nil {
		panic(fmt.Errorf("高特EMS配置解析失败：%s", err.Error()))
	}

	// 注册点位
	s.RegisterReadTask(s.ctx, GroupBasic, GroupController, GroupSetting)

	fmt.Printf("高特EMS初始化完毕！ 配置：%+v\n", s.essConfig)
}

func (s *sEssBoostGoldEss) Shutdown() {
	_ = s.SetPower(0)
	//s.set
}

func (s *sEssBoostGoldEss) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
}

func (s *sEssBoostGoldEss) GetCellMinTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetCellMaxTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetCellAvgTemp() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetCellMinVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetCellMaxVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetCellAvgVoltage() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetSoc() (float32, error) {
	return s.GetFloat32Value(Soc)
}

func (s *sEssBoostGoldEss) GetSoh() (float32, error) {
	return 100.0, nil
}

func (s *sEssBoostGoldEss) GetCapacity() (uint32, error) {
	return s.essConfig.Capacity, nil
}

func (s *sEssBoostGoldEss) GetCycleCount() (uint, error) {
	return 0, nil
}

func (s *sEssBoostGoldEss) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) SetReset() error {
	return s.WriteSingleRegister(CONTROL_ON_OFF, 3)
}

func (s *sEssBoostGoldEss) SetStatus(status c_base.EEnergyStoreStatus) error {
	switch status {
	case c_base.EPcsStatusStandby:
		return s.WriteSingleRegister(CONTROL_ON_OFF, 1)
	case c_base.EPcsStatusOff:
		return s.WriteSingleRegister(CONTROL_ON_OFF, 0)
	default:
		return nil
	}
}

func (s *sEssBoostGoldEss) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetStatus() (c_base.EEnergyStoreStatus, error) {
	status, err := s.GetIntValue(Status)
	if err != nil {
		return c_base.EPcsStatusUnknown, err
	}
	if status == 1 {
		return c_base.EPcsStatusStandby, nil
	}
	if status == 2 {
		power, err := s.GetPower()
		if err != nil {
			return c_base.EPcsStatusUnknown, err
		}
		if power > 0 {
			return c_base.EPcsStatusDischarge, nil
		}
		if power == 0 {
			return c_base.EPcsStatusStandby, nil
		}
		if power < 0 {
			return c_base.EPcsStatusCharge, nil
		}
	}

	fmt.Printf("获取逆变器状态：%d\n", status)
	return c_base.EPcsStatusUnknown, nil
}

func (s *sEssBoostGoldEss) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) SetPower(power int32) error {
	return s.WriteSingleRegister(Set_Ap_Power, power)
}

func (s *sEssBoostGoldEss) SetReactivePower(power int32) error {
	return s.WriteSingleRegister(Set_Rp_Power, power)
}

func (s *sEssBoostGoldEss) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetTargetPower() int32 {
	if s.essConfig.UsePcsData {
		return s.pcs.GetTargetPower()
	}
	power, err := s.GetFloat32Value(Set_Ap_Power)
	if err != nil {
		return 0
	}
	return int32(power)
}

func (s *sEssBoostGoldEss) GetTargetReactivePower() int32 {
	if s.essConfig.UsePcsData {
		return s.pcs.GetTargetReactivePower()
	}
	power, err := s.GetFloat32Value(Set_Rp_Power)
	if err != nil {
		return 0
	}
	return int32(power)
}

func (s *sEssBoostGoldEss) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetPower() (float64, error) {
	if s.essConfig.UsePcsData {
		return s.pcs.GetPower()
	}
	return s.GetFloat64Value(Set_Ap_Power)
}

func (s *sEssBoostGoldEss) GetApparentPower() (float64, error) {
	if s.essConfig.UsePcsData {
		return s.pcs.GetApparentPower()
	}
	return s.GetFloat64Value(Set_Ap_Power)
}

func (s *sEssBoostGoldEss) GetReactivePower() (float64, error) {
	if s.essConfig.UsePcsData {
		return s.pcs.GetReactivePower()
	}
	return s.GetFloat64Value(Set_Rp_Power)
}

func (s *sEssBoostGoldEss) GetRatedPower() int32 {
	return s.essConfig.RatedPower
}

func (s *sEssBoostGoldEss) GetMaxInputPower() (float32, error) {
	// TODO 改成动态
	return 100, nil
}

func (s *sEssBoostGoldEss) GetMaxOutputPower() (float32, error) {
	// TODO 改成动态
	return 100, nil
}

func (s *sEssBoostGoldEss) GetTodayIncomingQuantity() (float64, error) {
	if s.essConfig.UsePcsData {
		return s.pcs.GetTodayIncomingQuantity()
	}
	return 0, nil
}

func (s *sEssBoostGoldEss) GetHistoryIncomingQuantity() (float64, error) {
	if s.essConfig.UsePcsData {
		return s.pcs.GetHistoryIncomingQuantity()
	}
	return 0, nil
}

func (s *sEssBoostGoldEss) GetTodayOutgoingQuantity() (float64, error) {
	if s.essConfig.UsePcsData {
		return s.pcs.GetTodayOutgoingQuantity()
	}
	return 0, nil
}

func (s *sEssBoostGoldEss) GetHistoryOutgoingQuantity() (float64, error) {
	if s.essConfig.UsePcsData {
		return s.pcs.GetHistoryOutgoingQuantity()
	}
	return 0, nil
}

func (s *sEssBoostGoldEss) GetFireEnvTemperature() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) GetCarbonMonoxideConcentration() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostGoldEss) HasSmoke() (bool, error) {
	//TODO implement me
	panic("implement me")
}
