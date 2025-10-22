package c_base

import (
	"common/c_enum"
	"fmt"
)

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
	//ValueType c_enum.EValueType `json:"valueType,omitempty" dc:"返回值类型"`
	Factor float32 `json:"factor,omitempty" dc:"系数（默认: 0.0 不乘以系数）"`
	Offset int     `json:"offset,omitempty" dc:"偏移值（默认: 0）"`
}

func (s *SDataAccess) String() string {
	return fmt.Sprintf("DataAccess[bit:%d:%d, byte:%d:%d]",
		s.BitIndex, s.BitLength,
		s.ByteIndex, s.ByteLength)
}

// GetDecimalPlaces 根据Factor推断小数位数
// 该方法通过Factor的数值范围来确定结果应保留的小数位数
// 例如: Factor=0.1 返回1, Factor=0.01 返回2, Factor=2.5 返回1
func (s *SDataAccess) GetDecimalPlaces() int {
	// Factor为0时不进行转换，返回0位小数
	if s.Factor == 0.0 {
		return 0
	}

	// 使用绝对值处理负数因子
	factor := s.Factor
	if factor < 0 {
		factor = -factor
	}

	// 使用范围判断，避免浮点数精度问题
	const epsilon = 0.0000001 // float32精度容差

	// 0位小数: [1.0, +∞) 或者是 10, 100, 1000 等整数倍
	if factor >= 1.0-epsilon {
		// 检查是否为整数
		if factor-float32(int(factor)) < epsilon {
			return 0
		}
		// 1 < factor < 10，可能有小数位
		if factor < 10.0 {
			// 可能是 1.5, 2.5, 3.14 等
			// 继续往下判断
		} else {
			// factor >= 10，认为是整数倍数
			return 0
		}
	}

	// 1位小数: [0.1, 1.0)
	if factor >= 0.1-epsilon && factor < 1.0+epsilon {
		// 检查是否接近 0.1, 0.2, ..., 0.9
		scaled := factor * 10.0
		if scaled-float32(int(scaled+0.5)) < epsilon*10 {
			return 1
		}
	}

	// 2位小数: [0.01, 0.1)
	if factor >= 0.01-epsilon && factor < 0.1+epsilon {
		// 检查是否接近 0.01, 0.02, ..., 0.09
		scaled := factor * 100.0
		if scaled-float32(int(scaled+0.5)) < epsilon*100 {
			return 2
		}
	}

	// 3位小数: [0.001, 0.01)
	if factor >= 0.001-epsilon && factor < 0.01+epsilon {
		scaled := factor * 1000.0
		if scaled-float32(int(scaled+0.5)) < epsilon*1000 {
			return 3
		}
	}

	// 4位小数: [0.0001, 0.001)
	if factor >= 0.0001-epsilon && factor < 0.001+epsilon {
		scaled := factor * 10000.0
		if scaled-float32(int(scaled+0.5)) < epsilon*10000 {
			return 4
		}
	}

	// 5位小数: [0.00001, 0.0001)
	if factor >= 0.00001-epsilon && factor < 0.0001+epsilon {
		return 5
	}

	// 6位小数: [0.000001, 0.00001)
	if factor >= 0.000001-epsilon && factor < 0.00001+epsilon {
		return 6
	}

	// 默认返回2位小数（适用于大部分场景）
	return 2
}
