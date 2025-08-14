package util

// DecimalToBCD16 将一个最多四位的十进制数转换为16位BCD码
func DecimalToBCD16(n int) uint16 {
	var bcd uint16 = 0
	bcd |= uint16((n / 1000) << 12)      // 最高4位
	bcd |= uint16(((n / 100) % 10) << 8) // 次高4位
	bcd |= uint16(((n / 10) % 10) << 4)  // 次低4位
	bcd |= uint16(n % 10)                // 最低4位
	return bcd
}

// Bcd16ToDecimal 将16位BCD码转换为十进制数
func Bcd16ToDecimal(bcd uint16) int {
	// 通过移位和按位操作提取每个4位数
	thousands := int((bcd >> 12) & 0xF) // 最高4位
	hundreds := int((bcd >> 8) & 0xF)   // 次高4位
	tens := int((bcd >> 4) & 0xF)       // 次低4位
	ones := int(bcd & 0xF)              // 最低4位

	// 将每个提取的十进制数还原成完整的十进制数
	return thousands*1000 + hundreds*100 + tens*10 + ones
}

// DecimalToBCD16Bytes 将一个最多四位的十进制数转换为16位BCD码，并返回 []byte
func DecimalToBCD16Bytes(n int) []byte {
	var bcd uint16 = 0
	bcd |= uint16((n / 1000) << 12)      // 最高4位
	bcd |= uint16(((n / 100) % 10) << 8) // 次高4位
	bcd |= uint16(((n / 10) % 10) << 4)  // 次低4位
	bcd |= uint16(n % 10)                // 最低4位

	// 将16位的BCD拆分为2个字节
	highByte := byte(bcd >> 8)  // 高8位
	lowByte := byte(bcd & 0xFF) // 低8位

	return []byte{highByte, lowByte}
}

// BcdToDecimalMulti 处理多个字节的 BCD 编码
func BcdToDecimalMulti(bcd []byte) int {
	result := 0
	for _, byteVal := range bcd {
		result = result*100 + BcdToDecimal(byteVal)
	}
	return result
}

// BcdToDecimal 将 BCD 格式转换为十进制整数
func BcdToDecimal(bcd byte) int {
	high := int(bcd >> 4)  // 取出高 4 位
	low := int(bcd & 0x0F) // 取出低 4 位
	return high*10 + low
}
