package c_base

import (
	"common/c_enum"
	"time"

	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

type IPoint interface {
	GetKey() string
	GetName() string
	GetGroup() *MetaGroup
	GetLevel() c_enum.EAlarmLevel
	GetUnit() string
	GetDesc() string
	GetSort() int
	AlarmTrigger(value any) (bool, error)   // 判断触发或者消除告警
	ValueExplain(value any) (string, error) // 获取Value解释，一般为状态类型的解释
}

// SPoint 点位元数据
type SPoint struct {
	Key     string             `json:"key" v:"required"`  // 名称
	Name    string             `json:"name" v:"required"` // 名称
	Group   *MetaGroup         `json:"group" dc:"分组"`
	Unit    string             `json:"unit,omitempty"`    // 单位
	Desc    string             `json:"desc,omitempty"`    // 备注
	Sort    int                `json:"sort"`              // 排序
	Level   c_enum.EAlarmLevel `json:"level"`             // 点位级别
	Min     int64              `json:"min,omitempty"`     // 范围最小值
	Max     int64              `json:"max,omitempty"`     // 范围最大值
	Precise int                `json:"precise,omitempty"` // 设置浮点数精度（只是显示用）
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

func (s *SPoint) GetLevel() c_enum.EAlarmLevel {
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

type SModbusPoint struct {
	*SPoint
	Address      uint16
	Length       uint16
	IsBitAddress bool // 是否为位级别地址

	// -- 解析相关配置 --
	ByteEndian c_enum.EByteEndian `json:"byteEndian"` // 字节序 (默认: ByteEndianBig)
	WordOrder  EWordOrder         `json:"wordOrder"`  // 字序 (默认: WordOrderHighLow)
	DataFormat c_enum.EDataFormat `json:"dataFormat"` // 数据格式 (必须明确指定)

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

func (s *SModbusPointTask) GetName() string {
	return s.Name
}

func (s *SModbusPointTask) GetDescription() string {
	return s.Desc
}

func (s *SModbusPointTask) GetPoints() []IPoint {
	return s.Points
}

func (s *SModbusPointTask) GetLifeTime() time.Duration {
	return s.Lifetime
}

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
		// 根据IsBitAddress决定如何传递参数
		var isBit bool
		var index, length uint16

		if point.IsBitAddress {
			// 位级别：Address和Length都是位级别的
			isBit = true

			// 边界检查：确保点位地址在任务范围内
			// 每个寄存器占16位，所以任务的位地址范围是 [s.Addr*16, (s.Addr+s.Quantity)*16-1]
			taskStartBit := s.Addr * 16
			taskEndBit := (s.Addr+s.Quantity)*16 - 1
			// 检查起始地址和结束地址都在范围内
			bitEndAddress := point.Address + point.Length - 1
			if point.Address < taskStartBit || bitEndAddress > taskEndBit {
				return nil, errors.Errorf("bit address range [%d:%d] is out of task range [%d:%d]",
					point.Address, bitEndAddress, taskStartBit, taskEndBit)
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

func (s *SCanbusTask) GetName() string {
	return s.Name
}

func (s *SCanbusTask) GetDescription() string {
	return s.Desc
}

func (s *SCanbusTask) GetPoints() []IPoint {
	return s.Points
}

func (s *SCanbusTask) GetLifeTime() time.Duration {
	return s.Lifetime
}
