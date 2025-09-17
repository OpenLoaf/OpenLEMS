package entity

import (
	"reflect"
	"s_db/s_db_model"

	"github.com/pkg/errors"
)

// SSettingEntity 设置实体结构体
type SSettingEntity struct {
	Id        string `json:"id" dc:"设置ID"`
	Value     string `json:"value" dc:"设置值"`
	IsPublic  bool   `json:"isPublic" dc:"是否公开"`
	Enabled   bool   `json:"enabled" dc:"是否启用"`
	Remark    string `json:"remark" dc:"备注"`
	Sort      int    `json:"sort" dc:"排序"`
	Group     string `json:"group" dc:"分组"`
	CreatedAt string `json:"createdAt" dc:"创建时间"`
	UpdatedAt string `json:"updatedAt" dc:"更新时间"`
}

// UnmarshalValue 实现值转换接口
func (e *SSettingEntity) UnmarshalValue(value interface{}) error {
	if record, ok := value.(*s_db_model.SSettingModel); ok {
		*e = SSettingEntity{
			Id:        record.Id,
			Value:     record.Value,
			IsPublic:  record.IsPublic,
			Enabled:   record.Enabled,
			Remark:    record.Remark,
			Sort:      record.Sort,
			Group:     record.Group,
			CreatedAt: record.CreatedAt.String(),
			UpdatedAt: record.UpdatedAt.String(),
		}
		return nil
	}
	return errors.Errorf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}
