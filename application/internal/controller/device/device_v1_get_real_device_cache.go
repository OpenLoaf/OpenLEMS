package device

import (
	v1 "application/api/device/v1"
	"context"
)

func (c *ControllerV1) GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error) {
	//device := common.GetStorageInstance().GetDeviceById(req.DeviceId)
	//if device == nil {
	//	return nil, gerror.NewCode(gcode.CodeInvalidParameter, req.DeviceId, "设备不存在")
	//}
	//lastUpdateTime := ""
	//if device.GetLastUpdateTime() != nil {
	//	lastUpdateTime = device.GetLastUpdateTime().Format("2006-01-02 15:04:05")
	//}
	//res = &v1.GetRealDeviceCacheRes{
	//	DeviceId:       device.GetDeviceConfig().Id,
	//	DeviceType:     device.GetDriverType(),
	//	DeviceName:     device.GetDeviceConfig().Name,
	//	LastUpdateTime: lastUpdateTime,
	//}
	//var _ = make([]*entity.SSingleDeviceValue, 0)
	//
	//for _, wrapper := range device.GetMetaValueList() {
	//	if len(req.TelemetryKeyList) != 0 && c_util.Contains(req.TelemetryKeyList, wrapper.Meta.Name) == false {
	//		continue
	//	}
	//
	//	values = append(values, &entity.SSingleDeviceValue{
	//		DeviceId:   wrapper.DeviceId,
	//		DeviceType: wrapper.DeviceType,
	//		Meta:       wrapper.Meta,
	//		Value:      wrapper.Meta.ValueToString(wrapper.Value),
	//		HappenTime: wrapper.HappenTime.Format("2006-01-02 15:04:05"),
	//	})
	//}
	//
	//res.Values = values
	return res, nil
}
