package p_base

import (
	"common/c_base"
	"common/c_enum"
	"context"
	"p_base/internal"

	"github.com/gogf/gf/v2/os/gcache"
)

// NewGetProtocolCacheValue 创建协议缓存值获取器实例
func NewGetProtocolCacheValue(ctx context.Context, deviceId string, deviceType c_enum.EDeviceType, cache *gcache.Cache) c_base.IProtocolCacheValue {
	return internal.NewGetProtocolCacheValue(ctx, deviceId, deviceType, cache)
}

// DecoderBytes 通用字节解析函数，支持多种协议的数据解析
//
// 此函数是协议解析的核心，支持以下协议：
// - Modbus TCP/RTU: 支持所有标准数据类型和字节序
// - CANbus: 支持位级数据、BCD码、多字节序
// - IEC 61850: 支持IEEE 754浮点数、字符串
// - S7: 支持西门子PLC的数据格式
// - 其他工业协议: 通过自定义格式扩展
//
// 参数说明：
//   - bytes: 原始字节数据
//   - byteIndex: 字节起始索引
//   - byteLength: 字节长度（0表示纯位模式）
//   - bitIndex: 位起始索引
//   - bitLength: 位长度（0表示纯字节模式）
//   - byteEndian: 字节序（大端/小端）
//   - wordOrder: 字序（高字在前/低字在前）
//   - dataFormat: 数据格式（整数、浮点数、BCD、字符串等）
//   - returnFormat: 返回格式类型
//   - offset: 偏移量
//   - factor: 系数
//
// 三种读取模式：
//  1. 纯字节模式 (bitLength=0): 读取 byteIndex 开始的 byteLength 字节
//  2. 纯位模式 (byteLength=0): 读取 bitIndex 开始的 bitLength 位
//  3. 混合模式 (两者都不为0): 先读取字节，再从中提取指定位
//
// 使用示例：
//
//	// 模式1：Modbus 16位整数解析（纯字节模式）
//	result, err := DecoderBytes(data, 0, 2, 0, 0, ByteEndianBig, WordOrderHighLow, DataFormatUInt16, EInt16, 0, 1.0)
//	if err != nil {
//		// 处理错误
//	}
//
//	// 模式2：CANbus 位数据解析（纯位模式）
//	result, err := DecoderBytes(data, 0, 0, 5, 3, ByteEndianLittle, WordOrderHighLow, DataFormatBitRange, EInt32, 0, 1.0)
//	if err != nil {
//		// 处理错误
//	}
//
//	// 模式3：混合模式 - 从第2字节开始读取2字节，然后提取第3-7位
//	result, err := DecoderBytes(data, 2, 2, 3, 5, ByteEndianBig, WordOrderHighLow, DataFormatBitRange, EInt32, 0, 1.0)
//	if err != nil {
//		// 处理错误
//	}
func DecoderBytes(bytes []byte, byteIndex uint16, byteLength uint16, bitIndex uint16, bitLength uint16, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder, dataFormat c_enum.EDataFormat, returnFormat c_enum.EValueType, offset int, factor float32) (any, error) {
	return internal.DecoderBytes(bytes, byteIndex, byteLength, bitIndex, bitLength, byteEndian, wordOrder, dataFormat, returnFormat, offset, factor)
}

