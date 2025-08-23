package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"common/c_error"
	"common/c_log"
	"common/c_type"
	"context"
	"time"
)

const (
	IdButtonDischarge = "button-discharge" // 放电按钮
	IdButtonCharge    = "button-charge"    // 充电按钮
)

type sEssPylonCheckwatt struct {
	*c_base.SAlarmHandler
	*c_base.SDriverDescription
	deviceConfig *c_base.SDeviceConfig
	ctx          context.Context

	unitId  uint8           // modbus转发的id
	ammeter c_type.IAmmeter // 电表
	bms     c_type.IBms     // 电池
	pcs     c_type.IPcs     // 逆变器

	buttonScram     c_type.IGpio
	buttonDischarge c_type.IGpio
	buttonCharge    c_type.IGpio
	ledRunning      c_type.IGpio
	ledFault        c_type.IGpio
}

func (p *sEssPylonCheckwatt) ProtocolListen() {

}

func (p *sEssPylonCheckwatt) IsActivate() bool {
	return true
}

func (p *sEssPylonCheckwatt) IsPhysical() bool {
	return false
}

func (p *sEssPylonCheckwatt) InitDevice(deviceConfig *c_base.SDeviceConfig, _ c_base.IProtocol, childDevice []c_base.IDevice) {
	p.deviceConfig = deviceConfig

	// 从配置中获取电表、PCS、BMS的配置
	for _, dv := range childDevice {
		if dv.GetDriverType() == c_base.EDeviceAmmeter {
			p.ammeter = dv.(c_type.IAmmeter)
			c_log.BizInfof(p.ctx, "[%s] 电表初始化完毕!", dv.GetDeviceConfig().Name)
		}
		if dv.GetDriverType() == c_base.EDeviceBms {
			p.bms = dv.(c_type.IBms)
			c_log.BizInfof(p.ctx, "[%s] 电池初始化完毕!", dv.GetDeviceConfig().Name)
		}
		if dv.GetDriverType() == c_base.EDevicePcs {
			p.pcs = dv.(c_type.IPcs)
			c_log.BizInfof(p.ctx, "[%s] 逆变器初始化完毕!", dv.GetDeviceConfig().Name)
		}
		if dv.GetDriverType() == c_base.EDeviceGpio {
			if dv.GetDeviceConfig().Id == IdButtonDischarge {
				p.buttonDischarge = dv.(c_type.IGpio)
				c_log.BizInfof(p.ctx, "[%s] 放电按钮初始化完毕!", dv.GetDeviceConfig().Name)
			}
			if dv.GetDeviceConfig().Id == IdButtonCharge {
				p.buttonCharge = dv.(c_type.IGpio)
				c_log.BizInfof(p.ctx, "[%s] 充电按钮初始化完毕!", dv.GetDeviceConfig().Name)
			}
		}
	}

	c_log.BizInfof(p.ctx, "虚拟储能柜初始化完毕!")
}

func (p *sEssPylonCheckwatt) Shutdown() {
	_ = p.SetPower(0)
	_ = p.SetStatus(c_base.EPcsStatusOff)
}

func (p *sEssPylonCheckwatt) Destroy() {

}

func (p *sEssPylonCheckwatt) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
}

func (p *sEssPylonCheckwatt) GetDeviceConfig() *c_base.SDeviceConfig {
	return p.deviceConfig
}

