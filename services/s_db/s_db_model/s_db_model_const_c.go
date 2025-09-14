package s_db_model

import "github.com/gogf/gf/v2/os/gtime"

const (
	FieldId = "id"

	FieldSort = "sort"

	FieldName = "name"
	FieldType = "type"

	FieldEnable = "enable"

	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"

	FieldValue  = "value"
	FieldRemark = "remark"
	FieldGroup  = "group_name"

	// Automation 相关字段
	FieldStartTime      = "start_time"
	FieldEndTime        = "end_time"
	FieldTimeRangeType  = "time_range_type"
	FieldTimeRangeValue = "time_range_value"
	FieldTriggerRule    = "trigger_rule"
	FieldExecuteRule    = "execute_rule"
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
