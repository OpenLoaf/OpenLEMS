package util_binary

import "encoding/binary"

// _ReverseMiddleEndian 反中端序是 BA DC 正常大端是 AB CD 小端是 DC BA 中端序是 CD AB
var _ReverseMiddleEndian reverseMiddleEndian

type reverseMiddleEndian struct { // 2位以下的时候和大端是一样的，只有2位以上的时候，变成了。BA DC
}

func (c reverseMiddleEndian) String() string {
	return "_ReverseMiddleEndian"
}

func (c reverseMiddleEndian) GoString() string {
	return "util._ReverseMiddleEndian"
}

func (c reverseMiddleEndian) Uint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func (c reverseMiddleEndian) PutUint16(b []byte, v uint16) {
	binary.BigEndian.PutUint16(b, v)
}

func (c reverseMiddleEndian) AppendUint16(b []byte, v uint16) []byte {
	return binary.BigEndian.AppendUint16(b, v)
}

func (c reverseMiddleEndian) Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	// 下面就是和大端的区别了
	//return uint32(b[1]) | uint32(b[0])<<8 | uint32(b[3])<<16 | uint32(b[2])<<24 中端

	//return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24 大端
	//return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24 小端

	return uint32(b[2]) | uint32(b[3])<<8 | uint32(b[0])<<16 | uint32(b[1])<<24 // BA DC
}
func (c reverseMiddleEndian) PutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[1] = byte(v >> 24)
	b[0] = byte(v >> 16)
	b[3] = byte(v >> 8)
	b[2] = byte(v)
}

func (c reverseMiddleEndian) AppendUint32(b []byte, v uint32) []byte {
	return append(b,
		byte(v>>16),
		byte(v>>24),
		byte(v),
		byte(v>>8),
	)
}

func (c reverseMiddleEndian) Uint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	// GH EF CD AB
	/*
			 中 GH EF CD AB
			return uint64(b[1]) | uint64(b[0])<<8 | uint64(b[3])<<16 | uint64(b[2])<<24 |
					uint64(b[5])<<32 | uint64(b[4])<<40 | uint64(b[7])<<48 | uint64(b[6])<<56

			大 AB CD EF GH
		return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
				uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	*/

	return uint64(b[6]) | uint64(b[7])<<8 | uint64(b[4])<<16 | uint64(b[5])<<24 |
		uint64(b[2])<<32 | uint64(b[3])<<40 | uint64(b[0])<<48 | uint64(b[1])<<56
}

func (c reverseMiddleEndian) PutUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[1] = byte(v >> 56)
	b[0] = byte(v >> 48)
	b[3] = byte(v >> 40)
	b[2] = byte(v >> 32)
	b[5] = byte(v >> 24)
	b[4] = byte(v >> 16)
	b[7] = byte(v >> 8)
	b[6] = byte(v)
}

func (c reverseMiddleEndian) AppendUint64(b []byte, v uint64) []byte {
	return append(b,
		byte(v>>48),
		byte(v>>56),
		byte(v>>32),
		byte(v>>40),
		byte(v>>16),
		byte(v>>24),
		byte(v),
		byte(v>>8),
	)
}
