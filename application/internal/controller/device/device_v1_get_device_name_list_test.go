package device

import (
	v1 "application/api/device/v1"
	"common/c_base"
	"context"
	"testing"
)

func TestControllerV1_GetDeviceNameList(t *testing.T) {
	controller := &ControllerV1{}

	// 测试获取所有设备
	req := &v1.GetDeviceNameListReq{}
	res, err := controller.GetDeviceNameList(context.Background(), req)

	// 如果数据库连接正常，应该能获取到结果
	if err != nil {
		t.Logf("获取设备名称列表失败: %+v", err)
	} else {
		t.Logf("获取设备名称列表成功，设备数量: %d", len(res.DeviceNames))
		for deviceId, deviceName := range res.DeviceNames {
			t.Logf("设备ID: %s, 设备名称: %s", deviceId, deviceName)
		}
	}
}

func TestControllerV1_GetDeviceNameListWithTypeFilter(t *testing.T) {
	controller := &ControllerV1{}

	// 测试按设备类型过滤
	req := &v1.GetDeviceNameListReq{
		DeviceTypes: []c_enum.EDeviceType{c_enum.EDeviceBms, c_enum.EDevicePcs},
	}
	res, err := controller.GetDeviceNameList(context.Background(), req)

	// 如果数据库连接正常，应该能获取到结果
	if err != nil {
		t.Logf("按设备类型过滤获取设备名称列表失败: %+v", err)
	} else {
		t.Logf("按设备类型过滤获取设备名称列表成功，设备数量: %d", len(res.DeviceNames))
		for deviceId, deviceName := range res.DeviceNames {
			t.Logf("设备ID: %s, 设备名称: %s", deviceId, deviceName)
		}
	}
}
