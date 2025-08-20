package entity

import (
	"time"
)

type SSingleDeviceValue struct {
	Meta       *SSingleDeviceMeta `json:"meta,omitempty"`
	Value      string             `json:"value,omitempty"`
	HappenTime *time.Time         `json:"happenTime,omitempty"`
}

// Meta 点位元数据
type SSingleDeviceMeta struct {
	Name       string  `json:"name"`                 // 名称
	Cn         string  `json:"cn"`                   // 中文名称, TODO 以后改成I18N
	Addr       uint16  `json:"addr"`                 // 地址，索引
	BitLength  uint8   `json:"bitLength,omitempty"`  // 位长度, 可以和ReadType一起使用，表示位读取。 比如 RBit5，BigLength=3 代表读取第5位到第7位。0时忽略该参数！
	Endianness string  `json:"endianness,omitempty"` // 字节顺序，默认为大端
	ReadType   string  `json:"readType"`             // 数据类型，ECharSequence 字节顺序将会影响读取到的数据结果
	SystemType string  `json:"systemType"`           // 格式化类型,默认为SUseReadType!自动使用ReadType的类型。
	Level      string  `json:"level"`                // 点位级别
	Factor     float32 `json:"factor,omitempty"`     // 乘以系数，如果是0，自动会改成1。因为0无意义
	Offset     int     `json:"offset,omitempty"`     // 偏移值
	Min        int64   `json:"min,omitempty"`        // 范围最小值
	Max        int64   `json:"max,omitempty"`        // 范围最大值
	Precise    int     `json:"precise,omitempty"`    // 设置浮点数精度（只是显示用）
	Unit       string  `json:"unit,omitempty"`       // 单位
	Desc       string  `json:"desc,omitempty"`       // 备注
}
