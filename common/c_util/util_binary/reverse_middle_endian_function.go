package util_binary

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"math"
)

func RmeEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(values); i++ {
		if values[i] == nil {
			return buf.Bytes()
		}

		switch value := values[i].(type) {
		case int:
			buf.Write(RmeEncodeInt(value))
		case int8:
			buf.Write(RmeEncodeInt8(value))
		case int16:
			buf.Write(RmeEncodeInt16(value))
		case int32:
			buf.Write(RmeEncodeInt32(value))
		case int64:
			buf.Write(RmeEncodeInt64(value))
		case uint:
			buf.Write(RmeEncodeUint(value))
		case uint8:
			buf.Write(RmeEncodeUint8(value))
		case uint16:
			buf.Write(RmeEncodeUint16(value))
		case uint32:
			buf.Write(RmeEncodeUint32(value))
		case uint64:
			buf.Write(RmeEncodeUint64(value))
		case bool:
			buf.Write(RmeEncodeBool(value))
		case string:
			buf.Write(RmeEncodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(RmeEncodeFloat32(value))
		case float64:
			buf.Write(RmeEncodeFloat64(value))
		default:
			if err := binary.Write(buf, _ReverseMiddleEndian, value); err != nil {
				g.Log().Errorf(context.TODO(), "%+v", err)
				buf.Write(RmeEncodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

func RmeEncodeByLength(length int, values ...interface{}) []byte {
	b := RmeEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[0:length]
	}
	return b
}

func RmeDecode(b []byte, values ...interface{}) error {
	var (
		err error
		buf = bytes.NewBuffer(b)
	)
	for i := 0; i < len(values); i++ {
		if err = binary.Read(buf, _ReverseMiddleEndian, values[i]); err != nil {
			err = gerror.Wrap(err, `binary.Read failed`)
			return err
		}
	}
	return nil
}

func RmeEncodeString(s string) []byte {
	return []byte(s)
}

func RmeDecodeToString(b []byte) string {
	return string(b)
}

func RmeEncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

func RmeEncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return RmeEncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return RmeEncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return RmeEncodeInt32(int32(i))
	} else {
		return RmeEncodeInt64(int64(i))
	}
}

func RmeEncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return RmeEncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return RmeEncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return RmeEncodeUint32(uint32(i))
	} else {
		return RmeEncodeUint64(uint64(i))
	}
}

func RmeEncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func RmeEncodeUint8(i uint8) []byte {
	return []byte{i}
}

func RmeEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	_ReverseMiddleEndian.PutUint16(b, uint16(i))
	return b
}

func RmeEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	_ReverseMiddleEndian.PutUint16(b, i)
	return b
}

func RmeEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	_ReverseMiddleEndian.PutUint32(b, uint32(i))
	return b
}

func RmeEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	_ReverseMiddleEndian.PutUint32(b, i)
	return b
}

func RmeEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	_ReverseMiddleEndian.PutUint64(b, uint64(i))
	return b
}

func RmeEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	_ReverseMiddleEndian.PutUint64(b, i)
	return b
}

func RmeEncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	_ReverseMiddleEndian.PutUint32(b, bits)
	return b
}

func RmeEncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	_ReverseMiddleEndian.PutUint64(b, bits)
	return b
}

func RmeDecodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(RmeDecodeToUint8(b))
	} else if len(b) < 3 {
		return int(RmeDecodeToUint16(b))
	} else if len(b) < 5 {
		return int(RmeDecodeToUint32(b))
	} else {
		return int(RmeDecodeToUint64(b))
	}
}

func RmeDecodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(RmeDecodeToUint8(b))
	} else if len(b) < 3 {
		return uint(RmeDecodeToUint16(b))
	} else if len(b) < 5 {
		return uint(RmeDecodeToUint32(b))
	} else {
		return uint(RmeDecodeToUint64(b))
	}
}

func RmeDecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Equal(b, make([]byte, len(b))) {
		return false
	}
	return true
}

func RmeDecodeToInt8(b []byte) int8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return int8(b[0])
}

func RmeDecodeToUint8(b []byte) uint8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return b[0]
}

func RmeDecodeToInt16(b []byte) int16 {
	return int16(_ReverseMiddleEndian.Uint16(RmeFillUpSize(b, 2)))
}

func RmeDecodeToUint16(b []byte) uint16 {
	return _ReverseMiddleEndian.Uint16(RmeFillUpSize(b, 2))
}

func RmeDecodeToInt32(b []byte) int32 {
	return int32(_ReverseMiddleEndian.Uint32(RmeFillUpSize(b, 4)))
}

func RmeDecodeToUint32(b []byte) uint32 {
	return _ReverseMiddleEndian.Uint32(RmeFillUpSize(b, 4))
}

func RmeDecodeToInt64(b []byte) int64 {
	return int64(_ReverseMiddleEndian.Uint64(RmeFillUpSize(b, 8)))
}

func RmeDecodeToUint64(b []byte) uint64 {
	return _ReverseMiddleEndian.Uint64(RmeFillUpSize(b, 8))
}

func RmeDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(_ReverseMiddleEndian.Uint32(RmeFillUpSize(b, 4)))
}

func RmeDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(_ReverseMiddleEndian.Uint64(RmeFillUpSize(b, 8)))
}

// RmeFillUpSize fills up the bytes `b` to given length `l` using big BigEndian.
//
// Note that it creates a new bytes slice by copying the original one to avoid changing
// the original parameter bytes.
func RmeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c[l-len(b):], b)
	return c
}
