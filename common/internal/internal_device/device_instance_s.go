package internal_device

import (
	"ems-plan/c_base"
	"ems-plan/c_device"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"reflect"
	"sync"
)

var (
	Instances c_device.IDriverInstances
)

type sDeviceInstance struct {
	array *garray.SortedArray
	mutex sync.Mutex
}

func init() {
	Instances = &sDeviceInstance{
		mutex: sync.Mutex{},
		array: garray.NewSortedArray(func(v1, v2 interface{}) int {
			v1Info := v1.(c_base.IDriver).GetDeviceConfig().Id
			v2Info := v2.(c_base.IDriver).GetDeviceConfig().Id

			if v1Info > v2Info {
				return -1
			} else {
				return 1
			}
		}),
	}
}

// RegisterInstance 注册设备实例
func (d *sDeviceInstance) RegisterInstance(info c_base.IDriver) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if info.GetDeviceConfig().Id == "" {
		panic(gerror.Newf("类型：%s的Id不能为空！", reflect.TypeOf(info).String()))
	}
	// 不能重复注册
	if d.FindById(info.GetDeviceConfig().Id) != nil {
		panic(gerror.Newf("id[%s]已经注册到DeviceInstance中！", info.GetDeviceConfig().Id))
	}

	d.array.Add(info)
}

func (d *sDeviceInstance) FindByType(t c_base.EDeviceType) []c_base.IDriver {
	var result []c_base.IDriver
	for _, instance := range d.array.Slice() {
		if instance.(c_base.IDriver).GetDriverType() == t {
			result = append(result, instance.(c_base.IDriver))
		}
	}
	return result
}

// FindById 获取设备实例
func (d *sDeviceInstance) FindById(id string) c_base.IDriver {
	for _, instance := range d.array.Slice() {
		if instance.(c_base.IDriver).GetDeviceConfig().Id == id {
			return instance.(c_base.IDriver)
		}
	}
	return nil
}

func (d *sDeviceInstance) RemoveById(id string) {
	instance := d.FindById(id)
	if instance != nil {
		d.array.RemoveValue(instance)
	}
}

func (d *sDeviceInstance) FindAll(isVirtual ...bool) []c_base.IDriver {

	var result []c_base.IDriver
	for _, info := range d.array.Slice() {
		if isVirtual != nil && len(isVirtual) > 0 {
			if isVirtual[0] {
				if info.(c_base.IDriver).GetDeviceConfig().ProtocolId == "" {
					result = append(result, info.(c_base.IDriver))
				}
			} else {
				if info.(c_base.IDriver).GetDeviceConfig().ProtocolId != "" {
					result = append(result, info.(c_base.IDriver))
				}
			}
		} else {
			// 获取所有
			result = append(result, info.(c_base.IDriver))
		}

	}
	return result
}

func (d *sDeviceInstance) GetStationEnergyStore() c_device.IStationEnergyStore {
	driver := d.FindById(c_base.ConstStationEnergyStoreId)
	if driver == nil {
		return nil
	}
	return driver.(c_device.IStationEnergyStore)
}
