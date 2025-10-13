package c_base

import (
	"strconv"

	"github.com/shockerli/cvt"
)

// SProtocolPoint 协议点位基础结构
type SProtocolPoint struct {
	*SPoint                       // 嵌套基础点位信息
	DataAccess   *SDataAccess     `json:"dataAccess" v:"required" dc:"数据访问配置"`        // 数据访问配置
	ValueExplain []*SFieldExplain `json:"valueExplain,omitempty" yaml:"valueExplain"` // 值解释
}

// GetDataAccess 获取数据访问配置
func (s *SProtocolPoint) GetDataAccess() *SDataAccess {
	return s.DataAccess
}

func (s *SProtocolPoint) GetValueExplain() []*SFieldExplain {
	if len(s.ValueExplain) == 0 {
		return s.SPoint.GetValueExplain()
	}
	return s.ValueExplain
}

// GetValueExplainByValue 获取值解释，优先使用自身的 ValueExplain，其次回退到嵌入的 SPoint 逻辑
func (s *SProtocolPoint) GetValueExplainByValue(value any) (string, error) {
	// 1. 将 value 转换为字符串（与 SPoint 行为保持一致）
	var valueStr string
	var err error

	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, bool, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *bool:
		valueStr, err = cvt.StringE(value)
		if err != nil {
			return "", err
		}
	default:
		intVal, err := cvt.IntE(value)
		if err != nil {
			return "", err
		}
		valueStr, err = cvt.StringE(intVal)
		if err != nil {
			return "", err
		}
	}

	// 2. 先查找自身或回退后的 ValueExplain（通过 GetValueExplain 统一获取）
	explains := s.GetValueExplain()
	if len(explains) > 0 {
		for _, explain := range explains {
			if explain.Key == valueStr {
				return explain.Value, nil
			}
		}
	}

	// 3. 浮点数据格式化输出（与 SPoint 行为保持一致）
	if floatVal, err := cvt.Float64E(value); err == nil {
		formatted := strconv.FormatFloat(floatVal, 'f', int(s.SPoint.Precise), 64)
		return formatted, nil
	}

	// 4. 无法转换为浮点数，返回转换后的字符串
	return valueStr, nil
}
