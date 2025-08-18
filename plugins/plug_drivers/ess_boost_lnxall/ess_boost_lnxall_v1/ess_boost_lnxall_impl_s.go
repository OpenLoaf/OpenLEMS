package ess_boost_lnxall_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_error"
	"common/c_modbus"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

const (
	IdButtonDischarge = "button-discharge" // 放电按钮
	IdButtonCharge    = "button-charge"    // 充电按钮
	IdButtonScram     = "button-scram"     // 急停
	IdLedRunning      = "led-running"      // 运行指示灯
	IdLedFault        = "led-fault"        // 故障指示灯
)

type sEssBoostLnxallEss struct {
	c_modbus.IModbusProtocol
	*c_base.SDriverDescription
	deviceConfig *c_base.SDeviceConfig
	ctx          context.Context

	targetPower         int32 // 目标有功功率
	targetReactivePower int32 // 目标无功功率

	ammeter c_device.IAmmeter // 电表
	bms     c_device.IBms     // 电池
	pcs     c_device.IPcs     // 逆变器

	buttonScram     c_device.IGpio
	buttonDischarge c_device.IGpio
	buttonCharge    c_device.IGpio
	ledRunning      c_device.IGpio
	ledFault        c_device.IGpio
}

// 必须实现储能柜接口
var _ c_device.IEnergyStore = (*sEssBoostLnxallEss)(nil)

