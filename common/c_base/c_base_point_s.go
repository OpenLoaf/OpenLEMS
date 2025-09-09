package c_base

import (
	"encoding/binary"
	"math"
	"strings"
	"time"
	"unicode/utf16"

	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

type EPointAddressType int

const (
	EPointAddressTypeByte = iota
	EPointAddressTypeBit
)

// EByteEndian 字节序 (针对一个16位寄存器内部的字节顺序)
type EByteEndian int

const (
	ByteEndianBig    EByteEndian = iota // 大端字节序 (标准Modbus) [AB] -> 高字节A在前
	ByteEndianLittle                    // 小端字节序 [BA] -> 低字节B在前
)

// EWordOrder 字序 (针对多个16位寄存器之间的顺序)
type EWordOrder int

const (
	WordOrderHighLow EWordOrder = iota // 高字在前，低字在后 (标准Modbus) [AB CD]
	WordOrderLowHigh                   // 低字在前，高字在后 (字交换) [CD AB]
)

// EDataFormat 数据格式/编码
type EDataFormat int

const (
	DataFormatUInt16      EDataFormat = iota // 16位无符号整数
	DataFormatInt16                          // 16位有符号整数 (二进制补码)
	DataFormatUInt32                         // 32位无符号整数
	DataFormatInt32                          // 32位有符号整数
	DataFormatFloat32                        // 32位IEEE 754浮点数
	DataFormatFloat64                        // 64位IEEE 754浮点数
	DataFormatBCD                            // 二进制编码的十进制数 (16位寄存器)
	DataFormatBCD32                          // 二进制编码的十进制数 (32位，占2个寄存器)
	DataFormatStringASCII                    // ASCII字符串 (每个字节一个字符)
	DataFormatStringUTF16                    // UTF-16字符串 (每2字节一个字符)
	DataFormatBits                           // 位图(比特位)，用于线圈、状态字等
	DataFormatBitRange                       // 位范围，用于提取特定范围的位
	DataFormatCustom                         // 自定义格式，使用自定义解析函数
)

type IPoint interface {
	GetKey() string
	GetName() string
	GetGroup() *MetaGroup
	GetLevel() EAlarmLevel
	GetUnit() string
	GetDesc() string
	GetSort() int
	AlarmTrigger(value any) (bool, error)   // 判断触发或者消除告警
	ValueExplain(value any) (string, error) // 获取Value解释，一般为状态类型的解释
}

// SPoint 点位元数据
type SPoint struct {
	Key     string      `json:"key" v:"required"`  // 名称
	Name    string      `json:"name" v:"required"` // 名称
	Group   *MetaGroup  `json:"group" dc:"分组"`
	Unit    string      `json:"unit,omitempty"`    // 单位
	Desc    string      `json:"desc,omitempty"`    // 备注
	Sort    int         `json:"sort"`              // 排序
	Level   EAlarmLevel `json:"level"`             // 点位级别
	Min     int64       `json:"min,omitempty"`     // 范围最小值
	Max     int64       `json:"max,omitempty"`     // 范围最大值
	Precise int         `json:"precise,omitempty"` // 设置浮点数精度（只是显示用）
}

func (s *SPoint) GetKey() string {
	return s.Key
}

func (s *SPoint) GetName() string {
	return s.Name
}

func (s *SPoint) GetGroup() *MetaGroup {
	return s.Group
}

func (s *SPoint) GetLevel() EAlarmLevel {
	return s.Level
}

func (s *SPoint) GetUnit() string {
	return s.Unit
}

func (s *SPoint) GetDesc() string {
	return s.Desc
}

func (s *SPoint) GetSort() int {
	return s.Sort
}

type SPointValue struct {
	IPoint
	value      any
	happenTime time.Time
}

func (s *SPointValue) GetValue() any {
	return s.value
}

func (s *SPointValue) GetValueExplain() (string, error) {
	if s.IPoint != nil {
		if explainer, ok := s.IPoint.(interface {
			ValueExplain(value any) (string, error)
		}); ok {
			return explainer.ValueExplain(s.value)
		}
	}
	return cvt.StringE(s.value)
}

func (s *SPointValue) GetHappenTime() time.Time {
	return s.happenTime
}

type SModbusPoint struct {
	*SPoint
	Address     uint16
	Length      uint16
	AddressType EPointAddressType // byte或者bit

	// -- 解析相关配置 --
	ByteEndian EByteEndian `json:"byteEndian"` // 字节序 (默认: ByteEndianBig)
	WordOrder  EWordOrder  `json:"wordOrder"`  // 字序 (默认: WordOrderHighLow)
	DataFormat EDataFormat `json:"dataFormat"` // 数据格式 (必须明确指定)

	Type   ESystemType
	Factor float32 `json:"factor,omitempty"` // 乘以系数，如果是0，自动会改成1。因为0无意义
	Offset int     `json:"offset,omitempty"` // 偏移值

	StatusExplain func(value any) (string, error)       `json:"-"` // 状态解释，如果有的话就翻译一下状态
	Trigger       func(value interface{}) (bool, error) `json:"-"` // 触发告警警告故障信息
}

func (s *SModbusPoint) AlarmTrigger(value any) (bool, error) {
	if s.Trigger == nil {
		return false, nil
	}

	return s.Trigger(value)
}

func (s *SModbusPoint) ValueExplain(value any) (string, error) {
	if s.StatusExplain == nil {
		return cvt.String(value), nil
	}
	return s.StatusExplain(value)
}

type IPointTask interface {
	GetName() string
	GetDescription() string
	GetPoints() []IPoint // 获取点位
	GetLifeTime() time.Duration
}

type SModbusPointTask struct {
	Name     string
	Desc     string
	Addr     uint16
	Quantity uint16
	//Function       EModbusReadFunction
	CycleMill      int64
	Lifetime       time.Duration // lifetime 为0时候缓存永不过期，为负数时候不缓存并删除缓存的值
	Transitory     bool          // 是短暂的，查询一次后，需要再次调用查询才能查询，而且不会一直轮询。默认是永久查询
	TransitoryTime time.Duration
	Points         []IPoint
	CustomDecoder  func(bytes []byte, task *SModbusPointTask, point IPoint) (any, error) // 手动解析，空代表使用默认的协议解析器
}

type SCanbusTask struct {
	Name          string
	Desc          string
	GetCanbusID   func(params map[string]any) *uint32
	IDMatch       func(canId uint32) bool                       // 判断ID是否匹配，如果为空，直接判断是否和CanbusID相等
	Lifetime      time.Duration                                 // lifetime 为0时候缓存永不过期，为负数时候不缓存并删除缓存的值
	Points        []IPoint                                      // 点位列表
	IsRemote      bool                                          // 是否是远程帧（写重要）
	IsExtended    bool                                          // 是否是扩展帧（写重要）
	CustomDecoder func(bytes []byte, point IPoint) (any, error) // 手动解析，空代表使用默认的协议解析器
}

// var TestSwitchPoint = &SModbusPoint{Address: 0x0001, Length: 2, SPoint: &SPoint{Key: "TestSwitch", Name: "测试"}, Type: SBool, Trigger: cvt.BoolE}
// var TestSystemAlarmPoint = &SModbusPoint{Address: 0x0002, Length: 2, SPoint: &SPoint{Key: "TestSystemAlarmPoint", Name: "测试告警", Level: EAlarmLevelWarn}, Type: SInt8, Trigger: func(value interface{}) (bool, error) {
// 	v, e := cvt.Int8E(value)
// 	if e != nil {
// 		return false, e
// 	}
// 	return v == 2, nil
// }}
// var TestSystemAlarmPoint2 = &SModbusPoint{Address: 0x0002, Length: 2, SPoint: &SPoint{Key: "TestSystemAlarmPoint", Name: "测试告警", Level: EAlarmLevelError}, Type: SInt8, Trigger: func(value interface{}) (bool, error) {
// 	v, e := cvt.Int8E(value)
// 	if e != nil {
// 		return false, e
// 	}
// 	return v == 3, nil
// }}

// var TestTask = &SModbusPointTask{
// 	Addr: 0x0001,
// 	Points: []IPoint{
// 		TestSwitchPoint,
// 	},
// }

// func Test() {
// 	TestSwitchPoint.GetKey()
// }

func (s *SModbusPointTask) Decoder(bytes []byte, point *SModbusPoint) (*SPointValue, error) {
	var value any
	var err error

	if s.CustomDecoder != nil {
		value, err = s.CustomDecoder(bytes, s, point)
		if err != nil {
			return nil, errors.Wrap(err, "custom decoder failed")
		}
	} else {
		// 使用通用的DecoderBytes函数进行解析
		// 根据AddressType决定如何传递参数
		var isBit bool
		var index, length uint16

		if point.AddressType == EPointAddressTypeBit {
			// 位级别：Address和Length都是位级别的
			isBit = true

			// 边界检查：确保点位地址在任务范围内
			// 每个寄存器占16位，所以任务的位地址范围是 [s.Addr*16, (s.Addr+s.Quantity)*16-1]
			taskStartBit := s.Addr * 16
			taskEndBit := (s.Addr+s.Quantity)*16 - 1
			if point.Address < taskStartBit || point.Address > taskEndBit {
				return nil, errors.Errorf("bit address %d is out of task range [%d:%d]",
					point.Address, taskStartBit, taskEndBit)
			}

			// 计算位在字节数组中的索引
			// point.Address 是绝对位地址，s.Addr 是起始寄存器地址
			// 每个寄存器占16位(2字节)
			index = point.Address - taskStartBit // 相对位索引
			length = point.Length
		} else {
			// 字节级别：Address和Length都是寄存器级别的
			isBit = false

			// 边界检查：确保点位地址在任务范围内
			if point.Address < s.Addr || point.Address+point.Length > s.Addr+s.Quantity {
				return nil, errors.Errorf("register address range [%d:%d] is out of task range [%d:%d]",
					point.Address, point.Address+point.Length-1, s.Addr, s.Addr+s.Quantity-1)
			}

			// 计算数据在字节数组中的起始位置
			// point.Address 和 s.Addr 都是寄存器地址，每个寄存器2字节
			registerOffset := point.Address - s.Addr // 寄存器偏移量
			index = registerOffset * 2               // 转换为字节偏移量
			length = point.Length * 2                // 寄存器数量转换为字节数量
		}

		value, err = DecoderBytes(
			bytes,
			index,
			length,
			isBit,
			point.ByteEndian,
			point.WordOrder,
			point.DataFormat,
			point.Type,
			point.Offset,
			point.Factor,
			point.Min,
			point.Max,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decode point %s", point.GetKey())
		}
	}

	return &SPointValue{
		IPoint:     point,
		value:      value,
		happenTime: time.Now(),
	}, nil
}

//func ()  {
//
//}

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
//   - index: 起始字节索引
//   - length: 数据长度（字节数）
//   - byteEndian: 字节序（大端/小端）
//   - wordOrder: 字序（高字在前/低字在前）
//   - dataFormat: 数据格式（整数、浮点数、BCD、字符串等）
//   - returnFormat: 返回格式类型
//   - offset: 偏移量
//   - factor: 系数
//   - min: 最小值验证（数值类型）或最小长度验证（字符串类型，0表示不验证）
//   - max: 最大值验证（数值类型）或最大长度验证（字符串类型，0表示不验证）
//
// 使用示例：
//
//	// Modbus 16位整数解析（带范围验证）
//	result, err := DecoderBytes(data, 0, 2, ByteEndianBig, WordOrderHighLow, DataFormatUInt16, SUseReadType, 0, 1.0, 0, 1000)
//	if err != nil {
//		// 处理错误
//	}
//
//	// CANbus BCD码解析（无范围验证）
//	result, err := DecoderBytes(data, 2, 2, ByteEndianLittle, WordOrderHighLow, DataFormatBCD, SInt32, 0, 0.1, 0, 0)
//	if err != nil {
//		// 处理错误
//	}
//
//	// IEEE 754浮点数解析（带范围验证）
//	result, err := DecoderBytes(data, 0, 4, ByteEndianBig, WordOrderHighLow, DataFormatFloat32, SFloat64, 0, 1.0, -100, 100)
//	if err != nil {
//		// 处理错误
//	}
//
//	// ASCII字符串解析（带长度验证）
//	result, err := DecoderBytes(data, 0, 10, ByteEndianBig, WordOrderHighLow, DataFormatStringASCII, SString, 0, 1.0, 3, 20)
//	if err != nil {
//		// 处理错误
//	}
func DecoderBytes(bytes []byte, index uint16, length uint16, isBit bool, byteEndian EByteEndian, wordOrder EWordOrder, dataFormat EDataFormat, returnFormat ESystemType, offset int, factor float32, min, max int64) (any, error) {
	// 参数验证
	if len(bytes) == 0 {
		return nil, errors.New("empty input data")
	}

	// 确保系数不为0
	if factor == 0 {
		factor = 1.0
	}

	// 根据 isBit 参数计算实际需要的数据
	var data []byte
	var bitStart, bitLength uint8

	if isBit {
		// 位索引模式：index 和 length 都是位级别的
		bitStart = uint8(index % 8) // 在字节内的位偏移
		bitLength = uint8(length)   // 位长度

		// 计算需要的字节数
		byteIndex := index / 8
		requiredBytes := (index + length + 7) / 8

		if uint16(len(bytes)) < requiredBytes {
			return nil, errors.Errorf("insufficient data: need %d bytes for bit range, got %d", requiredBytes, len(bytes))
		}

		data = bytes[byteIndex:requiredBytes]
	} else {
		// 字节索引模式：index 和 length 都是字节级别的
		// 检查数据长度是否足够，如果不够则立即报错而不是截断
		if int(index)+int(length) > len(bytes) {
			return nil, errors.Errorf("insufficient data: requested range [%d:%d] exceeds data length %d", index, index+length, len(bytes))
		}

		data = bytes[index : index+length]
		bitStart = 0
		bitLength = 0
	}

	// 根据数据格式进行解析
	var rawValue any
	var err error

	switch dataFormat {
	case DataFormatUInt16:
		rawValue, err = decodeUInt16(data, byteEndian, wordOrder)
	case DataFormatInt16:
		rawValue, err = decodeInt16(data, byteEndian, wordOrder)
	case DataFormatUInt32:
		rawValue, err = decodeUInt32(data, byteEndian, wordOrder)
	case DataFormatInt32:
		rawValue, err = decodeInt32(data, byteEndian, wordOrder)
	case DataFormatFloat32:
		rawValue, err = decodeFloat32(data, byteEndian, wordOrder)
	case DataFormatFloat64:
		rawValue, err = decodeFloat64(data, byteEndian, wordOrder)
	case DataFormatBCD:
		rawValue, err = decodeBCD16(data, byteEndian, wordOrder)
	case DataFormatBCD32:
		rawValue, err = decodeBCD32(data, byteEndian, wordOrder)
	case DataFormatStringASCII:
		rawValue, err = decodeASCIIString(data)
	case DataFormatStringUTF16:
		rawValue, err = decodeUTF16String(data, byteEndian, wordOrder)
	case DataFormatBits:
		rawValue, err = decodeBits(data, byteEndian, wordOrder)
	case DataFormatBitRange:
		rawValue, err = decodeBitRange(data, bitStart, bitLength)
	case DataFormatCustom:
		// 自定义格式需要外部提供解析函数
		return nil, errors.New("custom data format not supported")
	default:
		return nil, errors.Errorf("unsupported data format: %v", dataFormat)
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to decode data")
	}

	// 应用系数和偏移量
	finalValue := applyFactorAndOffset(rawValue, factor, offset)

	// 转换为目标格式
	result := convertToReturnFormat(finalValue, returnFormat)

	// 数据范围验证
	if err := validateValueRange(result, min, max); err != nil {
		return nil, errors.Wrap(err, "value validation failed")
	}

	return result, nil
}

// 数据范围验证函数
func validateValueRange(value any, min, max int64) error {
	// 如果min和max都为0，表示不进行范围验证
	if min == 0 && max == 0 {
		return nil
	}

	// 字符串类型进行长度验证
	if str, ok := value.(string); ok {
		length := int64(len(str))
		if min > 0 && length < min {
			return errors.Errorf("string length %d is below minimum %d", length, min)
		}
		if max > 0 && length > max {
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
			return errors.Errorf("uint64 value %d exceeds int64 maximum %d", uintVal, math.MaxInt64)
		}
		val := int64(uintVal)
		if min > 0 && val < min {
			return errors.Errorf("value %d is below minimum %d", val, min)
		}
		if max > 0 && val > max {
			return errors.Errorf("value %d is above maximum %d", val, max)
		}
		return nil
	}

	// 处理其他数值类型
	val := cvt.Int64(value)

	if min > 0 && val < min {
		return errors.Errorf("value %d is below minimum %d", val, min)
	}

	if max > 0 && val > max {
		return errors.Errorf("value %d is above maximum %d", val, max)
	}

	return nil
}

// 字节序和字序处理辅助函数

// reorderBytes 根据字节序和字序重新排列字节
// 性能优化：预分配内存，减少内存分配次数
func reorderBytes(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) []byte {
	if len(data) < 2 {
		return data
	}

	// 预分配结果切片，避免多次内存分配
	result := make([]byte, len(data))

	// 处理字节序
	if byteEndian == ByteEndianLittle {
		// 小端字节序：反转每个16位字内的字节
		for i := 0; i < len(data); i += 2 {
			if i+1 < len(data) {
				result[i] = data[i+1]
				result[i+1] = data[i]
			} else {
				result[i] = data[i]
			}
		}
	} else {
		// 大端字节序：直接复制
		copy(result, data)
	}

	// 处理字序（仅对32位及以上数据有效）
	if len(result) >= 4 && wordOrder == WordOrderLowHigh {
		// 低字在前：交换16位字的顺序
		for i := 0; i < len(result)-3; i += 4 {
			result[i], result[i+2] = result[i+2], result[i]
			result[i+1], result[i+3] = result[i+3], result[i+1]
		}
	}

	return result
}

// 数值类型解析函数

func decodeUInt16(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for uint16")
	}

	reordered := reorderBytes(data[:2], byteEndian, wordOrder)
	return binary.BigEndian.Uint16(reordered), nil
}

