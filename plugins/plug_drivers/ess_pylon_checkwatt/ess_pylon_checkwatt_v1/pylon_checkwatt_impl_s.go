package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_error"
	"common/c_func"
	"common/c_log"
	"common/c_type"
	"github.com/shockerli/cvt"
	"gopkg.in/errgo.v2/fmt/errors"
)

type sEssPylonCheckwatt struct {
	*c_device.SVirtualDevice
	essConfig *sEssPylonCheckwattConfig

	//*c_base.SPolicyManager

	targetPower         int32
	targetReactivePower int32
	targetPowerFactor   float32
}

func (p *sEssPylonCheckwatt) ProtocolListen() {

}

func (p *sEssPylonCheckwatt) IsActivate() bool {
	return true
}

func (p *sEssPylonCheckwatt) IsPhysical() bool {
	return false
}

func (p *sEssPylonCheckwatt) Init() error {
	p.essConfig = &sEssPylonCheckwattConfig{}
	err := p.GetConfig().ScanParams(p.essConfig)
	if err != nil {
		return err
	}

	c_log.BizInfof(p.DeviceCtx, "虚拟储能柜初始化完毕!")
	return nil
}

func (p *sEssPylonCheckwatt) Shutdown() {
	_ = p.SetPower(0)
	_ = p.SetStatus(c_base.EPcsStatusOff)
}

func (p *sEssPylonCheckwatt) SetReset() error {

	return nil
}

func (p *sEssPylonCheckwatt) GetSoc() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float32, error) {
		return device.GetSoc()
	}, func(values []any) (float32, error) {
		return c_func.AggregateAvgFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetSoh() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float32, error) {
		return device.GetSoh()
	}, func(values []any) (float32, error) {
		return c_func.AggregateAvgFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetCapacity() (uint32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (uint32, error) {
		return device.GetCapacity()
	}, func(values []any) (uint32, error) {
		return c_func.AggregateSumUint32(values)
	})
}

func (p *sEssPylonCheckwatt) GetCycleCount() (uint, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (uint, error) {
		return device.GetCycleCount()
	}, func(values []any) (uint, error) {
		return c_func.AggregateSumUint(values)
	})
}

func (p *sEssPylonCheckwatt) GetRatedPower() (uint32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) (uint32, error) {
		return device.GetRatedPower()
	}, func(values []any) (uint32, error) {
		return c_func.AggregateSumUint32(values)
	})
}

func (p *sEssPylonCheckwatt) GetMaxInputPower() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) (float32, error) {
		return device.GetMaxInputPower()
	}, func(values []any) (float32, error) {
		return c_func.AggregateSumFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetMaxOutputPower() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) (float32, error) {
		return device.GetMaxOutputPower()
	}, func(values []any) (float32, error) {
		return c_func.AggregateSumFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetDcPower() (float64, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float64, error) {
		return device.GetDcPower()
	}, func(values []any) (float64, error) {
		return c_func.AggregateSumFloat64(values)
	})
}

func (p *sEssPylonCheckwatt) GetCellMinTemp() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float32, error) {
		return device.GetCellMinTemp()
	}, func(values []any) (float32, error) {
		return c_func.AggregateMinFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetCellMaxTemp() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float32, error) {
		return device.GetCellMaxTemp()
	}, func(values []any) (float32, error) {
		return c_func.AggregateMaxFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetCellAvgTemp() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float32, error) {
		return device.GetCellAvgTemp()
	}, func(values []any) (float32, error) {
		return c_func.AggregateAvgFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetCellMinVoltage() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float32, error) {
		return device.GetCellMinVoltage()
	}, func(values []any) (float32, error) {
		return c_func.AggregateMinFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetCellMaxVoltage() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float32, error) {
		return device.GetCellMaxVoltage()
	}, func(values []any) (float32, error) {
		return c_func.AggregateMaxFloat32(values)
	})
}

func (p *sEssPylonCheckwatt) GetCellAvgVoltage() (float32, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (float32, error) {
		return device.GetCellAvgVoltage()
	}, func(values []any) (float32, error) {
		return c_func.AggregateAvgFloat32(values)
	})
}

// GetAmmeterOrPcsSumData 从电表或者PCS获取数据聚合返回方法
func (p *sEssPylonCheckwatt) GetAmmeterOrPcsSumData(ammeterProcessFunction func(ammeter c_type.IAmmeter) (any, error), pcsProcessFunc func(pcs c_type.IPcs) (float64, error)) (float64, error) {
	v, err := p.GetFromChildAmmeterOrDeviceType(p.essConfig.AmmeterId, c_base.EDevicePcs,
		func(ammeter c_type.IAmmeter) (any, error) {
			return ammeterProcessFunction(ammeter)
		}, func(device c_base.IDevice) (any, error) {
			if pcs, ok := device.(c_type.IPcs); ok {
				return pcsProcessFunc(pcs)
			}
			return nil, errors.Newf("设备[%s]不是pcs类型!", device.GetConfig().Name)
		}, func(values []any) (any, error) {
			// 聚合
			return c_func.AggregateSumFloat64(values)
		})

	if err != nil {
		return 0, err
	}
	return cvt.Float64E(v)
}

