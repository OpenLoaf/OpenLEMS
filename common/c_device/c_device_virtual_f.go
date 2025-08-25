package c_device

import (
	"common/c_base"
	"common/c_error"
	"common/c_log"
	"gopkg.in/errgo.v2/fmt/errors"
)

// VirtualGetDataWithChildDeviceType 通过范型来获取数据 和 GetFromChildDeviceType 的区别在于这个是number
func VirtualGetDataWithChildDeviceType[T c_base.IDriver, V any](s *SVirtualDevice,
	processFunction func(device T) (V, error),
	aggregateFunction func(values []any) (V, error)) (V, error) {
	var results = make([]any, 0)
	var zero V
	for _, childDevice := range s.deviceConfig.ChildDeviceConfig {
		if childDevice.DriverInfo != nil {
			child := s.GetChildById(childDevice.Id)
			if child == nil {
				// 如果出现有配置，但是无实例。就认为是异常
				return zero, errors.Newf("设备[%s]未激活，获取数据失败！", childDevice.Name)
			}
			// 匹配类型
			if t, ok := child.(T); ok {
				result, err := processFunction(t)
				if err != nil {
					return zero, err
				}
				results = append(results, result)
			}
		}
	}
	if len(results) == 0 {
		return zero, c_error.NoData
	}
	return aggregateFunction(results)
}

// VirtualExecuteWithChildDeviceId 获取某个id的设备，并执行方法. 只能是虚拟设备才能使用。
func VirtualExecuteWithChildDeviceId[T c_base.IDriver, R c_base.Number](s *SVirtualDevice, id string, fc func(T) (R, error)) (R, error) {
	c := s.GetChildById(id)
	if c == nil {
		var zero R
		return zero, errors.Newf("child with id %s not found", id)
	}

	if typedChild, ok := c.(T); ok {
		return fc(typedChild)
	}
	var zero R
	return zero, errors.Newf("type assertion failed for id %s", id)
}

// VirtualExecuteWithChildDeviceType 执行某个类型的所有子设备的方法
func VirtualExecuteWithChildDeviceType[T c_base.IDriver](s *SVirtualDevice, fc func(device T) error) error {
	return VirtualExecuteAndRollbackWithChildDeviceType(s, fc, nil)
}

// VirtualExecuteAndRollbackWithChildDeviceType 执行某个类型的所有子设备的方法，如果执行失败了，一个个回滚
func VirtualExecuteAndRollbackWithChildDeviceType[T c_base.IDriver](s *SVirtualDevice, fc func(device T) error, rollbackFc func(device T) error) error {
	if fc == nil {
		return errors.Newf("VirtualExecuteAndRollbackWithChildDeviceType 方法参数错误！")
	}
	var childList = make([]T, 0)
	for _, childDevice := range s.deviceConfig.ChildDeviceConfig {
		if childDevice.DriverInfo != nil {
			child := s.GetChildById(childDevice.Id)
			if child == nil {
				// 如果出现有配置，但是无实例。就认为是异常
				return errors.Newf("设备[%s]未激活，执行方法失败！", childDevice.Name)
			}
			// 匹配类型
			if t, ok := child.(T); ok {
				childList = append(childList, t)
			}
		}
	}

	// 记录成功的
	var successChild = make([]T, 0)
	var err error
	// 再开一个循环去执行任务，防止多台设备，只执行了一两台
	for _, childDevice := range childList {
		err = fc(childDevice)
		if err != nil {
			break
		}
		successChild = append(successChild, childDevice)
	}

	if err != nil && rollbackFc != nil {
		// 执行回滚任务
		for _, child := range successChild {
			r := rollbackFc(child)
			if r != nil {
				c_log.BizInfof(s.DeviceCtx, "执行任务失败后，回滚也失败！失败原因:%s", r.Error())
			}
		}
	}

	return err
}
