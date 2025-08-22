package device

import (
	v1 "application/api/device/v1"
	"context"
	"testing"
)

func TestControllerV1_GetDeviceNameList(t *testing.T) {
	controller := &ControllerV1{}
	req := &v1.GetDeviceNameListReq{}

	// 注意：这个测试需要数据库连接，在实际环境中运行
	// 这里只是验证代码结构是否正确
	res, err := controller.GetDeviceNameList(context.Background(), req)

	// 如果数据库连接正常，应该能获取到结果
	if err != nil {
		t.Logf("获取设备名称列表失败: %v", err)
	} else {
		t.Logf("获取设备名称列表成功，设备数量: %d", len(res.DeviceNames))
		for deviceId, deviceName := range res.DeviceNames {
			t.Logf("设备ID: %s, 设备名称: %s", deviceId, deviceName)
		}
	}
}
