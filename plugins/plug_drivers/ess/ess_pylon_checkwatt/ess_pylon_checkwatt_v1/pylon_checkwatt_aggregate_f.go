package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"common/c_enum"
	"common/c_func"
	"common/c_type"

	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

func (p *sEssPylonCheckwatt) GetAmmeterOrPcsSumData(ammeterProcessFunction func(ammeter c_type.IAmmeter) (any, error), pcsProcessFunc func(pcs c_type.IPcs) (*float64, error)) (*float64, error) {
	return getAmmeterOrPcsSumData(p, ammeterProcessFunction, pcsProcessFunc, c_func.AggregateSumFloat64)
}

// GetAmmeterOrPcsSumData 从电表或者PCS获取数据聚合返回方法
func getAmmeterOrPcsSumData[T any](p *sEssPylonCheckwatt,
	ammeterProcessFunction func(ammeter c_type.IAmmeter) (any, error),
	pcsProcessFunc func(pcs c_type.IPcs) (T, error),
	aggregateFunc func(values []any) (T, error),
) (*float64, error) {
	v, err := p.GetFromChildAmmeterOrDeviceType(p.essConfig.AmmeterId, c_enum.EDevicePcs,
		func(ammeter c_type.IAmmeter) (any, error) {
			return ammeterProcessFunction(ammeter)
		}, func(device c_base.IDevice) (any, error) {
			if pcs, ok := device.(c_type.IPcs); ok {
				return pcsProcessFunc(pcs)
			}
			return nil, errors.Errorf("设备[%s]不是pcs类型!", device.GetConfig().Name)
		}, func(values []any) (any, error) {
			// 平均
			result, err := aggregateFunc(values)
			if err != nil {
				return nil, err
			}
			return result, nil
		})
	if err != nil {
		return nil, err
	}
	result, err := cvt.Float64E(v)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
