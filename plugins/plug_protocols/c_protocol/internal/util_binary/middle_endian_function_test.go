package util_binary

import (
	"common/c_util"
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"math"
	"testing"
)

var testData = map[string]interface{}{
	//"nil":         nil,
	"int":         int(123),
	"int8":        int8(-99),
	"int8.max":    math.MaxInt8,
	"int16":       int16(123),
	"int16.max":   math.MaxInt16,
	"int32":       int32(-199),
	"int32.max":   math.MaxInt32,
	"int64":       int64(123),
	"uint":        uint(123),
	"uint8":       uint8(123),
	"uint8.max":   math.MaxUint8,
	"uint16":      uint16(9999),
	"uint16.max":  math.MaxUint16,
	"uint32":      uint32(123),
	"uint64":      uint64(123),
	"bool.true":   true,
	"bool.false":  false,
	"string":      "hehe haha",
	"byte":        []byte("hehe haha"),
	"float32":     float32(123.456),
	"float32.max": math.MaxFloat32,
	"float64":     float64(123.456),
}

func Test_MeEncodeAndMeDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for k, v := range testData {
			ve := MeEncode(v)
			ve1 := MeEncodeByMength(len(ve), v)

			//t.Logf("%s:%v, encoded:%v\n", k, v, ve)
			switch v.(type) {
			case int:
				t.Assert(MeDecodeToInt(ve), v)
				t.Assert(MeDecodeToInt(ve1), v)
			case int8:
				t.Assert(MeDecodeToInt8(ve), v)
				t.Assert(MeDecodeToInt8(ve1), v)
			case int16:
				t.Assert(MeDecodeToInt16(ve), v)
				t.Assert(MeDecodeToInt16(ve1), v)
			case int32:
				t.Assert(MeDecodeToInt32(ve), v)
				t.Assert(MeDecodeToInt32(ve1), v)
			case int64:
				t.Assert(MeDecodeToInt64(ve), v)
				t.Assert(MeDecodeToInt64(ve1), v)
			case uint:
				t.Assert(MeDecodeToUint(ve), v)
				t.Assert(MeDecodeToUint(ve1), v)
			case uint8:
				t.Assert(MeDecodeToUint8(ve), v)
				t.Assert(MeDecodeToUint8(ve1), v)
			case uint16:
				t.Assert(MeDecodeToUint16(ve1), v)
				t.Assert(MeDecodeToUint16(ve), v)
			case uint32:
				t.Assert(MeDecodeToUint32(ve1), v)
				t.Assert(MeDecodeToUint32(ve), v)
			case uint64:
				t.Assert(MeDecodeToUint64(ve), v)
				t.Assert(MeDecodeToUint64(ve1), v)
			case bool:
				t.Assert(MeDecodeToBool(ve), v)
				t.Assert(MeDecodeToBool(ve1), v)
			case string:
				t.Assert(MeDecodeToString(ve), v)
				t.Assert(MeDecodeToString(ve1), v)
			case float32:
				t.Assert(MeDecodeToFloat32(ve), v)
				t.Assert(MeDecodeToFloat32(ve1), v)
			case float64:
				t.Assert(MeDecodeToFloat64(ve), v)
				t.Assert(MeDecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := MeDecode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func TestFuncRMEBinary(t *testing.T) {
	bytes := []byte{0x79, 0x2C, 0x0, 0x0}

	u := MeDecodeToInt32(bytes)
	if u != 31020 {
		t.Errorf("Uint32 failed! Value: %d, Uint32: %d", 746127360, u)
	}
	fmt.Println(u)

	size := MeFillUpSize(bytes, 8)
	c_util.PrintHex(size)
}
