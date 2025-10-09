package entity

import (
	"common/c_base"
	"fmt"
	"reflect"
	"time"
)

type SSingleDeviceGroup struct {
	GroupName string `json:"groupName" dc:"组名称"`
	GroupSort int    `json:"groupSort" dc:"组排序"`
	Values    []*SSingleDeviceValue
}

type SSingleDeviceValue struct {
	Meta          *SSingleDeviceMeta `json:"meta,omitempty"`
	Value         any                `json:"value,omitempty"`
	StatueExplain string             `json:"statueExplain,omitempty"`
	HappenTime    *time.Time         `json:"happenTime,omitempty"`
}

func (s *SSingleDeviceValue) UnmarshalValue(value interface{}) error {
	if record, ok := value.(*c_base.SPointValue); ok {
		// 转换点位值
		s.Value = record.GetValue()

		// 转换发生时间
		happenTime := record.GetHappenTime()
		s.HappenTime = &happenTime

		// 获取状态解释
		if explain, err := record.GetActualValueExplain(); err == nil {
			s.StatueExplain = explain
		}

		// 转换点位元数据
		if point := record.IPoint; point != nil {
			// 尝试获取具体的SPoint实例以访问Min、Before、Precise字段
			s.Meta = &SSingleDeviceMeta{
				Key:        point.GetKey(),
				Name:       point.GetName(),
				Cn:         point.GetName(),               // 暂时使用Name作为中文名称
				SystemType: point.GetValueType().String(), // 默认自动系统类型
				Min:        point.GetMin(),
				Max:        point.GetMax(),
				Precise:    point.GetPrecise(),
				Unit:       point.GetUnit(),
				Desc:       point.GetDesc(),
			}
		}

		return nil
	}

	return fmt.Errorf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}

// Meta 点位元数据
type SSingleDeviceMeta struct {
	Key        string `json:"key"`            // key
	Name       string `json:"name"`           // 名称
	Cn         string `json:"cn"`             // 中文名称, TODO 以后改成I18N
	SystemType string `json:"systemType"`     // 格式化类型,默认为SUseReadType!自动使用ReadType的类型。
	Min        int64  `json:"min,omitempty"`  // 范围最小值
	Max        int64  `json:"max,omitempty"`  // 范围最大值
	Precise    uint8  `json:"precise"`        // 设置浮点数精度（只是显示用）
	Unit       string `json:"unit,omitempty"` // 单位
	Desc       string `json:"desc,omitempty"` // 备注
}
