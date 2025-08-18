package telemetry

import (
	"application/api/telemetry/v1"
	"context"
)

func (c *ControllerV1) GetTelemetryGet(ctx context.Context, req *v1.GetTelemetryGetReq) (res *v1.GetTelemetryGetRes, err error) {
	//instance := common.GetStorageInstance().GetDeviceById(req.DeviceId)
	//if instance == nil {
	//	return nil, gerror.NewCode(gcode.CodeNotFound)
	//}
	//
	//// 反射执行方法
	//des := instance.GetDriverDescription()
	//value, err := des.GetTelemetry(req.TelemetryKey, instance)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &v1.GetTelemetryGetRes{
	//	TestJoinKey:      req.DeviceId + ":" + req.TelemetryKey,
	//	DeviceId:         req.DeviceId,
	//	TelemetryKey:     req.TelemetryKey,
	//	TelemetryKeyName: g.I18n().T(ctx, req.TelemetryKey),
	//	Value:            value,
	//}, nil

	return nil, nil
}