func decodeInt16(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for int16")
	}

	reordered := reorderBytes(data[:2], byteEndian, wordOrder)
	val := binary.BigEndian.Uint16(reordered)
	return int16(val), nil
}

func decodeUInt32(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for uint32")
	}

	reordered := reorderBytes(data[:4], byteEndian, wordOrder)
	return binary.BigEndian.Uint32(reordered), nil
}

func decodeInt32(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for int32")
	}

	reordered := reorderBytes(data[:4], byteEndian, wordOrder)
	val := binary.BigEndian.Uint32(reordered)
	return int32(val), nil
}

func decodeFloat32(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for float32")
	}

	reordered := reorderBytes(data[:4], byteEndian, wordOrder)
	bits := binary.BigEndian.Uint32(reordered)
	return math.Float32frombits(bits), nil
}

func decodeFloat64(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 8 {
		return nil, errors.New("insufficient data for float64")
	}

	reordered := reorderBytes(data[:8], byteEndian, wordOrder)
	bits := binary.BigEndian.Uint64(reordered)
	return math.Float64frombits(bits), nil
}

// BCD码解析函数

func decodeBCD16(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 2 {
		return nil, errors.New("insufficient data for BCD16")
	}

	reordered := reorderBytes(data[:2], byteEndian, wordOrder)

	// BCD解码：每个字节包含两个十进制数字
	high := int(reordered[0]>>4)*1000 + int(reordered[0]&0x0F)*100
	low := int(reordered[1]>>4)*10 + int(reordered[1]&0x0F)

	return high + low, nil
}