func (s *sEssBoostLnxallEss) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	s.IModbusProtocol = protocol.(c_modbus.IModbusProtocol)
	s.deviceConfig = deviceConfig
	// 从配置中获取电表、PCS、BMS的配置
	// 从配置中获取电表、PCS、BMS的配置
	for _, dv := range childDevice {
		if dv.GetDriverType() == c_base.EDeviceAmmeter {
			s.ammeter = dv.(c_device.IAmmeter)
			g.Log().Infof(s.ctx, "EssBoostLnxallEss 电表初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDeviceBms {
			s.bms = dv.(c_device.IBms)
			g.Log().Infof(s.ctx, "EssBoostLnxallEss 电池初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDevicePcs {
			s.pcs = dv.(c_device.IPcs)
			g.Log().Infof(s.ctx, "EssBoostLnxallEss 逆变器初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDeviceGpio {
			if dv.GetDeviceConfig().Id == IdButtonDischarge {
				s.buttonDischarge = dv.(c_device.IGpio)
				g.Log().Infof(s.ctx, "EssBoostLnxallEss 放电按钮初始化完毕!")
			}
			if dv.GetDeviceConfig().Id == IdButtonCharge {
				s.buttonCharge = dv.(c_device.IGpio)
				g.Log().Infof(s.ctx, "EssBoostLnxallEss 充电按钮初始化完毕!")
			}
			if dv.GetDeviceConfig().Id == IdButtonScram {
				s.buttonScram = dv.(c_device.IGpio)
				g.Log().Infof(s.ctx, "EssBoostLnxallEss 急停按钮初始化完毕!")
			}
			if dv.GetDeviceConfig().Id == IdLedRunning {
				s.ledRunning = dv.(c_device.IGpio)
				g.Log().Infof(s.ctx, "EssBoostLnxallEss 运行指示灯初始化完毕!")
			}
			if dv.GetDeviceConfig().Id == IdLedFault {
				s.ledFault = dv.(c_device.IGpio)
				g.Log().Infof(s.ctx, "EssBoostLnxallEss 故障指示灯初始化完毕!")
			}
		}
	}

	//s.Register

	s.RegisterRead(s.ctx,
		GroupController, GroupDetail,
	)

	g.Log().Infof(s.ctx, "EssBoostLnxallEss 储能柜初始化完毕!")
}

func (s *sEssBoostLnxallEss) Shutdown() {
	err := s.SetPower(0)
	if err != nil {
		g.Log().Errorf(s.ctx, "关闭储能柜失败! %v", err)
		return
	}
	// PCS停机
	err = s.SetPower(0)
	if err != nil {
		g.Log().Errorf(s.ctx, "关闭储能柜失败! 下发PCS零功率失败！%v", err)
		return
	}
}

func (s *sEssBoostLnxallEss) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
}

func (s *sEssBoostLnxallEss) GetLastUpdateTime() *time.Time {
	var lastUpdateTime *time.Time
	if s.ammeter != nil {
		lastUpdateTime = s.ammeter.GetLastUpdateTime()
	}
	if s.bms != nil {
		if lastUpdateTime == nil {
			lastUpdateTime = s.bms.GetLastUpdateTime()
		} else {
			if bmsTime := s.bms.GetLastUpdateTime(); bmsTime != nil && bmsTime.After(*lastUpdateTime) {
				lastUpdateTime = bmsTime
			}
		}
	}
	if s.pcs != nil {
		if lastUpdateTime == nil {
			lastUpdateTime = s.pcs.GetLastUpdateTime()
		} else {
			if pcsTime := s.pcs.GetLastUpdateTime(); pcsTime != nil && pcsTime.After(*lastUpdateTime) {
				lastUpdateTime = pcsTime
			}
		}
	}
	return lastUpdateTime
}

func (s *sEssBoostLnxallEss) GetMetaValueList() []*c_base.MetaValueWrapper {
	// 把电表、PCS、BMS、GPIO都所有的值都返回
	//var metaValueList []*c_base.MetaValueWrapper
	//if s.ammeter != nil {
	//	metaValueList = append(metaValueList, s.ammeter.GetMetaValueList()...)
	//}
	//if s.pcs != nil {
	//	metaValueList = append(metaValueList, s.pcs.GetMetaValueList()...)
	//}
	//if s.bms != nil {
	//	metaValueList = append(metaValueList, s.bms.GetMetaValueList()...)
	//}
	//if s.buttonScram != nil {
	//	metaValueList = append(metaValueList, s.buttonScram.GetMetaValueList()...)
	//}
	//if s.buttonCharge != nil {
	//	metaValueList = append(metaValueList, s.buttonCharge.GetMetaValueList()...)
	//}
	//if s.buttonDischarge != nil {
	//	metaValueList = append(metaValueList, s.buttonDischarge.GetMetaValueList()...)
	//}
	//if s.ledRunning != nil {
	//	metaValueList = append(metaValueList, s.ledRunning.GetMetaValueList()...)
	//}
	//if s.ledFault != nil {
	//	metaValueList = append(metaValueList, s.ledFault.GetMetaValueList()...)
	//}
	return GetMetaValueList(s.ammeter, s.pcs, s.bms, s.buttonCharge, s.buttonDischarge, s.buttonScram, s.ledRunning, s.ledFault)
}

func GetMetaValueList(driver ...c_base.IDevice) []*c_base.MetaValueWrapper {
	var metaValueList []*c_base.MetaValueWrapper
	if driver == nil {
		return nil
	}
	for _, d := range driver {
		if d == nil {
			continue
		}
		metaValueList = append(metaValueList, d.GetMetaValueList()...)
	}
	return metaValueList
}

func (s *sEssBoostLnxallEss) GetCellMinTemp() (float32, error) {
	return s.bms.GetCellMinTemp()
}

func (s *sEssBoostLnxallEss) GetCellMaxTemp() (float32, error) {
	return s.bms.GetCellMaxTemp()
}

func (s *sEssBoostLnxallEss) GetCellAvgTemp() (float32, error) {
	return s.bms.GetCellAvgTemp()
}

func (s *sEssBoostLnxallEss) GetCellMinVoltage() (float32, error) {
	return s.bms.GetCellMinVoltage()
}

func (s *sEssBoostLnxallEss) GetCellMaxVoltage() (float32, error) {
	return s.bms.GetCellMaxVoltage()
}

func (s *sEssBoostLnxallEss) GetCellAvgVoltage() (float32, error) {
	return s.bms.GetCellAvgVoltage()
}

func (s *sEssBoostLnxallEss) GetSoc() (float32, error) {
	return s.bms.GetSoc()
}

func (s *sEssBoostLnxallEss) GetSoh() (float32, error) {
	return s.bms.GetSoh()
}

func (s *sEssBoostLnxallEss) GetCapacity() (uint32, error) {
	return s.GetUint32Value(ESS_RATED_CAPACITY)
}

func (s *sEssBoostLnxallEss) GetCycleCount() (uint, error) {
	return s.bms.GetCycleCount()
}

func (s *sEssBoostLnxallEss) GetDcPower() (float64, error) {
	return s.bms.GetDcPower()
}

func (s *sEssBoostLnxallEss) SetReset() error {
	return s.bms.SetReset()
}

func (s *sEssBoostLnxallEss) SetStatus(status c_base.EEnergyStoreStatus) error {
	g.Log().Noticef(s.ctx, "SetStatus(%d)", status)
	switch status {
	case c_base.EPcsStatusStandby:
		return s.WriteSingleRegister(ESS_ON_OFF, 1)
	case c_base.EPcsStatusOff:
		return s.WriteSingleRegister(ESS_ON_OFF, 0)
	default:
		return c_error.NonSupportError
	}

}

func (s *sEssBoostLnxallEss) SetGridMode(mode c_base.EGridMode) error {
	return c_error.NonSupportError
}

func (s *sEssBoostLnxallEss) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return s.pcs.GetStatus()
}

