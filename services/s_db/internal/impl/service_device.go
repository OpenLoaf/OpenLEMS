package impl

import (
	"context"
	"s_db/s_db_interface"
	"s_db/s_db_model"
	"sync"

	"time"

	"github.com/google/uuid"
)

type sDeviceServiceImpl struct {
	ctx context.Context
}

var (
	deviceManageInstance s_db_interface.IDeviceService
	deviceManageOnce     sync.Once
)

func GetDeviceService() s_db_interface.IDeviceService {
	deviceManageOnce.Do(func() {
		deviceManageInstance = &sDeviceServiceImpl{
			ctx: context.Background(),
		}
	})
	return deviceManageInstance
}

func (s *sDeviceServiceImpl) GetDeviceList(ctx context.Context) ([]*s_db_model.Device, error) {
	devices, err := s_db_model.GetAllDevicesOrderBySortAndEnable(ctx, true)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (s *sDeviceServiceImpl) CreateDevice(ctx context.Context, deviceName string, pId string) (string, error) {
	// 生成带横杠的UUID
	deviceId := uuid.NewString()
	device := &s_db_model.Device{
		Id:            deviceId,
		Pid:           pId,
		Name:          deviceName,
		Enable:        true,
		Sort:          0,
		LogLevel:      "INFO",
		RetentionDays: 30,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	err := device.Create(ctx)
	if err != nil {
		return "", err
	}
	return device.Id, nil
}

// DeleteDevice 删除设备
func (s *sDeviceServiceImpl) DeleteDevice(ctx context.Context, deviceId string) error {
	device := &s_db_model.Device{
		Id: deviceId,
	}
	err := device.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
