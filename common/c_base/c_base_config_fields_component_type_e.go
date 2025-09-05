package c_base

type EConfigFieldsComponentType string

const (
	EConfigFieldsComponentTypeText         = "text"         // 文本
	EConfigFieldsComponentTypeNumber       = "number"       // 数字
	EConfigFieldsComponentTypeSwitch       = "switch"       // 开关
	EConfigFieldsComponentTypeSingleSelect = "singleSelect" // 单选
	EConfigFieldsComponentTypeMultiSelect  = "multiSelect"  // 多选
	EConfigFieldsComponentTypeDate         = "date"         // 日期
	EConfigFieldsComponentTypeTime         = "time"         // 时间
	EConfigFieldsComponentTypeDateTime     = "dateTime"     // 完整的日期+时间
)
