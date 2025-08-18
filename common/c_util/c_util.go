package c_util

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// TernaryOp 泛型三元运算符函数
func TernaryOp[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func ExistsInSlice(value any, slice []any) bool {
	if len(slice) == 0 {
		return false
	}
	for _, s := range slice {
		if value != s {
			return false
		}
	}
	return true
}

// 判断是否是数组或者切片
func isSliceOrArray(value any) bool {
	v := reflect.ValueOf(value)
	kind := v.Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

func Operate[T Number](a, b T) (sum, product T) {
	// 加法
	sum = a + b
	// 乘法
	product = a * b
	return
}

// Contains 函数检查任意类型的切片中是否包含某个值
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func IsNil[T interface{}](v T) bool {
	value := reflect.ValueOf(v)
	if !value.IsValid() || value.IsNil() {
		return true
	}
	return false
}

func PrintHex(b []byte) {
	fmt.Printf("[")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%X", b[i])
		if i != len(b)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("]\n")
}

func PrintBinary(b []byte, append ...string) {
	if len(append) == 0 {
		fmt.Printf("[")
	} else {
		fmt.Printf("%s[", append[0])
	}

	for i := 0; i < len(b); i++ {
		fmt.Printf("%b", b[i])
		if i != len(b)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("]\n")
}

func Float64ToString(value float64, pre int) string {
	if value == math.NaN() {
		return ""
	}
	return big.NewFloat(value).Text('f', pre)
}

func Float32ToString(value float32, pre int) string {
	return big.NewFloat(float64(value)).Text('f', pre)
}
