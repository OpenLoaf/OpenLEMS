package bms_pylon_tech_us108_v1

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

type sPylonTechUs108Bms struct {
	ctx context.Context
	p_modbus.IModbusProtocol
	*c_base.SDescription
	bmsConfig *PylonTechUs108BmsConfig
}

func (p *sPylonTechUs108Bms) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceBms
}

func (p *sPylonTechUs108Bms) Init(client c_base.IProtocol, cfg *c_base.SDriverConfig) {
	p.IModbusProtocol = client.(p_modbus.IModbusProtocol)

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

func (p *sPylonTechUs108Bms) Destroy() {

}

func (p *sPylonTechUs108Bms) GetRatedPower() uint32 {
	return p.bmsConfig.RatedPower
}

func (p *sPylonTechUs108Bms) GetMaxInputPower() (float32, error) {
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

func (p *sPylonTechUs108Bms) GetMaxOutputPower() (float32, error) {
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

func (p *sPylonTechUs108Bms) SetBmsStatus(status c_device.EBmsStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *sPylonTechUs108Bms) GetBmsStatus() (c_device.EBmsStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sPylonTechUs108Bms) GetSoc() (float32, error) {
	return p.GetFloat32Value(SOC)
}

func (p *sPylonTechUs108Bms) GetSoh() (float32, error) {
	return p.GetFloat32Value(SOH)
}

func (p *sPylonTechUs108Bms) GetDcPower() (float64, error) {
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

func (p *sPylonTechUs108Bms) GetDcVoltage() (float64, error) {
	return p.GetFloat64Value(DCVoltage)
}

func (p *sPylonTechUs108Bms) GetDcCurrent() (float64, error) {
	return p.GetFloat64Value(DCCurrent)
}

func (p *sPylonTechUs108Bms) GetCellMinTemp() (float32, error) {
	return p.GetFloat32Value(BatteryCellMinTemp)
}

func (p *sPylonTechUs108Bms) GetCellMaxTemp() (float32, error) {
	return p.GetFloat32Value(BatteryCellMaxTemp)
}

func (p *sPylonTechUs108Bms) GetCellAvgTemp() (float32, error) {
	minTemp, err := p.GetCellMinTemp()
	if err != nil {
		return 0, err
	}
	maxTemp, err := p.GetCellMaxTemp()
	if err != nil {
		return 0, err
	}
	return (minTemp + maxTemp) / 2.0, nil
}

func (p *sPylonTechUs108Bms) GetCellMinVoltage() (float32, error) {
	return p.GetFloat32Value(BatteryCellMinVoltage)
}

func (p *sPylonTechUs108Bms) GetCellMaxVoltage() (float32, error) {
	return p.GetFloat32Value(BatteryCellMaxVoltage)
}

func (p *sPylonTechUs108Bms) GetCellAvgVoltage() (float32, error) {
	minVoltage, err := p.GetCellMinVoltage()
	if err != nil {
		return 0, err
	}
	maxVoltage, err := p.GetCellMaxVoltage()
	if err != nil {
		return 0, err
	}
	return (minVoltage + maxVoltage) / 2.0, nil
}

func (p *sPylonTechUs108Bms) GetCycleCount() (uint, error) {
	return p.GetUintValue(CycleCount)
}

func (p *sPylonTechUs108Bms) GetTodayIncomingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, TodayCharge)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (p *sPylonTechUs108Bms) GetTodayOutgoingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, TodayDischarge)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (p *sPylonTechUs108Bms) GetHistoryIncomingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, HistoryCharge)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (p *sPylonTechUs108Bms) GetHistoryOutgoingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, HistoryDischarge)
	if err != nil {
		return 0, err
	}
	return read[0].Float64(), nil
}

func (p *sPylonTechUs108Bms) GetCapacity() (uint32, error) {
	return p.bmsConfig.Capacity, nil
}

func (p *sPylonTechUs108Bms) SetReset() error {
	return nil
}

func (p *sPylonTechUs108Bms) writeTime() {
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

func (p *sPylonTechUs108Bms) _syncTime() error {
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
