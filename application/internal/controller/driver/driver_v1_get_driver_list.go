package driver

import (
	"context"
	"s_driver"

	v1 "application/api/driver/v1"
	"application/internal/model/entity"
)

func (c *ControllerV1) GetDriverList(ctx context.Context, req *v1.GetDriverListReq) (res *v1.GetDriverListRes, err error) {
	driverManager := s_driver.GetDriverManager()
	driverList := driverManager.GetAllDriverNames()

	var driverInfoList []*entity.SDriver

	for _, driverName := range driverList {
		driverInfo, err := driverManager.GetDriverInfo(ctx, driverName)
		if err != nil {
			continue
		}

		// 转换为entity.SDriver结构
		sDriver := &entity.SDriver{
			DriverName:   driverInfo.Name,
			DriverType:   string(driverInfo.Type),
			DriverStatus: driverInfo.Available,
		}

		// 如果有描述信息，填充相关字段
		if driverInfo.Description != nil {
			sDriver.DriverVersion = driverInfo.Description.Version
			sDriver.DriverDescription = driverInfo.Description.Remark
			sDriver.DriverLastUpdate = driverInfo.Description.Create
			sDriver.ProtocolType = driverInfo.Description.ProtocolType
		}

		driverInfoList = append(driverInfoList, sDriver)
		// 按照驱动名称排序
		for i := 0; i < len(driverInfoList)-1; i++ {
			for j := i + 1; j < len(driverInfoList); j++ {
				if driverInfoList[i].DriverName > driverInfoList[j].DriverName {
					driverInfoList[i], driverInfoList[j] = driverInfoList[j], driverInfoList[i]
				}
			}
		}
	}

	return &v1.GetDriverListRes{
		DriverList: driverInfoList,
		Total:      len(driverInfoList),
	}, nil
}
