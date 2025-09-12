package c_func

import (
	"common/c_base"

	"github.com/shockerli/cvt"
)

// 求最小值
func minAggregate[T c_base.Number](values []any, convert func(interface{}) (T, error)) (*T, error) {
	if len(values) == 0 {
		return nil, nil
	}

	var mi T
	var isFirst = true
	var err error

	for _, v := range values {
		converted, e := convert(v)
		if e != nil {
			err = e
			break
		} else {
			if isFirst {
				mi = converted
				isFirst = false
			} else if converted < mi {
				mi = converted
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return &mi, nil
}

// 整数类型
var AggregateMinInt = func(values []any) (*int, error) {
	return minAggregate[int](values, cvt.IntE)
}

var AggregateMinInt8 = func(values []any) (*int8, error) {
	return minAggregate[int8](values, cvt.Int8E)
}

var AggregateMinInt16 = func(values []any) (*int16, error) {
	return minAggregate[int16](values, cvt.Int16E)
}

var AggregateMinInt32 = func(values []any) (*int32, error) {
	return minAggregate[int32](values, cvt.Int32E)
}

var AggregateMinInt64 = func(values []any) (*int64, error) {
	return minAggregate[int64](values, cvt.Int64E)
}

// 无符号整数类型
var AggregateMinUint = func(values []any) (*uint, error) {
	return minAggregate[uint](values, cvt.UintE)
}

var AggregateMinUint8 = func(values []any) (*uint8, error) {
	return minAggregate[uint8](values, cvt.Uint8E)
}

var AggregateMinUint16 = func(values []any) (*uint16, error) {
	return minAggregate[uint16](values, cvt.Uint16E)
}

var AggregateMinUint32 = func(values []any) (*uint32, error) {
	return minAggregate[uint32](values, cvt.Uint32E)
}

var AggregateMinUint64 = func(values []any) (*uint64, error) {
	return minAggregate[uint64](values, cvt.Uint64E)
}

// 浮点数类型
var AggregateMinFloat32 = func(values []any) (*float32, error) {
	return minAggregate[float32](values, cvt.Float32E)
}

var AggregateMinFloat64 = func(values []any) (*float64, error) {
	return minAggregate[float64](values, cvt.Float64E)
}
