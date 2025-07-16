package service

import (
	"context"
	"sqlite/model"

	"time"

	"github.com/google/uuid"
)

type IDeviceManage interface {
	GetDeviceList(ctx context.Context) ([]*model.Device, error)
	CreateDevice(ctx context.Context, deviceName string, pId string) (string, error)
	DeleteDevice(ctx context.Context, deviceId string) error
}

type sDeviceManage struct {
}

func NewDeviceManage(ctx context.Context) IDeviceManage {
	return &sDeviceManage{}
}

func (s *sDeviceManage) GetDeviceList(ctx context.Context) ([]*model.Device, error) {
	devices, err := model.GetAllDevicesOrderBySortAndEnable(ctx, true)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (s *sDeviceManage) CreateDevice(ctx context.Context, deviceName string, pId string) (string, error) {
	// 生成带横杠的UUID
	deviceId := uuid.NewString()
	device := &model.Device{
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
	device := &model.Device{
		Id: deviceId,
	}
	err := device.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
