package ess_pylon_checkwatt_v1

import (
	"common"
	"common/c_base"
	"common/c_device"
	"common/c_error"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

const (
	IdButtonDischarge = "button-discharge" // 放电按钮
	IdButtonCharge    = "button-charge"    // 充电按钮
)

type sPylonCheckwattEss struct {
	*c_base.SAlarmHandler
	*c_base.SDescription
	deviceConfig *c_base.SDriverConfig
	ctx          context.Context

	unitId  uint8             // modbus转发的id
	ammeter c_device.IAmmeter // 电表
	bms     c_device.IBms     // 电池
	pcs     c_device.IPcs     // 逆变器

	buttonScram     c_device.IGpio
	buttonDischarge c_device.IGpio
	buttonCharge    c_device.IGpio
	ledRunning      c_device.IGpio
	ledFault        c_device.IGpio
}

func (p *sPylonCheckwattEss) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	p.deviceConfig = deviceConfig

	// 从配置中获取电表、PCS、BMS的配置
	for _, child := range deviceConfig.DeviceChildren {
		dv := common.GetDeviceById(child.Id)
		if dv == nil {
			panic(gerror.Newf("sPylonCheckwattEss 设备Id: %s 不存在！", child.Id))
		}
		if dv.GetDriverType() == c_base.EDeviceAmmeter {
			p.ammeter = dv.(c_device.IAmmeter)
			g.Log().Infof(p.ctx, "sPylonCheckwattEss 电表初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDeviceBms {
			p.bms = dv.(c_device.IBms)
			g.Log().Infof(p.ctx, "sPylonCheckwattEss 电池初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDevicePcs {
			p.pcs = dv.(c_device.IPcs)
			g.Log().Infof(p.ctx, "sPylonCheckwattEss 逆变器初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDeviceGpio {
			if dv.GetDeviceConfig().Id == IdButtonDischarge {
				p.buttonDischarge = dv.(c_device.IGpio)
				g.Log().Infof(p.ctx, "sPylonCheckwattEss 放电按钮初始化完毕!")
			}
			if dv.GetDeviceConfig().Id == IdButtonCharge {
				p.buttonCharge = dv.(c_device.IGpio)
				g.Log().Infof(p.ctx, "sPylonCheckwattEss 充电按钮初始化完毕!")
			}
		}
	}

	g.Log().Infof(p.ctx, "sPylonCheckwattEss 虚拟储能柜初始化完毕!")
}

func (p *sPylonCheckwattEss) Destroy() {

}

func (p *sPylonCheckwattEss) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
}

func (p *sPylonCheckwattEss) GetDeviceConfig() *c_base.SDriverConfig {
	return p.deviceConfig
}

func (p *sPylonCheckwattEss) GetMetaValueList() []*c_base.MetaValueWrapper {
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

func (p *sPylonCheckwattEss) GetLastUpdateTime() *time.Time {
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

func (p *sPylonCheckwattEss) SetReset() error {
	return c_error.NonSupportError
}

func (p *sPylonCheckwattEss) SetBmsStatus(status c_device.EBmsStatus) error {
	return c_error.NonSupportError
}

func (p *sPylonCheckwattEss) GetBmsStatus() (c_device.EBmsStatus, error) {
	return p.bms.GetBmsStatus()
}

func (p *sPylonCheckwattEss) GetSoc() (float32, error) {
	return p.bms.GetSoc()
}

func (p *sPylonCheckwattEss) GetSoh() (float32, error) {
	return p.bms.GetSoh()
}

func (p *sPylonCheckwattEss) GetCapacity() (uint32, error) {
	return p.bms.GetCapacity()
}

func (p *sPylonCheckwattEss) GetCycleCount() (uint, error) {
	return p.bms.GetCycleCount()
}

func (p *sPylonCheckwattEss) GetRatedPower() uint32 {
	return p.pcs.GetRatedPower()
}

func (p *sPylonCheckwattEss) GetMaxInputPower() (float32, error) {
	return p.pcs.GetMaxInputPower()
}

func (p *sPylonCheckwattEss) GetMaxOutputPower() (float32, error) {
	return p.pcs.GetMaxOutputPower()
}

func (p *sPylonCheckwattEss) GetDcPower() (float64, error) {
	return p.bms.GetDcPower()
}

func (p *sPylonCheckwattEss) GetDcVoltage() (float64, error) {
	return p.bms.GetDcVoltage()
}

func (p *sPylonCheckwattEss) GetDcCurrent() (float64, error) {
	return p.bms.GetDcCurrent()
}

func (p *sPylonCheckwattEss) GetCellMinTemp() (float32, error) {
	return p.bms.GetCellMinTemp()
}

func (p *sPylonCheckwattEss) GetCellMaxTemp() (float32, error) {
	return p.bms.GetCellMaxTemp()
}

func (p *sPylonCheckwattEss) GetCellAvgTemp() (float32, error) {
	return p.bms.GetCellAvgTemp()
}

func (p *sPylonCheckwattEss) GetCellMinVoltage() (float32, error) {
	return p.bms.GetCellMinVoltage()
}

func (p *sPylonCheckwattEss) GetCellMaxVoltage() (float32, error) {
	return p.bms.GetCellMaxVoltage()
}

func (p *sPylonCheckwattEss) GetCellAvgVoltage() (float32, error) {
	return p.bms.GetCellAvgVoltage()
}

func (p *sPylonCheckwattEss) GetTodayIncomingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetTodayIncomingQuantity()
	}
	return p.pcs.GetTodayIncomingQuantity()
}

func (p *sPylonCheckwattEss) GetHistoryIncomingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetHistoryIncomingQuantity()
	}
	return p.pcs.GetHistoryIncomingQuantity()
}

