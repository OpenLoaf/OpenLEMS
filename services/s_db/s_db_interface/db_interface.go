package db_interface

import (
	"common/c_base"
	"context"
	"s_db/db_model"
)

// IDeviceService 设备服务
type IDeviceService interface {
	GetDeviceList(ctx context.Context) ([]*db_model.Device, error)
	CreateDevice(ctx context.Context, deviceName string, pId string) (string, error)
	DeleteDevice(ctx context.Context, deviceId string) error
}

type IConfigManage interface {
	GetDeviceConfig(ctx context.Context, activeDeviceRootId string) *c_base.SDriverConfig
	GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig
	GetSettingValueByName(ctx context.Context, name string) string
	SetSettingValueByName(ctx context.Context, name string, value string) error
}

type IProtocolManage interface {
	GetProtocolList(ctx context.Context, type_ string) ([]*db_model.Protocol, error)
	UpdateProtocol(ctx context.Context, protocolId string, data map[string]interface{}) error
	CreateProtocol(ctx context.Context, data map[string]interface{}) (string, error)
	DeleteProtocol(ctx context.Context, protocolId string) error
}
