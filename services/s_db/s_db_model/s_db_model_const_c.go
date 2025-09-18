package s_db_model

import "github.com/gogf/gf/v2/os/gtime"

const (
	// 通用字段 - 在2个以上表中都使用的字段
	FieldId        = "id"
	FieldName      = "name"
	FieldType      = "type"
	FieldEnabled   = "enabled"
	FieldSort      = "sort"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
	FieldValue     = "value"
	FieldRemark    = "remark"
	FieldGroup     = "group_name"
)

const ( // 特殊值
	NullValue  = "null"
	EmptyValue = ""
)

type SDatabaseBasic struct {
	Id        string      `json:"id,omitempty" orm:"id"` // 设备ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at"`
}
