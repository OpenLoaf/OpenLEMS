package c_base

import (
	"github.com/shockerli/cvt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

// BuildConfigStructFields 通过反射解析结构体字段，构建配置字段列表
// 支持的标签字段：
//   - json: JSON序列化字段名，必需标签，格式如 "fieldName" 或 "fieldName,omitempty"
//   - name: 字段显示名称，优先级高于desc和json字段名
//   - desc/dc: 字段描述信息，用于生成字段名称和描述
//   - ct: 组件类型，可选值：text(文本)、number(数字)、switch(开关)、singleSelect(单选)、multiSelect(多选)、date(日期)、time(时间)、dateTime(日期时间)
//   - vt: 值类型，可选值：string(字符串)、int(整数)、float(浮点数)、bool(布尔值)
//   - min: 最小值限制，仅对数值类型有效
//   - max: 最大值限制，仅对数值类型有效
//   - default: 默认值
//   - regex: 正则表达式验证规则
//   - regexFailedMessage: 正则表达式验证失败时的提示信息
//   - selectOptions: 选择项配置，格式为 "key1:value1,key2:value2"，用于单选和多选组件
//
// 示例：
//
//	type Config struct {
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

	t := v.Type()
	var fields []*SConfigStructFields

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
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

		// 解析组件类型标签(ct)
		if ct := field.Tag.Get("ct"); ct != "" {
			fieldConfig.ComponentType = EConfigFieldsComponentType(ct)
		}
		// 解析值类型标签(vt)
		if vt := field.Tag.Get("vt"); vt != "" {
			fieldConfig.ValueType = vt
		}

		// 解析数值范围限制标签
		if minStr := field.Tag.Get("min"); minStr != "" {
			fieldConfig.Min = cvt.Uint8(minStr)
		}
		if maxStr := field.Tag.Get("max"); maxStr != "" {
			fieldConfig.Max = cvt.Uint8(maxStr)
		}
		// 解析默认值标签
		if defaultVal := field.Tag.Get("default"); defaultVal != "" {
			fieldConfig.Default = defaultVal
		}
		// 解析正则表达式验证标签
		if regex := field.Tag.Get("regex"); regex != "" {
			fieldConfig.Regex = regex
		}
		// 解析正则表达式失败提示信息标签
		if regexFailedMsg := field.Tag.Get("regexFailedMessage"); regexFailedMsg != "" {
			fieldConfig.RegexFailedMessage = regexFailedMsg
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
// 参数：
//   - fieldType: 字段的反射类型
//
// 返回值：
//   - EConfigFieldsComponentType: 组件类型
//   - string: 值类型
//
// 类型映射规则：
//   - string -> text组件, string值类型
//   - int/int8/int16/int32/int64 -> number组件, int值类型
//   - uint/uint8/uint16/uint32/uint64 -> number组件, int值类型
//   - float32/float64 -> number组件, float值类型
//   - bool -> switch组件, bool值类型
//   - 其他类型 -> text组件, string值类型(默认)
func getFieldTypeInfo(fieldType reflect.Type) (EConfigFieldsComponentType, string) {
	switch fieldType.Kind() {
	case reflect.String:
		return EConfigFieldsComponentTypeText, "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return EConfigFieldsComponentTypeNumber, "int"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return EConfigFieldsComponentTypeNumber, "int"
	case reflect.Float32, reflect.Float64:
		return EConfigFieldsComponentTypeNumber, "float"
	case reflect.Bool:
		return EConfigFieldsComponentTypeSwitch, "bool"
	default:
		return EConfigFieldsComponentTypeText, "string"
	}
}
