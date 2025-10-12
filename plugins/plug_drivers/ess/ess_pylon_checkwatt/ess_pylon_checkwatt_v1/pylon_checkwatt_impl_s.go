package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_func"
	"common/c_log"
	"common/c_type"

	"github.com/pkg/errors"
)

type sEssPylonCheckwatt struct {
	*c_device.SVirtualDeviceImpl
	essConfig *sEssPylonCheckwattConfig

	//*c_base.SBasePolicyImpl

	targetPower         int32
	targetReactivePower int32
	targetPowerFactor   float32
}

var _ c_type.IEnergyStore = (*sEssPylonCheckwatt)(nil)

func (p *sEssPylonCheckwatt) Init() error {
	p.essConfig = &sEssPylonCheckwattConfig{}
	err := p.GetConfig().ScanParams(p.essConfig)
	if err != nil {
		return err
	}

	c_log.BizInfof(p.DeviceCtx, "虚拟储能柜初始化完毕!")

	p.RegisterAlarmHandlerFunc(c_enum.EAlarmActionLevelUp, func(alarm *c_base.SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, isFirstHandler bool) {
		c_log.BizInfof(p.DeviceCtx, "触发告警")
	})

	return nil
}

func (p *sEssPylonCheckwatt) GetFrequency() (*float64, error) {
	return getAmmeterOrPcsSumData(p, func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetFrequency()
	}, func(pcs c_type.IPcs) (*float64, error) {
		return pcs.GetFrequency()
	}, c_func.AggregateAvgFloat64)
}

func (p *sEssPylonCheckwatt) GetIGBTTemperature() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(pcs c_type.IPcs) (*float32, error) {
		return pcs.GetIGBTTemperature()
	}, c_func.AggregateAvgFloat32)
}

func (p *sEssPylonCheckwatt) Shutdown() {
	_ = p.SetPower(0)
	_ = p.SetStatus(c_enum.EPcsStatusOff)
}

func (p *sEssPylonCheckwatt) SetReset() error {

	return nil
}

func (p *sEssPylonCheckwatt) GetSoc() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float32, error) {
		return device.GetSoc()
	}, c_func.AggregateAvgFloat32)
}

func (p *sEssPylonCheckwatt) GetSoh() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float32, error) {
		return device.GetSoh()
	}, c_func.AggregateAvgFloat32)
}

func (p *sEssPylonCheckwatt) GetCapacity() (*uint32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*uint32, error) {
		return device.GetCapacity()
	}, c_func.AggregateSumUint32)
}

func (p *sEssPylonCheckwatt) GetCycleCount() (*uint, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*uint, error) {
		return device.GetCycleCount()
	}, c_func.AggregateSumUint)
}

func (p *sEssPylonCheckwatt) GetRatedPower() (*uint32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) (*uint32, error) {
		return device.GetRatedPower()
	}, c_func.AggregateSumUint32)
}

func (p *sEssPylonCheckwatt) GetMaxInputPower() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) (*float32, error) {
		return device.GetMaxInputPower()
	}, c_func.AggregateSumFloat32)
}

func (p *sEssPylonCheckwatt) GetMaxOutputPower() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) (*float32, error) {
		return device.GetMaxOutputPower()
	}, c_func.AggregateSumFloat32)
}

func (p *sEssPylonCheckwatt) GetDcPower() (*float64, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float64, error) {
		return device.GetDcPower()
	}, c_func.AggregateSumFloat64)
}

func (p *sEssPylonCheckwatt) GetCellMinTemp() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float32, error) {
		return device.GetCellMinTemp()
	}, c_func.AggregateMinFloat32)
}

func (p *sEssPylonCheckwatt) GetCellMaxTemp() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float32, error) {
		return device.GetCellMaxTemp()
	}, c_func.AggregateMaxFloat32)
}

func (p *sEssPylonCheckwatt) GetCellAvgTemp() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float32, error) {
		return device.GetCellAvgTemp()
	}, c_func.AggregateAvgFloat32)
}

func (p *sEssPylonCheckwatt) GetCellMinVoltage() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float32, error) {
		return device.GetCellMinVoltage()
	}, c_func.AggregateMinFloat32)
}

func (p *sEssPylonCheckwatt) GetCellMaxVoltage() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float32, error) {
		return device.GetCellMaxVoltage()
	}, c_func.AggregateMaxFloat32)
}

func (p *sEssPylonCheckwatt) GetCellAvgVoltage() (*float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*float32, error) {
		return device.GetCellAvgVoltage()
	}, c_func.AggregateAvgFloat32)
}

func (p *sEssPylonCheckwatt) GetTodayIncomingQuantity() (*float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetTodayIncomingQuantity()
	}, func(pcs c_type.IPcs) (*float64, error) {
		return pcs.GetTodayIncomingQuantity()
	})
}

func (p *sEssPylonCheckwatt) GetHistoryIncomingQuantity() (*float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetHistoryIncomingQuantity()
	}, func(pcs c_type.IPcs) (*float64, error) {
		return pcs.GetHistoryIncomingQuantity()
	})
}

func (p *sEssPylonCheckwatt) GetTodayOutgoingQuantity() (*float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetTodayOutgoingQuantity()
	}, func(pcs c_type.IPcs) (*float64, error) {
		return pcs.GetTodayOutgoingQuantity()
	})
}

func (p *sEssPylonCheckwatt) GetHistoryOutgoingQuantity() (*float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetHistoryOutgoingQuantity()
	}, func(pcs c_type.IPcs) (*float64, error) {
		return pcs.GetHistoryOutgoingQuantity()
	})
}