func (p *sEssPylonCheckwatt) GetTodayIncomingQuantity() (float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetTodayIncomingQuantity()
	}, func(pcs c_type.IPcs) (float64, error) {
		return pcs.GetTodayIncomingQuantity()
	})
}

func (p *sEssPylonCheckwatt) GetHistoryIncomingQuantity() (float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetHistoryIncomingQuantity()
	}, func(pcs c_type.IPcs) (float64, error) {
		return pcs.GetHistoryIncomingQuantity()
	})
}

func (p *sEssPylonCheckwatt) GetTodayOutgoingQuantity() (float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetTodayOutgoingQuantity()
	}, func(pcs c_type.IPcs) (float64, error) {
		return pcs.GetTodayOutgoingQuantity()
	})
}

func (p *sEssPylonCheckwatt) GetHistoryOutgoingQuantity() (float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetHistoryOutgoingQuantity()
	}, func(pcs c_type.IPcs) (float64, error) {
		return pcs.GetHistoryOutgoingQuantity()
	})
}

func (p *sEssPylonCheckwatt) SetStatus(status c_base.EEnergyStoreStatus) error {
	if status == c_base.EPcsStatusUnknown || status == c_base.EPcsStatusSync || status == c_base.EPcsStatusFault {
		return c_error.ErrorParam
	}
	bmsStatus, err := p.GetBmsStatus()
	if err != nil {
		return errors.Newf("获取BMS状态失败! 错误原因：%s", err.Error())
	}

	if bmsStatus == c_type.EBmsStatusOff {
		// 先去开机
		err = c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) error {
			return device.SetBmsStatus(c_type.EBmsStatusStandby) // 设为待机
		})
		if err != nil {
			return errors.Newf("设置BMS状态失败！错误原因: %s", err.Error())
		}

		bmsStatus, err = p.GetBmsStatus()
		if err != nil {
			return errors.Newf("设置BMS状态为开机后，仍然失败。指令[%s]放弃 原因:%s", status.String(), err.Error())
		}
	}
	// 设置PCS状态
	return c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) error {
		return device.SetStatus(status)
	})
}

func (p *sEssPylonCheckwatt) GetBmsStatus() (c_type.EBmsStatus, error) {
	// 判断电池是否上电，如果没有就先上电
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IBms) (c_type.EBmsStatus, error) {
		return device.GetBmsStatus()
	}, func(values []any) (c_type.EBmsStatus, error) {
		return c_func.EqualAggregate[c_type.EBmsStatus](values)
	})
}

func (p *sEssPylonCheckwatt) SetGridMode(mode c_base.EGridMode) error {
	return c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) error {
		return device.SetGridMode(mode)
	})
}

func (p *sEssPylonCheckwatt) GetStatus() (c_base.EEnergyStoreStatus, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) (c_base.EEnergyStoreStatus, error) {
		return device.GetStatus()
	}, func(values []any) (c_base.EEnergyStoreStatus, error) {
		return c_func.EqualAggregate[c_base.EEnergyStoreStatus](values)
	})
}

func (p *sEssPylonCheckwatt) GetGridMode() (c_base.EGridMode, error) {
	return c_device.VirtualGetDataWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) (c_base.EGridMode, error) {
		return device.GetGridMode()
	}, func(values []any) (c_base.EGridMode, error) {
		return c_func.EqualAggregate[c_base.EGridMode](values)
	})
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
	//return c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) error {
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
	return c_device.VirtualExecuteWithChildDeviceType(p.SVirtualDevice, func(device c_type.IPcs) error {
		return device.SetPowerFactor(factor)
	})
}

func (p *sEssPylonCheckwatt) GetTargetPower() int32 {
	return p.targetPower
}

func (p *sEssPylonCheckwatt) GetTargetReactivePower() int32 {
	return p.targetReactivePower
}

func (p *sEssPylonCheckwatt) GetTargetPowerFactor() float32 {
	return p.targetPowerFactor
}

func (p *sEssPylonCheckwatt) GetPower() (float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetPTotal()
	}, func(pcs c_type.IPcs) (float64, error) {
		return pcs.GetPower()
	})
}

func (p *sEssPylonCheckwatt) GetApparentPower() (float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetSTotal()
	}, func(pcs c_type.IPcs) (float64, error) {
		return pcs.GetApparentPower()
	})
}

func (p *sEssPylonCheckwatt) GetReactivePower() (float64, error) {
	return p.GetAmmeterOrPcsSumData(func(ammeter c_type.IAmmeter) (any, error) {
		return ammeter.GetQTotal()
	}, func(pcs c_type.IPcs) (float64, error) {
		return pcs.GetReactivePower()
	})
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