func (p *sEssPylonCheckwatt) GetMetaValueList() []*c_base.MetaValueWrapper {
	// 把电表、PCS、BMS、GPIO都所有的值都返回
	var metaValueList []*c_base.MetaValueWrapper
	if p.ammeter != nil {
		metaValueList = append(metaValueList, p.ammeter.GetMetaValueList()...)
	}
	if p.pcs != nil {
		metaValueList = append(metaValueList, p.pcs.GetMetaValueList()...)
	}
	if p.bms != nil {
		metaValueList = append(metaValueList, p.bms.GetMetaValueList()...)
	}
	if p.buttonScram != nil {
		metaValueList = append(metaValueList, p.buttonScram.GetMetaValueList()...)
	}
	if p.buttonCharge != nil {
		metaValueList = append(metaValueList, p.buttonCharge.GetMetaValueList()...)
	}
	if p.buttonDischarge != nil {
		metaValueList = append(metaValueList, p.buttonDischarge.GetMetaValueList()...)
	}
	if p.ledRunning != nil {
		metaValueList = append(metaValueList, p.ledRunning.GetMetaValueList()...)
	}
	if p.ledFault != nil {
		metaValueList = append(metaValueList, p.ledFault.GetMetaValueList()...)
	}
	return metaValueList
}

func (p *sEssPylonCheckwatt) GetLastUpdateTime() *time.Time {
	var lastUpdateTime *time.Time
	if p.ammeter != nil {
		lastUpdateTime = p.ammeter.GetLastUpdateTime()
	}
	if p.bms != nil {
		if lastUpdateTime == nil {
			lastUpdateTime = p.bms.GetLastUpdateTime()
		} else {
			if bmsTime := p.bms.GetLastUpdateTime(); bmsTime != nil && bmsTime.After(*lastUpdateTime) {
				lastUpdateTime = bmsTime
			}
		}
	}
	if p.pcs != nil {
		if lastUpdateTime == nil {
			lastUpdateTime = p.pcs.GetLastUpdateTime()
		} else {
			if pcsTime := p.pcs.GetLastUpdateTime(); pcsTime != nil && pcsTime.After(*lastUpdateTime) {
				lastUpdateTime = pcsTime
			}
		}
	}
	return lastUpdateTime
}

func (p *sEssPylonCheckwatt) SetReset() error {
	return c_error.NonSupportError
}

func (p *sEssPylonCheckwatt) SetBmsStatus(status c_type.EBmsStatus) error {
	return c_error.NonSupportError
}

func (p *sEssPylonCheckwatt) GetBmsStatus() (c_type.EBmsStatus, error) {
	return p.bms.GetBmsStatus()
}

func (p *sEssPylonCheckwatt) GetSoc() (float32, error) {
	return p.bms.GetSoc()
}

func (p *sEssPylonCheckwatt) GetSoh() (float32, error) {
	return p.bms.GetSoh()
}

func (p *sEssPylonCheckwatt) GetCapacity() (uint32, error) {
	return p.bms.GetCapacity()
}

func (p *sEssPylonCheckwatt) GetCycleCount() (uint, error) {
	return p.bms.GetCycleCount()
}

func (p *sEssPylonCheckwatt) GetRatedPower() int32 {
	return p.pcs.GetRatedPower()
}

func (p *sEssPylonCheckwatt) GetMaxInputPower() (float32, error) {
	return p.pcs.GetMaxInputPower()
}

func (p *sEssPylonCheckwatt) GetMaxOutputPower() (float32, error) {
	return p.pcs.GetMaxOutputPower()
}

func (p *sEssPylonCheckwatt) GetDcPower() (float64, error) {
	return p.bms.GetDcPower()
}

func (p *sEssPylonCheckwatt) GetDcVoltage() (float64, error) {
	return p.bms.GetDcVoltage()
}

func (p *sEssPylonCheckwatt) GetDcCurrent() (float64, error) {
	return p.bms.GetDcCurrent()
}

func (p *sEssPylonCheckwatt) GetCellMinTemp() (float32, error) {
	return p.bms.GetCellMinTemp()
}

func (p *sEssPylonCheckwatt) GetCellMaxTemp() (float32, error) {
	return p.bms.GetCellMaxTemp()
}

func (p *sEssPylonCheckwatt) GetCellAvgTemp() (float32, error) {
	return p.bms.GetCellAvgTemp()
}

