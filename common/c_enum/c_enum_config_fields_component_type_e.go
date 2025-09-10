package c_enum

type EConfigFieldsComponentType string

const (
	EConfigFieldsComponentTypeText         EConfigFieldsComponentType = "text"         // 文本
	EConfigFieldsComponentTypeNumber       EConfigFieldsComponentType = "number"       // 数字
	EConfigFieldsComponentTypeSwitch       EConfigFieldsComponentType = "switch"       // 开关
	EConfigFieldsComponentTypeSingleSelect EConfigFieldsComponentType = "singleSelect" // 单选
	EConfigFieldsComponentTypeMultiSelect  EConfigFieldsComponentType = "multiSelect"  // 多选
	EConfigFieldsComponentTypeDate         EConfigFieldsComponentType = "date"         // 日期
	EConfigFieldsComponentTypeTime         EConfigFieldsComponentType = "time"         // 时间
	EConfigFieldsComponentTypeDateTime     EConfigFieldsComponentType = "dateTime"     // 完整的日期+时间
)
