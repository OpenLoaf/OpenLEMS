package internal

import (
	"common/c_enum"
	"math"

	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

// 数据范围验证函数（公共函数）
func ValidateValueRange(value any, min, max int64) error {
	// 如果min和max都为0，表示不进行范围验证
	if min == 0 && max == 0 {
		return nil
	}

	// 字符串类型进行长度验证
	if str, ok := value.(string); ok {
		length := int64(len(str))
		// 对于字符串长度，min和max应该是非负数，但我们仍然正确处理
		if min != 0 && length < min {
			return errors.Errorf("string length %d is below minimum %d", length, min)
		}
		if max != 0 && length > max {
			return errors.Errorf("string length %d is above maximum %d", length, max)
		}
		return nil
	}

	// 布尔类型不进行数值范围验证
	if _, ok := value.(bool); ok {
		return nil
	}

	// 处理uint64类型，避免溢出风险
	if uintVal, ok := value.(uint64); ok {
		// 检查uint64值是否在int64范围内
		if uintVal > math.MaxInt64 {
			return errors.Errorf("uint64 value %d exceeds int64 maximum %d", uintVal, int64(math.MaxInt64))
		}
		val := int64(uintVal)
		// 修复：正确处理负数范围验证
		if min != 0 && val < min {
			return errors.Errorf("value %d is below minimum %d", val, min)
		}
		if max != 0 && val > max {
			return errors.Errorf("value %d is above maximum %d", val, max)
		}
		return nil
	}

	// 处理其他数值类型
	val := cvt.Int64(value)

	// 修复：正确处理负数范围验证
	// 只有当min或max不为0时才进行相应的检查
	if min != 0 && val < min {
		return errors.Errorf("value %d is below minimum %d", val, min)
	}

	if max != 0 && val > max {
		return errors.Errorf("value %d is above maximum %d", val, max)
	}

	return nil
}

// 系数和偏移量处理（解码时使用）
func ApplyFactorAndOffset(value any, factor float32, offset int) any {
	// 如果系数为1且偏移为0，直接返回原值
	if factor == 1.0 && offset == 0 {
		return value
	}

	// 字符串类型不应用系数和偏移量
	if _, ok := value.(string); ok {
		return value
	}

	// 使用cvt库转换为float64，然后应用系数和偏移量
	val := cvt.Float64(value)
	return val*float64(factor) + float64(offset)
}

// 系数和偏移量处理（编码时使用，反向应用）
func ApplyFactorAndOffsetForEncoding(value any, factor float32, offset int) any {
	// 如果系数为1且偏移为0，直接返回原值
	if factor == 1.0 && offset == 0 {
		return value
	}

	// 字符串类型不应用系数和偏移量
	if _, ok := value.(string); ok {
		return value
	}

	// 使用cvt库转换为float64，然后反向应用系数和偏移量
	val := cvt.Float64(value)
	// 编码时反向应用：(value - offset) / factor
	return (val - float64(offset)) / float64(factor)
}

// 转换为目标格式（公共函数）
func ConvertToReturnFormat(value any, returnFormat c_enum.EValueType) any {
	if value == nil {
		return nil
	}

	switch returnFormat {
	case c_enum.EAuto:
		return value
	case c_enum.EBool:
		return cvt.Bool(value)
	case c_enum.EInt8:
		return cvt.Int8(value)
	case c_enum.EUint8:
		return cvt.Uint8(value)
	case c_enum.EInt16:
		return cvt.Int16(value)
	case c_enum.EUint16:
		return cvt.Uint16(value)
	case c_enum.EInt32:
		return cvt.Int32(value)
	case c_enum.EUint32:
		return cvt.Uint32(value)
	case c_enum.EInt64:
		return cvt.Int64(value)
	case c_enum.EUint64:
		return cvt.Uint64(value)
	case c_enum.EFloat32:
		return cvt.Float32(value)
	case c_enum.EFloat64:
		return cvt.Float64(value)
	case c_enum.EString:
		return cvt.String(value)
	default:
		return value
	}
}

// 字节序和字序处理辅助函数

// ReorderBytes 根据字节序和字序重新排列字节（解码时使用）
// 遵循Modbus协议标准：先处理字序（Word Order），再处理字节序（Byte Endian）
// 性能优化：预分配内存，减少内存分配次数
func ReorderBytes(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) []byte {
	if len(data) < 2 {
		return data
	}

	// 预分配结果切片，避免多次内存分配
	result := make([]byte, len(data))
	copy(result, data)

	// 第一步：处理字序（Word Order）- 交换16位字的顺序
	// 这在Modbus协议中首先发生
	if wordOrder == c_enum.WordOrderLowHigh {
		switch len(result) {
		case 4:
			// 32位数据：[AB CD EF GH] -> [EF GH AB CD]
			result[0], result[2] = result[2], result[0]
			result[1], result[3] = result[3], result[1]
		case 8:
			// 64位数据：[AB CD EF GH IJ KL MN OP] -> [IJ KL MN OP AB CD EF GH]
			// 交换前4字节和后4字节
			for i := 0; i < 4; i++ {
				result[i], result[i+4] = result[i+4], result[i]
			}
		default:
			// 对于其他长度，按4字节分组交换（保持向后兼容）
			if len(result) >= 4 {
				for i := 0; i < len(result)-3; i += 4 {
					result[i], result[i+2] = result[i+2], result[i]
					result[i+1], result[i+3] = result[i+3], result[i+1]
				}
			}
		}
	}

	// 第二步：处理字节序（Byte Endian）- 在每个16位字内交换字节
	// 这在Modbus协议中第二步发生
	if byteEndian == c_enum.ByteEndianLittle {
		// 小端字节序：反转每个16位字内的字节
		// 例如：[AB CD] -> [BA DC]
		for i := 0; i < len(result); i += 2 {
			if i+1 < len(result) {
				result[i], result[i+1] = result[i+1], result[i]
			}
		}
	}

	return result
}

// ReorderBytesForEncoding 根据字节序和字序重新排列字节（编码时使用）
// 与解码时的ReorderBytes相反的操作
func ReorderBytesForEncoding(data []byte, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder) []byte {
	if len(data) < 2 {
		return data
	}

	// 预分配结果切片，避免多次内存分配
	result := make([]byte, len(data))
	copy(result, data)

	// 第一步：处理字节序（Byte Endian）- 在每个16位字内交换字节
	// 编码时先处理字节序
	if byteEndian == c_enum.ByteEndianLittle {
		// 小端字节序：反转每个16位字内的字节
		// 例如：[BA DC] -> [AB CD]
		for i := 0; i < len(result); i += 2 {
			if i+1 < len(result) {
				result[i], result[i+1] = result[i+1], result[i]
			}
		}
	}

	// 第二步：处理字序（Word Order）- 交换16位字的顺序
	// 编码时后处理字序
	if wordOrder == c_enum.WordOrderLowHigh {
		switch len(result) {
		case 4:
			// 32位数据：[EF GH AB CD] -> [AB CD EF GH]
			result[0], result[2] = result[2], result[0]
			result[1], result[3] = result[3], result[1]
		case 8:
			// 64位数据：[IJ KL MN OP AB CD EF GH] -> [AB CD EF GH IJ KL MN OP]
			// 交换前4字节和后4字节
			for i := 0; i < 4; i++ {
				result[i], result[i+4] = result[i+4], result[i]
			}
		default:
			// 对于其他长度，按4字节分组交换（保持向后兼容）
			if len(result) >= 4 {
				for i := 0; i < len(result)-3; i += 4 {
					result[i], result[i+2] = result[i+2], result[i]
					result[i+1], result[i+3] = result[i+3], result[i+1]
				}
			}
		}
	}

	return result
}

// 确保系数不为0的辅助函数
func EnsureFactorNotZero(factor float32) float32 {
	if factor == 0 {
		return 1.0
	}
	return factor
}
