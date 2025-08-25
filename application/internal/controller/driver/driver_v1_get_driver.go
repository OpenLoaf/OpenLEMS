package driver

import (
	v1 "application/api/driver/v1"
	"application/internal/model/entity"
	"common"
	"context"
)

func (c *ControllerV1) GetDriver(ctx context.Context, req *v1.GetDriverReq) (res *v1.GetDriverRes, err error) {
	driverManager := common.GetDeviceManager()
	driverInfo, err := driverManager.GetDriverInfo(req.DriverName)
	if err != nil {
		return nil, err
	}

	// 映射到响应实体
	sDriver := &entity.SDriver{
		DriverName:   driverInfo.Name,
		DriverType:   string(driverInfo.Type),
		DriverStatus: true, // todo 修改成设备状态
	}

	sDriver.DriverVersion = driverInfo.Version
	sDriver.DriverDescription = driverInfo.Remark
	sDriver.DriverLastUpdate = driverInfo.Create
	sDriver.ProtocolType = string(driverInfo.ProtocolType)
	// 扩展字段映射
	sDriver.Brand = driverInfo.Brand
	sDriver.Model = driverInfo.Model
	sDriver.BuildTime = driverInfo.BuildTime
	sDriver.CommitHash = driverInfo.CommitHash
	sDriver.Author = driverInfo.Author
	if len(driverInfo.Telemetry) > 0 {
		telemetry := make([]*entity.DriverTelemetry, 0, len(driverInfo.Telemetry))
		for _, t := range driverInfo.Telemetry {
			telemetry = append(telemetry, &entity.DriverTelemetry{
				Name:        t.Name,
				DisplayName: t.DisplayName,
				Unit:        t.Unit,
				Remark:      t.Remark,
			})
		}
		sDriver.Telemetry = telemetry
	}
	return &v1.GetDriverRes{Driver: sDriver}, nil
}
