package c_base

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"math/big"
	"reflect"
	"strconv"
)

// Meta 点位元数据
type Meta struct {
	Debug bool `json:"-"` // 调试打印
	// TODO 这里再加一个系统点位类型
	Name       string         `json:"name"`                 // 名称
	Cn         string         `json:"cn"`                   // 中文名称, TODO 以后改成I18N
	Addr       uint16         `json:"addr"`                 // 地址，索引
	BitLength  uint8          `json:"bitLength,omitempty"`  // 位长度, 可以和ReadType一起使用，表示位读取。 比如 RBit5，BigLength=3 代表读取第5位到第7位。0时忽略该参数！
	Endianness ECharSequence  `json:"endianness,omitempty"` // 字节顺序，默认为大端
	ReadType   EReadType      `json:"readType"`             // 数据类型，ECharSequence 字节顺序将会影响读取到的数据结果
	SystemType ESystemType    `json:"systemType"`           // 格式化类型,默认为SUseReadType!自动使用ReadType的类型。
	Level      EAlarmLevel    `json:"level"`                // 点位级别
	Factor     float32        `json:"factor,omitempty"`     // 乘以系数，如果是0，自动会改成1。因为0无意义
	Offset     int            `json:"offset,omitempty"`     // 偏移值
	Min        int64          `json:"min,omitempty"`        // 范围最小值
	Max        int64          `json:"max,omitempty"`        // 范围最大值
	Precise    int            `json:"precise,omitempty"`    // 设置浮点数精度（只是显示用）
	Unit       string         `json:"unit,omitempty"`       // 单位
	Desc       string         `json:"desc,omitempty"`       // 备注
	Trigger    func(any) bool `json:"-"`                    // 触发告警警告故障信息
}

func (p *Meta) AddrDecString() string {
	return strconv.FormatUint(uint64(p.Addr), 10)
}

//func (p *Meta) GetValueReflectKind() reflect.Kind {
//	return p.SystemType.GetReflectKind(p.ReadType, p.BitLength)
//}

func (p *Meta) ValueToString(value *gvar.Var) string {
	switch p.SystemType.GetReflectKind(p.ReadType, p.BitLength) {
	case reflect.Bool:
		b := value.Bool()
		if b {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.String:
		return fmt.Sprintf("%v", value.String())
	case reflect.Float32, reflect.Float64:
		return big.NewFloat(value.Float64()).Text('f', p.Precise)
	default:
		return fmt.Sprintf("%v", value.String())
	}
}
