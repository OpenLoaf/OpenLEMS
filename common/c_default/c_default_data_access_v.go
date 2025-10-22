package c_default

import (
	"common/c_base"
	"common/c_enum"
)

// Modbus 常用默认数据访问配置
var (
	VDataAccessUInt16      = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}      // 16位无符号整数 (标准Modbus)
	VDataAccessInt16       = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}       // 16位有符号整数
	VDataAccessUInt32      = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}      // 32位无符号整数 (2个寄存器)
	VDataAccessInt32       = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}       // 32位有符号整数 (2个寄存器)
	VDataAccessFloat32     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatFloat32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}     // 32位浮点数 (2个寄存器)
	VDataAccessFloat64     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatFloat64, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}     // 64位浮点数 (4个寄存器)
	VDataAccessBCD16       = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBCD, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}         // 16位BCD码
	VDataAccessBCD32       = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBCD32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}       // 32位BCD码 (2个寄存器)
	VDataAccessBits        = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBits, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}        // 位图/状态字
	VDataAccessBitRange    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}    // 位范围 (用于提取特定位)
	VDataAccessStringASCII = &c_base.SDataAccess{DataFormat: c_enum.DataFormatStringASCII, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // ASCII字符串
	VDataAccessStringUTF16 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatStringUTF16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // UTF-16字符串
)

// 小端字节序变体 (适用于某些设备)
var (
	VDataAccessUInt16Little  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianLittle, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}  // 16位无符号整数 (小端字节序)
	VDataAccessUInt32Little  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianLittle, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}  // 32位无符号整数 (小端字节序)
	VDataAccessFloat32Little = &c_base.SDataAccess{DataFormat: c_enum.DataFormatFloat32, ByteEndian: c_enum.ByteEndianLittle, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // 32位浮点数 (小端字节序)
)

// 字序交换变体 (适用于某些设备)
var (
	VDataAccessUInt32WordSwap  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderLowHigh, Factor: 1.0, Offset: 0}  // 32位无符号整数 (字序交换)
	VDataAccessFloat32WordSwap = &c_base.SDataAccess{DataFormat: c_enum.DataFormatFloat32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderLowHigh, Factor: 1.0, Offset: 0} // 32位浮点数 (字序交换)
)

// 带缩放因子的常用配置
var (
	VDataAccessUInt16Scale00001 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.0001, Offset: 0} // 16位无符号整数
	VDataAccessUInt16Scale0001  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0}  // 16位无符号整数
	VDataAccessUInt16Scale001   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.01, Offset: 0}   // 16位无符号整数
	VDataAccessUInt16Scale01    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}    // 16位无符号整数 (缩放因子0.1)
	VDataAccessUInt16Scale10    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 10.0, Offset: 0}   // 16位无符号整数 (缩放因子10)
	VDataAccessInt16Scale00001  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.0001, Offset: 0}  // 16位有符号整数
	VDataAccessInt16Scale0001   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0}   // 16位有符号整数
	VDataAccessInt16Scale001    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.01, Offset: 0}    // 16位有符号整数
	VDataAccessInt16Scale01     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}     // 16位有符号整数 (缩放因子0.1)
	VDataAccessInt16Scale10     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 10.0, Offset: 0}    // 16位有符号整数 (缩放因子10)
	VDataAccessUInt32Scale00001 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.0001, Offset: 0} // 32位无符号整数
	VDataAccessUInt32Scale0001  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0}  // 32位无符号整数
	VDataAccessUInt32Scale001   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.01, Offset: 0}   // 32位无符号整数
	VDataAccessUInt32Scale01    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}    // 32位无符号整数 (缩放因子0.1)
	VDataAccessUInt32Scale10    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 10.0, Offset: 0}   // 32位无符号整数 (缩放因子10)
	VDataAccessInt32Scale00001  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.0001, Offset: 0}  // 32位有符号整数
	VDataAccessInt32Scale0001   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0}   // 32位有符号整数
	VDataAccessInt32Scale001    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.01, Offset: 0}    // 32位有符号整数
	VDataAccessInt32Scale01     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}     // 32位有符号整数 (缩放因子0.1)
	VDataAccessInt32Scale10     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 10.0, Offset: 0}    // 32位有符号整数 (缩放因子10)
)

