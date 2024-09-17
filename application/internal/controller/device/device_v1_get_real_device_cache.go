package device

import (
	"application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/util"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error) {

	device := common.GetDeviceById(req.DeviceId)
	if device == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, req.DeviceId, "设备不存在")
	}
	lastUpdateTime := ""
	if device.GetLastUpdateTime() != nil {
		lastUpdateTime = device.GetLastUpdateTime().Format("2006-01-02 15:04:05")
	}
	res = &v1.GetRealDeviceCacheRes{
		DeviceId:       device.GetDeviceConfig().Id,
		DeviceType:     device.GetDriverType(),
		DeviceName:     device.GetDeviceConfig().Name,
		LastUpdateTime: lastUpdateTime,
	}
	var values = make([]*entity.SSingleDeviceValue, 0)

	for _, wrapper := range device.GetMetaValueList() {
		if len(req.TelemetryKeyList) != 0 && util.Contains(req.TelemetryKeyList, wrapper.Meta.Name) == false {
			continue
		}

		values = append(values, &entity.SSingleDeviceValue{
			DeviceId:   wrapper.DeviceId,
			DeviceType: wrapper.DeviceType,
			Meta:       wrapper.Meta,
			Value:      wrapper.Meta.ValueToString(wrapper.Value),
			HappenTime: wrapper.HappenTime.Format("2006-01-02 15:04:05"),
		})
	}

	res.Values = values
	return res, nil
}
