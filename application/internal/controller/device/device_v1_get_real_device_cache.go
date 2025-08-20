package device

import (
	v1 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error) {

	deviceWrapper := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if deviceWrapper == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}
	deviceInstance := deviceWrapper.GetDeviceInstance()
	if deviceInstance == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}
	res = &v1.GetRealDeviceCacheRes{
		DeviceServerState: deviceWrapper.GetDeviceState().String(),
	}
	res.LastUpdateTime = deviceInstance.GetLastUpdateTime()

	//driverDescription := deviceInstance.GetDriverDescription()

	//for _, t := range driverDescription.Telemetry {
	//
	//}
	//
	//driverDescription.GetAllTelemetry(deviceInstance)

	res.Values = make([]*entity.SSingleDeviceValue, 0)

	for _, v := range deviceInstance.GetMetaValueList() {
		d := &entity.SSingleDeviceValue{}

		_ = gconv.Scan(v, d)
		if v.Meta.SystemType == c_base.SUseReadType {
			d.Meta.SystemType = d.Meta.ReadType
		}
		if v.Meta.StatusExplain != nil {
			d.StatueExplain = v.Meta.StatusExplain(v.Value)
		}

		res.Values = append(res.Values, d)
	}

	return res, nil
}
