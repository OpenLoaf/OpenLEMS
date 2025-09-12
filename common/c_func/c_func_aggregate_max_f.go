package c_func

import (
	"common/c_base"

	"github.com/shockerli/cvt"
)

// 求最大值
func maxAggregate[T c_base.Number](values []any, convert func(interface{}) (T, error)) (*T, error) {
	if len(values) == 0 {
		return nil, nil
	}

	var mx T
	var isFirst = true
	var err error

	for _, v := range values {
		converted, e := convert(v)
		if e != nil {
			err = e
			break
		} else {
			if isFirst {
				mx = converted
				isFirst = false
			} else if converted > mx {
				mx = converted
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return &mx, nil
}

// 整数类型
var AggregateMaxInt = func(values []any) (*int, error) {
	return maxAggregate[int](values, cvt.IntE)
}

var AggregateMaxInt8 = func(values []any) (*int8, error) {
	return maxAggregate[int8](values, cvt.Int8E)
}

var AggregateMaxInt16 = func(values []any) (*int16, error) {
	return maxAggregate[int16](values, cvt.Int16E)
}

var AggregateMaxInt32 = func(values []any) (*int32, error) {
	return maxAggregate[int32](values, cvt.Int32E)
}

var AggregateMaxInt64 = func(values []any) (*int64, error) {
	return maxAggregate[int64](values, cvt.Int64E)
}

// 无符号整数类型
var AggregateMaxUint = func(values []any) (*uint, error) {
	return maxAggregate[uint](values, cvt.UintE)
}

var AggregateMaxUint8 = func(values []any) (*uint8, error) {
	return maxAggregate[uint8](values, cvt.Uint8E)
}

var AggregateMaxUint16 = func(values []any) (*uint16, error) {
	return maxAggregate[uint16](values, cvt.Uint16E)
}

var AggregateMaxUint32 = func(values []any) (*uint32, error) {
	return maxAggregate[uint32](values, cvt.Uint32E)
}

var AggregateMaxUint64 = func(values []any) (*uint64, error) {
	return maxAggregate[uint64](values, cvt.Uint64E)
}

// 浮点数类型
var AggregateMaxFloat32 = func(values []any) (*float32, error) {
	return maxAggregate[float32](values, cvt.Float32E)
}

var AggregateMaxFloat64 = func(values []any) (*float64, error) {
	return maxAggregate[float64](values, cvt.Float64E)
}
