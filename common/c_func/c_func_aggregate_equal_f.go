package c_func

import (
	"common/c_base"
	"errors"
	"reflect"

	"github.com/shockerli/cvt"
)

// 判断所有值是否相等（数字类型）
func equalAggregate[T c_base.Number](values []any, convert func(interface{}) (T, error)) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, errors.New("empty values")
	}

	var firstValue T
	var isFirst = true
	var err error

	for _, v := range values {
		converted, e := convert(v)
		if e != nil {
			err = e
			break
		} else {
			if isFirst {
				firstValue = converted
				isFirst = false
			} else if converted != firstValue {
				var zero T
				return zero, errors.New("values are not equal")
			}
		}
	}
	if err != nil {
		var zero T
		return zero, err
	}
	return firstValue, nil
}

// 通用相等性判断函数（支持任意类型）
func equalAggregateAny(values []any) (any, error) {
	if len(values) == 0 {
		return nil, errors.New("empty values")
	}

	firstValue := values[0]
	for i := 1; i < len(values); i++ {
		if !reflect.DeepEqual(firstValue, values[i]) {
			return nil, errors.New("values are not equal")
		}
	}
	return firstValue, nil
}
func EqualAggregate[T any](values []any) (T, error) {
	var none T
	if len(values) == 0 {
		return none, errors.New("empty values")
	}

	firstValue := values[0]
	for i := 1; i < len(values); i++ {
		if !reflect.DeepEqual(firstValue, values[i]) {
			return none, errors.New("values are not equal")
		}
	}
	if value, ok := firstValue.(T); ok {
		return value, nil
	}
	return none, errors.New("first values are not equal")
}

// 通用相等性判断函数
var AggregateEqual = func(values []any) (any, error) {
	return equalAggregateAny(values)
}

// 整数类型
var AggregateEqualInt = func(values []any) (int, error) {
	return equalAggregate[int](values, cvt.IntE)
}

var AggregateEqualInt8 = func(values []any) (int8, error) {
	return equalAggregate[int8](values, cvt.Int8E)
}

var AggregateEqualInt16 = func(values []any) (int16, error) {
	return equalAggregate[int16](values, cvt.Int16E)
}

var AggregateEqualInt32 = func(values []any) (int32, error) {
	return equalAggregate[int32](values, cvt.Int32E)
}

var AggregateEqualInt64 = func(values []any) (int64, error) {
	return equalAggregate[int64](values, cvt.Int64E)
}

// 无符号整数类型
var AggregateEqualUint = func(values []any) (uint, error) {
	return equalAggregate[uint](values, cvt.UintE)
}

var AggregateEqualUint8 = func(values []any) (uint8, error) {
	return equalAggregate[uint8](values, cvt.Uint8E)
}

var AggregateEqualUint16 = func(values []any) (uint16, error) {
	return equalAggregate[uint16](values, cvt.Uint16E)
}

var AggregateEqualUint32 = func(values []any) (uint32, error) {
	return equalAggregate[uint32](values, cvt.Uint32E)
}

var AggregateEqualUint64 = func(values []any) (uint64, error) {
	return equalAggregate[uint64](values, cvt.Uint64E)
}

// 浮点数类型
var AggregateEqualFloat32 = func(values []any) (float32, error) {
	return equalAggregate[float32](values, cvt.Float32E)
}

var AggregateEqualFloat64 = func(values []any) (float64, error) {
	return equalAggregate[float64](values, cvt.Float64E)
}
