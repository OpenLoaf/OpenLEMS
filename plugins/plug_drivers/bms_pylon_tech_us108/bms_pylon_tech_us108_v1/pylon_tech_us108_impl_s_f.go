package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_log"
	"common/c_modbus"
	"common/c_util"
	"context"
	"fmt"
	"math"
	"time"
)

type sBmsPylonTechUs108 struct {
	ctx context.Context
	c_modbus.IModbusProtocol
	*c_base.SDriverDescription
	bmsConfig *PylonTechUs108BmsConfig
}

var _ c_device.IBms = (*sBmsPylonTechUs108)(nil)

func (p *sBmsPylonTechUs108) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceBms
}

func (p *sBmsPylonTechUs108) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	p.IModbusProtocol = protocol.(c_modbus.IModbusProtocol)

	bmsConfig := &PylonTechUs108BmsConfig{}
	err := deviceConfig.ScanParams(bmsConfig)
	if err != nil {
		panic(fmt.Errorf("BMS配置解析失败：内容:%v 原因: %s", deviceConfig.Params, err.Error()))
	}

	// 注册
	p.IModbusProtocol.RegisterRead(p.ctx, GroupHeart, GroupInfo, GroupTime, GroupStatistics)

	if bmsConfig.SyncTime {
		p.writeTime()
		c_log.Infof(p.ctx, "syncTime配置为：true！同步时间已开启！")
	} else {
		c_log.Infof(p.ctx, "syncTime配置为：false！时间不同步！")
	}
}

func (p *sBmsPylonTechUs108) Shutdown() {

}

func (p *sBmsPylonTechUs108) GetRatedPower() int32 {
	return p.bmsConfig.RatedPower
}

func (p *sBmsPylonTechUs108) GetMaxInputPower() (float32, error) {
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
	c_log.Debugf(p.ctx, "最大充电 电压：%f, 电流：%f, 功率：%f", values[0], values[1], power)

	if p.bmsConfig.MaxInputPower != 0 && power < float32(p.bmsConfig.MaxInputPower) {
		return float32(p.bmsConfig.MaxInputPower), nil
	}
	return power, nil
}

func (p *sBmsPylonTechUs108) GetMaxOutputPower() (float32, error) {
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
	c_log.Debugf(p.ctx, "最大放电 电压：%f, 电流：%f, 功率：%f, 配置功率：%f", values[0], values[1], power, p.bmsConfig.MaxOutputPower)
	return power, nil
}

func (p *sBmsPylonTechUs108) SetBmsStatus(status c_device.EBmsStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *sBmsPylonTechUs108) GetBmsStatus() (c_device.EBmsStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sBmsPylonTechUs108) GetSoc() (float32, error) {
	return p.GetFloat32Value(SOC)
}

func (p *sBmsPylonTechUs108) GetSoh() (float32, error) {
	return p.GetFloat32Value(SOH)
}

func (p *sBmsPylonTechUs108) GetDcPower() (float64, error) {
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

func (p *sBmsPylonTechUs108) GetDcVoltage() (float64, error) {
	return p.GetFloat64Value(DCVoltage)
}

func (p *sBmsPylonTechUs108) GetDcCurrent() (float64, error) {
	return p.GetFloat64Value(DCCurrent)
}

func (p *sBmsPylonTechUs108) GetCellMinTemp() (float32, error) {
	return p.GetFloat32Value(BatteryCellMinTemp)
}

func (p *sBmsPylonTechUs108) GetCellMaxTemp() (float32, error) {
	return p.GetFloat32Value(BatteryCellMaxTemp)
}

func (p *sBmsPylonTechUs108) GetCellAvgTemp() (float32, error) {
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

func (p *sBmsPylonTechUs108) GetCellMinVoltage() (float32, error) {
	return p.GetFloat32Value(BatteryCellMinVoltage)
}

func (p *sBmsPylonTechUs108) GetCellMaxVoltage() (float32, error) {
	return p.GetFloat32Value(BatteryCellMaxVoltage)
}

func (p *sBmsPylonTechUs108) GetCellAvgVoltage() (float32, error) {
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

func (p *sBmsPylonTechUs108) GetCycleCount() (uint, error) {
	return p.GetUintValue(CycleCount)
}

func (p *sBmsPylonTechUs108) GetTodayIncomingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, TodayCharge)
	if err != nil {
		return 0, err
	}
	return c_util.ToFloat64First(read)
}

func (p *sBmsPylonTechUs108) GetTodayOutgoingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, TodayDischarge)
	if err != nil {
		return 0, err
	}
	//return read[0].Float64(), nil
	return c_util.ToFloat64First(read)
}

func (p *sBmsPylonTechUs108) GetHistoryIncomingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, HistoryCharge)
	if err != nil {
		return 0, err
	}
	//return read[0].Float64(), nil
	return c_util.ToFloat64First(read)
}

func (p *sBmsPylonTechUs108) GetHistoryOutgoingQuantity() (float64, error) {
	read, err := p.ReadGroupSync(GroupStatistics, true, HistoryDischarge)
	if err != nil {
		return 0, err
	}
	//return read[0].Float64(), nil
	return c_util.ToFloat64First(read)
}

func (p *sBmsPylonTechUs108) GetCapacity() (uint32, error) {
	return p.bmsConfig.Capacity, nil
}

func (p *sBmsPylonTechUs108) SetReset() error {
	return nil
}

func (p *sBmsPylonTechUs108) writeTime() {
	_ = p._syncTime()
	go func() {
		ticker := time.NewTicker(12 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-p.ctx.Done():
				c_log.Noticef(p.ctx, "writeTime() 关闭!")
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

func (p *sBmsPylonTechUs108) _syncTime() error {
	if !p.IsActivate() {
		return fmt.Errorf("modbus client is not activate")
	}
	now := time.Now()

	err := p.WriteMultipleRegisters(GroupTime, []int64{int64(now.Year() - 2000), int64(now.Month()),
		int64(now.Day()), int64(now.Hour()), int64(now.Minute()), int64(now.Second())})
	if err != nil {
		return err
	}
	c_log.Infof(p.ctx, "同步时间成功！")
	return nil
}
