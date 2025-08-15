package s_db_interface

import (
	"common/c_base"
	"context"
	"s_db/s_db_model"
)

// IDeviceService 设备服务
type IDeviceService interface {
	GetDeviceList(ctx context.Context, parentId string) ([]*s_db_model.SDeviceModel, error)       // 获取所有设备列表
	CreateDevice(ctx context.Context, deviceName string, pId string) (string, error)              // 创建设备
	DeleteDevice(ctx context.Context, deviceId string) error                                      // 删除设备
	GetDeviceById(ctx context.Context, id string) (*s_db_model.SDeviceModel, error)               // 根据ID 获取设备
	GetRecursiveDevicesByPid(ctx context.Context, pid string) ([]*s_db_model.SDeviceModel, error) // 递归获取该设备ID下的所有设备
	GetRootDeviceId() (string, error)                                                             // 获取根设备ID
}

type IConfigService interface {
	GetDeviceConfig(ctx context.Context, activeDeviceRootId string) *c_base.SDriverConfig
	GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig
	GetSettingValueByName(ctx context.Context, name string) string
	SetSettingValueByName(ctx context.Context, name string, value string) error
}

type IProtocolService interface {
	GetProtocolList(ctx context.Context, type_ string) ([]*s_db_model.SProtocolModel, error)
	UpdateProtocol(ctx context.Context, protocolId string, data map[string]interface{}) error
	CreateProtocol(ctx context.Context, data map[string]interface{}) (string, error)
	DeleteProtocol(ctx context.Context, protocolId string) error
	GetAllProtocols(ctx context.Context) ([]*s_db_model.SProtocolModel, error) // 获取协议列表
}
