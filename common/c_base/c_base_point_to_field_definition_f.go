package c_base

import "common/c_enum"

// ConvertIPointToFieldDefinition 将IPoint转换为SFieldDefinition
// 参考 entity/device.go 中的转换逻辑，抽取为通用函数以供跨层复用
func ConvertIPointToFieldDefinition(point IPoint) *SFieldDefinition {
	if point == nil {
		return nil
	}

	// 转换值类型
	valueType := convertEValueTypeToConfigFieldsValueType(point.GetValueType())

	// 根据值类型推断组件类型
	componentType := inferComponentTypeFromValueType(valueType)

	// 处理指针类型字段
	var unit *string
	if unitStr := point.GetUnit(); unitStr != "" {
		unit = &unitStr
	}

	var min, max *int64
	if minVal := point.GetMin(); minVal != 0 {
		min = &minVal
	}
	if maxVal := point.GetMax(); maxVal != 0 {
		max = &maxVal
	}

	// 获取分组信息
	var group *SPointGroup
	if groupInfo := point.GetGroup(); groupInfo != nil {
		// 复制一份，避免外部修改影响原对象
		group = &SPointGroup{
			GroupKey:  groupInfo.GroupKey,
			GroupName: groupInfo.GroupName,
			GroupSort: groupInfo.GroupSort,
			Disable:   groupInfo.Disable,
		}
	}

	// 创建SFieldDefinition
	fieldDef := &SFieldDefinition{
		Key:           point.GetKey(),
		Name:          point.GetName(),
		Group:         group,
		ValueType:     valueType,
		ComponentType: componentType,
		Unit:          unit,
		Min:           min,
		Max:           max,
		Description:   point.GetDesc(),
		Required:      false, // 遥测/点位定义通常不是必填
		ValueExplain:  point.GetValueExplain(),
	}

	return fieldDef
}

// convertEValueTypeToConfigFieldsValueType 将EValueType转换为EConfigFieldsValueType
func convertEValueTypeToConfigFieldsValueType(valueType c_enum.EValueType) c_enum.EConfigFieldsValueType {
	switch valueType {
	case c_enum.EBool:
		return c_enum.EConfigFieldsValueTypeBoolean
	case c_enum.EInt8, c_enum.EUint8, c_enum.EInt16, c_enum.EUint16, c_enum.EInt32, c_enum.EUint32, c_enum.EInt64, c_enum.EUint64:
		return c_enum.EConfigFieldsValueTypeInt
	case c_enum.EFloat32, c_enum.EFloat64:
		return c_enum.EConfigFieldsValueTypeFloat
	case c_enum.EString:
		return c_enum.EConfigFieldsValueTypeString
	default:
		return c_enum.EConfigFieldsValueTypeString
	}
}

// inferComponentTypeFromValueType 根据值类型推断组件类型
// 注意：label组件类型需要通过ct标签显式指定，不会自动推断
func inferComponentTypeFromValueType(valueType c_enum.EConfigFieldsValueType) c_enum.EConfigFieldsComponentType {
	switch valueType {
	case c_enum.EConfigFieldsValueTypeBoolean:
		return c_enum.EConfigFieldsComponentTypeSwitch
	case c_enum.EConfigFieldsValueTypeInt, c_enum.EConfigFieldsValueTypeFloat:
		return c_enum.EConfigFieldsComponentTypeNumber
	default:
		return c_enum.EConfigFieldsComponentTypeText
	}
}
