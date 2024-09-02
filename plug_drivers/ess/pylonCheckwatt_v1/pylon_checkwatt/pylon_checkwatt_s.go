package pylon_checkwatt

import (
	"context"
	common "ems-plan"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"ems-plan/c_error"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

const (
	IdButtonDischarge = "button-discharge" // 放电按钮
	IdButtonCharge    = "button-charge"    // 充电按钮
)

type PylonCheckwattEss struct {
	*c_base.SAlarmHandler
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

func NewPlugin(ctx context.Context) c_device.IEnergyStore {
	return &PylonCheckwattEss{ctx: ctx, SAlarmHandler: &c_base.SAlarmHandler{}}
}

func (p *PylonCheckwattEss) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	p.deviceConfig = deviceConfig
	deviceConfig.IsVirtual = true

	// 从配置中获取电表、PCS、BMS的配置
	for _, child := range deviceConfig.DeviceChildren {
		dv := common.DeviceInstance.FindById(child.Id)
		if dv == nil {
			panic(fmt.Sprintf("设备Id: %s 不存在！", child.Id))
		}
		if dv.GetDriverType() == c_base.EDeviceAmmeter {
			p.ammeter = dv.(c_device.IAmmeter)
			g.Log().Infof(p.ctx, "PylonCheckwattEss 电表初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDeviceBms {
			p.bms = dv.(c_device.IBms)
			g.Log().Infof(p.ctx, "PylonCheckwattEss 电池初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDevicePcs {
			p.pcs = dv.(c_device.IPcs)
			g.Log().Infof(p.ctx, "PylonCheckwattEss 逆变器初始化完毕!")
		}
		if dv.GetDriverType() == c_base.EDeviceGpio {
			if dv.GetDeviceConfig().Id == IdButtonDischarge {
				p.buttonDischarge = dv.(c_device.IGpio)
				g.Log().Infof(p.ctx, "PylonCheckwattEss 放电按钮初始化完毕!")
			}
			if dv.GetDeviceConfig().Id == IdButtonCharge {
				p.buttonCharge = dv.(c_device.IGpio)
				g.Log().Infof(p.ctx, "PylonCheckwattEss 充电按钮初始化完毕!")
			}
		}
	}

	g.Log().Infof(p.ctx, "PylonCheckwattEss 虚拟储能柜初始化完毕!")
}

func (p *PylonCheckwattEss) GetDescription() *c_base.SDescription {
	return &c_base.SDescription{
		Brand:  "Plyon",
		Model:  "Checkwatt",
		Remark: "虚拟派能柜，整合PCS与BMS",
		Telemetry: []*c_base.STelemetry{
			{Name: "soc", Unit: "%", Remark: "SOC"},
			{Name: "power", Unit: "kW", Remark: "功率"},
			{Name: "apparentPower", Unit: "kVA", Remark: "视在功率"},
			{Name: "reactivePower", Unit: "kVar", Remark: "无功功率"},
			{Name: "todayIncomingQuantity", I18nKey: "essTodayCharge", Unit: "kWh", Remark: "当日充电量"},
			{Name: "todayOutgoingQuantity", I18nKey: "essTodayDischarge", Unit: "kWh", Remark: "当日放电量"},
			{Name: "historyIncomingQuantity", I18nKey: "essHistoryCharge", Unit: "kWh", Remark: "历史充电量"},
			{Name: "historyOutgoingQuantity", I18nKey: "essHistoryDischarge", Unit: "kWh", Remark: "历史放电量"},
		},
	}
}

func (p *PylonCheckwattEss) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
}

func (p *PylonCheckwattEss) GetDeviceConfig() *c_base.SDriverConfig {
	return p.deviceConfig
}

func (p *PylonCheckwattEss) GetFunctionList() []*c_base.STelemetry {
	return []*c_base.STelemetry{
		{Name: "soc", Unit: "%", Remark: "SOC"},
		{Name: "power", Unit: "kW", Remark: "功率"},
		{Name: "apparentPower", Unit: "kVA", Remark: "视在功率"},
		{Name: "reactivePower", Unit: "kVar", Remark: "无功功率"},
		{Name: "todayIncomingQuantity", I18nKey: "essTodayCharge", Unit: "kWh", Remark: "当日充电量"},
		{Name: "todayOutgoingQuantity", I18nKey: "essTodayDischarge", Unit: "kWh", Remark: "当日放电量"},
		{Name: "historyIncomingQuantity", I18nKey: "essHistoryCharge", Unit: "kWh", Remark: "历史充电量"},
		{Name: "historyOutgoingQuantity", I18nKey: "essHistoryDischarge", Unit: "kWh", Remark: "历史放电量"},
	}
}

func (p *PylonCheckwattEss) GetMetaValueList() []*c_base.MetaValueWrapper {
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

func (p *PylonCheckwattEss) GetLastUpdateTime() *time.Time {
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

func (p *PylonCheckwattEss) SetReset() error {
	return c_error.NonSupportError
}

func (p *PylonCheckwattEss) SetBmsStatus(status c_device.EBmsStatus) error {
	return c_error.NonSupportError
}

func (p *PylonCheckwattEss) GetBmsStatus() (c_device.EBmsStatus, error) {
	return p.bms.GetBmsStatus()
}

func (p *PylonCheckwattEss) GetSoc() (float32, error) {
	return p.bms.GetSoc()
}

func (p *PylonCheckwattEss) GetSoh() (float32, error) {
	return p.bms.GetSoh()
}

func (p *PylonCheckwattEss) GetCellTemp() (float32, float32, float32, error) {
	return p.bms.GetCellTemp()
}

func (p *PylonCheckwattEss) GetCellVoltage() (float32, float32, float32, error) {
	return p.bms.GetCellVoltage()
}

func (p *PylonCheckwattEss) GetCapacity() (uint32, error) {
	return p.bms.GetCapacity()
}

func (p *PylonCheckwattEss) GetCycleCount() (uint, error) {
	return p.bms.GetCycleCount()
}

func (p *PylonCheckwattEss) GetRatedPower() uint32 {
	return p.pcs.GetRatedPower()
}

func (p *PylonCheckwattEss) GetMaxInputPower() (float32, error) {
	return p.pcs.GetMaxInputPower()
}

func (p *PylonCheckwattEss) GetMaxOutputPower() (float32, error) {
	return p.pcs.GetMaxOutputPower()
}

func (p *PylonCheckwattEss) GetDcPower() (float64, error) {
	return p.bms.GetDcPower()
}

func (p *PylonCheckwattEss) GetDcVoltage() (float64, error) {
	return p.bms.GetDcVoltage()
}

func (p *PylonCheckwattEss) GetDcCurrent() (float64, error) {
	return p.bms.GetDcCurrent()
}

func (p *PylonCheckwattEss) GetTodayIncomingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetTodayIncomingQuantity()
	}
	return p.pcs.GetTodayIncomingQuantity()
}

func (p *PylonCheckwattEss) GetHistoryIncomingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetHistoryIncomingQuantity()
	}
	return p.pcs.GetHistoryIncomingQuantity()
}