func (p *sEssPylonCheckwatt) SetStatus(status c_enum.EEnergyStoreStatus) error {
	if status == c_enum.EPcsStatusUnknown || status == c_enum.EPcsStatusSync || status == c_enum.EPcsStatusFault {
		return errors.New("参数错误")
	}
	bmsStatus, err := p.GetBmsStatus()
	if err != nil {
		return errors.Errorf("获取BMS状态失败! 错误原因：%s", err.Error())
	}

	if bmsStatus != nil && *bmsStatus == c_enum.EBmsStatusOff {
		// 先去开机
		err = c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) error {
			return device.SetBmsStatus(c_enum.EBmsStatusStandby) // 设为待机
		})
		if err != nil {
			return errors.Errorf("设置BMS状态失败！错误原因: %s", err.Error())
		}

		// _, err = p.GetBmsStatus()
		// if err != nil {
		// 	return errors.Errorf("设置BMS状态为开机后，仍然失败。指令[%s]放弃 原因:%s", status.String(), err.Error())
		// }
	}
	// 设置PCS状态
	return c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) error {
		return device.SetStatus(status)
	})
}

func (p *sEssPylonCheckwatt) SetGridMode(mode c_enum.EGridMode) error {
	return c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) error {
		return device.SetGridMode(mode)
	})
}

func (p *sEssPylonCheckwatt) GetBmsStatus() (*c_enum.EBmsStatus, error) {
	// 判断电池是否上电，如果没有就先上电
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IBms) (*c_enum.EBmsStatus, error) {
		return device.GetBmsStatus()
	}, c_func.EqualAggregate)
}

func (p *sEssPylonCheckwatt) GetStatus() (*c_enum.EEnergyStoreStatus, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) (*c_enum.EEnergyStoreStatus, error) {
		return device.GetStatus()
	}, c_func.EqualAggregate)
}

func (p *sEssPylonCheckwatt) GetGridMode() (*c_enum.EGridMode, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) (*c_enum.EGridMode, error) {
		return device.GetGridMode()
	}, c_func.EqualAggregate)
}

func (p *sEssPylonCheckwatt) SetPower(power int32) error {
	// 判断一下防止超限
	if power > 0 {
		maxOutputPower, err := p.GetMaxOutputPower()
		if err != nil {
			return err
		}
		if maxOutputPower != nil && power > int32(*maxOutputPower) {
			return errors.New("数值超过限制")
		}
	} else {
		maxInputPower, err := p.GetMaxInputPower()
		if err != nil {
			return err
		}
		if maxInputPower != nil && power < int32(-*maxInputPower) {
			return errors.New("数值超过限制")
		}
	}
	//return c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) error {
	//	return device.SetPower(power)
	//})

	// todo 此处功率需要分配后，给每个设备赋值
	return nil
}

func (p *sEssPylonCheckwatt) SetReactivePower(power int32) error {
	// todo 此处功率需要分配后，给每个设备赋值
	return nil
}

func (p *sEssPylonCheckwatt) SetPowerFactor(factor float32) error {
	return c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDeviceImpl, func(device c_type.IPcs) error {
		return device.SetPowerFactor(factor)
	})
}

func (p *sEssPylonCheckwatt) GetTargetPower() (*int32, error) {
	return &p.targetPower, nil
}

func (p *sEssPylonCheckwatt) GetTargetReactivePower() (*int32, error) {
	return &p.targetReactivePower, nil
}

func (p *sEssPylonCheckwatt) GetTargetPowerFactor() (*float32, error) {
	return &p.targetPowerFactor, nil
}

func (p *sEssPylonCheckwatt) GetPower() (*float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetPTotal()
	}, func(pcs c_type.IPcs) (*float64, error) {
		return pcs.GetPower()
	})
}

func (p *sEssPylonCheckwatt) GetApparentPower() (*float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetSTotal()
	}, func(pcs c_type.IPcs) (*float64, error) {
		return pcs.GetApparentPower()
	})
}

func (p *sEssPylonCheckwatt) GetReactivePower() (*float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetQTotal()
	}, func(pcs c_type.IPcs) (*float64, error) {
		return pcs.GetReactivePower()
	})
}

func (p *sEssPylonCheckwatt) GetFireEnvTemperature() (*float64, error) {
	return nil, errors.New("不支持的操作")
}

func (p *sEssPylonCheckwatt) GetCarbonMonoxideConcentration() (*float64, error) {
	return nil, errors.New("不支持的操作")
}

func (p *sEssPylonCheckwatt) HasSmoke() (*bool, error) {
	return nil, errors.New("不支持的操作")
}

// 实现新的IDevice接口方法
func (p *sEssPylonCheckwatt) GetDevicePoints() []c_base.IPoint {
	// 虚拟ESS驱动没有点位，返回空列表
	return []c_base.IPoint{}
}

// GetTelemetryPoints 获取主要遥测点位列表（只返回关键点位）
func (p *sEssPylonCheckwatt) GetTelemetryPoints() []c_base.IPoint {
	// 虚拟ESS驱动没有点位，返回空列表
	return []c_base.IPoint{}
}

// GetExportModbusPoints 获取暴露出去的modbus点位
func (p *sEssPylonCheckwatt) GetExportModbusPoints() []c_base.IPoint {
	// 虚拟ESS驱动没有点位，返回空列表
	return []c_base.IPoint{}
}
