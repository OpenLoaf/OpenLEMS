package device

import (
	v1 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetVirtualDeviceCache(ctx context.Context, req *v1.GetVirtualDeviceCacheReq) (res *v1.GetVirtualDeviceCacheRes, err error) {
	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(req.DeviceId)
	if deviceConfig == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}

	// 汇总所有子设备的遥测点位
	var allValues []*entity.SSingleDeviceValue

	// 遍历所有子设备
	for _, childConfig := range deviceConfig.ChildDeviceConfig {
		childDevice := common.GetDeviceManager().GetDeviceById(childConfig.Id)
		if childDevice == nil {
			continue
		}

		// 获取子设备的遥测点位
		telemetryPoints := childDevice.GetTelemetryPoints()
		deviceName := childConfig.Name

		for _, point := range telemetryPoints {
			if point == nil || point.IsHidden() {
				continue
			}

			// 创建点位值，使用NewPointValue构造函数
			pointValue := c_base.NewPointValue(childConfig.Id, point, nil)

			// 转换为SSingleDeviceValue
			d := &entity.SSingleDeviceValue{}
			_ = gconv.Scan(pointValue, d)

			// 设置分组信息：设备名:汇总
			if d.Meta != nil {
				d.Meta.Name = fmt.Sprintf("%s:汇总", deviceName)
			}

			allValues = append(allValues, d)
		}
	}

	return &v1.GetVirtualDeviceCacheRes{
		DeviceServerState: "",
		AlarmLevel:        "",
		LastUpdateTime:    nil,
		Values:            allValues,
	}, nil
}
