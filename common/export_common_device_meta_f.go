package common

import (
	"common/c_base"
	"common/c_device"
	"common/internal/internal_device"
	"common/internal/internal_meta"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

// internal_device.Instances 所有设备实例

// MetaTransformAndCache 元数据转换并缓存
func MetaTransformAndCache(ctx context.Context, deviceId string, deviceType c_base.EDeviceType, protocol c_base.IProtocol, meta *c_base.Meta, value any, cache *gcache.Cache, lifetime time.Duration) (*gvar.Var, error) {
	return internal_meta.MetaProcess(ctx, deviceId, deviceType, protocol, meta, value, cache, lifetime)
}

// RegisterDevice 注册设备实例
func RegisterDevice(info c_base.IDriver) {
	internal_device.Instances.RegisterInstance(info)
}

// GetDeviceAll 获取所有设备实例, 参数为空时获取所有设备实例, 参数为true时获取虚拟设备实例, 参数为false时获取实体设备实例
func GetDeviceAll(isVirtual ...bool) []c_base.IDriver {
	return internal_device.Instances.FindAll(isVirtual...)
}

// GetDeviceById 通过id获取设备实例
func GetDeviceById(id string) c_base.IDriver {
	return internal_device.Instances.FindById(id)
}

// GetDeviceByType 通过类型获取设备实例
func GetDeviceByType(t c_base.EDeviceType) []c_base.IDriver {
	return internal_device.Instances.FindByType(t)
}

// RemoveDeviceById 通过id删除设备实例
func RemoveDeviceById(id string) {
	internal_device.Instances.RemoveById(id)
}

// GetStationEnergyStore 获取站点能量存储
func GetStationEnergyStore() c_device.IStationEnergyStore {
	return internal_device.Instances.GetStationEnergyStore()
}
