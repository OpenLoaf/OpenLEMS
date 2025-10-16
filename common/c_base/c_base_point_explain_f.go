package c_base

import (
	"fmt"
	"strconv"

	"github.com/shockerli/cvt"
)

// ExplainPointValue 根据点位配置和值生成解释文本
func ExplainPointValue(point IPoint, value any) (string, error) {
	if point == nil {
		return "", nil
	}

	if point.GetKey() == "ChargeForbiddenMark" {
		fmt.Printf("")
	}

	// 获取点位的解释配置和精度
	explains := point.GetValueExplain()
	precise := point.GetPrecise()

	return explainByValueCommon(value, explains, precise)
}

// explainByValueCommon 公共的值解释逻辑：根据给定的 explains 列表匹配并返回解释
func explainByValueCommon(value any, explains []*SFieldExplain, precise uint8) (string, error) {
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
