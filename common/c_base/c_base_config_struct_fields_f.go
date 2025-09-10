package c_base

import (
	"common/c_enum"
	"reflect"
	"strings"

	"github.com/shockerli/cvt"

	"github.com/pkg/errors"
)

// BuildConfigStructFields 通过反射解析结构体字段，构建配置字段列表
// 支持嵌套结构体的平铺展开，包括匿名嵌入字段
// 支持的标签字段：
//   - json: JSON序列化字段名，必需标签，格式如 "fieldName" 或 "fieldName,omitempty"
//   - name: 字段显示名称，优先级高于desc和json字段名
//   - desc/dc: 字段描述信息，用于生成字段名称和描述
//   - ct: 组件类型，可选值：text(文本)、number(数字)、switch(开关)、singleSelect(单选)、multiSelect(多选)、date(日期)、time(时间)、dateTime(日期时间)
//   - vt: 值类型，可选值：string(字符串)、int(整数)、float(浮点数)、bool(布尔值)
//   - min: 最小值限制，仅对数值类型有效
//   - max: 最大值限制，仅对数值类型有效
//   - default: 默认值
//   - required: 是否必填
//   - unit: 单位信息
//   - regex: 正则表达式验证规则
//   - regexFailedMessage: 正则表达式验证失败时的提示信息
//   - selectOptions: 选择项配置，格式为 "key1:value1,key2:value2"，用于单选和多选组件
//
// 示例：
//
//	type BaseConfig struct {
//	    UnitId uint8 `json:"unitId" name:"ModbusID" min:"1" max:"255" default:"1"`
//	}
//
//	type Config struct {
//	    BaseConfig                    // 匿名嵌入，字段会被平铺展开
//	    Name string `json:"name" name:"设备名称" desc:"设备的显示名称" ct:"text" vt:"string" regex:"^[a-zA-Z0-9_-]+$" regexFailedMessage:"只能包含字母、数字、下划线和连字符"`
//	    Port int    `json:"port" name:"端口号" desc:"设备通信端口" ct:"number" vt:"int" min:"1" max:"65535" default:"8080"`
//	    Enable bool `json:"enable" name:"启用状态" desc:"是否启用设备" ct:"switch" vt:"bool" default:"true"`
//	}
func BuildConfigStructFields(config any) ([]*SConfigStructFields, error) {
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

	return buildConfigStructFieldsRecursive(v.Type(), "")
}

