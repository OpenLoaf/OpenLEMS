package pylon_tech_us108

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"plug_protocol_modbus/p_modbus"
	"time"
)

type PylonTechUs108Bms struct {
	c_base.IDriverConfig
	p_modbus.IModbusProtocol
	ctx         context.Context
	log         *glog.Logger
	description c_base.SDescription
}

func (p *PylonTechUs108Bms) GetDescription() c_base.SDescription {
	return c_base.SDescription{
		Brand:  "Plyon",
		Model:  "TechUs108",
		Type:   c_base.EDeviceBms,
		Remark: "派能108kWh风冷电池MBMS",
	}
}

func (p *PylonTechUs108Bms) Init(ctx context.Context, client c_base.IProtocol, cfg any) error {
	log := g.Log()

	p.log = log
	p.ctx = ctx
	p.IModbusProtocol = client.(p_modbus.IModbusProtocol)

	// 注册
	p.IModbusProtocol.RegisterRead(ctx, GroupHeart, GroupInfo, GroupTime, GroupStatistics)

	var (
		config *p_modbus.SModbusDeviceConfig
		ok     bool
	)
	if config, ok = cfg.(*p_modbus.SModbusDeviceConfig); !ok || config == nil {
		panic("配置文件转换失败！请检查配置文件！")
	}
	p.IDriverConfig = config

	p.log.Noticef(ctx, "配置信息:%+v", config)

	/*	if v, ok := configMap["syncTime"]; ok && v == "true" {
			p.writeTime()
			p.log.Infof(ctx, "syncTime配置为：true！同步时间已开启！")
		} else {
			p.log.Infof(ctx, "syncTime配置为：false！时间不同步！")
		}
	*/
	p.IModbusProtocol.Init(p.GetType())
	return nil
}

func (p *PylonTechUs108Bms) GetType() c_base.EDeviceType {
	return c_base.EDeviceBms
}

func (p *PylonTechUs108Bms) HasAlarm() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonTechUs108Bms) GetRatedPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonTechUs108Bms) GetMaxInputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PylonTechUs108Bms) GetMaxOutputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
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

func (p *PylonTechUs108Bms) GetCapacity() (uint16, error) {
	return uint16(108), nil
}

func (p *PylonTechUs108Bms) SetReset() error {
	return nil
}

func (p *PylonTechUs108Bms) writeTime() {
	err := p._syncTime()
	if err != nil {
		go func() {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-p.ctx.Done():
					p.log.Noticef(p.ctx, "writeTime() 关闭!")
				case <-ticker.C:
					if !p.IsActivate() {
						continue
					}
					err := p._syncTime()
					if err == nil {
						p.log.Infof(p.ctx, "同步时间成功！")
						//break
					}
				}

			}
		}()
	}
}

func (p *PylonTechUs108Bms) _syncTime() error {
	//if !p.IsActivate() {
	//	return fmt.Errorf("modbus client is not activate")
	//}
	//now := time.Now()
	//
	//err := p.WriteMultipleRegisters(info.GroupTime, []int64{int64(now.Year() - 2000), int64(now.Month()),
	//	int64(now.Day()), int64(now.Hour()), int64(now.Minute()), int64(now.Second())})
	//if err != nil {
	//	return err
	//}
	return nil
}
