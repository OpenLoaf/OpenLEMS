package c_base

import (
	"strconv"

	"github.com/shockerli/cvt"
)

// ExplainPointValue 根据点位配置和值生成解释文本
func ExplainPointValue(point IPoint, value any) (string, error) {
	if point == nil {
		return "", nil
	}

	//if point.GetKey() == "ChargeForbiddenMark" {
	//	fmt.Printf("")
	//}

	// 获取点位的解释配置和精度
	explains := point.GetValueExplain()
	precise := point.GetPrecise()

	valueStr, _, err := ExplainValueWithColor(value, explains, precise)
	return valueStr, err
}

// ExplainValueWithColor 公共的值解释逻辑：根据给定的 explains 列表匹配并返回解释和颜色
// 返回值：(解释文本, 颜色代码, 错误)
func ExplainValueWithColor(value any, explains []*SFieldExplain, precise uint8) (string, string, error) {
	// 1. 判断是否是浮点数类型，如果是则直接格式化返回（浮点数不会有explain）
	if floatVal, err := cvt.Float64E(value); err == nil {
		// 检查是否确实是浮点数类型（排除整数和布尔值）
		switch value.(type) {
		case float32, float64, *float32, *float64:
			formatted := strconv.FormatFloat(floatVal, 'f', int(precise), 64)
			return formatted, "", nil
		}
	}

	// 2. 将value转换为字符串，如果是枚举之类的，转为int的字符串
	var valueStr string
	var err error

	// 检查值是否为数值类型（整数或浮点数）
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, bool, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *bool:
		// 数值类型直接转换为字符串
		valueStr, err = cvt.StringE(value)
		if err != nil {
			return "", "", err
		}
	default:
		// 非数值类型（如枚举）先转为int再转为字符串
		intVal, err := cvt.IntE(value)
		if err != nil {
			return "", "", err
		}
		valueStr, err = cvt.StringE(intVal)
		if err != nil {
			return "", "", err
		}
	}

	// 3. 从给定的 explains 中判断是否和value匹配
	if len(explains) > 0 {
		for _, explain := range explains {
			if explain.Key == valueStr {
				return explain.Value, explain.Color, nil
			}
		}
	}

	// 如果没有匹配的解释，返回转换后的字符串
	return valueStr, "", nil
}
