package impl

import (
	"context"
	"database/db_model"

	"time"

	"github.com/google/uuid"
)

type sDeviceManage struct {
}

func NewDeviceManage(ctx context.Context) *sDeviceManage {
	return &sDeviceManage{}
}

func (s *sDeviceManage) GetDeviceList(ctx context.Context) ([]*db_model.Device, error) {
	devices, err := db_model.GetAllDevicesOrderBySortAndEnable(ctx, true)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (s *sDeviceManage) CreateDevice(ctx context.Context, deviceName string, pId string) (string, error) {
	// 生成带横杠的UUID
	deviceId := uuid.NewString()
	device := &db_model.Device{
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
func (s *sDeviceManage) DeleteDevice(ctx context.Context, deviceId string) error {
	device := &db_model.Device{
		Id: deviceId,
	}
	err := device.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
