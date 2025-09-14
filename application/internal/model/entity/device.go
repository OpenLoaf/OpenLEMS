package entity

import (
	"common/c_base"
	"common/c_enum"
	"p_modbus"
	"reflect"

	"github.com/pkg/errors"

	"github.com/gogf/gf/v2/util/gconv"
)

type SDeviceTree struct {
	Id         string `json:"id,omitempty" dc:"设备ID"`
	Pid        string `json:"pid,omitempty" dc:"父设备Id"`
	Name       string `json:"name,omitempty" dc:"设备名称"`
	ProtocolId string `json:"protocolId,omitempty" dc:"协议配置ID,如果配置了肯定是实体设备"`
	Driver     string `json:"driver,omitempty" dc:"驱动名称"`
	LogLevel   string `json:"logLevel,omitempty" dc:"日志等级"`
	//Strategy           string         `json:"strategy,omitempty" dc:"策略名称"`
	ManualMode         bool           `json:"manualMode" dc:"手动模式"`
	StorageEnable      bool           `json:"StorageEnable" dc:"是否存储"`
	StorageIntervalSec int32          `json:"storageIntervalSec" dc:"存储间隔(秒),0代表默认1分钟，负数代表不存储"`
	Sort               int            `json:"sort" dc:"排序"`
	Enabled            bool           `json:"enabled" dc:"是否启用"`
	Params             map[string]any `json:"params,omitempty" dc:"设备参数"`
	CreatedAt          string         `json:"created_at" dc:"创建时间"`
	UpdatedAt          string         `json:"updated_at" dc:"更新时间"`

	ConfigFields    []*c_base.SConfigStructFields `json:"configFields" yaml:"fields" dc:"配置字段信息"`
	IsVirtualDevice bool                          `json:"isVirtualDevice" dc:"是否是虚拟设备"`
	DriverType      string                        `json:"driverType" dc:"驱动类型"`
	DriverBrand     string                        `json:"driverBrand,omitempty" dc:"驱动品牌"`
	DriverModel     string                        `json:"driverModel" yaml:"model" dc:"驱动型号"`
	DriverVersion   string                        `json:"driverVersion" yaml:"version" v:"required" dc:"驱动版本"`
	DriverTelemetry []*c_base.STelemetry          `json:"driverTelemetry" yaml:"telemetry" dc:"遥测"`
	DriverService   []*c_base.SDriverService      `json:"driverService" yaml:"customService" dc:"自定义服务"`

	ProtocolName    string `json:"protocolName" dc:"协议名称"`
	ProtocolType    string `json:"protocolType,omitempty" dc:"协议类型"`
	ProtocolAddress string `json:"protocolAddress" dc:"协议地址"`

	Children []*SDeviceTree `json:"children" dc:"子设备列表"`
}

func (t *SDeviceTree) UnmarshalValue(value interface{}) error {
	if record, ok := value.(*c_base.SDeviceConfig); ok {

		*t = SDeviceTree{
			Id:         record.Id,
			Pid:        record.Pid,
			Name:       record.Name,
			ProtocolId: record.ProtocolId,
			Driver:     record.Driver,
			LogLevel:   record.LogLevel,
			//Strategy:           record.Strategy,
			ManualMode:         record.ManualMode,
			StorageEnable:      record.StorageEnable,
			StorageIntervalSec: record.StorageIntervalSec,
			Sort:               record.Sort,
			Enabled:            record.Enabled,
			Params:             record.Params,
			CreatedAt:          record.CreatedAt,
			UpdatedAt:          record.UpdatedAt,
		}
		driverInfo := record.DriverInfo
		if driverInfo != nil {
			t.DriverModel = driverInfo.Model
			t.DriverBrand = driverInfo.Brand
			t.DriverVersion = driverInfo.Version
			t.DriverTelemetry = driverInfo.Telemetry
			t.DriverService = driverInfo.Service
			t.DriverType = string(driverInfo.Type)

			if driverInfo.GetConfigStructFields() != nil {
				t.ConfigFields = driverInfo.GetConfigStructFields()
			}
		}
		protocolConfig := record.ProtocolConfig
		if protocolConfig != nil {
			t.ProtocolName = protocolConfig.Name
			t.ProtocolType = string(protocolConfig.Type)
			t.ProtocolAddress = protocolConfig.Address
			if t.ConfigFields == nil {
				t.ConfigFields = make([]*c_base.SConfigStructFields, 0)
				switch protocolConfig.GetProtocol() {
				// 添加modbus的设备配置
				case c_enum.EModbusTcp, c_enum.EModbusRtu:
					fields := p_modbus.GetModbusDeviceConfigFields()
					if fields != nil {
						t.ConfigFields = append(t.ConfigFields, fields...)
					}
				}
			}
		} else {
			t.IsVirtualDevice = true
		}

		if len(record.ChildDeviceConfig) > 0 {
			children := make([]*SDeviceTree, 0, len(record.ChildDeviceConfig))
			for _, child := range record.ChildDeviceConfig {
				if child == nil {
					continue
				}
				var childTree *SDeviceTree
				if err := gconv.Struct(child, &childTree); err != nil {
					return err
				}
				children = append(children, childTree)
			}
			t.Children = children
		}
		return nil
	}
	return errors.Errorf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}
