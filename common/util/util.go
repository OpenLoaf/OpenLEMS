package util

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
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

func GetNow() string {
	return gtime.Now().Format("Y-m-d H:i:s.u")
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

func PrintBit(b []gbinary.Bit, append ...string) {
	if len(append) == 0 {
		fmt.Printf("[")
	} else {
		fmt.Printf("%s[", append[0])
	}

	for i := 0; i < len(b); i++ {
		fmt.Printf("%b", b[i])
		if (i+1)%4 == 0 {
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

// ExecuteFunction 执行函数
func ExecuteFunction(info any, telemetryKey string) (any, error) {
	if telemetryKey == "" {
		return 0, gerror.Newf("遥测点位名称不能为空")
	}
	if info == nil {
		return 0, gerror.Newf("对象为空！")
	}

	functionName := fmt.Sprintf("Get%s", gstr.UcFirst(telemetryKey))
	method := reflect.ValueOf(info).MethodByName(functionName)
	if !method.IsValid() {
		return 0, gerror.Newf("method %s not found", telemetryKey)
	}

	// 空参数调用
	value := method.Call(nil)
	if len(value) == 1 {
		return value[0].Interface(), nil
	}

	if len(value) != 2 {
		return 0, gerror.Newf("function %s return value length is not 2", telemetryKey)
	}
	if value[1].Interface() != nil {
		return 0, value[1].Interface().(error)
	}
	return value[0].Interface(), nil
}
