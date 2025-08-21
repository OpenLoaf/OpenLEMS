// Package internal_meta 提供CANbus数据解析和转换功能。
//
// 此包主要用于工业设备通信中的CANbus协议数据解析，支持多种数据格式：
//   - 位级数据读取：精确读取指定字节中的特定位，支持单位布尔值和多位整数
//   - BCD码解析：二进制编码十进制数的解析，支持多种字节序
//   - 整数类型：8/16/32/64位有符号和无符号整数
//   - 浮点数类型：IEEE754标准的32/64位浮点数
//   - 多字节序支持：大端、小端、中端序和反中端序
//
// 主要应用场景：
//   - 工业设备监控系统
//   - 汽车电子ECU通信
//   - 传感器数据采集
//   - 设备状态诊断
//
// 核心函数 ParseCanbusData 根据元数据配置解析原始CANbus数据帧，
// 自动处理数据类型转换、字节序调整和数值校正（系数、偏移量）。
package internal_meta

import (
	"common/c_base"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
)

// ParseCanbusData 根据 Meta 定义解析 CANbus 原始数据。
//
// 此函数是 CANbus 数据解析的核心入口，支持多种数据类型的解析，包括：
// - 位级读取 (RBit0-RBit15): 读取指定字节中的特定位
// - BCD码解析 (RBcd16): 解析16位BCD编码数据
// - 整数类型 (Int8/16/32/64, Uint8/16/32/64): 有符号和无符号整数
// - 浮点数类型 (Float32/64): IEEE754格式浮点数
//
// 参数:
//   - canData: CANbus 消息的原始数据部分，通常为8字节的数据帧
//   - meta: 包含解析规则的元数据，定义了地址、数据类型、字节序等信息
//
// 返回值:
//   - any: 解析后的数据，类型根据 meta.ReadType 而定
//   - error: 解析过程中的错误信息
//
// 使用示例:
//
//	meta := &c_base.Meta{
//	    Addr:      2,              // 字节地址
//	    ReadType:  c_base.RBit0,   // 读取第0位
//	    BitLength: 1,              // 位长度
//	    Factor:    1.0,            // 系数
//	    Offset:    0,              // 偏移量
//	}
//	result, err := ParseCanbusData([]byte{0x01, 0x00, 0x07}, meta)
func ParseCanbusData(canData []byte, meta *c_base.Meta) (any, error) {
	if meta == nil {
		return nil, gerror.New("Meta cannot be nil")
	}

	// 确保 Factor 不为 0，以避免除零或意外清零
	if meta.Factor == 0 {
		meta.Factor = 1.0 // 默认设置为 1
	}

	// 检查 Meta.Addr 是否超出 canData 范围（对于单字节读取类型）
	if int(meta.Addr) >= len(canData) && !(meta.ReadType >= c_base.RBit0 && meta.ReadType <= c_base.RBit15) {
		return nil, fmt.Errorf("SourceAddress %d out of bounds for data length %d for non-bit read type", meta.Addr, len(canData))
	}

	switch meta.ReadType {
	case c_base.RBit0, c_base.RBit1, c_base.RBit2, c_base.RBit3,
		c_base.RBit4, c_base.RBit5, c_base.RBit6, c_base.RBit7,
		c_base.RBit8, c_base.RBit9, c_base.RBit10, c_base.RBit11,
		c_base.RBit12, c_base.RBit13, c_base.RBit14, c_base.RBit15:
		// 位级读取
		return parseBitValue(canData, meta)

	case c_base.RBcd16:
		// BCD 码解析 (假设为 2 字节，即 16 位 BCD)
		if int(meta.Addr)+2 > len(canData) {
			return nil, fmt.Errorf("BCD16 read out of bounds. Addr: %d, Data Len: %d", meta.Addr, len(canData))
		}
		dataBytes := canData[meta.Addr : meta.Addr+2]

		var finalBCDVal float64
		// BCD 值的解释强烈依赖于具体的规范。
		// 这里假设一个通用的 4 位 BCD 格式 (一个字节两个 BCD 数字)。
		// 例如，0x12 0x34 -> 1234
		if meta.Endianness == c_base.EBigEndian {
			// 大端序：dataBytes[0] 是最高位的两个 BCD 数字，dataBytes[1] 是最低位的两个 BCD 数字
			finalBCDVal = float64(
				(int(dataBytes[0]>>4)&0x0F)*1000 + // 将 uint8 转换为 int 再乘以 1000
					(int(dataBytes[0]&0x0F))*100 + // 将 uint8 转换为 int 再乘以 100
					(int(dataBytes[1]>>4)&0x0F)*10 + // 将 uint8 转换为 int 再乘以 10
					(int(dataBytes[1]&0x0F))*1) // 将 uint8 转换为 int 再乘以 1
		} else if meta.Endianness == c_base.ELittleEndian {
			// 小端序：dataBytes[1] 是最高位的两个 BCD 数字，dataBytes[0] 是最低位的两个 BCD 数字
			finalBCDVal = float64(
				(int(dataBytes[1]>>4)&0x0F)*1000 +
					(int(dataBytes[1]&0x0F))*100 +
					(int(dataBytes[0]>>4)&0x0F)*10 +
					(int(dataBytes[0]&0x0F))*1)
		} else {
			// 对于中端序等非标准 BCD 字节序，通常需要 `FillUpSize` 预处理。
			// 这里简化处理，直接使用大端序的逻辑作为默认。
			// 如果您的 BCD 规范有特殊字节序，请确保 c_base.ECharSequence.FillUpSize 能正确处理。

			processedBytes := ECharSequenceFillUpSize(meta.Endianness, dataBytes, 2)
			finalBCDVal = float64(
				(int(processedBytes[0]>>4)&0x0F)*1000 +
					(int(processedBytes[0]&0x0F))*100 +
					(int(processedBytes[1]>>4)&0x0F)*10 +
					(int(processedBytes[1]&0x0F))*1)
		}
		return finalBCDVal*float64(meta.Factor) + float64(meta.Offset), nil

	case c_base.RInt8:
		val := int8(canData[meta.Addr])
		return float32(val)*meta.Factor + float32(meta.Offset), nil
	case c_base.RUint8:
		val := uint8(canData[meta.Addr])
		return float32(val)*meta.Factor + float32(meta.Offset), nil

	case c_base.RInt16:
		b, err := checkAndGetBytes(canData, meta, 2)
		if err != nil {
			return nil, err
		}

		v := ECharSequenceDecodeToInt16(meta.Endianness, b)
		return float32(v)*meta.Factor + float32(meta.Offset), nil
	case c_base.RUint16:
		b, err := checkAndGetBytes(canData, meta, 2)
		if err != nil {
			return nil, err
		}
		v := ECharSequenceDecodeToUint16(meta.Endianness, b)
		return float32(v)*meta.Factor + float32(meta.Offset), nil
	case c_base.RInt32:
		b, err := checkAndGetBytes(canData, meta, 4)
		if err != nil {
			return nil, err
		}
		v := ECharSequenceDecodeToInt32(meta.Endianness, b)
		return float32(v)*meta.Factor + float32(meta.Offset), nil
	case c_base.RUint32:
		b, err := checkAndGetBytes(canData, meta, 4)
		if err != nil {
			return nil, err
		}
		v := ECharSequenceDecodeToUint32(meta.Endianness, b)
		return float32(v)*meta.Factor + float32(meta.Offset), nil

	case c_base.RInt64:
		b, err := checkAndGetBytes(canData, meta, 8)
		if err != nil {
			return nil, err
		}
		v := ECharSequenceDecodeToInt64(meta.Endianness, b)
		return float32(v)*meta.Factor + float32(meta.Offset), nil
	case c_base.RUint64:
		b, err := checkAndGetBytes(canData, meta, 8)
		if err != nil {
			return nil, err
		}
		v := ECharSequenceDecodeToUint64(meta.Endianness, b)
		return float32(v)*meta.Factor + float32(meta.Offset), nil
	case c_base.RFloat32:
		b, err := checkAndGetBytes(canData, meta, 8)
		if err != nil {
			return nil, err
		}
		v := ECharSequenceDecodeToFloat32(meta.Endianness, b)
		return float32(v)*float32(meta.Factor) + float32(meta.Offset), nil
	case c_base.RFloat64:
		b, err := checkAndGetBytes(canData, meta, 8)
		if err != nil {
			return nil, err
		}
		v := ECharSequenceDecodeToFloat64(meta.Endianness, b)
		return float32(v)*float32(meta.Factor) + float32(meta.Offset), nil

	default:
		return nil, fmt.Errorf("Unsupported ReadType: %v", meta.ReadType)
	}
}

