package bms_pylon_tech_us108_v1

import (
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"common/c_type"
	"math"
	"time"

	"github.com/pkg/errors"

	"github.com/shockerli/cvt"
)

type sBmsPylonTechUs108 struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
	bmsConfig *sPylonTechUs108BmsConfig
}

func (p *sBmsPylonTechUs108) Init() error {
	p.bmsConfig = &sPylonTechUs108BmsConfig{}
	err := p.GetConfig().ScanParams(p.bmsConfig)
	if err != nil {
		return err
	}

	// 2025-08-19 删除了GroupTime
	_ = p.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		protocol.RegisterTask(GroupHeart, GroupInfo, GroupStatistics)
		return nil
	})

	if p.bmsConfig.SyncTime {
		p.startWriteTimeTask()
		c_log.BizInfof(p.DeviceCtx, "syncTime配置为：true！同步时间已开启！")
	} else {
		c_log.BizInfof(p.DeviceCtx, "syncTime配置为：false！时间同步任务未启动！")
	}

	return nil
}

var _ c_type.IBms = (*sBmsPylonTechUs108)(nil)

func (p *sBmsPylonTechUs108) Shutdown() {

}

func (p *sBmsPylonTechUs108) GetRatedPower() (*uint32, error) {
	if p.bmsConfig.RatedPower == nil {
		return nil, errors.New("参数错误")
	}
	return p.bmsConfig.RatedPower, nil
}

func (p *sBmsPylonTechUs108) GetMaxInputPower() (*float32, error) {
	return p.GetFromProtocolFloat32(func(protocol c_proto.IModbusProtocol) (any, error) {
		chargeForbiddenMark, err := protocol.GetBool(ChargeForbiddenMark)
		if err != nil {
			return 0, err
		}
		if chargeForbiddenMark != nil && *chargeForbiddenMark {
			// 禁止充电
			return 0, nil
		}

		// 通过电压和电流来计算功率
		values, err := protocol.GetFloat32Values(PileMaxV, PileMaxI)
		if err != nil {
			return 0, err
		}
		if values[0] == nil || values[1] == nil {
			return 0, nil // 没有采集到数据
		}
		power := *values[0] * *values[1]
		c_log.Debugf(p.DeviceCtx, "最大充电 电压：%f, 电流：%f, 功率：%f", *values[0], *values[1], power)

		v := p.bmsConfig.getMaxInputPower(power)
		if v != nil {
			return v, nil
		}
		return power, nil
	})
}

func (p *sBmsPylonTechUs108) GetMaxOutputPower() (*float32, error) {
	return p.GetFromProtocolFloat32(func(protocol c_proto.IModbusProtocol) (any, error) {
		dischargeForbiddenMark, err := protocol.GetBool(DischargeForbiddenMark)
		if err != nil {
			return 0, err
		}
		if dischargeForbiddenMark != nil && *dischargeForbiddenMark {
			// 禁止放电
			return 0, nil
		}
		// 通过电压和电流来计算功率
		values, err := protocol.GetFloat32Values(PileMinV, PileMaxDI)
		if err != nil {
			return 0, err
		}
		if values[0] == nil || values[1] == nil {
			return 0, nil // 没有采集到数据
		}
		power := *values[0] * *values[1]
		power = float32(math.Abs(float64(power)))

		v := p.bmsConfig.getMaxOutputPower(power)
		if v != nil {
			return v, nil
		}

		c_log.Debugf(p.DeviceCtx, "最大放电 电压：%f, 电流：%f, 功率：%f, 配置功率：%f", *values[0], *values[1], power, p.bmsConfig.MaxOutputPower)
		return power, nil
	})
}

func (p *sBmsPylonTechUs108) GetBmsStatus() (*c_enum.EBmsStatus, error) {
	statusValue, err := p.GetFromProtocolUint8(func(protocol c_proto.IModbusProtocol) (any, error) {
		value, err := protocol.GetIntValue(BasicStatus)
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil // 没有采集到数据
		}
		return *value, nil
	})
	if err != nil {
		status := c_enum.EBmsStatusUnknown
		return &status, err
	}
	if statusValue == nil {
		status := c_enum.EBmsStatusUnknown
		return &status, nil
	}
	switch *statusValue {
	case 0:
		status := c_enum.EBmsStatusOff
		return &status, nil
	case 1:
		status := c_enum.EBmsStatusCharge
		return &status, nil
	case 2:
		status := c_enum.EBmsStatusDischarge
		return &status, nil
	case 3:
		status := c_enum.EBmsStatusStandby
		return &status, nil
	}
	status := c_enum.EBmsStatusUnknown
	return &status, nil
}

func (p *sBmsPylonTechUs108) SetBmsStatus(status c_enum.EBmsStatus) error {
	var err error
	return p.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		switch status {
		case c_enum.EBmsStatusOff:
			err = protocol.WriteSingleRegister(BasicStatus, 0) // 确保可以休眠
		case c_enum.EBmsStatusUnknown:
			err = errors.New("参数错误")
		case c_enum.EBmsStatusStandby:
			err = protocol.WriteSingleRegister(BasicStatus, 3) // 确保可以待机
		case c_enum.EBmsStatusCharge:
		case c_enum.EBmsStatusDischarge:
		case c_enum.EBmsStatusFault:
			// 这些虽然不支持，但是不会返回错误。确保其他的服务调用的时候能够正常
		}

		if err != nil {
			return err
		}

		// 阻塞3秒，等待上电
		time.Sleep(3 * time.Second)

		return nil
	})
}

