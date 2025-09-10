package c_base

import (
	"common/c_enum"

	"github.com/shockerli/cvt"
)

// SPoint 点位元数据
type SPoint struct {
	Key     string             `json:"key" v:"required"`  // 名称
	Name    string             `json:"name" v:"required"` // 名称
	Group   *SPointGroup       `json:"group" dc:"分组"`
	Unit    string             `json:"unit,omitempty"`    // 单位
	Desc    string             `json:"desc,omitempty"`    // 备注
	Sort    int                `json:"sort"`              // 排序
	Level   c_enum.EAlarmLevel `json:"level"`             // 点位级别
	Min     int64              `json:"min,omitempty"`     // 范围最小值
	Max     int64              `json:"max,omitempty"`     // 范围最大值
	Precise int                `json:"precise,omitempty"` // 设置浮点数精度（只是显示用）
}

func (s *SPoint) AlarmTrigger(value any) (bool, error) {
	return false, nil
}

func (s *SPoint) ValueExplain(value any) (string, error) {
	return cvt.StringE(value)
}

func (s *SPoint) GetKey() string {
	return s.Key
}

func (s *SPoint) GetName() string {
	return s.Name
}

func (s *SPoint) GetGroup() *SPointGroup {
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
