package service

import (
	"context"
	"sqlite/model"
)

type IDeviceManage interface {
	GetDeviceList(ctx context.Context) ([]*model.Device, error)
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
