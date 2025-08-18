package c_util

import (
	"fmt"
	"time"
)

// ToBoolFirst 将 []any 转换为 bool，返回第一个值
func ToBoolFirst(val []any) (bool, error) {
	if len(val) == 0 {
		return false, fmt.Errorf("empty slice")
	}
	return ToBool(val[0])
}

// ToIntFirst 将 []any 转换为 int，返回第一个值
func ToIntFirst(val []any) (int, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToInt(val[0])
}

// ToInt8First 将 []any 转换为 int8，返回第一个值
func ToInt8First(val []any) (int8, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToInt8(val[0])
}

// ToInt16First 将 []any 转换为 int16，返回第一个值
func ToInt16First(val []any) (int16, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToInt16(val[0])
}

// ToInt32First 将 []any 转换为 int32，返回第一个值
func ToInt32First(val []any) (int32, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToInt32(val[0])
}

// ToInt64First 将 []any 转换为 int64，返回第一个值
func ToInt64First(val []any) (int64, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToInt64(val[0])
}

// ToUintFirst 将 []any 转换为 uint，返回第一个值
func ToUintFirst(val []any) (uint, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToUint(val[0])
}

// ToUint8First 将 []any 转换为 uint8，返回第一个值
func ToUint8First(val []any) (uint8, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToUint8(val[0])
}

// ToUint16First 将 []any 转换为 uint16，返回第一个值
func ToUint16First(val []any) (uint16, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToUint16(val[0])
}

// ToUint32First 将 []any 转换为 uint32，返回第一个值
func ToUint32First(val []any) (uint32, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToUint32(val[0])
}

// ToUint64First 将 []any 转换为 uint64，返回第一个值
func ToUint64First(val []any) (uint64, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToUint64(val[0])
}

// ToFloat32First 将 []any 转换为 float32，返回第一个值
func ToFloat32First(val []any) (float32, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToFloat32(val[0])
}

// ToFloat64First 将 []any 转换为 float64，返回第一个值
func ToFloat64First(val []any) (float64, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	return ToFloat64(val[0])
}

// ToStringFirst 将 []any 转换为 string，返回第一个值
func ToStringFirst(val []any) (string, error) {
	if len(val) == 0 {
		return "", fmt.Errorf("empty slice")
	}
	return ToString(val[0])
}

// ToTimeFirst 将 []any 转换为 time.Time，返回第一个值
func ToTimeFirst(val []any) (time.Time, error) {
	if len(val) == 0 {
		return time.Time{}, fmt.Errorf("empty slice")
	}
	return ToTime(val[0])
}
