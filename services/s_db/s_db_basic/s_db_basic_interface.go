package s_db_basic

import (
	"common/c_base"
	"context"
	"s_db/s_db_model"
)

// IDeviceService 设备服务
type IDeviceService interface {
	GetEnableDeviceConfigsWithRecursion(ctx context.Context, parentId string) ([]*c_base.SDeviceConfig, error) // 获取所有设备列表
	GetRootDeviceId() string                                                                                   // 获取根设备ID
	GetAllDevices(ctx context.Context) ([]*s_db_model.SDeviceModel, error)                                     // 获取所有设备
	UpdateDevice(ctx context.Context, deviceId string, data map[string]interface{}) error
}

type ISettingService interface {
	GetAllSettings(ctx context.Context) ([]*s_db_model.SSettingModel, error)
	GetSettingValueByKey(ctx context.Context, name string, defaultValue ...string) string // 获取设置，如果获取不到，就设置为默认值
	SetSettingValueByName(ctx context.Context, name string, value string) error
}

type IProtocolService interface {
	GetProtocolList(ctx context.Context, type_ string) ([]*s_db_model.SProtocolModel, error)
	UpdateProtocol(ctx context.Context, protocolId string, data map[string]interface{}) error
	CreateProtocol(ctx context.Context, data map[string]interface{}) (string, error)
	DeleteProtocol(ctx context.Context, protocolId string) error
	GetAllProtocolConfigs(ctx context.Context) ([]*c_base.SProtocolConfig, error) // 获取协议列表
}
