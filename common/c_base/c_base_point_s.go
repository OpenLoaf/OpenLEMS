package c_base

import (
	"common/c_enum"
	"strconv"

	"github.com/shockerli/cvt"
)

// SPoint 点位元数据
type SPoint struct {
	Key          string                                                                      `json:"key" v:"required" ` // 名称
	Name         string                                                                      `json:"name" v:"required"` // 名称
	ValueType    c_enum.EValueType                                                           `json:"value_type,omitempty" v:"required"`
	Group        *SPointGroup                                                                `json:"group" dc:"分组"`
	Unit         string                                                                      `json:"unit,omitempty"`             // 单位
	Desc         string                                                                      `json:"desc,omitempty"`             // 备注
	Sort         int                                                                         `json:"sort"`                       // 排序
	Min          int64                                                                       `json:"min,omitempty" dc:"正常范围最小值"` // 范围最小值,  默认为0
	Max          int64                                                                       `json:"max,omitempty" dc:"正常范围最大值"` // 范围最大值，默认为0
	Precise      uint8                                                                       `json:"precise,omitempty"`          // 设置浮点数精度（只是显示用）
	Trigger      func(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) `json:"-" dc:"告警触发函数"`
	Hidden       bool                                                                        `json:"hidden" dc:"是否显示"`
	ValueExplain []*SFieldExplain                                                            `json:"valueExplain,omitempty" yaml:"valueExplain"` // 值解释
}

func (s *SPoint) IsHidden() bool {
	return s.Hidden
}

func (s *SPoint) IsAlarmPoint() bool {
	return s.Trigger != nil
}

func (s *SPoint) GetValueType() c_enum.EValueType {
	return s.ValueType
}

func (s *SPoint) TriggerAlarm(value any) (trigger bool, level c_enum.EAlarmLevel, err error) {
	if s.Trigger != nil {
		return s.Trigger(value)
	}
	return false, c_enum.EAlarmLevelNone, nil
}

func (s *SPoint) GetValueExplain() []*SFieldExplain {
	if s.ValueExplain == nil {
		s.ValueExplain = []*SFieldExplain{}
	}
	return s.ValueExplain
}

// explainByValueCommon 公共的值解释逻辑：根据给定的 explains 列表匹配并返回解释
func (s *SPoint) explainByValueCommon(value any, explains []*SFieldExplain, precise uint8) (string, error) {
	// 1. 将value转换为字符串，如果是枚举之类的，转为int的字符串
	var valueStr string
	var err error

	// 检查值是否为数值类型（整数或浮点数）
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, bool, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *bool:
		// 数值类型直接转换为字符串
		valueStr, err = cvt.StringE(value)
		if err != nil {
			return "", err
		}
	default:
		// 非数值类型（如枚举）先转为int再转为字符串
		intVal, err := cvt.IntE(value)
		if err != nil {
			return "", err
		}
		valueStr, err = cvt.StringE(intVal)
		if err != nil {
			return "", err
		}
	}

	// 2. 从给定的 explains 中判断是否和value匹配
	if len(explains) > 0 {
		for _, explain := range explains {
			if explain.Key == valueStr {
				return explain.Value, nil
			}
		}
	}

	// 3. 浮点数据进行格式化输出
	if floatVal, err := cvt.Float64E(value); err == nil {
		formatted := strconv.FormatFloat(floatVal, 'f', int(precise), 64)
		return formatted, nil
	}

	// 如果无法转换为浮点数，返回转换后的字符串
	return valueStr, nil
}

func (s *SPoint) GetKey() string {
	return s.Key
}

func (s *SPoint) GetName() string {
	return s.Name
}

func (s *SPoint) GetGroup() *SPointGroup {
	return s.Group
}

func (s *SPoint) GetUnit() string {
	return s.Unit
}

func (s *SPoint) GetDesc() string {
	return s.Desc
}

func (s *SPoint) GetSort() int {
	return s.Sort
}

func (s *SPoint) GetMin() int64 {
	return s.Min
}

func (s *SPoint) GetMax() int64 {
	return s.Max
}

func (s *SPoint) GetPrecise() uint8 {
	return s.Precise
}

// IsProtocolPoint 判断是否为协议点位，SPoint 不是协议点位，返回 false
func (s *SPoint) IsProtocolPoint() bool {
	return false
}
