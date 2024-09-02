package internal_device

import (
	"ems-plan/c_base"
	"ems-plan/c_device"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
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
		panic(fmt.Sprintf("类型：%s的Id不能为空！", reflect.TypeOf(info).String()))
	}
	// 不能重复注册
	if d.FindById(info.GetDeviceConfig().Id) != nil {
		panic("the id '" + info.GetDeviceConfig().Id + "' has been registered")
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

// GetInstance 获取设备实例
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
				if info.(c_base.IDriver).GetDeviceConfig().IsVirtual {
					result = append(result, info.(c_base.IDriver))
				}
			} else {
				if !info.(c_base.IDriver).GetDeviceConfig().IsVirtual {
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

//func (d *sDeviceInstance) FindByConfig(config *c_base.SDriverConfig) c_base.IDriver {
//	if config.Id == "" || config.RefId == "" {
//		// Id 或者 refId不能同时为空
//		panic(fmt.Sprintf("Id: %s, RefId: %s", config.Id, config.RefId))
//	}
//
//
//}
