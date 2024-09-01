package device

import (
	"application/api/device/v1"
	"application/internal/model/entity"
	"context"
	common "ems-plan"
)

func (c *ControllerV1) GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error) {
	var devices = make([]*entity.SDevice, 0)
	for _, driver := range common.DeviceInstance.FindAll() {
		lastUpdateTime := ""
		if driver.GetLastUpdateTime() != nil {
			lastUpdateTime = driver.GetLastUpdateTime().Format("2006-01-02 15:04:05")
		}
		devices = append(devices, &entity.SDevice{
			DeviceId:       driver.GetId(),
			DeviceType:     string(driver.GetType()),
			DeviceName:     driver.GetName(),
			IsMaster:       driver.GetIsMaster(),
			IsActivate:     driver.IsActivate(),
			LastUpdateTime: lastUpdateTime,
			AlarmLevel:     driver.GetAlarmLevel(),
		})
	}

	return &v1.GetRealDeviceListRes{Devices: devices}, nil
}