// parseBitValue 专门处理 RBitX 类型的位级数据读取。
//
// 此函数负责从 CANbus 数据中精确读取指定位置和长度的位数据。
// 支持单位布尔值读取和多位无符号整数读取，正确处理位索引转换和字节序。
//
// 工作原理:
//  1. 计算绝对位偏移量 (字节索引 * 8 + 位索引)
//  2. 确定需要读取的字节范围
//  3. **关键改进：先根据字节序重排字节**
//  4. 将重排后的字节转换为位数组
//  5. 直接根据位偏移提取目标位（无需复杂索引转换）
//
// 参数:
//   - canData: CANbus 原始数据
//   - meta: 包含位读取配置的元数据，关键字段包括:
//   - Addr: 字节地址索引
//   - ReadType: 位偏移 (RBit0-RBit15)
//   - BitLength: 读取的位数 (0或1表示单位，>1表示多位)
//   - Endianness: 字节序，影响多位数据的解释
//
// 返回值:
//   - any: 单位返回bool，多位返回uint8/16/32/64
//   - error: 读取范围越界或其他错误
//
// 使用示例:
//
//	// 读取第3个字节的第0位 (最低位)
//	meta := &c_base.Meta{
//	    Addr: 2, ReadType: c_base.RBit0, BitLength: 1
//	}
//	result, err := parseBitValue([]byte{0x01, 0x00, 0x07}, meta)
//	// result = true (因为0x07的第0位是1)
func parseBitValue(canData []byte, meta *c_base.Meta) (any, error) {
	// 计算从 canData 开始的绝对比特偏移量。
	// Meta.Addr 是字节索引。ReadType (RBitX) 表示在该字节起始的"概念性 16 位字"中的比特偏移。
	// 例如：Meta.Addr = 2, ReadType = RBit7 -> 意味着从 canData[2] 的第 7 位开始读取。
	// 绝对比特起始位置 = 字节索引 * 8 + 字节内的比特偏移量
	absoluteBitStart := int(meta.Addr)*8 + int(meta.ReadType) // 例如：Addr 2, RBit7 -> 2*8 + 7 = 23 (绝对比特 23)

	actualBitLength := int(meta.BitLength)
	if actualBitLength == 0 {
		actualBitLength = 1 // 如果 BitLength 为 0，则默认读取 1 位
	}

	absoluteBitEnd := absoluteBitStart + actualBitLength

	// 确定从 canData 中需要读取的字节范围，以覆盖所有必需的比特。
	startByteIndex := absoluteBitStart / 8
	endByteIndex := (absoluteBitEnd + 7) / 8 // +7 用于向上取整，确保完全覆盖相关比特所需的字节

	if endByteIndex > len(canData) {
		return nil, fmt.Errorf("Bit read data range out of bounds. Start byte: %d, End byte (exclusive): %d, Data Len: %d", startByteIndex, endByteIndex, len(canData))
	}

	// 提取相关字节
	relevantBytes := canData[startByteIndex:endByteIndex]

	// **关键改进：先根据字节序重排字节**
	// 这样可以确保多字节数据在不同字节序下的正确解释
	// 例如：[0x12, 0x34] 在小端序下应该变成 [0x34, 0x12]
	reorderedBytes := ECharSequenceFillUpSize(meta.Endianness, relevantBytes, len(relevantBytes))

	// 将重排后的字节转换为平面比特切片。
	// gbinary.DecodeBytesToBits 按字节顺序处理，且在每个字节内从 MSB 到 LSB。
	// 例如：字节 0x86 (10000110) -> [1, 0, 0, 0, 0, 1, 1, 0]
	allBits := gbinary.DecodeBytesToBits(reorderedBytes)

	// 计算在重排后的位数组中的起始比特位置
	// 由于已经进行了字节序重排，现在可以直接使用位偏移
	bitOffsetInReorderedBytes := absoluteBitStart - (startByteIndex * 8)

	// **简化的位提取：由于字节已经重排，现在可以直接使用 LSB 偏移**
	// 但仍需要处理 gbinary.DecodeBytesToBits 的 MSB-first 特性
	totalBitsInReorderedBytes := len(allBits)

	// 转换位索引：从LSB索引转换为MSB数组索引
	// 例如：总共8位，要取第0位(LSB) -> 数组索引 = 8-1-0 = 7
	msbStartIndex := totalBitsInReorderedBytes - bitOffsetInReorderedBytes - actualBitLength
	msbEndIndex := msbStartIndex + actualBitLength

	// 边界检查
	if msbStartIndex < 0 || msbEndIndex > totalBitsInReorderedBytes {
		return nil, fmt.Errorf("Bit extraction within reordered bits out of bounds. MSB start: %d, MSB end: %d, Total bits: %d", msbStartIndex, msbEndIndex, totalBitsInReorderedBytes)
	}

	// 切片 `allBits` 以获取目标比特。
	targetBits := allBits[msbStartIndex:msbEndIndex]

	// 如果只请求一个比特，返回布尔值。
	if actualBitLength == 1 {
		return targetBits[0] == 1, nil
	}

	// 对于多比特值：
	// 1. 根据字节序填充比特 (FillUpSizeBit 处理此逻辑)。
	// 2. 将填充后的比特编码为字节 (gbinary.EncodeBitsToBytes 通常按大端字节序从比特流生成字节)。
	// 3. 将字节转换为最终的无符号整型。
	filledBits := ECharSequenceFillUpSizeBit(meta.Endianness, targetBits, actualBitLength)
	toBytes := gbinary.EncodeBitsToBytes(filledBits) // 这将根据 'filledBits' 的顺序生成字节

	// 现在将生成的字节转换为目标无符号整型。
	// `convertBitsToUint` 将根据长度处理最终的解释。
	return convertBitsToUint(toBytes, actualBitLength, meta.Endianness)
}

