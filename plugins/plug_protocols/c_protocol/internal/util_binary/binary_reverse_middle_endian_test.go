package util_binary

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func BenchmarkRMEBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value := rand.Int32()
		b4 := make([]byte, 0, 4)
		b4 = _ReverseMiddleEndian.AppendUint32(b4, uint32(value))

		u := _ReverseMiddleEndian.Uint32(b4)
		if u != uint32(value) {
			b.Errorf("Uint32 failed! Value: %d, Uint32: %d", value, u)
		}
	}
}

func TestRMEBinary(t *testing.T) {
	value := 746127360
	b4 := make([]byte, 0, 4)
	b4 = _ReverseMiddleEndian.AppendUint32(b4, uint32(value))

	printHex(b4)

	u := _ReverseMiddleEndian.Uint32(b4)
	if u != uint32(value) {
		t.Errorf("Uint32 failed! Value: %d, Uint32: %d", value, u)
	}
}

func TestEncode(t *testing.T) {
	bytes := []byte{0x79, 0x2C, 0x0, 0x0}

	u := _ReverseMiddleEndian.Uint32(bytes)
	if u != 746127360 {
		t.Errorf("Uint32 failed! Value: %d, Uint32: %d", 746127360, u)
	}
	fmt.Println(u)
}
