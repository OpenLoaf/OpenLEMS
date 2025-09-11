package internal

import (
	"common/c_base"
	"common/c_enum"
)

// GetEffectiveByteLength 获取有效的字节长度
func GetEffectiveByteLength(byteLength uint16, dataFormat c_enum.EDataFormat) uint16 {
	// 如果明确指定了字节长度，使用指定值
	if byteLength > 0 {
		return byteLength
	}

	// 否则根据数据格式返回默认长度
	switch dataFormat {
	case c_enum.DataFormatUInt8, c_enum.DataFormatInt8:
		return 1
	case c_enum.DataFormatUInt16, c_enum.DataFormatInt16, c_enum.DataFormatBCD:
		return 2
	case c_enum.DataFormatUInt32, c_enum.DataFormatInt32, c_enum.DataFormatFloat32, c_enum.DataFormatBCD32:
		return 4
	case c_enum.DataFormatFloat64:
		return 8
	case c_enum.DataFormatStringASCII, c_enum.DataFormatStringUTF16, c_enum.DataFormatBits, c_enum.DataFormatBitRange, c_enum.DataFormatCustom:
		// 可变长度格式，如果没有指定长度，返回0表示需要外部指定
		return 0
	default:
		return 0
	}
}

// GetEffectiveBitLength 获取有效的位长度
func GetEffectiveBitLength(bitLength uint16, dataFormat c_enum.EDataFormat) uint16 {
	if bitLength > 0 {
		return bitLength
	}

	// 如果是位相关格式且没有指定位长度，返回默认值
	switch dataFormat {
	case c_enum.DataFormatBits, c_enum.DataFormatBitRange:
		return 1 // 默认1位
	default:
		return 0 // 非位格式
	}
}

// IsBitMode 判断是否为位模式
func IsBitMode(bitLength uint16, dataFormat c_enum.EDataFormat) bool {
	return bitLength > 0 || dataFormat == c_enum.DataFormatBits || dataFormat == c_enum.DataFormatBitRange
}

// IsByteMode 判断是否为字节模式
func IsByteMode(byteLength uint16, dataFormat c_enum.EDataFormat) bool {
	return byteLength > 0 || GetEffectiveByteLength(byteLength, dataFormat) > 0
}

// GetQuantity 计算需要的寄存器数量
func GetQuantity(byteLength uint16, bitLength uint16, dataFormat c_enum.EDataFormat) uint16 {
	if IsBitMode(bitLength, dataFormat) {
		// 位模式：1 quantity = 16 bits
		effectiveBitLength := GetEffectiveBitLength(bitLength, dataFormat)
		if effectiveBitLength == 0 {
			effectiveBitLength = 1
		}
		return (effectiveBitLength + 15) / 16 // 向上取整
	} else {
		// 字节模式：1 quantity = 2 bytes
		effectiveByteLength := GetEffectiveByteLength(byteLength, dataFormat)
		if effectiveByteLength == 0 {
			effectiveByteLength = 2 // 至少需要1个quantity
		}
		return (effectiveByteLength + 1) / 2 // 向上取整
	}
}

// GetQuantityFromDataAccess 使用 SDataAccess 结构体计算需要的寄存器数量
// 此函数提供向后兼容性，内部调用新的 GetQuantity 函数
func GetQuantityFromDataAccess(dataAccess *c_base.SDataAccess) uint16 {
	return GetQuantity(dataAccess.ByteLength, dataAccess.BitLength, dataAccess.DataFormat)
}