// checkAndGetBytes 解析多字节整数和浮点数的通用辅助函数。

func checkAndGetBytes(canData []byte, meta *c_base.Meta, size int) ([]byte, error) {
	if int(meta.Addr)+size > len(canData) {
		return nil, fmt.Errorf("Read out of bounds for ReadType %v. Addr: %d, Size: %d, Data Len: %d", meta.ReadType, meta.Addr, size, len(canData))
	}
	return canData[meta.Addr : meta.Addr+uint16(size)], nil
}

// convertBitsToUint 将位数据转换为相应的无符号整数类型。
//
// 此函数专门处理从位级读取得到的数据，根据位长度自动选择合适的整数类型。
// 正确处理字节序转换，确保位数据能够正确解释为数值。
//
// 工作原理:
//  1. 根据字节序决定是否需要字节反转
//  2. 根据位长度选择相应的整数类型 (uint8/16/32/64)
//  3. 使用 gbinary 解码函数完成最终转换
//
// 参数:
//   - dataBytes: 从位数据编码得到的字节切片
//   - bitLength: 原始位数据的长度 (1-64)
//   - endianness: 字节序，影响多字节数据的解释
//
// 返回值:
//   - any: 根据位长度返回 uint8/16/32/64 类型的无符号整数
//   - error: 数据不足或位长度超出支持范围的错误
//
// 类型选择规则:
//   - 1-8 位 -> uint8
//   - 9-16 位 -> uint16
//   - 17-32 位 -> uint32
//   - 33-64 位 -> uint64
//
// 字节序处理:
//   - 小端序: 需要反转字节顺序
//   - 其他字节序: 保持原有顺序
func convertBitsToUint(dataBytes []byte, bitLength int, endianness c_base.ECharSequence) (any, error) {
	// gbinary.EncodeBitsToBytes 通常按大端字节序从比特流生成字节。
	// 如果目标 `endianness` 是小端，且 `DecodeToUintX` 期望小端字节，
	// 那么 `dataBytes` 可能需要在这里反转。
	// 这是将比特流转换为小端整型的常见模式。
	if endianness == c_base.ELittleEndian {
		// 如果需要小端解释，反转字节
		for i, j := 0, len(dataBytes)-1; i < j; i, j = i+1, j-1 {
			dataBytes[i], dataBytes[j] = dataBytes[j], dataBytes[i]
		}
	}

	var val any
	switch {
	case bitLength <= 8:
		if len(dataBytes) == 0 {
			return nil, gerror.New("Not enough bytes for Uint8 conversion after bit encoding")
		}
		// 假设 gbinary.DecodeToUint8 接收 []byte 并进行转换。
		// 如果它只接收一个 byte 类型，则需要 dataBytes[0]。
		val = gbinary.DecodeToUint8(dataBytes)
	case bitLength <= 16:
		if len(dataBytes) < 2 {
			return nil, gerror.New("Not enough bytes for Uint16 conversion after bit encoding")
		}
		val = gbinary.DecodeToUint16(dataBytes)
	case bitLength <= 32:
		if len(dataBytes) < 4 {
			return nil, gerror.New("Not enough bytes for Uint32 conversion after bit encoding")
		}
		val = gbinary.DecodeToUint32(dataBytes)
	case bitLength <= 64:
		if len(dataBytes) < 8 {
			return nil, gerror.New("Not enough bytes for Uint64 conversion after bit encoding")
		}
		val = gbinary.DecodeToUint64(dataBytes)
	default:
		return nil, fmt.Errorf("Bit length %d exceeds max supported 64-bit integer", bitLength)
	}

	return val, nil
}
