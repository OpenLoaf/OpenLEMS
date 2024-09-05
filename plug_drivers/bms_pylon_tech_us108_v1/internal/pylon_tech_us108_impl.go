package internal

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"math"
	"plug_protocol_modbus/p_modbus"
	"time"
)

type PylonTechUs108Bms struct {
	ctx context.Context
	p_modbus.IModbusProtocol
	description *c_base.SDescription
	bmsConfig   *PylonTechUs108BmsConfig
}

func NewPlugin(ctx context.Context) c_device.IBms {
	return &PylonTechUs108Bms{
		ctx: ctx,
	}
}

func (p *PylonTechUs108Bms) GetDescription() *c_base.SDescription {
	return p.description
}

func (p *PylonTechUs108Bms) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceBms
}

func (p *PylonTechUs108Bms) Init(client c_base.IProtocol, cfg *c_base.SDriverConfig) {
	p.IModbusProtocol = client.(p_modbus.IModbusProtocol)
	p.description = &c_base.SDescription{
		Brand:  "Plyon",
		Model:  "TechUs108",
		Remark: "派能108kWh风冷电池MBMS",
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

	bmsConfig := &PylonTechUs108BmsConfig{}
	err := gconv.Scan(cfg.Params, bmsConfig)
	if err != nil {
		panic(gerror.Newf("BMS配置解析失败：%s", err.Error()))
	}

	// 注册
	p.IModbusProtocol.RegisterRead(p.ctx, GroupHeart, GroupInfo, GroupTime, GroupStatistics)

	if bmsConfig.SyncTime {
		p.writeTime()
		g.Log().Infof(p.ctx, "syncTime配置为：true！同步时间已开启！")
	} else {
		g.Log().Infof(p.ctx, "syncTime配置为：false！时间不同步！")
	}
}

func (p *PylonTechUs108Bms) Destroy() {

}

func (p *PylonTechUs108Bms) GetRatedPower() uint32 {
	return p.bmsConfig.RatedPower
}

func (p *PylonTechUs108Bms) GetMaxInputPower() (float32, error) {
	chargeForbiddenMark, err := p.GetBool(ChargeForbiddenMark)
	if err != nil && chargeForbiddenMark {
		// 禁止充电
		return 0, nil
	}
	// 通过电压和电流来计算功率
	values, err := p.GetFloat32Values(PileMaxV, PileMaxI)
	if err != nil {
		return 0, err
	}
	power := values[0] * values[1]
	g.Log().Debugf(p.ctx, "最大充电 电压：%f, 电流：%f, 功率：%f", values[0], values[1], power)

	if p.bmsConfig.MaxInputPower != 0 && power < float32(p.bmsConfig.MaxInputPower) {
		return float32(p.bmsConfig.MaxInputPower), nil
	}
	return power, nil
}

func (p *PylonTechUs108Bms) GetMaxOutputPower() (float32, error) {
	dischargeForbiddenMark, err := p.GetBool(DischargeForbiddenMark)
	if err != nil && dischargeForbiddenMark {
		// 禁止放电
		return 0, nil
	}
	// 通过电压和电流来计算功率
	values, err := p.GetFloat32Values(PileMinV, PileMaxDI)
	if err != nil {
		return 0, nil
	}
	power := values[0] * values[1]
	power = float32(math.Abs(float64(power)))

	if p.bmsConfig.MaxOutputPower != 0 && power < float32(p.bmsConfig.MaxOutputPower) {
		return float32(p.bmsConfig.MaxOutputPower), nil
	}
	g.Log().Debugf(p.ctx, "最大放电 电压：%f, 电流：%f, 功率：%f, 配置功率：%f", values[0], values[1], power, p.bmsConfig.MaxOutputPower)
	return power, nil
}

func (p *PylonTechUs108Bms) SetBmsStatus(status c_device.EBmsStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *PylonTechUs108Bms) GetBmsStatus() (c_device.EBmsStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonTechUs108Bms) GetSoc() (float32, error) {
	return p.GetFloat32Value(SOC)
}

func (p *PylonTechUs108Bms) GetSoh() (float32, error) {
	return p.GetFloat32Value(SOH)
}

func (p *PylonTechUs108Bms) GetDcPower() (float64, error) {
	current, err := p.GetDcCurrent()
	if err != nil {
		return 0, err
	}
	voltage, err := p.GetDcVoltage()
	if err != nil {
		return 0, err
	}
	// kW
	return current * voltage / 1000, nil
}

func (p *PylonTechUs108Bms) GetDcVoltage() (float64, error) {
	return p.GetFloat64Value(DCVoltage)
}

func (p *PylonTechUs108Bms) GetDcCurrent() (float64, error) {
	return p.GetFloat64Value(DCCurrent)
}

// GetCellTemp 电芯最低温度, 电芯最高温度, 电芯平均温度
func (p *PylonTechUs108Bms) GetCellTemp() (float32, float32, float32, error) {
	values, err := p.GetFloat32Values(BatteryCellMinTemp, BatteryCellMaxTemp)
	if err != nil {
		return 0, 0, 0, err
	}
	return values[0], values[1], (values[0] + values[1]) / 2, err
}

func (p *PylonTechUs108Bms) GetCellVoltage() (float32, float32, float32, error) {
	values, err := p.GetFloat32Values(BatteryCellMinVoltage, BatteryCellMaxVoltage)
	if err != nil {
		return 0, 0, 0, err
	}
	return values[0], values[1], (values[0] + values[1]) / 2, err
}

func (p *PylonTechUs108Bms) GetCycleCount() (uint, error) {
	return p.GetUintValue(CycleCount)
}

func (p *PylonTechUs108Bms) GetTodayIncomingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, TodayCharge)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (p *PylonTechUs108Bms) GetTodayOutgoingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, TodayDischarge)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (p *PylonTechUs108Bms) GetHistoryIncomingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, HistoryCharge)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (p *PylonTechUs108Bms) GetHistoryOutgoingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, HistoryDischarge)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (p *PylonTechUs108Bms) GetCapacity() (uint32, error) {
	return p.bmsConfig.Capacity, nil
}

func (p *PylonTechUs108Bms) SetReset() error {
	return nil
}

func (p *PylonTechUs108Bms) writeTime() {
	_ = p._syncTime()
	go func() {
		ticker := time.NewTicker(12 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-p.ctx.Done():
				g.Log().Noticef(p.ctx, "writeTime() 关闭!")
				return
			case <-ticker.C:
				if !p.IsActivate() {
					continue
				}
				_ = p._syncTime()
			}

		}
	}()
}

func (p *PylonTechUs108Bms) _syncTime() error {
	if !p.IsActivate() {
		return gerror.Newf("modbus client is not activate")
	}
	now := time.Now()

	err := p.WriteMultipleRegisters(GroupTime, []int64{int64(now.Year() - 2000), int64(now.Month()),
		int64(now.Day()), int64(now.Hour()), int64(now.Minute()), int64(now.Second())})
	if err != nil {
		return err
	}
	g.Log().Infof(p.ctx, "同步时间成功！")
	return nil
}