func (p *sPylonCheckwattEss) GetTodayOutgoingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetTodayOutgoingQuantity()
	}
	return p.pcs.GetTodayOutgoingQuantity()
}

func (p *sPylonCheckwattEss) GetHistoryOutgoingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetHistoryOutgoingQuantity()
	}
	return p.pcs.GetHistoryOutgoingQuantity()
}

func (p *sPylonCheckwattEss) SetStatus(status c_base.EEnergyStoreStatus) error {
	if status == c_base.EPcsStatusUnknown || status == c_base.EPcsStatusSync || status == c_base.EPcsStatusFault {
		return c_error.ErrorParam
	}

	return p.pcs.SetStatus(status)
}

func (p *sPylonCheckwattEss) SetGridMode(mode c_base.EGridMode) error {
	return p.pcs.SetGridMode(mode)
}

func (p *sPylonCheckwattEss) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return p.pcs.GetStatus()
}

func (p *sPylonCheckwattEss) GetGridMode() (c_base.EGridMode, error) {
	return p.pcs.GetGridMode()
}

func (p *sPylonCheckwattEss) SetPower(power int32) error {
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

func (p *sPylonCheckwattEss) SetReactivePower(power int32) error {
	return p.pcs.SetReactivePower(power)
}

func (p *sPylonCheckwattEss) SetPowerFactor(factor float32) error {
	return p.pcs.SetPowerFactor(factor)
}

func (p *sPylonCheckwattEss) GetTargetPower() int32 {
	return p.pcs.GetTargetPower()
}

func (p *sPylonCheckwattEss) GetTargetReactivePower() int32 {
	return p.pcs.GetTargetReactivePower()
}

func (p *sPylonCheckwattEss) GetTargetPowerFactor() float32 {
	return p.pcs.GetTargetPowerFactor()
}

func (p *sPylonCheckwattEss) GetPower() (float64, error) {
	return p.pcs.GetPower()
}

func (p *sPylonCheckwattEss) GetApparentPower() (float64, error) {
	return p.pcs.GetApparentPower()
}

func (p *sPylonCheckwattEss) GetReactivePower() (float64, error) {
	return p.pcs.GetReactivePower()
}

func (p *sPylonCheckwattEss) GetFireEnvTemperature() (float64, error) {
	return -1, c_error.NonSupportError
}

func (p *sPylonCheckwattEss) GetCarbonMonoxideConcentration() (float64, error) {
	return -1, c_error.NonSupportError
}

func (p *sPylonCheckwattEss) HasSmoke() (bool, error) {
	return false, c_error.NonSupportError
}
