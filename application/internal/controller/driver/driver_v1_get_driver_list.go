package driver

import (
	v1 "application/api/driver/v1"
	"application/internal/model/entity"
	"common"
	"context"
)

func (c *ControllerV1) GetDriverList(ctx context.Context, req *v1.GetDriverListReq) (res *v1.GetDriverListRes, err error) {
	driverManager := common.GetDeviceManager()
	driverList := driverManager.GetAllDriversInfo()

	var driverInfoList []*entity.SDriver

	for _, driverInfo := range driverList {
		//driverInfo, err := driverManager.GetDriverInfo(driver.Name)
		//if err != nil {
		//	continue
		//}

		// 转换为entity.SDriver结构
		sDriver := &entity.SDriver{
			DriverName: driverInfo.Name,
			DriverType: string(driverInfo.Type),
			//DriverStatus:       driverInfo.Available,
			//DriverIsEmbedded:   driverInfo.Embedded,
			DriverPath:         driverInfo.Path,
			DriverHash:         driverInfo.HashCode,
			DriverFileSizeByte: driverInfo.FileSizeByte,
		}

		// 如果有描述信息，填充相关字段
		sDriver.DriverVersion = driverInfo.Version
		sDriver.DriverDescription = driverInfo.Remark
		sDriver.DriverLastUpdate = driverInfo.Create
		sDriver.ProtocolType = string(driverInfo.ProtocolType)
		sDriver.CommitHash = driverInfo.CommitHash
		sDriver.Remark = driverInfo.Remark
		sDriver.Brand = driverInfo.Brand
		sDriver.Model = driverInfo.Model
		sDriver.DriverImage = driverInfo.Image

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
