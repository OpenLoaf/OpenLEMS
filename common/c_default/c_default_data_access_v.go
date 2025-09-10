package c_default

import (
	"common/c_base"
	"common/c_enum"
)

// Modbus 常用默认数据访问配置
var (
	VDataAccessUInt16      = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint16, Factor: 1.0, Offset: 0}      // 16位无符号整数 (标准Modbus)
	VDataAccessInt16       = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EInt16, Factor: 1.0, Offset: 0}        // 16位有符号整数
	VDataAccessUInt32      = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32, Factor: 1.0, Offset: 0}      // 32位无符号整数 (2个寄存器)
	VDataAccessInt32       = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EInt32, Factor: 1.0, Offset: 0}        // 32位有符号整数 (2个寄存器)
	VDataAccessFloat32     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatFloat32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 1.0, Offset: 0}    // 32位浮点数 (2个寄存器)
	VDataAccessFloat64     = &c_base.SDataAccess{DataFormat: c_enum.DataFormatFloat64, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 1.0, Offset: 0}    // 64位浮点数 (4个寄存器)
	VDataAccessBCD16       = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBCD, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint16, Factor: 1.0, Offset: 0}         // 16位BCD码
	VDataAccessBCD32       = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBCD32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32, Factor: 1.0, Offset: 0}       // 32位BCD码 (2个寄存器)
	VDataAccessBits        = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBits, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint16, Factor: 1.0, Offset: 0}        // 位图/状态字
	VDataAccessBitRange    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool, Factor: 1.0, Offset: 0}      // 位范围 (用于提取特定位)
	VDataAccessStringASCII = &c_base.SDataAccess{DataFormat: c_enum.DataFormatStringASCII, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EString, Factor: 1.0, Offset: 0} // ASCII字符串
	VDataAccessStringUTF16 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatStringUTF16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EString, Factor: 1.0, Offset: 0} // UTF-16字符串
)

// 小端字节序变体 (适用于某些设备)
var (
	VDataAccessUInt16Little  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianLittle, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint16, Factor: 1.0, Offset: 0}   // 16位无符号整数 (小端字节序)
	VDataAccessUInt32Little  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianLittle, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint32, Factor: 1.0, Offset: 0}   // 32位无符号整数 (小端字节序)
	VDataAccessFloat32Little = &c_base.SDataAccess{DataFormat: c_enum.DataFormatFloat32, ByteEndian: c_enum.ByteEndianLittle, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 1.0, Offset: 0} // 32位浮点数 (小端字节序)
)

// 字序交换变体 (适用于某些设备)
var (
	VDataAccessUInt32WordSwap  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderLowHigh, ValueType: c_enum.EUint32, Factor: 1.0, Offset: 0}   // 32位无符号整数 (字序交换)
	VDataAccessFloat32WordSwap = &c_base.SDataAccess{DataFormat: c_enum.DataFormatFloat32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderLowHigh, ValueType: c_enum.EFloat32, Factor: 1.0, Offset: 0} // 32位浮点数 (字序交换)
)

// 带缩放因子的常用配置
var (
	VDataAccessUInt16Scale0001 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 0.001, Offset: 0} // 16位无符号整数 (缩放因子0.001)
	VDataAccessUInt16Scale01   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 0.1, Offset: 0}   // 16位无符号整数 (缩放因子0.1)
	VDataAccessUInt16Scale001  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 0.01, Offset: 0}  // 16位无符号整数 (缩放因子0.01)
	VDataAccessUInt16Scale10   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 10.0, Offset: 0}  // 16位无符号整数 (缩放因子10)
	VDataAccessInt16Scale01    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 0.1, Offset: 0}    // 16位有符号整数 (缩放因子0.1)
	VDataAccessInt16Scale001   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 0.01, Offset: 0}   // 16位有符号整数 (缩放因子0.01)
	VDataAccessInt16Scale10    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat32, Factor: 10.0, Offset: 0}   // 16位有符号整数 (缩放因子10)
	VDataAccessUInt32Scale01   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.1, Offset: 0}   // 32位无符号整数 (缩放因子0.1)
	VDataAccessUInt32Scale001  = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.01, Offset: 0}  // 32位无符号整数 (缩放因子0.01)
	VDataAccessUInt32Scale10   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 10.0, Offset: 0}  // 32位无符号整数 (缩放因子10)
	VDataAccessInt32Scale01    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.1, Offset: 0}    // 32位有符号整数 (缩放因子0.1)
	VDataAccessInt32Scale001   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 0.01, Offset: 0}   // 32位有符号整数 (缩放因子0.01)
	VDataAccessInt32Scale10    = &c_base.SDataAccess{DataFormat: c_enum.DataFormatInt32, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EFloat64, Factor: 10.0, Offset: 0}   // 32位有符号整数 (缩放因子10)
)

// 位范围配置 (用于状态位)
var (
	VDataAccessBitRangeUInt16 = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EUint16} // 位范围 (16位无符号整数)
	VDataAccessBitRangeBool   = &c_base.SDataAccess{DataFormat: c_enum.DataFormatBitRange, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool}   // 位范围 (布尔值)
)

// 布尔类型配置 (UInt16转Bool)
var (
	VDataAccessUInt16ToBool = &c_base.SDataAccess{DataFormat: c_enum.DataFormatUInt16, ByteEndian: c_enum.ByteEndianBig, WordOrder: c_enum.WordOrderHighLow, ValueType: c_enum.EBool} // 16位无符号整数转布尔值
)
