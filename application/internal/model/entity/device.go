package entity

import (
	"common"
	"common/c_base"
	"common/c_enum"
	"p_modbus"
	"reflect"

	"github.com/pkg/errors"

	"github.com/gogf/gf/v2/util/gconv"
)

type SDeviceTree struct {
	Id          string `json:"id,omitempty" dc:"设备ID"`
	Pid         string `json:"pid,omitempty" dc:"父设备Id"`
	Name        string `json:"name,omitempty" dc:"设备名称"`
	ProtocolId  string `json:"protocolId,omitempty" dc:"协议配置ID,如果配置了肯定是实体设备"`
	Driver      string `json:"driver,omitempty" dc:"驱动名称"`
	EnableDebug bool   `json:"enableDebug" dc:"启用调试模式"`
	//Strategy           string         `json:"strategy,omitempty" dc:"策略名称"`
	ManualMode         bool                   `json:"manualMode" dc:"手动模式"`
	StorageEnable      bool                   `json:"storageEnable" dc:"是否存储"`
	StorageIntervalSec int32                  `json:"storageIntervalSec" dc:"存储间隔(秒),0代表默认1分钟，负数代表不存储"`
	ExternalParam      *c_base.SExternalParam `json:"externalParam,omitempty" dc:"对外参数，存储JSON格式数据"`
	Sort               int                    `json:"sort" dc:"排序"`
	Enabled            bool                   `json:"enabled" dc:"是否启用"`
	Params             map[string]any         `json:"params,omitempty" dc:"设备参数"`
	CreatedAt          string                 `json:"created_at" dc:"创建时间"`
	UpdatedAt          string                 `json:"updated_at" dc:"更新时间"`

	ConfigFields    []*c_base.SFieldDefinition `json:"configFields" yaml:"fields" dc:"配置字段信息"`
	IsVirtualDevice bool                       `json:"isVirtualDevice" dc:"是否是虚拟设备"`
	DriverType      string                     `json:"driverType" dc:"驱动类型"`
	DriverBrand     string                     `json:"driverBrand,omitempty" dc:"驱动品牌"`
	DriverModel     string                     `json:"driverModel" yaml:"model" dc:"驱动型号"`
	DriverVersion   string                     `json:"driverVersion" yaml:"version" v:"required" dc:"驱动版本"`
	DriverTelemetry []*c_base.SFieldDefinition `json:"driverTelemetry" yaml:"telemetry" dc:"遥测"`
	DriverService   []*c_base.SDriverService   `json:"driverService" yaml:"customService" dc:"自定义服务"`

	ProtocolName    string `json:"protocolName" dc:"协议名称"`
	ProtocolType    string `json:"protocolType,omitempty" dc:"协议类型"`
	ProtocolAddress string `json:"protocolAddress" dc:"协议地址"`

	AlarmLevel     string `json:"alarmLevel" dc:"告警等级"`
	ProtocolStatus string `json:"protocolStatus" dc:"协议状态"`

	Children []*SDeviceTree `json:"children" dc:"子设备列表"`
}

