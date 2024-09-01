package pylon_checkwatt

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"ems-plan/c_error"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"plug_protocol_gpio_sysfs/p_gpio_sysfs"
	"time"
)

const (
	IdButtonDischarge = "button-discharge" // 放电按钮
	IdButtonCharge    = "button-charge"    // 充电按钮
)

type PylonCheckwattEss struct {
	*c_base.SAlarmHandler
	deviceConfig *c_base.SDriverConfig
	Ctx          context.Context
	CabinetId    uint8 // 属于哪个柜子

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

func NewEss(ctx context.Context, cabinetId uint8, drivers []c_base.IDriver, gpioMap map[string]c_device.IGpio) (*PylonCheckwattEss, error) {

	ess := &PylonCheckwattEss{
		CabinetId: cabinetId,
	}
	ess.Ctx = context.WithValue(ctx, "DeviceName", ess.GetId())

	ess.SAlarmHandler = &c_base.SAlarmHandler{
		Ctx: ess.Ctx,
		AlarmHappened: func(alarm *c_base.SAlarmDetail) {
			g.Log().Warningf(ess.Ctx, "柜子：%s,发生告警：%s", ess.GetId(), alarm.Meta.Cn)
		},
	}

	var (
		pcsCount     int
		bmsCount     int
		ammeterCount int
	)
	for _, driver := range drivers {
		switch driver.GetDriverType() {
		case c_base.EDeviceAmmeter:
			ammeterCount++
			ess.ammeter = driver.(c_device.IAmmeter)
		case c_base.EDevicePcs:
			pcsCount++
			ess.pcs = driver.(c_device.IPcs)
			ess.pcs.RegisterMonitorChan(ess.GetMonitorChan())
		case c_base.EDeviceBms:
			bmsCount++
			ess.bms = driver.(c_device.IBms)
			ess.bms.RegisterMonitorChan(ess.GetMonitorChan())
		}
	}
	if ess.pcs == nil || ess.bms == nil {
		panic(fmt.Sprintf("一个柜子需要电池和PCS组成，但是现在有一项不存在，请检查配置！ PCS数量:%d, BMS数量:%d，电表数量:%d, cabinetId:%d", pcsCount, bmsCount, ammeterCount, cabinetId))
	}

	if pcsCount > 1 || bmsCount > 1 || ammeterCount > 1 {
		panic(fmt.Sprintf("当前柜子ID：%d加载的设备数量不正确！PCS数量:%d, BMS数量:%d，电表数量:%d, 请检查配置！", cabinetId, pcsCount, bmsCount, ammeterCount))
	}

	// 注册输入输出
	if output, exist := gpioMap[p_gpio_sysfs.IdLedRunning]; exist {
		ess.ledRunning = output
		g.Log().Infof(ess.Ctx, "注册LED运行灯成功！")
	}
	if output, exist := gpioMap[p_gpio_sysfs.IdLedFault]; exist {
		ess.ledFault = output
		g.Log().Infof(ess.Ctx, "注册LED故障灯成功！")
	}

	if input, exist := gpioMap[p_gpio_sysfs.IdButtonScram]; exist {
		input.RegisterHandler(func(ctx context.Context, status bool) {
			if status {
				g.Log().Warningf(ess.Ctx, "触发急停！")
				// 紧急停机
				if ess.ledFault != nil {
					_ = ess.ledFault.SetHigh()
					g.Log().Warningf(ess.Ctx, "触发急停！触发故障灯！")
				}
			} else {
				_ = ess.ledFault.SetLow()
				g.Log().Infof(ess.Ctx, "解除急停！故障灯熄灭！")
			}

		})
		ess.buttonScram = input
		g.Log().Infof(ess.Ctx, "注册急停按钮成功！")
	}
	return ess, nil
}

func (p *PylonCheckwattEss) GetDescription() *c_base.SDescription {
	return &c_base.SDescription{
		Brand:  "Plyon",
		Model:  "Checkwatt",
		Remark: "虚拟派能柜，整合PCS与BMS",
	}
}

func (p *PylonCheckwattEss) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
}

func (p *PylonCheckwattEss) GetDeviceConfig() *c_base.SDriverConfig {
	return p.deviceConfig
}

func (p *PylonCheckwattEss) GetFunctionList() []*c_base.SFunction {
	return []*c_base.SFunction{
		{FunctionName: "soc", Unit: "%", Remark: "SOC"},
		{FunctionName: "power", Unit: "kW", Remark: "功率"},
		{FunctionName: "apparentPower", Unit: "kVA", Remark: "视在功率"},
		{FunctionName: "reactivePower", Unit: "kVar", Remark: "无功功率"},
		{FunctionName: "todayIncomingQuantity", FunctionNameI18nOverwrite: "essTodayCharge", Unit: "kWh", Remark: "当日充电量"},
		{FunctionName: "todayOutgoingQuantity", FunctionNameI18nOverwrite: "essTodayDischarge", Unit: "kWh", Remark: "当日放电量"},
		{FunctionName: "historyIncomingQuantity", FunctionNameI18nOverwrite: "essHistoryCharge", Unit: "kWh", Remark: "历史充电量"},
		{FunctionName: "historyOutgoingQuantity", FunctionNameI18nOverwrite: "essHistoryDischarge", Unit: "kWh", Remark: "历史放电量"},
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

func (p *PylonCheckwattEss) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	p.deviceConfig = deviceConfig
	g.Log().Infof(p.Ctx, "PylonCheckwattEss Init!CabinetId:%d", p.CabinetId)
}

func (p *PylonCheckwattEss) GetId() string {
	return fmt.Sprintf("pylonCheckwattEss_%d", p.CabinetId)
}

func (p *PylonCheckwattEss) GetType() c_base.EDeviceType {
	return c_base.EDeviceEnergyStore
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
	return nil
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

func (p *PylonCheckwattEss) GetRatedPower() (int32, error) {
	return p.pcs.GetRatedPower()
}

func (p *PylonCheckwattEss) GetMaxInputPower() (float64, error) {
	return p.pcs.GetMaxInputPower()
}

func (p *PylonCheckwattEss) GetMaxOutputPower() (float64, error) {
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
