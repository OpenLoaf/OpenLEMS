package util_binary

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func BenchmarkBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value := rand.Int32()
		b4 := make([]byte, 0, 4)
		b4 = _MiddleEndian.AppendUint32(b4, uint32(value))

		u := _MiddleEndian.Uint32(b4)
		if u != uint32(value) {
			b.Errorf("Uint32 failed! Value: %d, Uint32: %d", value, u)
		}
	}
}

func printHex(b []byte) {
	fmt.Printf("[")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%X", b[i])
		if i != len(b)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("]\n")
}
