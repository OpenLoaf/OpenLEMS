package c_util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// ToBool 将 any 转换为 bool
func ToBool(val any) (bool, error) {
	switch v := val.(type) {
	case bool:
		return v, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(val).Int() != 0, nil
	case float32, float64:
		return reflect.ValueOf(val).Float() != 0, nil
	case string:
		return strconv.ParseBool(v)
	case json.Number:
		// 尝试转换为 int64
		if intVal, err := v.Int64(); err == nil {
			return intVal != 0, nil
		}
		// 尝试转换为 float64
		if floatVal, err := v.Float64(); err == nil {
			return floatVal != 0, nil
		}
		return false, fmt.Errorf("cannot convert json.Number %s to bool", string(v))
	default:
		return false, fmt.Errorf("cannot convert %T to bool", val)
	}
}

// ToInt 将 any 转换为 int
func ToInt(val any) (int, error) {
	v, err := ToInt64(val)
	return int(v), err
}

// ToInt8 将 any 转换为 int8
func ToInt8(val any) (int8, error) {
	v, err := ToInt64(val)
	return int8(v), err
}

// ToInt16 将 any 转换为 int16
func ToInt16(val any) (int16, error) {
	v, err := ToInt64(val)
	return int16(v), err
}

// ToInt32 将 any 转换为 int32
func ToInt32(val any) (int32, error) {
	v, err := ToInt64(val)
	return int32(v), err
}

// ToInt64 将 any 转换为 int64
func ToInt64(val any) (int64, error) {
	switch v := val.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	case json.Number:
		// 首先尝试转换为 int64
		if intVal, err := v.Int64(); err == nil {
			return intVal, nil
		}
		// 如果失败，尝试转换为 float64 然后转为 int64
		if floatVal, err := v.Float64(); err == nil {
			return int64(floatVal), nil
		}
		// 最后尝试直接解析字符串
		return strconv.ParseInt(string(v), 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to int64", val)
	}
}

// ToUint 将 any 转换为 uint
func ToUint(val any) (uint, error) {
	v, err := ToUint64(val)
	return uint(v), err
}

// ToUint8 将 any 转换为 uint8
func ToUint8(val any) (uint8, error) {
	v, err := ToUint64(val)
	return uint8(v), err
}

// ToUint16 将 any 转换为 uint16
func ToUint16(val any) (uint16, error) {
	v, err := ToUint64(val)
	return uint16(v), err
}

// ToUint32 将 any 转换为 uint32
func ToUint32(val any) (uint32, error) {
	v, err := ToUint64(val)
	return uint32(v), err
}

// ToUint64 将 any 转换为 uint64
func ToUint64(val any) (uint64, error) {
	switch v := val.(type) {
	case int:
		return uint64(v), nil
	case int8:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case float32:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseUint(v, 10, 64)
	case json.Number:
		// 首先尝试转换为 int64，然后转为 uint64
		if intVal, err := v.Int64(); err == nil {
			return uint64(intVal), nil
		}
		// 如果 int64 转换失败，尝试转换为 float64 然后转为 uint64
		if floatVal, err := v.Float64(); err == nil {
			return uint64(floatVal), nil
		}
		// 最后尝试直接解析字符串为 uint64
		return strconv.ParseUint(string(v), 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to uint64", val)
	}
}

// ToFloat32 将 any 转换为 float32
func ToFloat32(val any) (float32, error) {
	v, err := ToFloat64(val)
	return float32(v), err
}

// ToFloat64 将 any 转换为 float64
func ToFloat64(val any) (float64, error) {
	switch v := val.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseFloat(v, 64)
	case json.Number:
		return v.Float64()
	default:
		return 0, fmt.Errorf("cannot convert %T to float64", val)
	}
}

// ToString 将 any 转换为 string
func ToString(val any) (string, error) {
	switch v := val.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		return strconv.FormatBool(v), nil
	case time.Time:
		return v.String(), nil
	case error:
		return v.Error(), nil
	case json.Number:
		return string(v), nil
	case fmt.Stringer:
		return v.String(), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}

// ToTime 将 any 转换为 time.Time
func ToTime(val any) (time.Time, error) {
	switch v := val.(type) {
	case time.Time:
		return v, nil
	case string:
		// 尝试常见的时间格式
		formats := []string{
			time.RFC3339,
			"2006-01-02 15:04:05",
			"2006-01-02",
			time.RFC1123,
		}

		for _, format := range formats {
			if t, err := time.Parse(format, v); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("cannot parse time from string: %s", v)
	case int, int64:
		// 假设是Unix时间戳
		return time.Unix(reflect.ValueOf(val).Int(), 0), nil
	default:
		return time.Time{}, fmt.Errorf("cannot convert %T to time.Time", val)
	}
}
