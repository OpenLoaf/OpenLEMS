//go:generate stringer -type=EConfigStructFieldsComponentType -trimprefix=EConfigStructFieldsComponentType -output=c_base_config_struct_fields_component_type_e_string.go
package c_base

type EConfigStructFieldsComponentType int

const (
	EConfigStructFieldsComponentTypeText         EConfigStructFieldsComponentType = iota // 文本
	EConfigStructFieldsComponentTypeNumber                                               // 数字
	EConfigStructFieldsComponentTypeSwitch                                               // 开关
	EConfigStructFieldsComponentTypeSingleSelect                                         // 单选
	EConfigStructFieldsComponentTypeMultiSelect                                          // 多选
	EConfigStructFieldsComponentTypeDate                                                 // 日期
	EConfigStructFieldsComponentTypeTime                                                 // 时间
	EConfigStructFieldsComponentTypeDateTime                                             // 完整的日期+时间
)
