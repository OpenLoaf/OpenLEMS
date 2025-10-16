package c_proto

import (
	"common/c_base"
	"common/c_enum"
)

// NewDidioPoint 创建基础 DIDO 点位
func NewDidioPoint(pin, chipIndex uint8, key, name string, valueType c_enum.EValueType, dataAccess *c_base.SDataAccess) *SDidioPoint {
	return &SDidioPoint{
		Pin:       pin,
		ChipIndex: chipIndex,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       key,
				Name:      name,
				ValueType: valueType,
			},
			DataAccess: dataAccess,
		},
	}
}

// NewDidioPointWithUnit 创建带单位的 DIDO 点位
func NewDidioPointWithUnit(pin, chipIndex uint8, key, name string, valueType c_enum.EValueType, unit string, dataAccess *c_base.SDataAccess) *SDidioPoint {
	return &SDidioPoint{
		Pin:       pin,
		ChipIndex: chipIndex,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       key,
				Name:      name,
				ValueType: valueType,
				Unit:      unit,
			},
			DataAccess: dataAccess,
		},
	}
}

// NewDidioPointWithDesc 创建带描述的 DIDO 点位
func NewDidioPointWithDesc(pin, chipIndex uint8, key, name string, valueType c_enum.EValueType, unit, desc string, dataAccess *c_base.SDataAccess) *SDidioPoint {
	return &SDidioPoint{
		Pin:       pin,
		ChipIndex: chipIndex,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       key,
				Name:      name,
				ValueType: valueType,
				Unit:      unit,
				Desc:      desc,
			},
			DataAccess: dataAccess,
		},
	}
}

// NewDidioPointFromPreset 使用预定义 SPoint 创建点位
func NewDidioPointFromPreset(pin, chipIndex uint8, preset *c_base.SPoint, dataAccess *c_base.SDataAccess) *SDidioPoint {
	return &SDidioPoint{
		Pin:       pin,
		ChipIndex: chipIndex,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     preset,
			DataAccess: dataAccess,
		},
	}
}

// DidioPointOption DIDO 点位选项函数类型
type DidioPointOption func(*SDidioPoint)

// NewDidioPointExt 使用选项模式创建 DIDO 点位（复杂场景）
func NewDidioPointExt(pin, chipIndex uint8, opts ...DidioPointOption) *SDidioPoint {
	point := &SDidioPoint{
		Pin:       pin,
		ChipIndex: chipIndex,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{},
		},
	}
	for _, opt := range opts {
		opt(point)
	}
	return point
}

// DIDO 选项函数定义
func WithDidioKey(key string) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint.Key = key
	}
}

func WithDidioName(name string) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint.Name = name
	}
}

func WithDidioValueType(valueType c_enum.EValueType) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint.ValueType = valueType
	}
}

func WithDidioUnit(unit string) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint.Unit = unit
	}
}

func WithDidioDesc(desc string) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint.Desc = desc
	}
}

func WithDidioDataAccess(dataAccess *c_base.SDataAccess) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.DataAccess = dataAccess
	}
}

func WithDidioValueExplain(explains []*c_base.SFieldExplain) DidioPointOption {
	return func(p *SDidioPoint) {
		p.ValueExplain = explains
	}
}

func WithDidioTrigger(trigger func(value interface{}) (bool, c_enum.EAlarmLevel, error)) DidioPointOption {
	return func(p *SDidioPoint) {
		p.Trigger = trigger
	}
}

func WithDidioGroup(group *c_base.SPointGroup) DidioPointOption {
	return func(p *SDidioPoint) {
		p.Group = group
	}
}

func WithDidioSort(sort int) DidioPointOption {
	return func(p *SDidioPoint) {
		p.Sort = sort
	}
}

func WithDidioMinMax(min, max int64) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint.Min = min
		p.SProtocolPoint.SPoint.Max = max
	}
}

func WithDidioPrecise(precise uint8) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint.Precise = precise
	}
}

func WithDidioHidden(hidden bool) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint.Hidden = hidden
	}
}

// 复合选项
func WithDidioPresetPoint(preset *c_base.SPoint) DidioPointOption {
	return func(p *SDidioPoint) {
		p.SProtocolPoint.SPoint = preset
	}
}
