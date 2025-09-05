package c_base

type EConfigFieldsComponentType string

const (
	EConfigStructFieldsComponentTypeText         = "text"         // 文本
	EConfigStructFieldsComponentTypeNumber       = "number"       // 数字
	EConfigStructFieldsComponentTypeSwitch       = "switch"       // 开关
	EConfigStructFieldsComponentTypeSingleSelect = "singleSelect" // 单选
	EConfigStructFieldsComponentTypeMultiSelect  = "multiSelect"  // 多选
	EConfigStructFieldsComponentTypeDate         = "date"         // 日期
	EConfigStructFieldsComponentTypeTime         = "time"         // 时间
	EConfigStructFieldsComponentTypeDateTime     = "dateTime"     // 完整的日期+时间
)