func decodeBCD32(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 4 {
		return nil, errors.New("insufficient data for BCD32")
	}

	reordered := reorderBytes(data[:4], byteEndian, wordOrder)

	// BCD解码：4个字节包含8个十进制数字
	result := 0
	multiplier := 10000000

	for i := 0; i < 4; i++ {
		high := int(reordered[i]>>4) * multiplier
		low := int(reordered[i]&0x0F) * (multiplier / 10)
		result += high + low
		multiplier /= 100
	}

	return result, nil
}

// 字符串解析函数

func decodeASCIIString(data []byte) (any, error) {
	// 移除尾部的null字符
	str := strings.TrimRight(string(data), "\x00")
	return str, nil
}

func decodeUTF16String(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) < 2 {
		return "", errors.New("insufficient data for UTF16 string")
	}

	// 确保数据长度是偶数
	if len(data)%2 != 0 {
		data = data[:len(data)-1]
	}

	reordered := reorderBytes(data, byteEndian, wordOrder)

	// 转换为UTF-16代码点
	codePoints := make([]uint16, len(reordered)/2)
	for i := 0; i < len(reordered); i += 2 {
		codePoints[i/2] = binary.BigEndian.Uint16(reordered[i : i+2])
	}

	// 转换为UTF-8字符串
	runes := utf16.Decode(codePoints)
	return string(runes), nil
}

