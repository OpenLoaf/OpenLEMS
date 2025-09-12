package c_func

import (
	"common/c_base"

	"github.com/shockerli/cvt"
)

// 求总和
func sumAggregate[T c_base.Number](values []any, convert func(interface{}) (T, error)) (*T, error) {
	var value T
	var err error
	for _, v := range values {
		converted, e := convert(v)
		if e != nil {
			err = e
			break
		} else {
			value += converted
		}
	}
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// 整数类型
var AggregateSumInt = func(values []any) (*int, error) {
	return sumAggregate[int](values, cvt.IntE)
}

var AggregateSumInt8 = func(values []any) (*int8, error) {
	return sumAggregate[int8](values, cvt.Int8E)
}

var AggregateSumInt16 = func(values []any) (*int16, error) {
	return sumAggregate[int16](values, cvt.Int16E)
}

var AggregateSumInt32 = func(values []any) (*int32, error) {
	return sumAggregate[int32](values, cvt.Int32E)
}

var AggregateSumInt64 = func(values []any) (*int64, error) {
	return sumAggregate[int64](values, cvt.Int64E)
}

// 无符号整数类型
var AggregateSumUint = func(values []any) (*uint, error) {
	return sumAggregate[uint](values, cvt.UintE)
}

var AggregateSumUint8 = func(values []any) (*uint8, error) {
	return sumAggregate[uint8](values, cvt.Uint8E)
}

var AggregateSumUint16 = func(values []any) (*uint16, error) {
	return sumAggregate[uint16](values, cvt.Uint16E)
}

var AggregateSumUint32 = func(values []any) (*uint32, error) {
	return sumAggregate[uint32](values, cvt.Uint32E)
}

var AggregateSumUint64 = func(values []any) (*uint64, error) {
	return sumAggregate[uint64](values, cvt.Uint64E)
}

// 浮点数类型
var AggregateSumFloat32 = func(values []any) (*float32, error) {
	return sumAggregate[float32](values, cvt.Float32E)
}

var AggregateSumFloat64 = func(values []any) (*float64, error) {
	return sumAggregate[float64](values, cvt.Float64E)
}
