package device

import (
	"application/api/device/v1"
	"application/internal/model/entity"
	"context"
	common "ems-plan"
)

func (c *ControllerV1) GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error) {
	var devices = make([]*entity.SDevice, 0)
	var isVirtual []bool
	switch req.ShowType {
	case 1:
		isVirtual = append(isVirtual, false)
	case 2:
		isVirtual = append(isVirtual, true)
	}

	for _, driver := range common.GetDeviceAll(isVirtual...) {
		lastUpdateTime := ""
		if driver.GetLastUpdateTime() != nil {
			lastUpdateTime = driver.GetLastUpdateTime().Format("2006-01-02 15:04:05")
		}
		devices = append(devices, &entity.SDevice{
			DeviceId:       driver.GetDeviceConfig().Id,
			DeviceType:     string(driver.GetDriverType()),
			DeviceName:     driver.GetDeviceConfig().Name,
			IsMaster:       driver.GetDeviceConfig().IsMaster,
			LastUpdateTime: lastUpdateTime,
			IsVirtual:      driver.GetDeviceConfig().IsVirtual,
			AlarmLevel:     driver.GetAlarmLevel(),
		})
	}

	return &v1.GetRealDeviceListRes{Devices: devices}, nil
}
