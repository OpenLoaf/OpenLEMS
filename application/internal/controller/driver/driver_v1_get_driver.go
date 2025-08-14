package driver

import (
	"context"
	"s_driver"

	v1 "application/api/driver/v1"
	"application/internal/model/entity"
)

func (c *ControllerV1) GetDriver(ctx context.Context, req *v1.GetDriverReq) (res *v1.GetDriverRes, err error) {
	driverManager := s_driver.GetDriverManager()
	info, err := driverManager.GetDriverInfo(ctx, req.DriverName)
	if err != nil {
		return nil, err
	}

	// 映射到响应实体
	sDriver := &entity.SDriver{
		DriverName:   info.Name,
		DriverType:   string(info.Type),
		DriverStatus: info.Available,
	}
	if info.Description != nil {
		sDriver.DriverVersion = info.Description.Version
		sDriver.DriverDescription = info.Description.Remark
		sDriver.DriverLastUpdate = info.Description.Create
		sDriver.ProtocolType = info.Description.ProtocolType
		// 扩展字段映射
		sDriver.Brand = info.Description.Brand
		sDriver.Model = info.Description.Model
		sDriver.BuildTime = info.Description.BuildTime
		sDriver.CommitHash = info.Description.CommitHash
		sDriver.Author = info.Description.Author
		if len(info.Description.Telemetry) > 0 {
			telemetry := make([]*entity.DriverTelemetry, 0, len(info.Description.Telemetry))
			for _, t := range info.Description.Telemetry {
				telemetry = append(telemetry, &entity.DriverTelemetry{
					Name:                t.Name,
					NationalizationName: t.NationalizationName,
					Unit:                t.Unit,
					Remark:              t.Remark,
				})
			}
			sDriver.Telemetry = telemetry
		}
	}

	return &v1.GetDriverRes{Driver: sDriver}, nil
}
