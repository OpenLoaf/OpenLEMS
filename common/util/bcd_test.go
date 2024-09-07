package util

import (
	"fmt"
	"testing"
)

func TestBcd16ToDecimal(t *testing.T) {
	number := 2024
	bcd := DecimalToBCD16(number)
	fmt.Printf("Decimal: %d, BCD: %04x\n", number, bcd)

	decimal := Bcd16ToDecimal(bcd)
	fmt.Printf("Decoded BCD: %04x, Decimal: %d\n", bcd, decimal)

	b := []byte{0x12, 0x34} // 表示 1234
	d := BcdToDecimalMulti(b)
	fmt.Printf("BCD: %x, Decimal: %d\n", b, d) // 输出: BCD: 1234 -> 1234

}
