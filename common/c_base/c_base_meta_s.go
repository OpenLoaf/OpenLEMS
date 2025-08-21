package c_base

// Meta 点位元数据
type Meta struct {
	Debug bool `json:"-"` // 调试打印
	// TODO 这里再加一个系统点位类型
	Name  string     `json:"name"` // 名称
	Group *MetaGroup `json:"group" dc:"分组"`

	Cn            string                 `json:"cn"`                   // 中文名称, TODO 以后改成I18N
	Addr          uint16                 `json:"addr"`                 // 地址，索引
	BitLength     uint8                  `json:"bitLength,omitempty"`  // 位长度, 可以和ReadType一起使用，表示位读取。 比如 RBit5，BigLength=3 代表读取第5位到第7位。0时忽略该参数！
	Endianness    ECharSequence          `json:"endianness,omitempty"` // 字节顺序，默认为大端
	ReadType      EReadType              `json:"readType"`             // 数据类型，ECharSequence 字节顺序将会影响读取到的数据结果
	SystemType    ESystemType            `json:"systemType"`           // 格式化类型,默认为SUseReadType!自动使用ReadType的类型。
	Level         EAlarmLevel            `json:"level"`                // 点位级别
	Factor        float32                `json:"factor,omitempty"`     // 乘以系数，如果是0，自动会改成1。因为0无意义
	Offset        int                    `json:"offset,omitempty"`     // 偏移值
	Min           int64                  `json:"min,omitempty"`        // 范围最小值
	Max           int64                  `json:"max,omitempty"`        // 范围最大值
	Precise       int                    `json:"precise,omitempty"`    // 设置浮点数精度（只是显示用）
	Unit          string                 `json:"unit,omitempty"`       // 单位
	Desc          string                 `json:"desc,omitempty"`       // 备注
	Sort          int                    `json:"sort"`                 // 排序
	StatusExplain func(value any) string `json:"-"`                    // 状态解释，如果有的话就翻译一下状态
	Trigger       func(any) bool         `json:"-"`                    // 触发告警警告故障信息
}

type MetaGroup struct {
	GroupName string `json:"groupName" dc:"组名称"`
	GroupSort int    `json:"groupSort" dc:"组排序"`
	Display   bool   `json:"display" dc:"是否显示"`
}