func (t *SDeviceTree) UnmarshalValue(value interface{}) error {
	if record, ok := value.(*c_base.SDeviceConfig); ok {

		*t = SDeviceTree{
			Id:          record.Id,
			Pid:         record.Pid,
			Name:        record.Name,
			ProtocolId:  record.ProtocolId,
			Driver:      record.Driver,
			EnableDebug: record.EnableDebug,
			//Strategy:           record.Strategy,
			ManualMode:         record.ManualMode,
			StorageEnable:      record.StorageEnable,
			StorageIntervalSec: record.StorageIntervalSec,
			ExternalParam:      record.ExternalParam,
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
			// t.DriverTelemetry = driverInfo.Telemetry // 已移除Telemetry字段
			t.DriverService = driverInfo.Service
			t.DriverType = string(driverInfo.Type)

			// 处理驱动配置点位
			if driverInfo.ConfigPoints != nil {
				t.ConfigFields = make([]*c_base.SFieldDefinition, 0, len(driverInfo.ConfigPoints))
				for _, configPoint := range driverInfo.ConfigPoints {
					if configPoint != nil {
						fieldDef := configPoint.ToFieldDefinition()
						if fieldDef != nil {
							t.ConfigFields = append(t.ConfigFields, fieldDef)
						}
					}
				}
			}

		}
		protocolConfig := record.ProtocolConfig
		if protocolConfig != nil {
			t.ProtocolName = protocolConfig.Name
			t.ProtocolType = string(protocolConfig.Type)
			t.ProtocolAddress = protocolConfig.Address
			if t.ConfigFields == nil {
				t.ConfigFields = make([]*c_base.SFieldDefinition, 0)

				//c_base.BuildConfigStructFields()

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

		// 获取设备的告警等级和协议状态
		if device := common.GetDeviceManager().GetDeviceById(record.Id); device != nil {
			t.AlarmLevel = device.GetAlarmLevel().String()
			t.ProtocolStatus = device.GetProtocolStatus().String()

			// 从telemetry中获取所有遥测点位并转换为DriverTelemetry
			telemetryPoints := device.GetTelemetryPoints()
			if len(telemetryPoints) > 0 {
				t.DriverTelemetry = make([]*c_base.SFieldDefinition, 0, len(telemetryPoints))
				for _, point := range telemetryPoints {
					if point != nil {
						fieldDef := convertIPointToFieldDefinition(point)
						if fieldDef != nil {
							t.DriverTelemetry = append(t.DriverTelemetry, fieldDef)
						}
					}
				}
			}
		}

		return nil
	}
	return errors.Errorf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}

// convertIPointToFieldDefinition 将IPoint转换为SFieldDefinition
// 参考SConfigPoint.ToFieldDefinition()方法的实现逻辑
func convertIPointToFieldDefinition(point c_base.IPoint) *c_base.SFieldDefinition {
	if point == nil {
		return nil
	}

	// 转换值类型
	valueType := convertEValueTypeToConfigFieldsValueType(point.GetValueType())

	// 根据值类型推断组件类型
	componentType := inferComponentTypeFromValueType(valueType)

	// 处理指针类型字段
	var unit *string
	if unitStr := point.GetUnit(); unitStr != "" {
		unit = &unitStr
	}

	var min, max *int64
	if minVal := point.GetMin(); minVal != 0 {
		min = &minVal
	}
	if maxVal := point.GetMax(); maxVal != 0 {
		max = &maxVal
	}

	// 获取分组信息

	// 创建SFieldDefinition
	fieldDef := &c_base.SFieldDefinition{
		Key:           point.GetKey(),
		Name:          point.GetName(),
		Group:         point.GetGroup(),
		ValueType:     valueType,
		ComponentType: componentType,
		Unit:          unit,
		Min:           min,
		Max:           max,
		Description:   point.GetDesc(),
		Required:      false, // 遥测点位通常不是必填的
		ValueExplain:  point.GetValueExplain(),
	}

	return fieldDef
}

// convertEValueTypeToConfigFieldsValueType 将EValueType转换为EConfigFieldsValueType
func convertEValueTypeToConfigFieldsValueType(valueType c_enum.EValueType) c_enum.EConfigFieldsValueType {
	switch valueType {
	case c_enum.EBool:
		return c_enum.EConfigFieldsValueTypeBoolean
	case c_enum.EInt8, c_enum.EUint8, c_enum.EInt16, c_enum.EUint16, c_enum.EInt32, c_enum.EUint32, c_enum.EInt64, c_enum.EUint64:
		return c_enum.EConfigFieldsValueTypeInt
	case c_enum.EFloat32, c_enum.EFloat64:
		return c_enum.EConfigFieldsValueTypeFloat
	case c_enum.EString:
		return c_enum.EConfigFieldsValueTypeString
	default:
		return c_enum.EConfigFieldsValueTypeString
	}
}

// inferComponentTypeFromValueType 根据值类型推断组件类型
func inferComponentTypeFromValueType(valueType c_enum.EConfigFieldsValueType) c_enum.EConfigFieldsComponentType {
	switch valueType {
	case c_enum.EConfigFieldsValueTypeBoolean:
		return c_enum.EConfigFieldsComponentTypeSwitch
	case c_enum.EConfigFieldsValueTypeInt, c_enum.EConfigFieldsValueTypeFloat:
		return c_enum.EConfigFieldsComponentTypeNumber
	default:
		return c_enum.EConfigFieldsComponentTypeText
	}
}
