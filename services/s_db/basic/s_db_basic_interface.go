package basic

import (
	"common/c_base"
	"context"
	"s_db/model"
)

// IDeviceService 设备服务
type IDeviceService interface {
	GetDeviceConfigs(ctx context.Context, parentId string) ([]*c_base.SDeviceConfig, error) // 获取所有设备列表
	GetRootDeviceId() string                                                                // 获取根设备ID
	//GetDeviceById(ctx context.Context, deviceId string) *model.SDeviceModel
}

type ISettingService interface {
	GetSettingValueByKey(ctx context.Context, name string, defaultValue ...string) string // 获取设置，如果获取不到，就设置为默认值
	SetSettingValueByName(ctx context.Context, name string, value string) error
}

type IProtocolService interface {
	GetProtocolList(ctx context.Context, type_ string) ([]*model.SProtocolModel, error)
	UpdateProtocol(ctx context.Context, protocolId string, data map[string]interface{}) error
	CreateProtocol(ctx context.Context, data map[string]interface{}) (string, error)
	DeleteProtocol(ctx context.Context, protocolId string) error
	GetAllProtocolConfigs(ctx context.Context) ([]*c_base.SProtocolConfig, error) // 获取协议列表
}
