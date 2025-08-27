package c_device

import (
	"common"
	"common/c_base"
	"common/c_type"
	"context"
	"github.com/pkg/errors"
	"time"
)

type SVirtualDeviceImpl struct { // 虚拟设备
	c_base.IAlarm
	DeviceCtx    context.Context
	cancel       context.CancelFunc
	deviceConfig *c_base.SDeviceConfig
}

var _ c_base.IDevice = (*SVirtualDeviceImpl)(nil)

func NewVirtualDevice(ctx context.Context, deviceConfig *c_base.SDeviceConfig) *SVirtualDeviceImpl {
	deviceCtx, cancel := context.WithCancel(ctx)

	return &SVirtualDeviceImpl{
		IAlarm:       NewAlarmImpl(deviceCtx, deviceConfig.Id, deviceConfig.Pid),
		DeviceCtx:    deviceCtx,
		cancel:       cancel,
		deviceConfig: deviceConfig,
	}
}

func (s *SVirtualDeviceImpl) Reset() error {
	for _, childDevice := range s.deviceConfig.ChildDeviceConfig {
		d := common.GetDeviceManager().GetDeviceById(childDevice.Id)
		if d != nil {
			d.ResetAlarm()
		}
	}
	return nil
}

func (s *SVirtualDeviceImpl) GetMetaValueList() []*c_base.MetaValueWrapper {
	var list = make([]*c_base.MetaValueWrapper, 0)
	for _, childDevice := range s.deviceConfig.ChildDeviceConfig {
		child := common.GetDeviceManager().GetDeviceById(childDevice.Id)
		if child == nil {
			continue
		}
		childList := child.GetMetaValueList()
		if len(childList) > 0 {
			list = append(list, childList[:]...)
		}
	}
	return list
}

func (s *SVirtualDeviceImpl) GetLastUpdateTime() *time.Time {
	var lastUpdateTime *time.Time
	for _, childDevice := range s.deviceConfig.ChildDeviceConfig {
		child := s.GetChildById(childDevice.Id)
		if child == nil {
			continue
		}
		if t := child.GetLastUpdateTime(); t != nil {
			if lastUpdateTime == nil || lastUpdateTime.Before(*t) {
				lastUpdateTime = t
			}
		}
	}
	return lastUpdateTime
}

func (s *SVirtualDeviceImpl) GetChildById(childDeviceId string) c_base.IDevice {
	return common.GetDeviceManager().GetDeviceById(childDeviceId)
}

func (s *SVirtualDeviceImpl) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *SVirtualDeviceImpl) GetFromChildDeviceId(childDeviceId string, processFunction func(device c_base.IDevice) (any, error)) (any, error) {
	child := s.GetChildById(childDeviceId)
	if child == nil {
		return nil, errors.New("数据不存在")
	}
	return processFunction(child)
}

// GetFromChildDeviceType 使用设备类型来获取数据，和VirtualDataFromChildDeviceType的区别在于，这个方法能处理所有类型，而VirtualDataFromChildDeviceType 只能处理数字
func (s *SVirtualDeviceImpl) GetFromChildDeviceType(childDeviceType c_base.EDeviceType,
	processFunction func(device c_base.IDevice) (any, error), // 处理函数
	aggregateFunction func(values []any) (any, error)) (any, error) { // 聚合函数
	var results = make([]any, 0)
	for _, childDevice := range s.deviceConfig.ChildDeviceConfig {
		if childDevice.DriverInfo != nil && childDevice.GetType() == childDeviceType {
			child := s.GetChildById(childDevice.Id)
			if child == nil {
				// 如果出现有配置，但是无实例。就认为是异常
				return nil, errors.Errorf("设备[%s]未激活，获取数据失败！", childDevice.Name)
			}
			// 匹配类型
			res, err := processFunction(child)
			if err != nil {
				return nil, err
			}
			results = append(results, res)
		}
	}
	if len(results) == 0 {
		return nil, errors.New("数据不存在")
	}
	return aggregateFunction(results)
}

// GetFromChildAmmeterOrDeviceType 根据电表id或者设备类型来获取数据，优先使用电表，如果电表id为空，才会使用type
func (s *SVirtualDeviceImpl) GetFromChildAmmeterOrDeviceType(ammeterId string, childDeviceType c_base.EDeviceType,
	ammeterProcessFunction func(ammeter c_type.IAmmeter) (any, error),
	processFunction func(device c_base.IDevice) (any, error),
	aggregateFunction func(values []any) (any, error)) (any, error) {
	if ammeterId != "" {
		// 如果不为空，那么必须是电表实例才行
		device := s.GetChildById(ammeterId)
		if device == nil {
			return nil, errors.Errorf("电表ID：[%s]未激活，获取数据失败！", ammeterId)
		}
		if ammeter, ok := device.(c_type.IAmmeter); ok {
			return ammeterProcessFunction(ammeter)
		} else {
			return nil, errors.Errorf("设备ID：[%s] 并不是电表，获取数据失败！", ammeterId)
		}
	}

	return s.GetFromChildDeviceType(childDeviceType, processFunction, aggregateFunction)
}
