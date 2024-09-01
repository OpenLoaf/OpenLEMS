package telemetry

import (
	"context"

	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"time"
)

func (s *sTelemetry) GetValueByTelemetryKey(deviceId string, telemetryKey string) (any, error) {
	instance, err := device.GetInstance(deviceId)
	if err != nil {
		return nil, err
	}

	return station.ExecuteFunction(instance, telemetryKey)
}

func (s *sTelemetry) GetValueWithCache(deviceId string) ([]*entity.MetaValue, *time.Time, error) {
	instance, err := device.GetInstance(deviceId)
	if err != nil {
		return nil, nil, err
	}

	info := instance.GetInfo()
	if info.IsVirtual {
		return nil, nil, fmt.Errorf("设备 %s 为虚拟设备，请使用真实设备！", deviceId)
	}

	// 排序
	_sortValues := garray.NewSortedArray(func(v1, v2 interface{}) int {
		return int(v1.(*entity.MetaValue).Meta.Addr - v2.(*entity.MetaValue).Meta.Addr)
	})

	ctx := context.Background()
	if i, ok := instance.(device.IDriver); ok {
		if i.GetCache() == nil {
			return nil, nil, fmt.Errorf("设备 %s 未初始化缓存！", deviceId)
		}
		keys, err := i.GetCache().Keys(ctx)
		if err != nil {
			return nil, nil, err
		}
		for _, k := range keys {
			if meta, ok := k.(*point.Meta); ok {
				_varValue, err := i.GetCache().Get(ctx, meta)
				if err != nil {
					return nil, nil, err
				}
				_sortValues.Add(&entity.MetaValue{
					Meta:  meta,
					Value: meta.ValueToString(_varValue),
				})
			}
		}
	}

	result := make([]*entity.MetaValue, _sortValues.Len())
	for i, v := range _sortValues.Slice() {
		result[i] = v.(*entity.MetaValue)
	}

	return result, instance.GetLastUpdateTime(), nil
}
