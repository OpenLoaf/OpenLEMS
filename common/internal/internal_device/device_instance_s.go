package internal_device

import (
	"ems-plan/c_device"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"reflect"
	"sync"
)

var (
	Instances c_device.IInstances
)

type sDeviceInstance struct {
	array *garray.SortedArray
	mutex sync.Mutex
}

func init() {
	Instances = &sDeviceInstance{
		mutex: sync.Mutex{},
		array: garray.NewSortedArray(func(v1, v2 interface{}) int {
			v1Info := v1.(c_device.IInfo).GetId()
			v2Info := v2.(c_device.IInfo).GetId()

			if v1Info > v2Info {
				return -1
			} else {
				return 1
			}
		}),
	}
}

// RegisterInstance 注册设备实例
func (d *sDeviceInstance) RegisterInstance(info c_device.IInfo) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if info.GetId() == "" {
		panic(fmt.Sprintf("类型：%s的Id不能为空！", reflect.TypeOf(info).String()))
	}
	// 不能重复注册
	if d.GetInstance(info.GetId()) != nil {
		panic("the id '" + info.GetId() + "' has been registered")
	}

	d.array.Add(info)
}

// GetInstance 获取设备实例
func (d *sDeviceInstance) GetInstance(id string) c_device.IInfo {
	for _, instance := range d.array.Slice() {
		if instance.(c_device.IInfo).GetId() == id {
			return instance.(c_device.IInfo)
		}
	}
	return nil
}

func (d *sDeviceInstance) DelInstance(id string) {
	instance := d.GetInstance(id)
	if instance != nil {
		d.array.RemoveValue(instance)
	}
}

func (d *sDeviceInstance) GetList() []c_device.IInfo {
	var result []c_device.IInfo
	for _, info := range d.array.Slice() {
		result = append(result, info.(c_device.IInfo))
	}
	return result
}
