package c_base

type EPointAddressType int

const (
	EPointAddressTypeByte = iota
	EPointAddressTypeBit
)

// Meta 点位元数据
type Point[T Value] struct {
	Debug bool `json:"-"` // 调试打印

	Key         string     `json:"key"`  // 名称
	Name        string     `json:"name"` // 名称
	Group       *MetaGroup `json:"group" dc:"分组"`
	Address     uint16
	Length      uint16
	AddressType EPointAddressType // byte或者bit

	Endianness ECharSequence `json:"endianness,omitempty"` // 字节顺序，默认为大端

	Level EAlarmLevel `json:"level"` // 点位级别

	Factor  float32 `json:"factor,omitempty"`  // 乘以系数，如果是0，自动会改成1。因为0无意义
	Offset  int     `json:"offset,omitempty"`  // 偏移值
	Min     int64   `json:"min,omitempty"`     // 范围最小值
	Max     int64   `json:"max,omitempty"`     // 范围最大值
	Precise int     `json:"precise,omitempty"` // 设置浮点数精度（只是显示用）

	Unit          string               `json:"unit,omitempty"` // 单位
	Desc          string               `json:"desc,omitempty"` // 备注
	Sort          int                  `json:"sort"`           // 排序
	StatusExplain func(value T) string `json:"-"`              // 状态解释，如果有的话就翻译一下状态
	Trigger       func(T) bool         `json:"-"`              // 触发告警警告故障信息
}

//
//func (m *Meta) GetValueExplain(value any) any {
//	if value == nil {
//		return nil
//	}
//	if m.StatusExplain == nil {
//		return value
//	}
//	explainValue := m.StatusExplain(value)
//	if explainValue == "" {
//		return value
//	}
//	return explainValue
//}
//
//type MetaGroup struct {
//	GroupName string `json:"groupName" dc:"组名称"`
//	GroupSort int    `json:"groupSort" dc:"组排序"`
//	Display   bool   `json:"display" dc:"是否显示"`
//}
//
//type MetaValue struct {
//	Value      any        `json:"value,omitempty" dc:"数值"`
//	HappenTime *time.Time `json:"happenTime,omitempty" dc:"发生时间"`
//}
//
//type MetaValueWrapper struct {
//	DeviceId   string      `json:"deviceId,omitempty" dc:"设备ID"`
//	DeviceType EDeviceType `json:"deviceType,omitempty" dc:"设备类型"`
//	Level      EAlarmLevel `json:"level" dc:"告警级别"`
//	Meta       *Meta       `json:"meta,omitempty" dc:"点位信息"`
//	Value      any         `json:"value,omitempty" dc:"数值"`
//	HappenTime *time.Time  `json:"happenTime,omitempty" dc:"发生时间"`
//}
