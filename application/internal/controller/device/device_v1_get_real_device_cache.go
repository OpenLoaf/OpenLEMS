package device

import (
	v1 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"common/c_enum"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error) {
	if common.GetDeviceManager().Status() == c_enum.EStateInit {
		// 系统还在初始化中
		return &v1.GetRealDeviceCacheRes{}, nil
	}

	device := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if device == nil {
		return &v1.GetRealDeviceCacheRes{}, nil
	}

	list := c_base.GetPointValueList(device)

	// 根据showTelemetryOnly参数决定是否追加遥测点位
	if !req.ShowTelemetryOnly {
		list = append(list, c_base.GetAllTelemetryPoint(device)...)
	}

	values := make([]*entity.SSingleDeviceValue, 0, len(list))
	for _, v := range list {
		if v.IPoint == nil || v.IsHidden() {
			continue
		}
		d := &entity.SSingleDeviceValue{}
		_ = gconv.Scan(v, d)
		if explain, err := c_base.ExplainPointValue(v.IPoint, v.GetValue()); err == nil && explain != "" {
			d.StatueExplain = explain
		}
		values = append(values, d)
	}

	return &v1.GetRealDeviceCacheRes{
		DeviceServerState: device.GetProtocolStatus().String(),
		AlarmLevel:        device.GetAlarmLevel().String(),
		LastUpdateTime:    device.GetLastUpdateTime(),
		Values:            values,
	}, nil
}
