package common

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"ems-plan/internal/internal_device"
	"ems-plan/internal/internal_meta"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

// internal_device.Instances 所有设备实例

// MetaTransformAndCache 元数据转换并缓存
func MetaTransformAndCache(ctx context.Context, protocol c_base.IProtocol, meta *c_base.Meta, value any, cache *gcache.Cache, lifetime time.Duration) (*gvar.Var, error) {
	return internal_meta.MetaProcess("", "", ctx, protocol, meta, value, cache, lifetime)
}

func RegisterDevice(info c_base.IDriver) {
	internal_device.Instances.RegisterInstance(info)
}

// GetDeviceAll 获取所有设备实例, 参数为空时获取所有设备实例, 参数为true时获取虚拟设备实例, 参数为false时获取实体设备实例
func GetDeviceAll(isVirtual ...bool) []c_base.IDriver {
	return internal_device.Instances.FindAll(isVirtual...)
}

func GetDeviceById(id string) c_base.IDriver {
	return internal_device.Instances.FindById(id)
}

func GetDeviceByType(t c_base.EDeviceType) []c_base.IDriver {
	return internal_device.Instances.FindByType(t)
}

func RemoveDeviceById(id string) {
	internal_device.Instances.RemoveById(id)
}

func GetStationEnergyStore() c_device.IStationEnergyStore {
	return internal_device.Instances.GetStationEnergyStore()
}
