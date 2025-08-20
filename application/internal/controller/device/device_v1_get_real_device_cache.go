package device

import (
	v1 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error) {

	deviceWrapper := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	deviceInstance := deviceWrapper.GetDeviceInstance()

	if deviceInstance == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}
	res = &v1.GetRealDeviceCacheRes{}
	res.LastUpdateTime = deviceInstance.GetLastUpdateTime()

	res.Values = make([]*entity.SSingleDeviceValue, 0)

	for _, v := range deviceInstance.GetMetaValueList() {
		d := &entity.SSingleDeviceValue{}
		_ = gconv.Scan(v, d)
		res.Values = append(res.Values, d)
	}

	return res, nil
}
