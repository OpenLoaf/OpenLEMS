package c_base

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SDriverConfig 基础设备配置
type SDriverConfig struct {
	gmeta.Meta         `orm:"table:device"`
	Id                 string            `json:"id,omitempty" orm:"id"`                            // 设备ID
	Name               string            `json:"name,omitempty" orm:"name"`                        // 设备名称
	ProtocolId         string            `json:"protocolId,omitempty" orm:"protocolId"`            // 协议配置ID,如果配置了肯定是实体设备
	Driver             string            `json:"driver,omitempty" orm:"driver"`                    // 驱动名称
	IsEnable           bool              `json:"enable" orm:"isEnable"`                            // 是否启用
	StorageEnable      bool              `json:"StorageEnable" orm:"StorageEnable"`                // 是否存储
	StorageIntervalSec int32             `json:"storageIntervalSec" orm:"storageIntervalSec"`      // 存储间隔(秒),0代表默认1分钟，负数代表不存储
	LogLevel           string            `json:"logLevel,omitempty" orm:"logLevel"`                // 日志等级
	Params             map[string]string `json:"params,omitempty" orm:"params"`                    // 额外参数
	DeviceChildren     []*SDriverConfig  `json:"deviceChildren,omitempty" orm:"with:parent_id=id"` // 子设备
}

func (s *SDriverConfig) Check() error {
	if s.Id == "" {
		return gerror.Newf("设备ID不能为空")
	}
	if s.Name == "" {
		return gerror.Newf("设备名称不能为空")
	}
	return nil
}

func (s *SDriverConfig) IsVirtual() bool {
	return s.ProtocolId == ""
}

// 现在暂时从配置文件中读取，后续可以从数据库中读取
//func (s *SDriverConfig) UnmarshalValue(value interface{}) error {
//	if record, ok := value.(gdb.Record); ok {
//		*s = SDriverConfig{
//			Id: record["id"].String(),
//			//Name:               record[],
//			ProtocolId:         "",
//			Driver:             "",
//			IsEnable:           false,
//			StorageIntervalSec: 0,
//			LogLevel:           "",
//			Params:             nil,
//			DeviceChildren:     nil,
//		}
//		return nil
//	}
//	return gerror.Newf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
//}