// buildConfigStructFieldsRecursive 递归处理结构体字段，支持嵌套结构体的平铺展开
func buildConfigStructFieldsRecursive(structType reflect.Type, prefix string) ([]*SConfigStructFields, error) {
	var fields []*SConfigStructFields

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if !field.IsExported() {
			continue
		}

		// 处理匿名嵌入字段（嵌套结构体）
		if field.Anonymous {
			// 递归处理嵌入的结构体
			embeddedFields, err := buildConfigStructFieldsRecursive(field.Type, prefix)
			if err != nil {
				return nil, err
			}
			fields = append(fields, embeddedFields...)
			continue
		}

		// 获取json标签，用于确定字段的JSON序列化名称
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// 解析json标签，获取字段名（忽略omitempty等选项）
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "" {
			jsonName = field.Name
		}

		// 如果有前缀，则组合字段名
		if prefix != "" {
			jsonName = prefix + "." + jsonName
		}

		// 获取描述信息，优先使用dc标签，其次使用desc标签
		desc := field.Tag.Get("dc")
		if desc == "" {
			desc = field.Tag.Get("desc")
		}

		// 根据字段类型确定组件类型和值类型
		componentType, valueType := getFieldTypeInfo(field.Type)

		fieldConfig := &SConfigStructFields{
			Name:          field.Tag.Get("name"),
			Code:          jsonName,
			ValueType:     valueType,
			ComponentType: componentType,
			Description:   desc,
		}

		// 设置字段名称，优先级：name标签 > desc标签 > json字段名
		if name := field.Tag.Get("name"); name != "" {
			fieldConfig.Name = name
		} else if dc := field.Tag.Get("desc"); dc != "" {
			fieldConfig.Name = dc
		} else {
			fieldConfig.Name = jsonName
		}

		// 解析值类型标签(vt)
		if g := field.Tag.Get("group"); g != "" {
			fieldConfig.Group = g
		}

		// 解析组件类型标签(ct)
		if ct := field.Tag.Get("ct"); ct != "" {
			fieldConfig.ComponentType = c_enum.EConfigFieldsComponentType(ct)
		}
		// 解析值类型标签(vt)
		if vt := field.Tag.Get("vt"); vt != "" {
			fieldConfig.ValueType = vt
		}
		// 解析组件类型标签(ct)
		if req := field.Tag.Get("required"); req != "" {
			fieldConfig.Required = cvt.Bool(req)
		}

		if unit := field.Tag.Get("unit"); unit != "" {
			fieldConfig.Unit = &unit
		}

		// 解析数值范围限制标签
		if minStr := field.Tag.Get("min"); minStr != "" {
			minVal := cvt.Int64(minStr)
			fieldConfig.Min = &minVal
		}
		if maxStr := field.Tag.Get("max"); maxStr != "" {
			maxVal := cvt.Int64(maxStr)
			fieldConfig.Max = &maxVal
		}
		// 解析默认值标签
		if defaultVal := field.Tag.Get("default"); defaultVal != "" {
			fieldConfig.Default = &defaultVal
		}
		// 解析正则表达式验证标签
		if regex := field.Tag.Get("regex"); regex != "" {
			fieldConfig.Regex = &regex
		}
		// 解析正则表达式失败提示信息标签
		if regexFailedMsg := field.Tag.Get("regexFailedMessage"); regexFailedMsg != "" {
			fieldConfig.RegexFailedMessage = &regexFailedMsg
		}
		// 解析选择项配置标签
		if selectOptions := field.Tag.Get("selectOptions"); selectOptions != "" {
			// 解析格式：key1:value1,key2:value2
			options := make(map[string]string)
			pairs := strings.Split(selectOptions, ",")
			for _, pair := range pairs {
				kv := strings.Split(strings.TrimSpace(pair), ":")
				if len(kv) == 2 {
					options[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
				}
			}
			fieldConfig.SelectOptions = options
		}

		fields = append(fields, fieldConfig)
	}

	return fields, nil
}

// getFieldTypeInfo 根据反射类型自动推断组件类型和值类型
// 支持指针类型，会自动获取指针指向的原始数据类型
// 参数：
//   - fieldType: 字段的反射类型
//
// 返回值：
//   - EConfigFieldsComponentType: 组件类型
//   - string: 值类型
//
// 类型映射规则：
//   - string/*string -> text组件, string值类型
//   - int/int8/int16/int32/int64/*int/*int8/*int16/*int32/*int64 -> number组件, int值类型
//   - uint/uint8/uint16/uint32/uint64/*uint/*uint8/*uint16/*uint32/*uint64 -> number组件, int值类型
//   - float32/float64/*float32/*float64 -> number组件, float值类型
//   - bool/*bool -> switch组件, bool值类型
//   - 其他类型 -> text组件, string值类型(默认)
func getFieldTypeInfo(fieldType reflect.Type) (c_enum.EConfigFieldsComponentType, string) {
	// 处理指针类型，获取指针指向的原始类型
	originalType := fieldType
	if fieldType.Kind() == reflect.Ptr {
		originalType = fieldType.Elem()
	}

	switch originalType.Kind() {
	case reflect.String:
		return c_enum.EConfigFieldsComponentTypeText, "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return c_enum.EConfigFieldsComponentTypeNumber, "int"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return c_enum.EConfigFieldsComponentTypeNumber, "int"
	case reflect.Float32, reflect.Float64:
		return c_enum.EConfigFieldsComponentTypeNumber, "float"
	case reflect.Bool:
		return c_enum.EConfigFieldsComponentTypeSwitch, "bool"
	default:
		return c_enum.EConfigFieldsComponentTypeText, "string"
	}
}