// EncoderBytes 通用字节编码函数，支持多种协议的数据编码
//
// 此函数用于生成Modbus、CANbus、IEC 61850、S7等协议的发送指令，支持以下协议：
// - Modbus TCP/RTU: 支持所有标准数据类型和字节序
// - CANbus: 支持位级数据、BCD码、多字节序
// - IEC 61850: 支持IEEE 754浮点数、字符串
// - S7: 支持西门子PLC的数据格式
// - 其他工业协议: 通过自定义格式扩展
//
// 参数说明：
//   - value: 要编码的值
//   - dataFormat: 数据格式（整数、浮点数、BCD、字符串等）
//   - byteEndian: 字节序（大端/小端）
//   - wordOrder: 字序（高字在前/低字在前）
//   - offset: 偏移量（编码前应用）
//   - factor: 系数（编码前应用）
//
// 使用示例：
//
//	// Modbus 16位整数编码
//	result, err := EncoderBytes(1234, DataFormatUInt16, ByteEndianBig, WordOrderHighLow, 0, 1.0)
//	if err != nil {
//		// 处理错误
//	}
//
//	// CANbus BCD码编码
//	result, err := EncoderBytes(1234, DataFormatBCD, ByteEndianLittle, WordOrderHighLow, 0, 0.1)
//	if err != nil {
//		// 处理错误
//	}
//
//	// IEEE 754浮点数编码
//	result, err := EncoderBytes(3.14, DataFormatFloat32, ByteEndianBig, WordOrderHighLow, 0, 1.0)
//	if err != nil {
//		// 处理错误
//	}
//
//	// ASCII字符串编码
//	result, err := EncoderBytes("Hello", DataFormatStringASCII, ByteEndianBig, WordOrderHighLow, 0, 1.0)
//	if err != nil {
//		// 处理错误
//	}
func EncoderBytes(value any, dataFormat c_enum.EDataFormat, byteEndian c_enum.EByteEndian, wordOrder c_enum.EWordOrder, offset int, factor float32) ([]byte, error) {
	return internal.EncoderBytes(value, dataFormat, byteEndian, wordOrder, offset, factor)
}

// ValidateValueRange 数据范围验证函数
//
// 此函数用于验证数据的范围，支持以下验证类型：
// - 数值类型：验证数值是否在指定范围内
// - 字符串类型：验证字符串长度是否在指定范围内
// - 布尔类型：自动跳过范围验证
//
// 参数说明：
//   - value: 要验证的值
//   - min: 最小值验证（数值类型）或最小长度验证（字符串类型，0表示不验证）
//   - max: 最大值验证（数值类型）或最大长度验证（字符串类型，0表示不验证）
//
// 验证规则：
//  1. 如果min和max都为0，表示不进行范围验证
//  2. 数值类型：验证值是否在[min, max]范围内
//  3. 字符串类型：验证长度是否在[min, max]范围内
//  4. 布尔类型：自动跳过验证
//
// 使用示例：
//
//	// 验证数值范围
//	err := ValidateValueRange(1234, 0, 10000)
//	if err != nil {
//		// 处理验证失败
//	}
//
//	// 验证字符串长度
//	err := ValidateValueRange("Hello", 3, 20)
//	if err != nil {
//		// 处理验证失败
//	}
//
//	// 不进行验证
//	err := ValidateValueRange(1234, 0, 0)
//	// err 为 nil
func ValidateValueRange(value any, min, max int64) error {
	return internal.ValidateValueRange(value, min, max)
}

// GetEffectiveByteLength 获取有效的字节长度
//
// 此函数用于根据数据访问配置获取有效的字节长度，支持以下逻辑：
// - 如果明确指定了 byteLength，使用指定值
// - 否则根据 dataFormat 返回默认长度
// - 对于可变长度格式，返回0表示需要外部指定
//
// 参数说明：
//   - byteLength: 指定的字节长度
//   - dataFormat: 数据格式
//
// 返回值：
//   - uint16: 有效的字节长度
//
// 使用示例：
//
//	// 获取16位整数的有效字节长度
//	length := GetEffectiveByteLength(0, DataFormatUInt16)
//	if length == 0 {
//		// 需要外部指定长度
//	}
func GetEffectiveByteLength(byteLength uint16, dataFormat c_enum.EDataFormat) uint16 {
	return internal.GetEffectiveByteLength(byteLength, dataFormat)
}