func (s *sEssBoostLnxallEss) GetGridMode() (c_base.EGridMode, error) {
	return s.pcs.GetGridMode()
}

func (s *sEssBoostLnxallEss) SetPower(power int32) error {
	s.targetPower = power
	g.Log().Infof(s.ctx, "SetPower(%d)", power)
	return s.WriteSingleRegister(ESS_SET_AP_POWER, power)
	//return s.pcs.SetPower(power)
}

func (s *sEssBoostLnxallEss) SetReactivePower(power int32) error {
	s.targetReactivePower = power
	return s.WriteSingleRegister(ESS_SET_RP_POWER, power)
}

func (s *sEssBoostLnxallEss) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetTargetPower() int32 {
	value, err := s.GetInt32Value(ESS_SET_AP_POWER)
	if err != nil {
		return 0
	}
	return value
}

func (s *sEssBoostLnxallEss) GetTargetReactivePower() int32 {
	value, err := s.GetInt32Value(ESS_SET_RP_POWER)
	if err != nil {
		return 0
	}
	return value
}

func (s *sEssBoostLnxallEss) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sEssBoostLnxallEss) GetPower() (float64, error) {
	return s.pcs.GetPower()
}

func (s *sEssBoostLnxallEss) GetApparentPower() (float64, error) {
	return s.pcs.GetApparentPower()
}

func (s *sEssBoostLnxallEss) GetReactivePower() (float64, error) {
	return s.pcs.GetReactivePower()
}

func (s *sEssBoostLnxallEss) GetRatedPower() int32 {
	//value, err := s.GetInt32Value(ESS_RATED_POWER)
	//if err != nil {
	//	return 0
	//}
	return 100
}

func (s *sEssBoostLnxallEss) GetMaxInputPower() (float32, error) {
	//return s.GetFloat32Value(ESS_MAX_CHARGE_POWER)
	return 100.0, nil
}

func (s *sEssBoostLnxallEss) GetMaxOutputPower() (float32, error) {
	//return s.GetFloat32Value(ESS_MAX_DISCHARGE_POWER)
	return 100.0, nil
}

func (s *sEssBoostLnxallEss) GetTodayIncomingQuantity() (float64, error) {
	return s.pcs.GetTodayIncomingQuantity()
}

func (s *sEssBoostLnxallEss) GetHistoryIncomingQuantity() (float64, error) {
	return s.pcs.GetHistoryIncomingQuantity()
}

func (s *sEssBoostLnxallEss) GetTodayOutgoingQuantity() (float64, error) {
	return s.pcs.GetTodayOutgoingQuantity()
}

func (s *sEssBoostLnxallEss) GetHistoryOutgoingQuantity() (float64, error) {
	return s.pcs.GetHistoryOutgoingQuantity()
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
