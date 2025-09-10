package c_base

import (
	"common/c_enum"
	"encoding/json"

	"github.com/pkg/errors"
)

type SDeviceConfig struct { // 设备配置
	Id         string `json:"id,omitempty" orm:"id"`                  // 设备ID
	Pid        string `json:"pid,omitempty" orm:"pid"`                // 父设备Id
	Name       string `json:"name,omitempty" orm:"name"`              // 设备名称
	ProtocolId string `json:"protocolId,omitempty" orm:"protocol_id"` // 协议配置ID,如果配置了肯定是实体设备
	Driver     string `json:"driver,omitempty" orm:"driver"`          // 驱动名称

	LogLevel           string         `json:"logLevel,omitempty" orm:"log_level"`            // 日志等级
	Strategy           string         `json:"strategy,omitempty" orm:"strategy"`             // 	策略名称
	StorageEnable      bool           `json:"StorageEnable" orm:"storage_enable"`            // 是否存储
	StorageIntervalSec int32          `json:"storageIntervalSec" orm:"storage_interval_sec"` // 存储间隔(秒),0代表默认1分钟，负数代表不存储
	Sort               int            `json:"sort" orm:"sort"`
	Enabled            bool           `json:"enabled" orm:"enabled"`         // 是否启用
	Params             map[string]any `json:"params,omitempty" orm:"params"` // 额外参数
	CreatedAt          string         `json:"created_at" orm:"created_at"`
	UpdatedAt          string         `json:"updated_at" orm:"updated_at"`

	// 后面的都是动态更新的数据
	DriverInfo        *SDriverInfo     `json:"driverInfo,omitempty" orm:"driver_info"`         // 驱动信息
	ProtocolConfig    *SProtocolConfig `json:"protocolConfig,omitempty" orm:"protocol_config"` // todo 协议配置
	ChildDeviceConfig []*SDeviceConfig `json:"childDeviceConfig,omitempty" orm:"child_device_config"`
	//InitStatus Status
	FailedMessage string `json:"failedMessage,omitempty" orm:"failed_message" dc:"初始化失败原因"`
}

func (s *SDeviceConfig) ScanParams(target any) error {
	if target == nil {
		return errors.New("config target cannot be nil")
	}
	if s.Params == nil {
		return nil
	}

	// 先将 map[string]string 转换为 JSON 字节
	jsonBytes, err := json.Marshal(s.Params)
	if err != nil {
		return errors.Errorf("failed to marshal params to json: %+v", err)
	}

	// 再将 JSON 字节转换为目标结构体
	if err := json.Unmarshal(jsonBytes, target); err != nil {
		return errors.Errorf("failed to unmarshal json to target struct: %+v", err)
	}

	return nil
}

func (s *SDeviceConfig) GetType() c_enum.EDeviceType {
	if s.DriverInfo == nil {
		return c_enum.EDeviceNone
	}
	return s.DriverInfo.Type
}

//func (s *SDeviceConfig) GetTelemetry(key string, instance any) (any, error) {
//	if s.DriverInfo == nil {
//		return nil, errors.New("device not found")
//	}
//	return s.DriverInfo.GetTelemetry(key, instance)
//}

func (s *SDeviceConfig) GetAllTelemetry(instance IDevice) map[string]any {
	if s.DriverInfo == nil {
		return nil
	}
	return s.DriverInfo.GetAllTelemetry(instance)
}

func (s *SDeviceConfig) ExecuteCustomService(functionName string, instance IDevice, params any) error {
	if s.DriverInfo == nil {
		return nil
	}
	return s.DriverInfo.ExecuteCustomService(functionName, instance, params)
}