// GetEffectiveBitLength 获取有效的位长度
//
// 此函数用于根据数据访问配置获取有效的位长度，支持以下逻辑：
// - 如果明确指定了 bitLength，使用指定值
// - 否则根据 dataFormat 返回默认位长度
// - 对于非位格式，返回0
//
// 参数说明：
//   - bitLength: 指定的位长度
//   - dataFormat: 数据格式
//
// 返回值：
//   - uint16: 有效的位长度
//
// 使用示例：
//
//	// 获取位数据的有效位长度
//	bitLength := GetEffectiveBitLength(0, DataFormatBitRange)
//	if bitLength > 0 {
//		// 处理位数据
//	}
func GetEffectiveBitLength(bitLength uint16, dataFormat c_enum.EDataFormat) uint16 {
	return internal.GetEffectiveBitLength(bitLength, dataFormat)
}

// IsBitMode 判断是否为位模式
//
// 此函数用于判断数据访问配置是否为位模式，支持以下判断逻辑：
// - bitLength > 0
// - dataFormat 为 DataFormatBits 或 DataFormatBitRange
//
// 参数说明：
//   - bitLength: 位长度
//   - dataFormat: 数据格式
//
// 返回值：
//   - bool: 是否为位模式
//
// 使用示例：
//
//	// 判断是否为位模式
//	if IsBitMode(3, DataFormatBitRange) {
//		// 处理位数据
//	} else {
//		// 处理字节数据
//	}
func IsBitMode(bitLength uint16, dataFormat c_enum.EDataFormat) bool {
	return internal.IsBitMode(bitLength, dataFormat)
}

// IsByteMode 判断是否为字节模式
//
// 此函数用于判断数据访问配置是否为字节模式，支持以下判断逻辑：
// - byteLength > 0
// - 或者 GetEffectiveByteLength 返回 > 0
//
// 参数说明：
//   - byteLength: 字节长度
//   - dataFormat: 数据格式
//
// 返回值：
//   - bool: 是否为字节模式
//
// 使用示例：
//
//	// 判断是否为字节模式
//	if IsByteMode(2, DataFormatUInt16) {
//		// 处理字节数据
//	} else {
//		// 处理位数据
//	}
func IsByteMode(byteLength uint16, dataFormat c_enum.EDataFormat) bool {
	return internal.IsByteMode(byteLength, dataFormat)
}

// GetQuantity 计算需要的寄存器数量
//
// 此函数用于计算数据访问配置需要的寄存器数量，支持以下计算逻辑：
// - 位模式：1 quantity = 16 bits，向上取整
// - 字节模式：1 quantity = 2 bytes，向上取整
//
// 参数说明：
//   - byteLength: 字节长度
//   - bitLength: 位长度
//   - dataFormat: 数据格式
//
// 返回值：
//   - uint16: 需要的寄存器数量
//
// 使用示例：
//
//	// 计算需要的寄存器数量
//	quantity := GetQuantity(2, 0, DataFormatUInt16)
//	// 用于 Modbus 协议的数据读取
func GetQuantity(byteLength uint16, bitLength uint16, dataFormat c_enum.EDataFormat) uint16 {
	return internal.GetQuantity(byteLength, bitLength, dataFormat)
}

// GetQuantityFromDataAccess 使用 SDataAccess 结构体计算需要的寄存器数量
//
// 此函数提供向后兼容性，用于计算数据访问配置需要的寄存器数量，支持以下计算逻辑：
// - 位模式：1 quantity = 16 bits，向上取整
// - 字节模式：1 quantity = 2 bytes，向上取整
//
// 参数说明：
//   - dataAccess: 数据访问配置结构体
//
// 返回值：
//   - uint16: 需要的寄存器数量
//
// 使用示例：
//
//	// 计算需要的寄存器数量（向后兼容版本）
//	quantity := GetQuantityFromDataAccess(&dataAccess)
//	// 用于 Modbus 协议的数据读取
func GetQuantityFromDataAccess(dataAccess *c_base.SDataAccess) uint16 {
	return internal.GetQuantityFromDataAccess(dataAccess)
}
