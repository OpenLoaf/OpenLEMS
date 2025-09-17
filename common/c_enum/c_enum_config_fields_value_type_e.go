package c_enum

type EConfigFieldsValueType string

const (
	EConfigFieldsValueTypeString  EConfigFieldsValueType = "string" // 文本
	EConfigFieldsValueTypeInt     EConfigFieldsValueType = "int"    // 整数
	EConfigFieldsValueTypeFloat   EConfigFieldsValueType = "float"  // 浮点数
	EConfigFieldsValueTypeBoolean EConfigFieldsValueType = "bool"   //  布尔值
)
