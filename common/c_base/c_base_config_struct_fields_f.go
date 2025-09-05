package c_base

import (
	"github.com/shockerli/cvt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func BuildConfigStructFields(config any) ([]*SConfigFields, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	v := reflect.ValueOf(config)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("not struct")
	}

	t := v.Type()
	var fields []*SConfigFields

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		// 获取字段标签
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// 解析json标签，获取字段名
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "" {
			jsonName = field.Name
		}

		// 获取描述信息
		desc := field.Tag.Get("dc")
		if desc == "" {
			desc = field.Tag.Get("desc")
		}

		// 根据字段类型确定组件类型和值类型
		componentType, valueType := getFieldTypeInfo(field.Type)

		fieldConfig := &SConfigFields{
			Name:          field.Tag.Get("name"),
			Code:          jsonName,
			ValueType:     valueType,
			ComponentType: componentType,
			Description:   desc,
		}

		if name := field.Tag.Get("name"); name != "" {
			fieldConfig.Name = name
		} else if dc := field.Tag.Get("desc"); dc != "" {
			fieldConfig.Name = dc
		} else {
			fieldConfig.Name = jsonName
		}

		if ct := field.Tag.Get("ct"); ct != "" {
			fieldConfig.ComponentType = EConfigFieldsComponentType(ct)
		}
		if vt := field.Tag.Get("vt"); vt != "" {
			fieldConfig.ValueType = vt
		}

		// 解析其他标签
		if minStr := field.Tag.Get("min"); minStr != "" {
			fieldConfig.Min = cvt.Uint8(minStr)
		}
		if maxStr := field.Tag.Get("max"); maxStr != "" {
			fieldConfig.Max = cvt.Uint8(maxStr)
		}
		if defaultVal := field.Tag.Get("default"); defaultVal != "" {
			fieldConfig.Default = defaultVal
		}
		if regex := field.Tag.Get("regex"); regex != "" {
			fieldConfig.Regex = regex
		}

		fields = append(fields, fieldConfig)
	}

	return fields, nil
}

// getFieldTypeInfo 根据反射类型返回组件类型和值类型
func getFieldTypeInfo(fieldType reflect.Type) (EConfigFieldsComponentType, string) {
	switch fieldType.Kind() {
	case reflect.String:
		return EConfigStructFieldsComponentTypeText, "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return EConfigStructFieldsComponentTypeNumber, "int"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return EConfigStructFieldsComponentTypeNumber, "int"
	case reflect.Float32, reflect.Float64:
		return EConfigStructFieldsComponentTypeNumber, "float"
	case reflect.Bool:
		return EConfigStructFieldsComponentTypeSwitch, "bool"
	default:
		return EConfigStructFieldsComponentTypeText, "string"
	}
}