func (p *PylonCheckwattEss) GetTodayOutgoingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetTodayOutgoingQuantity()
	}
	return p.pcs.GetTodayOutgoingQuantity()
}

func (p *PylonCheckwattEss) GetHistoryOutgoingQuantity() (float64, error) {
	if p.ammeter != nil {
		return p.ammeter.GetHistoryOutgoingQuantity()
	}
	return p.pcs.GetHistoryOutgoingQuantity()
}

func (p *PylonCheckwattEss) SetStatus(status c_base.EEnergyStoreStatus) error {
	return p.pcs.SetStatus(status)
}

func (p *PylonCheckwattEss) SetGridMode(mode c_base.EGridMode) error {
	return p.pcs.SetGridMode(mode)
}

func (p *PylonCheckwattEss) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return p.pcs.GetStatus()
}

func (p *PylonCheckwattEss) GetGridMode() (c_base.EGridMode, error) {
	return p.pcs.GetGridMode()
}

func (p *PylonCheckwattEss) SetPower(power int32) error {
	return p.pcs.SetPower(power)
}

func (p *PylonCheckwattEss) SetReactivePower(power int32) error {
	return p.pcs.SetReactivePower(power)
}

func (p *PylonCheckwattEss) SetPowerFactor(factor float32) error {
	return p.pcs.SetPowerFactor(factor)
}

func (p *PylonCheckwattEss) GetTargetPower() int32 {
	return p.pcs.GetTargetPower()
}

func (p *PylonCheckwattEss) GetTargetReactivePower() int32 {
	return p.pcs.GetTargetReactivePower()
}

func (p *PylonCheckwattEss) GetTargetPowerFactor() float32 {
	return p.pcs.GetTargetPowerFactor()
}

func (p *PylonCheckwattEss) GetPower() (float64, error) {
	return p.pcs.GetPower()
}

func (p *PylonCheckwattEss) GetApparentPower() (float64, error) {
	return p.pcs.GetApparentPower()
}

func (p *PylonCheckwattEss) GetReactivePower() (float64, error) {
	return p.pcs.GetReactivePower()
}

func (p *PylonCheckwattEss) GetFireEnvTemperature() (float64, error) {
	return -1, c_error.NonSupportError
}

func (p *PylonCheckwattEss) GetCarbonMonoxideConcentration() (float64, error) {
	return -1, c_error.NonSupportError
}

func (p *PylonCheckwattEss) HasSmoke() (bool, error) {
	return false, c_error.NonSupportError
}
