// Code generated manually for EConfigFieldsComponentType string enum; DO NOT EDIT.

package c_enum

import "fmt"

// String returns the string representation of the EConfigFieldsComponentType
func (i EConfigFieldsComponentType) String() string {
	switch i {
	case EConfigFieldsComponentTypeText:
		return "text"
	case EConfigFieldsComponentTypeNumber:
		return "number"
	case EConfigFieldsComponentTypeSwitch:
		return "switch"
	case EConfigFieldsComponentTypeSingleSelect:
		return "singleSelect"
	case EConfigFieldsComponentTypeMultiSelect:
		return "multiSelect"
	case EConfigFieldsComponentTypeDate:
		return "date"
	case EConfigFieldsComponentTypeTime:
		return "time"
	case EConfigFieldsComponentTypeDateTime:
		return "dateTime"
	default:
		return fmt.Sprintf("EConfigFieldsComponentType(%s)", string(i))
	}
}