// 位图解析函数

func decodeBits(data []byte, byteEndian EByteEndian, wordOrder EWordOrder) (any, error) {
	if len(data) == 0 {
		return nil, errors.New("insufficient data for bits")
	}

	// 对于位图，通常返回原始字节数组或转换为整数
	if len(data) == 1 {
		return uint8(data[0]), nil
	} else if len(data) == 2 {
		reordered := reorderBytes(data, byteEndian, wordOrder)
		return binary.BigEndian.Uint16(reordered), nil
	} else if len(data) == 4 {
		reordered := reorderBytes(data, byteEndian, wordOrder)
		return binary.BigEndian.Uint32(reordered), nil
	} else {
		// 对于更长的位图，返回字节数组
		return data, nil
	}
}

// 位范围解析函数
// 从字节数组中提取特定范围的位，支持跨字节操作
func decodeBitRange(data []byte, bitStart, bitLength uint8) (any, error) {
	if len(data) == 0 {
		return nil, errors.New("insufficient data for bit range")
	}

	// 参数验证
	if bitLength == 0 || bitLength > 64 {
		return nil, errors.New("bit length must be 1-64")
	}

	// 如果只需要1个字节且位长度不超过8，使用简单方法
	if bitLength <= 8 && bitStart < 8 {
		if bitStart+bitLength <= 8 {
			// 单字节内操作
			byteValue := data[0]
			mask := uint8((1 << bitLength) - 1)
			extractedBits := (byteValue >> bitStart) & mask
			return extractedBits, nil
		}
	}

	// 跨字节操作：构建结果值
	var result uint64 = 0

	for i := uint8(0); i < bitLength; i++ {
		currentBitIndex := bitStart + i
		byteIndex := currentBitIndex / 8
		bitOffset := currentBitIndex % 8

		// 检查是否超出数据范围
		if byteIndex >= uint8(len(data)) {
			break
		}

		// 获取当前位
		bitValue := (data[byteIndex] >> bitOffset) & 1

		// 将位添加到结果中（从低位开始）
		result |= uint64(bitValue) << i
	}

	// 根据位长度返回适当类型
	if bitLength <= 8 {
		return uint8(result), nil
	} else if bitLength <= 16 {
		return uint16(result), nil
	} else if bitLength <= 32 {
		return uint32(result), nil
	} else {
		return result, nil
	}
}

// 系数和偏移量处理
// 使用cvt库简化类型转换

func applyFactorAndOffset(value any, factor float32, offset int) any {
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

// 转换为目标格式

func convertToReturnFormat(value any, returnFormat ESystemType) any {
	if value == nil {
		return nil
	}

	switch returnFormat {
	case SUseReadType:
		// 使用原始类型
		return value
	case SBool:
		return cvt.Bool(value)
	case SInt8:
		return cvt.Int8(value)
	case SUint8:
		return cvt.Uint8(value)
	case SInt16:
		return cvt.Int16(value)
	case SUint16:
		return cvt.Uint16(value)
	case SInt32:
		return cvt.Int32(value)
	case SUint32:
		return cvt.Uint32(value)
	case SInt64:
		return cvt.Int64(value)
	case SUint64:
		return cvt.Uint64(value)
	case SFloat32:
		return cvt.Float32(value)
	case SFloat64:
		return cvt.Float64(value)
	case SString:
		return cvt.String(value)
	default:
		return value
	}
}
