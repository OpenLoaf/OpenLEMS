package c_func

import (
	"common/c_base"

	"github.com/shockerli/cvt"
)

// 求平均值
func avgAggregate[T c_base.Number](values []any, convert func(interface{}) (T, error)) (*T, error) {
	if len(values) == 0 {
		return nil, nil
	}

	var sum T
	var err error
	for _, v := range values {
		converted, e := convert(v)
		if e != nil {
			err = e
			break
		} else {
			sum += converted
		}
	}
	if err != nil {
		return nil, err
	}
	result := sum / T(len(values))
	return &result, nil
}

// 整数类型
var AggregateAvgInt = func(values []any) (*int, error) {
	return avgAggregate[int](values, cvt.IntE)
}

var AggregateAvgInt8 = func(values []any) (*int8, error) {
	return avgAggregate[int8](values, cvt.Int8E)
}

var AggregateAvgInt16 = func(values []any) (*int16, error) {
	return avgAggregate[int16](values, cvt.Int16E)
}

var AggregateAvgInt32 = func(values []any) (*int32, error) {
	return avgAggregate[int32](values, cvt.Int32E)
}

var AggregateAvgInt64 = func(values []any) (*int64, error) {
	return avgAggregate[int64](values, cvt.Int64E)
}

// 无符号整数类型
var AggregateAvgUint = func(values []any) (*uint, error) {
	return avgAggregate[uint](values, cvt.UintE)
}

var AggregateAvgUint8 = func(values []any) (*uint8, error) {
	return avgAggregate[uint8](values, cvt.Uint8E)
}

var AggregateAvgUint16 = func(values []any) (*uint16, error) {
	return avgAggregate[uint16](values, cvt.Uint16E)
}

var AggregateAvgUint32 = func(values []any) (*uint32, error) {
	return avgAggregate[uint32](values, cvt.Uint32E)
}

var AggregateAvgUint64 = func(values []any) (*uint64, error) {
	return avgAggregate[uint64](values, cvt.Uint64E)
}

// 浮点数类型
var AggregateAvgFloat32 = func(values []any) (*float32, error) {
	return avgAggregate[float32](values, cvt.Float32E)
}

var AggregateAvgFloat64 = func(values []any) (*float64, error) {
	return avgAggregate[float64](values, cvt.Float64E)
}
