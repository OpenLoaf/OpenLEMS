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
	v := internal_meta.MetaProcess(meta, value)
	return internal_meta.CacheValue(ctx, deviceId, deviceType, protocol, meta, v, cache, lifetime)
}

// RegisterRunningDevice 注册设备实例
func RegisterRunningDevice(info c_base.IDevice) {
	internal_device.Instances.RegisterInstance(info)
}

// GetRunningDeviceAll 获取所有设备实例, 参数为空时获取所有设备实例, 参数为true时获取虚拟设备实例, 参数为false时获取实体设备实例
func GetRunningDeviceAll(isVirtual ...bool) []c_base.IDevice {
	return internal_device.Instances.FindAll(isVirtual...)
}

// GetRunningDeviceById 通过id获取设备实例
func GetRunningDeviceById(id string) c_base.IDevice {
	return internal_device.Instances.FindById(id)
}

// GetRunningDeviceByType 通过类型获取设备实例
func GetRunningDeviceByType(t c_base.EDeviceType) []c_base.IDevice {
	return internal_device.Instances.FindByType(t)
}

// RemoveRunningDeviceById 通过id删除设备实例
func RemoveRunningDeviceById(id string) {
	internal_device.Instances.RemoveById(id)
}

// GetStationEnergyStore 获取站点能量存储
func GetStationEnergyStore() c_device.IStationEnergyStore {
	return internal_device.Instances.GetStationEnergyStore()
}

// MetaTransformCanbus 解析can的数据
func MetaTransformCanbus(ctx context.Context, deviceId string, deviceType c_base.EDeviceType, protocol c_base.IProtocol, meta *c_base.Meta, canData []byte, cache *gcache.Cache, lifetime time.Duration) (any, error) {
	v, err := internal_meta.ParseCanbusData(canData, meta)
	if err != nil {
		return nil, err
	}

	return internal_meta.CacheValue(ctx, deviceId, deviceType, protocol, meta, v, cache, lifetime)
}