// 位范围配置 (用于状态位)
var (
	VDataAccessBitRangeUInt16 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow} // 位范围 (16位无符号整数)
	VDataAccessBitRangeBool   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow} // 位范围 (布尔值)
)

// 布尔类型配置 (UInt16转Bool)
var (
	VDataAccessUInt16ToBool = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow} // 16位无符号整数转布尔值
)

// Canbus常用
var (
	VDataAccessFloat32Byte0Scale01  = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessFloat32Byte1Scale01  = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessFloat32Byte2Scale01  = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessFloat32Byte3Scale01  = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessFloat32Byte0Scale001 = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessFloat32Byte1Scale001 = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessFloat32Byte2Scale001 = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessFloat32Byte3Scale001 = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}

	VDataAccessFloat64Byte0Scale01  = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessFloat64Byte1Scale01  = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessFloat64Byte2Scale01  = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessFloat64Byte3Scale01  = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessFloat64Byte0Scale001 = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessFloat64Byte1Scale001 = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessFloat64Byte2Scale001 = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessFloat64Byte3Scale001 = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}

	// 更多缩放因子变体
	VDataAccessFloat32Byte0Scale1   = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 1.0}
	VDataAccessFloat32Byte1Scale1   = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 1.0}
	VDataAccessFloat32Byte2Scale1   = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 1.0}
	VDataAccessFloat32Byte3Scale1   = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 1.0}
	VDataAccessFloat32Byte0Scale10  = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 10.0}
	VDataAccessFloat32Byte1Scale10  = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 10.0}
	VDataAccessFloat32Byte2Scale10  = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 10.0}
	VDataAccessFloat32Byte3Scale10  = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 10.0}
	VDataAccessFloat32Byte0Scale100 = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 100.0}
	VDataAccessFloat32Byte1Scale100 = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 100.0}
	VDataAccessFloat32Byte2Scale100 = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 100.0}
	VDataAccessFloat32Byte3Scale100 = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 100.0}

	// 无符号整数变体
	VDataAccessUInt16Byte0Scale01  = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatUInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessUInt16Byte1Scale01  = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatUInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessUInt16Byte2Scale01  = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatUInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessUInt16Byte3Scale01  = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatUInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.1}
	VDataAccessUInt16Byte0Scale001 = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 2, DataFormat: c_enum.DataFormatUInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessUInt16Byte1Scale001 = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 2, DataFormat: c_enum.DataFormatUInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessUInt16Byte2Scale001 = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 2, DataFormat: c_enum.DataFormatUInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}
	VDataAccessUInt16Byte3Scale001 = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 2, DataFormat: c_enum.DataFormatUInt16, WordOrder: c_enum.WordOrderLowHigh, Factor: 0.001}

	// 单字节配置
	VDataAccessUInt8Byte0 = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessUInt8Byte1 = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessUInt8Byte2 = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessUInt8Byte3 = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessUInt8Byte4 = &c_base.SDataAccess{ByteIndex: 4, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessUInt8Byte5 = &c_base.SDataAccess{ByteIndex: 5, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessUInt8Byte6 = &c_base.SDataAccess{ByteIndex: 6, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessUInt8Byte7 = &c_base.SDataAccess{ByteIndex: 7, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}

	VDataAccessInt8Byte0 = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 1, DataFormat: c_enum.DataFormatInt8, Factor: 1.0}
	VDataAccessInt8Byte1 = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 1, DataFormat: c_enum.DataFormatInt8, Factor: 1.0}
	VDataAccessInt8Byte2 = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 1, DataFormat: c_enum.DataFormatInt8, Factor: 1.0}
	VDataAccessInt8Byte3 = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 1, DataFormat: c_enum.DataFormatInt8, Factor: 1.0}
	VDataAccessInt8Byte4 = &c_base.SDataAccess{ByteIndex: 4, ByteLength: 1, DataFormat: c_enum.DataFormatInt8, Factor: 1.0}
	VDataAccessInt8Byte5 = &c_base.SDataAccess{ByteIndex: 5, ByteLength: 1, DataFormat: c_enum.DataFormatInt8, Factor: 1.0}
	VDataAccessInt8Byte6 = &c_base.SDataAccess{ByteIndex: 6, ByteLength: 1, DataFormat: c_enum.DataFormatInt8, Factor: 1.0}
	VDataAccessInt8Byte7 = &c_base.SDataAccess{ByteIndex: 7, ByteLength: 1, DataFormat: c_enum.DataFormatInt8, Factor: 1.0}

	// 布尔值配置 (单字节转布尔)
	VDataAccessBoolByte0 = &c_base.SDataAccess{ByteIndex: 0, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessBoolByte1 = &c_base.SDataAccess{ByteIndex: 1, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessBoolByte2 = &c_base.SDataAccess{ByteIndex: 2, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessBoolByte3 = &c_base.SDataAccess{ByteIndex: 3, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessBoolByte4 = &c_base.SDataAccess{ByteIndex: 4, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessBoolByte5 = &c_base.SDataAccess{ByteIndex: 5, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessBoolByte6 = &c_base.SDataAccess{ByteIndex: 6, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
	VDataAccessBoolByte7 = &c_base.SDataAccess{ByteIndex: 7, ByteLength: 1, DataFormat: c_enum.DataFormatUInt8, Factor: 1.0}
)

// 温度相关配置 (常用缩放因子)
var (
	VDataAccessTempCelsiusScale01    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}   // 温度 (摄氏度, 缩放0.1)
	VDataAccessTempCelsiusScale001   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0} // 温度 (摄氏度, 缩放0.001)
	VDataAccessTempCelsiusScale1     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}   // 温度 (摄氏度, 无缩放)
	VDataAccessTempFahrenheitScale01 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 32}  // 温度 (华氏度, 缩放0.1, 偏移32)
)

// 电压电流功率相关配置
var (
	VDataAccessVoltageScale01  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}    // 电压 (缩放0.1V)
	VDataAccessVoltageScale001 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0}  // 电压 (缩放0.001V)
	VDataAccessVoltageScale1   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}    // 电压 (无缩放V)
	VDataAccessCurrentScale01  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}    // 电流 (缩放0.1A)
	VDataAccessCurrentScale001 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0}  // 电流 (缩放0.001A)
	VDataAccessCurrentScale1   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}    // 电流 (无缩放A)
	VDataAccessPowerScale01    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}    // 功率 (缩放0.1W)
	VDataAccessPowerScale001   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0}  // 功率 (缩放0.001W)
	VDataAccessPowerScale1     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}    // 功率 (无缩放W)
	VDataAccessPowerScale1000  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1000.0, Offset: 0} // 功率 (缩放1000W)
)

// 频率相关配置
var (
	VDataAccessFreqScale01  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}   // 频率 (缩放0.1Hz)
	VDataAccessFreqScale001 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0} // 频率 (缩放0.001Hz)
	VDataAccessFreqScale1   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}   // 频率 (无缩放Hz)
)

// 百分比相关配置
var (
	VDataAccessPercentScale01  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.1, Offset: 0}   // 百分比 (缩放0.1%)
	VDataAccessPercentScale001 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 0.001, Offset: 0} // 百分比 (缩放0.001%)
	VDataAccessPercentScale1   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}   // 百分比 (无缩放%)
)

// 时间相关配置
var (
	VDataAccessTimeSecond = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // 时间 (秒)
	VDataAccessTimeMinute = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // 时间 (分钟)
	VDataAccessTimeHour   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // 时间 (小时)
	VDataAccessTimeDay    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // 时间 (天)
)

// 计数器相关配置
var (
	VDataAccessCounter32 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // 32位计数器
	VDataAccessCounter16 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0} // 16位计数器
	VDataAccessCounter8  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt8, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, Factor: 1.0, Offset: 0}  // 8位计数器
)
