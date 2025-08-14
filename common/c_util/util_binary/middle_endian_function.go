package util_binary

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"math"
)

func MeEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(values); i++ {
		if values[i] == nil {
			return buf.Bytes()
		}
		switch value := values[i].(type) {
		case int:
			buf.Write(MeEncodeInt(value))
		case int8:
			buf.Write(MeEncodeInt8(value))
		case int16:
			buf.Write(MeEncodeInt16(value))
		case int32:
			buf.Write(MeEncodeInt32(value))
		case int64:
			buf.Write(MeEncodeInt64(value))
		case uint:
			buf.Write(MeEncodeUint(value))
		case uint8:
			buf.Write(MeEncodeUint8(value))
		case uint16:
			buf.Write(MeEncodeUint16(value))
		case uint32:
			buf.Write(MeEncodeUint32(value))
		case uint64:
			buf.Write(MeEncodeUint64(value))
		case bool:
			buf.Write(MeEncodeBool(value))
		case string:
			buf.Write(MeEncodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(MeEncodeFloat32(value))
		case float64:
			buf.Write(MeEncodeFloat64(value))

		default:
			if err := binary.Write(buf, _MiddleEndian, value); err != nil {
				g.Log().Errorf(context.TODO(), "%+v", err)
				buf.Write(MeEncodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

func MeEncodeByMength(length int, values ...interface{}) []byte {
	b := MeEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[0:length]
	}
	return b
}

func MeDecode(b []byte, values ...interface{}) error {
	var (
		err error
		buf = bytes.NewBuffer(b)
	)
	for i := 0; i < len(values); i++ {
		if err = binary.Read(buf, _MiddleEndian, values[i]); err != nil {
			err = gerror.Wrap(err, `binary.Read failed`)
			return err
		}
	}
	return nil
}

func MeEncodeString(s string) []byte {
	return []byte(s)
}

func MeDecodeToString(b []byte) string {
	return string(b)
}

func MeEncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

func MeEncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return MeEncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return MeEncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return MeEncodeInt32(int32(i))
	} else {
		return MeEncodeInt64(int64(i))
	}
}

func MeEncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return MeEncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return MeEncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return MeEncodeUint32(uint32(i))
	} else {
		return MeEncodeUint64(uint64(i))
	}
}

func MeEncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func MeEncodeUint8(i uint8) []byte {
	return []byte{i}
}

func MeEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	_MiddleEndian.PutUint16(b, uint16(i))
	return b
}

func MeEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	_MiddleEndian.PutUint16(b, i)
	return b
}

func MeEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	_MiddleEndian.PutUint32(b, uint32(i))
	return b
}

func MeEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	_MiddleEndian.PutUint32(b, i)
	return b
}

func MeEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	_MiddleEndian.PutUint64(b, uint64(i))
	return b
}

func MeEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	_MiddleEndian.PutUint64(b, i)
	return b
}

func MeEncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	_MiddleEndian.PutUint32(b, bits)
	return b
}

func MeEncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	_MiddleEndian.PutUint64(b, bits)
	return b
}

func MeDecodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(MeDecodeToUint8(b))
	} else if len(b) < 3 {
		return int(MeDecodeToUint16(b))
	} else if len(b) < 5 {
		return int(MeDecodeToUint32(b))
	} else {
		return int(MeDecodeToUint64(b))
	}
}

func MeDecodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(MeDecodeToUint8(b))
	} else if len(b) < 3 {
		return uint(MeDecodeToUint16(b))
	} else if len(b) < 5 {
		return uint(MeDecodeToUint32(b))
	} else {
		return uint(MeDecodeToUint64(b))
	}
}

func MeDecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Equal(b, make([]byte, len(b))) {
		return false
	}
	return true
}

func MeDecodeToInt8(b []byte) int8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return int8(b[0])
}

func MeDecodeToUint8(b []byte) uint8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return b[0]
}

func MeDecodeToInt16(b []byte) int16 {
	return int16(_MiddleEndian.Uint16(MeFillUpSize(b, 2)))
}

func MeDecodeToUint16(b []byte) uint16 {
	return _MiddleEndian.Uint16(MeFillUpSize(b, 2))
}

func MeDecodeToInt32(b []byte) int32 {
	return int32(_MiddleEndian.Uint32(MeFillUpSize(b, 4)))
}

func MeDecodeToUint32(b []byte) uint32 {
	return _MiddleEndian.Uint32(MeFillUpSize(b, 4))
}

func MeDecodeToInt64(b []byte) int64 {
	return int64(_MiddleEndian.Uint64(MeFillUpSize(b, 8)))
}

func MeDecodeToUint64(b []byte) uint64 {
	return _MiddleEndian.Uint64(MeFillUpSize(b, 8))
}

func MeDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(_MiddleEndian.Uint32(MeFillUpSize(b, 4)))
}

func MeDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(_MiddleEndian.Uint64(MeFillUpSize(b, 8)))
}

// MeFillUpSize fills up the bytes `b` to given length `l` using LittleEndian.
//
// Note that it creates a new bytes slice by copying the original one to avoid changing
// the original parameter bytes.
func MeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}

// FillUpBitSizeRight 和小端、中端一致
func FillUpBitSizeRight(b []gbinary.Bit, l int) []gbinary.Bit {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]gbinary.Bit, l)
	copy(c, b)
	return c
}

// FillUpBitSizeLeft 和大端，反中端一致
func FillUpBitSizeLeft(b []gbinary.Bit, l int) []gbinary.Bit {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]gbinary.Bit, l)
	copy(c[l-len(b):], b)
	return c
}