func (p *sEssPylonCheckwatt) GetCellMinVoltage() (float32, error) {
	return p.bms.GetCellMinVoltage()
}

func (p *sEssPylonCheckwatt) GetCellMaxVoltage() (float32, error) {
	return p.bms.GetCellMaxVoltage()
}

func (p *sEssPylonCheckwatt) GetCellAvgVoltage() (float32, error) {
	return p.bms.GetCellAvgVoltage()
}

func (p *sEssPylonCheckwatt) GetTodayIncomingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetTodayIncomingQuantity()
	}
	return p.pcs.GetTodayIncomingQuantity()
}

func (p *sEssPylonCheckwatt) GetHistoryIncomingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetHistoryIncomingQuantity()
	}
	return p.pcs.GetHistoryIncomingQuantity()
}

func (p *sEssPylonCheckwatt) GetTodayOutgoingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetTodayOutgoingQuantity()
	}
	return p.pcs.GetTodayOutgoingQuantity()
}

func (p *sEssPylonCheckwatt) GetHistoryOutgoingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetHistoryOutgoingQuantity()
	}
	return p.pcs.GetHistoryOutgoingQuantity()
}

func (p *sEssPylonCheckwatt) SetStatus(status c_base.EEnergyStoreStatus) error {
	if status == c_base.EPcsStatusUnknown || status == c_base.EPcsStatusSync || status == c_base.EPcsStatusFault {
		return c_error.ErrorParam
	}

	return p.pcs.SetStatus(status)
}

func (p *sEssPylonCheckwatt) SetGridMode(mode c_base.EGridMode) error {
	return p.pcs.SetGridMode(mode)
}

func (p *sEssPylonCheckwatt) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return p.pcs.GetStatus()
}

func (p *sEssPylonCheckwatt) GetGridMode() (c_base.EGridMode, error) {
	return p.pcs.GetGridMode()
}

func (p *sEssPylonCheckwatt) SetPower(power int32) error {
	// 判断一下防止超限
	if power > 0 {
		maxOutputPower, err := p.GetMaxOutputPower()
		if err != nil {
			return err
		}
		if power > int32(maxOutputPower) {
			return c_error.OverLimitError
		}
	} else {
		maxInputPower, err := p.GetMaxInputPower()
		if err != nil {
			return err
		}
		if power < int32(-maxInputPower) {
			return c_error.OverLimitError
		}
	}
	return p.pcs.SetPower(power)
}

func (p *sEssPylonCheckwatt) SetReactivePower(power int32) error {
	return p.pcs.SetReactivePower(power)
}

func (p *sEssPylonCheckwatt) SetPowerFactor(factor float32) error {
	return p.pcs.SetPowerFactor(factor)
}

func (p *sEssPylonCheckwatt) GetTargetPower() int32 {
	return p.pcs.GetTargetPower()
}

func (p *sEssPylonCheckwatt) GetTargetReactivePower() int32 {
	return p.pcs.GetTargetReactivePower()
}

func (p *sEssPylonCheckwatt) GetTargetPowerFactor() float32 {
	return p.pcs.GetTargetPowerFactor()
}

func (p *sEssPylonCheckwatt) GetPower() (float64, error) {
	return p.pcs.GetPower()
}

func (p *sEssPylonCheckwatt) GetApparentPower() (float64, error) {
	return p.pcs.GetApparentPower()
}

func (p *sEssPylonCheckwatt) GetReactivePower() (float64, error) {
	return p.pcs.GetReactivePower()
}

func (p *sEssPylonCheckwatt) GetFireEnvTemperature() (float64, error) {
	return -1, c_error.NonSupportError
}

func (p *sEssPylonCheckwatt) GetCarbonMonoxideConcentration() (float64, error) {
	return -1, c_error.NonSupportError
}

func (p *sEssPylonCheckwatt) HasSmoke() (bool, error) {
	return false, c_error.NonSupportError
}
