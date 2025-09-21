package c_enum

type ESettingFieldType string

const (
	// ESettingFieldTypeJson JSON对象类型
	ESettingFieldTypeJson ESettingFieldType = "json"

	// ESettingFieldTypeJsonArray JSON数组类型
	ESettingFieldTypeJsonArray ESettingFieldType = "json_array"

	// ESettingFieldTypeText 文本类型
	ESettingFieldTypeText ESettingFieldType = "text"

	// ESettingFieldTypeNumber 数字类型
	ESettingFieldTypeNumber ESettingFieldType = "number"

	// ESettingFieldTypeBoolean 布尔类型
	ESettingFieldTypeBoolean ESettingFieldType = "boolean"
)
