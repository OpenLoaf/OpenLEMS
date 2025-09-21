package s_db_basic

import "common/c_enum"

// SSystemSettingDefine 系统设置定义结构体
type SSystemSettingDefine struct {
	Id           string                   `json:"id"`           // 设置ID
	Group        string                   `json:"group"`        // 设置分组
	DefaultValue string                   `json:"defaultValue"` // 默认值
	IsPublic     bool                     `json:"isPublic"`     // 是否公开
	Remark       string                   `json:"remark"`       // 备注
	FieldType    c_enum.ESettingFieldType `json:"fieldType"`    // 存储字段类型
}