func (p *sBmsPylonTechUs108) GetSoc() (*float32, error) {
	return p.GetFromPointFloat32(SOC)
}

func (p *sBmsPylonTechUs108) GetSoh() (*float32, error) {
	return p.GetFromPointFloat32(SOH)
}

func (p *sBmsPylonTechUs108) GetDcPower() (*float64, error) {
	current, err := p.GetDcCurrent()
	if err != nil {
		return nil, err
	}
	voltage, err := p.GetDcVoltage()
	if err != nil {
		return nil, err
	}
	if current == nil || voltage == nil {
		return nil, nil
	}
	// kW
	power := *current * *voltage / 1000
	return &power, nil
}

func (p *sBmsPylonTechUs108) GetDcVoltage() (*float64, error) {
	return p.GetFromPointFloat64(DCVoltage)
}

func (p *sBmsPylonTechUs108) GetDcCurrent() (*float64, error) {
	return p.GetFromPointFloat64(DCCurrent)
}

func (p *sBmsPylonTechUs108) GetCellMinTemp() (*float32, error) {
	return p.GetFromPointFloat32(BatteryCellMinTemp)
}

func (p *sBmsPylonTechUs108) GetCellMaxTemp() (*float32, error) {
	return p.GetFromPointFloat32(BatteryCellMaxTemp)
}

func (p *sBmsPylonTechUs108) GetCellAvgTemp() (*float32, error) {
	minTemp, err := p.GetCellMinTemp()
	if err != nil {
		return nil, err
	}
	maxTemp, err := p.GetCellMaxTemp()
	if err != nil {
		return nil, err
	}
	if minTemp == nil || maxTemp == nil {
		return nil, nil
	}
	avgTemp := (*minTemp + *maxTemp) / 2.0
	return &avgTemp, nil
}

func (p *sBmsPylonTechUs108) GetCellMinVoltage() (*float32, error) {
	return p.GetFromPointFloat32(BatteryCellMinVoltage)
}

func (p *sBmsPylonTechUs108) GetCellMaxVoltage() (*float32, error) {
	return p.GetFromPointFloat32(BatteryCellMaxVoltage)
}

func (p *sBmsPylonTechUs108) GetCellAvgVoltage() (*float32, error) {
	minVoltage, err := p.GetCellMinVoltage()
	if err != nil {
		return nil, err
	}
	maxVoltage, err := p.GetCellMaxVoltage()
	if err != nil {
		return nil, err
	}
	if minVoltage == nil || maxVoltage == nil {
		return nil, nil
	}
	avgVoltage := (*minVoltage + *maxVoltage) / 2.0
	return &avgVoltage, nil
}

func (p *sBmsPylonTechUs108) GetCycleCount() (*uint, error) {
	return p.GetFromPointUint(CycleCount)
}

func (p *sBmsPylonTechUs108) GetTodayIncomingQuantity() (*float64, error) {
	return p.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		v, err := protocol.ReadSingleSync(TodayCharge, c_enum.EMqHoldingRegisters, 3*time.Second, true)
		return cvt.Float64(v), err
	})
}

func (p *sBmsPylonTechUs108) GetTodayOutgoingQuantity() (*float64, error) {
	return p.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		v, err := protocol.ReadSingleSync(TodayDischarge, c_enum.EMqHoldingRegisters, 3*time.Second, true)
		return cvt.Float64(v), err
	})
}

func (p *sBmsPylonTechUs108) GetHistoryIncomingQuantity() (*float64, error) {
	return p.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		v, err := protocol.ReadSingleSync(HistoryCharge, c_enum.EMqHoldingRegisters, 3*time.Second, true)
		return cvt.Float64(v), err
	})
}

func (p *sBmsPylonTechUs108) GetHistoryOutgoingQuantity() (*float64, error) {
	return p.GetFromProtocolFloat64(func(protocol c_proto.IModbusProtocol) (any, error) {
		v, err := protocol.ReadSingleSync(HistoryDischarge, c_enum.EMqHoldingRegisters, 3*time.Second, true)
		return cvt.Float64(v), err
	})
}

func (p *sBmsPylonTechUs108) GetCapacity() (*uint32, error) {
	if p.bmsConfig.Capacity == nil {
		return nil, errors.New("不支持的操作")
	}
	return p.bmsConfig.Capacity, nil
}

func (p *sBmsPylonTechUs108) SetReset() error {
	return nil
}

func (p *sBmsPylonTechUs108) startWriteTimeTask() {
	_ = p._syncTime()
	go func() {
		ticker := time.NewTicker(12 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-p.DeviceCtx.Done():
				c_log.Debugf(p.DeviceCtx, "startWriteTimeTask() 关闭!")
				return
			case <-ticker.C:
				if p.GetProtocolStatus() != c_enum.EProtocolConnected {
					continue
				}
				_ = p._syncTime()
			}

		}
	}()
}

func (p *sBmsPylonTechUs108) _syncTime() error {
	return p.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		now := time.Now()
		err := protocol.WriteMultipleRegisters(GroupTime, []int64{int64(now.Year() - 2000), int64(now.Month()),
			int64(now.Day()), int64(now.Hour()), int64(now.Minute()), int64(now.Second())})
		if err != nil {
			return err
		}
		c_log.Infof(p.DeviceCtx, "同步时间成功！")
		return nil
	})
}
