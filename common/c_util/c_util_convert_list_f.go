package c_util

import (
	"fmt"
	"time"
)

// ToBoolSlice 将 []any 转换为 []bool
func ToBoolSlice(val []any) ([]bool, error) {
	result := make([]bool, len(val))
	for i, item := range val {
		b, err := ToBool(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = b
	}
	return result, nil
}

// ToIntSlice 将 []any 转换为 []int
func ToIntSlice(val []any) ([]int, error) {
	result := make([]int, len(val))
	for i, item := range val {
		n, err := ToInt(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToInt64Slice 将 []any 转换为 []int64
func ToInt64Slice(val []any) ([]int64, error) {
	result := make([]int64, len(val))
	for i, item := range val {
		n, err := ToInt64(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToFloat64Slice 将 []any 转换为 []float64
func ToFloat64Slice(val []any) ([]float64, error) {
	result := make([]float64, len(val))
	for i, item := range val {
		f, err := ToFloat64(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = f
	}
	return result, nil
}

// ToStringSlice 将 []any 转换为 []string
func ToStringSlice(val []any) ([]string, error) {
	result := make([]string, len(val))
	for i, item := range val {
		s, err := ToString(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = s
	}
	return result, nil
}

// ToTimeSlice 将 []any 转换为 []time.Time
func ToTimeSlice(val []any) ([]time.Time, error) {
	result := make([]time.Time, len(val))
	for i, item := range val {
		t, err := ToTime(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = t
	}
	return result, nil
}

// ToInt8Slice 将 []any 转换为 []int8
func ToInt8Slice(val []any) ([]int8, error) {
	result := make([]int8, len(val))
	for i, item := range val {
		n, err := ToInt8(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToInt16Slice 将 []any 转换为 []int16
func ToInt16Slice(val []any) ([]int16, error) {
	result := make([]int16, len(val))
	for i, item := range val {
		n, err := ToInt16(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToInt32Slice 将 []any 转换为 []int32
func ToInt32Slice(val []any) ([]int32, error) {
	result := make([]int32, len(val))
	for i, item := range val {
		n, err := ToInt32(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToUintSlice 将 []any 转换为 []uint
func ToUintSlice(val []any) ([]uint, error) {
	result := make([]uint, len(val))
	for i, item := range val {
		n, err := ToUint(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToUint8Slice 将 []any 转换为 []uint8
func ToUint8Slice(val []any) ([]uint8, error) {
	result := make([]uint8, len(val))
	for i, item := range val {
		n, err := ToUint8(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToUint16Slice 将 []any 转换为 []uint16
func ToUint16Slice(val []any) ([]uint16, error) {
	result := make([]uint16, len(val))
	for i, item := range val {
		n, err := ToUint16(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToUint32Slice 将 []any 转换为 []uint32
func ToUint32Slice(val []any) ([]uint32, error) {
	result := make([]uint32, len(val))
	for i, item := range val {
		n, err := ToUint32(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToUint64Slice 将 []any 转换为 []uint64
func ToUint64Slice(val []any) ([]uint64, error) {
	result := make([]uint64, len(val))
	for i, item := range val {
		n, err := ToUint64(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = n
	}
	return result, nil
}

// ToFloat32Slice 将 []any 转换为 []float32
func ToFloat32Slice(val []any) ([]float32, error) {
	result := make([]float32, len(val))
	for i, item := range val {
		f, err := ToFloat32(item)
		if err != nil {
			return nil, fmt.Errorf("index %d: %v", i, err)
		}
		result[i] = f
	}
	return result, nil
}
