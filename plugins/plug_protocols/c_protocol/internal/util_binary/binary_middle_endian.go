package util_binary

import "encoding/binary"

// _MiddleEndian 中端序是 CD AB 正常大端是 AB CD 小端是 DC BA
var _MiddleEndian middleEndian

type middleEndian struct { // 2位以下的时候和小端是一样的，只有2位以上的时候，变成了。CD AB
}

func (c middleEndian) String() string {
	return "_MiddleEndian"
}

func (c middleEndian) GoString() string {
	return "util._MiddleEndian"
}

func (c middleEndian) Uint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func (c middleEndian) PutUint16(b []byte, v uint16) {
	binary.LittleEndian.PutUint16(b, v)
}

func (c middleEndian) AppendUint16(b []byte, v uint16) []byte {
	return binary.LittleEndian.AppendUint16(b, v)
}

func (c middleEndian) Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	// 下面就是和小端的区别了
	return uint32(b[1]) | uint32(b[0])<<8 | uint32(b[3])<<16 | uint32(b[2])<<24
}
func (c middleEndian) PutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[2] = byte(v >> 24)
	b[3] = byte(v >> 16)
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}

func (c middleEndian) AppendUint32(b []byte, v uint32) []byte {
	return append(b,
		byte(v>>8),
		byte(v),
		byte(v>>24),
		byte(v>>16),
	)
}

func (c middleEndian) Uint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	// GH EF CD AB
	return uint64(b[1]) | uint64(b[0])<<8 | uint64(b[3])<<16 | uint64(b[2])<<24 |
		uint64(b[5])<<32 | uint64(b[4])<<40 | uint64(b[7])<<48 | uint64(b[6])<<56
}

func (c middleEndian) PutUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[1] = byte(v)
	b[0] = byte(v >> 8)
	b[3] = byte(v >> 16)
	b[2] = byte(v >> 24)
	b[5] = byte(v >> 32)
	b[4] = byte(v >> 40)
	b[7] = byte(v >> 48)
	b[6] = byte(v >> 56)
}

func (c middleEndian) AppendUint64(b []byte, v uint64) []byte {
	return append(b,
		byte(v>>8),
		byte(v),
		byte(v>>24),
		byte(v>>16),
		byte(v>>40),
		byte(v>>32),
		byte(v>>56),
		byte(v>>48),
	)
}
