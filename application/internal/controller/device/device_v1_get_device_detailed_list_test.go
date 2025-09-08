package device

import (
	v1 "application/api/device/v1"
	"common/c_base"
	"context"
	"testing"
)

func TestControllerV1_GetDeviceDetailedList(t *testing.T) {
	controller := &ControllerV1{}

	// 测试获取所有设备详细信息
	req := &v1.GetDeviceDetailedListReq{}
	res, err := controller.GetDeviceDetailedList(context.Background(), req)

	// 如果数据库连接正常，应该能获取到结果
	if err != nil {
		t.Logf("获取设备详细列表失败: %+v", err)
	} else {
		t.Logf("获取设备详细列表成功，设备数量: %d", res.Total)
		for _, device := range res.Devices {
			t.Logf("设备ID: %s, 设备名称: %s, 设备类型: %s, 启用状态: %t",
				device.Id, device.Name, device.DeviceType, device.Enabled)
		}
	}
}

func TestControllerV1_GetDeviceDetailedListWithTypeFilter(t *testing.T) {
	controller := &ControllerV1{}

	// 测试按设备类型过滤
	req := &v1.GetDeviceDetailedListReq{
		DeviceTypes: []c_base.EDeviceType{c_base.EDeviceBms, c_base.EDevicePcs},
	}
	res, err := controller.GetDeviceDetailedList(context.Background(), req)

	// 如果数据库连接正常，应该能获取到结果
	if err != nil {
		t.Logf("按设备类型过滤获取设备详细列表失败: %+v", err)
	} else {
		t.Logf("按设备类型过滤获取设备详细列表成功，设备数量: %d", res.Total)
		for _, device := range res.Devices {
			t.Logf("设备ID: %s, 设备名称: %s, 设备类型: %s, 驱动品牌: %s, 驱动型号: %s",
				device.Id, device.Name, device.DeviceType, device.DriverBrand, device.DriverModel)
		}
	}
}

func TestControllerV1_GetDeviceDetailedListWithEnabledFilter(t *testing.T) {
	controller := &ControllerV1{}

	// 测试按启用状态过滤
	enabled := true
	req := &v1.GetDeviceDetailedListReq{
		Enabled: &enabled,
	}
	res, err := controller.GetDeviceDetailedList(context.Background(), req)

	// 如果数据库连接正常，应该能获取到结果
	if err != nil {
		t.Logf("按启用状态过滤获取设备详细列表失败: %+v", err)
	} else {
		t.Logf("按启用状态过滤获取设备详细列表成功，设备数量: %d", res.Total)
		for _, device := range res.Devices {
			t.Logf("设备ID: %s, 设备名称: %s, 启用状态: %t",
				device.Id, device.Name, device.Enabled)
		}
	}
}

func TestControllerV1_GetDeviceDetailedListWithCombinedFilter(t *testing.T) {
	controller := &ControllerV1{}

	// 测试组合过滤条件
	enabled := true
	req := &v1.GetDeviceDetailedListReq{
		DeviceTypes: []c_base.EDeviceType{c_base.EDeviceBms, c_base.EDeviceEss},
		Enabled:     &enabled,
	}
	res, err := controller.GetDeviceDetailedList(context.Background(), req)

	// 如果数据库连接正常，应该能获取到结果
	if err != nil {
		t.Logf("组合过滤条件获取设备详细列表失败: %+v", err)
	} else {
		t.Logf("组合过滤条件获取设备详细列表成功，设备数量: %d", res.Total)
		for _, device := range res.Devices {
			t.Logf("设备ID: %s, 设备名称: %s, 设备类型: %s, 启用状态: %t, 协议类型: %s",
				device.Id, device.Name, device.DeviceType, device.Enabled, device.ProtocolType)
		}
	}
}
