package c_base

import "common/c_enum"

// SDataAccess 数据访问配置
type SDataAccess struct {
	// 位范围配置
	BitIndex  uint16 `json:"bitIndex,omitempty" v:"min:0" dc:"位起始索引"`
	BitLength uint16 `json:"bitLength,omitempty" v:"max:64" dc:"位长度（0表示纯字节模式）"`

	// 数据范围配置
	ByteIndex  uint16 `json:"byteIndex,omitempty" v:"min:0" dc:"字节起始索引"`
	ByteLength uint16 `json:"byteLength,omitempty" v:"max:255" dc:"字节长度（0表示使用DataFormat默认长度）"`

	// 数据格式配置
	DataFormat c_enum.EDataFormat `json:"dataFormat" v:"required" dc:"数据格式"`
	ByteEndian c_enum.EByteEndian `json:"byteEndian,omitempty" dc:"字节序（默认: ByteEndianBig）"`
	WordOrder  c_enum.EWordOrder  `json:"wordOrder,omitempty" dc:"字序（默认: WordOrderHighLow）"`

	// 数据转换配置
	ValueType c_enum.EValueType `json:"valueType,omitempty" dc:"返回值类型"`
	Factor    float32           `json:"factor,omitempty" dc:"系数（默认: 0.0 不乘以系数）"`
	Offset    int               `json:"offset,omitempty" dc:"偏移值（默认: 0）"`

	// 验证配置
	MinValue int64 `json:"minValue,omitempty" v:"gte:0" dc:"最小值验证（0表示不验证）"`
	MaxValue int64 `json:"maxValue,omitempty" v:"gte:0|lte:MinValue" dc:"最大值验证（0表示不验证）"`
}
