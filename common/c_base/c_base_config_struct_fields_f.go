package c_base

import (
	"common/c_enum"
	"reflect"
	"strings"

	"github.com/shockerli/cvt"
	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"
)

// TagMapping 标签映射结构，存储字段名到标签值的映射
type TagMapping struct {
	FieldName string // SConfigStructFields中的字段名
	JsonTag   string // json标签值
	YamlTag   string // yaml标签值
	ShortTag  string // short标签值
}

var (
	// 缓存SConfigStructFields的标签映射，避免重复反射
	configFieldsTagCache map[string]*TagMapping
)

func BuildDescriptionFromYaml(yamlData []byte, deviceConfig ...any) *SDriverInfo {
	info := &SDriverInfo{}
	err := yaml.Unmarshal(yamlData, info)
	if err != nil {
		panic(errors.Errorf("解析版本信息失败！请检查build.yaml文件!%+v", err))
	}

	if len(deviceConfig) > 0 && deviceConfig[0] != nil {
		f, err := BuildConfigStructFields(deviceConfig[0])
		if err != nil {
			panic(errors.Errorf("配置对象中的Fileds解析失败!%+v", err))
		}
		info.SetConfigStructFields(f)
	}
	return info
}

// initConfigFieldsTagCache 初始化SConfigStructFields的标签映射缓存
func initConfigFieldsTagCache() {
	if configFieldsTagCache != nil {
		return
	}

	configFieldsTagCache = make(map[string]*TagMapping)
	structType := reflect.TypeOf(SConfigStructFields{})

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if !field.IsExported() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		yamlTag := field.Tag.Get("yaml")
		shortTag := field.Tag.Get("short")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// 解析json标签，获取字段名（忽略omitempty等选项）
		jsonName := strings.Split(jsonTag, ",")[0]

		mapping := &TagMapping{
			FieldName: field.Name,
			JsonTag:   jsonName,
			YamlTag:   yamlTag,
			ShortTag:  shortTag,
		}

		// 使用多个key来索引同一个映射，支持json、yaml、short三种标签查找
		configFieldsTagCache[jsonName] = mapping
		if yamlTag != "" && yamlTag != jsonName {
			configFieldsTagCache[yamlTag] = mapping
		}
		if shortTag != "" && shortTag != jsonName && shortTag != yamlTag {
			configFieldsTagCache[shortTag] = mapping
		}
	}
}

// getTagMappingByTag 根据标签值获取对应的标签映射
func getTagMappingByTag(tagValue string) *TagMapping {
	initConfigFieldsTagCache()
	return configFieldsTagCache[tagValue]
}

// BuildConfigStructFields 通过反射解析结构体字段，构建配置字段列表
// 支持嵌套结构体的平铺展开，包括匿名嵌入字段
//
// 动态标签解析功能：
// 该方法会动态读取SConfigStructFields结构体中定义的标签映射(json、yaml、short)，
// 然后在被解析的结构体字段中查找匹配的标签值。
//
// 标签优先级：short > json > yaml
//
// 支持的标签字段（基于SConfigStructFields定义）：
//   - name/name/name: 字段显示名称
//   - code/code/code: JSON序列化字段名，必需标签
//   - group/group/group: 字段分组
//   - valueType/value_type/vt: 值类型，可选值：string、int、float、bool
//   - componentType/component_type/ct: 组件类型，如text、number、switch等
//   - step/step/step: 数值步长
//   - required/required/req: 是否必填
//   - unit/unit/unit: 单位信息
//   - min/min/min: 最小值限制
//   - max/max/max: 最大值限制
//   - default/default/def: 默认值
//   - selectOptions/select_options/opts: 选择项配置
//   - regex/regex/regex: 正则表达式验证规则
//   - regexFailedMessage/regex_failed_message/rfm: 正则验证失败提示
//   - description/description/desc: 字段描述信息
//
// 示例：
//
//	type BaseConfig struct {
//	    UnitId uint8 `code:"unitId" name:"ModbusID" min:"1" max:"255" def:"1"`
//	}
//
//	type GpioDeviceConfig struct {
//	    BaseConfig                    // 匿名嵌入，字段会被平铺展开
//	    Key string `code:"name" name:"设备名称" desc:"设备的显示名称" ct:"text" vt:"string" regex:"^[a-zA-Z0-9_-]+$" rfm:"只能包含字母、数字、下划线和连字符"`
//	    Port int    `code:"port" name:"端口号" desc:"设备通信端口" ct:"number" vt:"int" min:"1" max:"65535" def:"8080"`
//	    Enabled bool `code:"enable" name:"启用状态" desc:"是否启用设备" ct:"switch" vt:"bool" def:"true"`
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

// parseFieldTagValue 动态解析字段标签值
// 根据SConfigStructFields中定义的标签映射，从字段标签中提取对应的值
func parseFieldTagValue(field reflect.StructField, targetFieldName string) string {
	initConfigFieldsTagCache()

	// 遍历字段的所有标签，查找匹配的标签值
	for _, mapping := range configFieldsTagCache {
		if mapping.FieldName != targetFieldName {
			continue
		}

		// 按优先级检查标签：short > json > yaml
		if mapping.ShortTag != "" {
			if value := field.Tag.Get(mapping.ShortTag); value != "" {
				return strings.Split(value, ",")[0] // 移除omitempty等选项
			}
		}
		if mapping.JsonTag != "" {
			if value := field.Tag.Get(mapping.JsonTag); value != "" {
				return strings.Split(value, ",")[0] // 移除omitempty等选项
			}
		}
		if mapping.YamlTag != "" {
			if value := field.Tag.Get(mapping.YamlTag); value != "" {
				return strings.Split(value, ",")[0] // 移除omitempty等选项
			}
		}
		break
	}

	return ""
}

// populateFieldConfigFromTags 动态填充字段配置
// 通过反射SConfigStructFields结构体，动态解析所有支持的标签
func populateFieldConfigFromTags(field reflect.StructField, fieldConfig *SConfigStructFields) {
	structType := reflect.TypeOf(SConfigStructFields{})
	fieldConfigValue := reflect.ValueOf(fieldConfig).Elem()

	// 遍历SConfigStructFields的所有字段
	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)
		if !structField.IsExported() {
			continue
		}

		// 跳过Code字段，因为它已经在外部处理了
		if structField.Name == "Code" {
			continue
		}

		// 获取标签值
		tagValue := parseFieldTagValue(field, structField.Name)
		if tagValue == "" {
			continue
		}

		// 获取目标字段
		targetField := fieldConfigValue.FieldByName(structField.Name)
		if !targetField.IsValid() || !targetField.CanSet() {
			continue
		}

		// 根据字段类型设置值
		switch structField.Type.Kind() {
		case reflect.String:
			targetField.SetString(tagValue)

		case reflect.Bool:
			targetField.SetBool(cvt.Bool(tagValue))

		case reflect.Ptr:
			// 处理指针类型
			elemType := structField.Type.Elem()
			switch elemType.Kind() {
			case reflect.String:
				value := tagValue
				targetField.Set(reflect.ValueOf(&value))

			case reflect.Int64:
				value := cvt.Int64(tagValue)
				targetField.Set(reflect.ValueOf(&value))

			case reflect.Float32:
				value := cvt.Float32(tagValue)
				targetField.Set(reflect.ValueOf(&value))
			}

		case reflect.Slice:
			// 处理SelectOptions字段
			if structField.Name == "SelectOptions" {
				pairs := strings.Split(tagValue, ",")
				options := make([]string, 0, len(pairs))
				for _, pair := range pairs {
					trimmedPair := strings.TrimSpace(pair)
					if trimmedPair != "" {
						options = append(options, trimmedPair)
					}
				}
				targetField.Set(reflect.ValueOf(options))
			}

		default:
			// 处理枚举类型 (EConfigFieldsComponentType)
			if structField.Name == "ComponentType" {
				enumValue := c_enum.EConfigFieldsComponentType(tagValue)
				targetField.Set(reflect.ValueOf(enumValue))
			}
		}
	}
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

		// 动态获取json标签，用于确定字段的JSON序列化名称
		jsonName := parseFieldTagValue(field, "Code") // Code字段对应json标签
		if jsonName == "" {
			// 如果没有找到对应的标签，尝试使用传统的json标签作为后备
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" || jsonTag == "-" {
				continue
			}
			jsonName = strings.Split(jsonTag, ",")[0]
		}
		if jsonName == "" {
			jsonName = field.Name
		}

		// 如果有前缀，则组合字段名
		if prefix != "" {
			jsonName = prefix + "." + jsonName
		}

		// 根据字段类型确定组件类型和值类型
		componentType, valueType := getFieldTypeInfo(field.Type)

		fieldConfig := &SConfigStructFields{
			Code:          jsonName,
			ValueType:     valueType,
			ComponentType: componentType,
		}

		// 动态解析所有SConfigStructFields中定义的字段
		populateFieldConfigFromTags(field, fieldConfig)

		// 设置字段名称，优先级：动态解析的name > desc > json字段名
		if fieldConfig.Name == "" {
			if fieldConfig.Description != "" {
				fieldConfig.Name = fieldConfig.Description
			} else {
				fieldConfig.Name = jsonName
			}
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
func getFieldTypeInfo(fieldType reflect.Type) (c_enum.EConfigFieldsComponentType, c_enum.EConfigFieldsValueType) {
	// 处理指针类型，获取指针指向的原始类型
	originalType := fieldType
	if fieldType.Kind() == reflect.Ptr {
		originalType = fieldType.Elem()
	}

	switch originalType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return c_enum.EConfigFieldsComponentTypeNumber, c_enum.EConfigFieldsValueTypeInt
	case reflect.Float32, reflect.Float64:
		return c_enum.EConfigFieldsComponentTypeNumber, c_enum.EConfigFieldsValueTypeFloat
	case reflect.Bool:
		return c_enum.EConfigFieldsComponentTypeSwitch, c_enum.EConfigFieldsValueTypeBoolean
	default:
		return c_enum.EConfigFieldsComponentTypeText, c_enum.EConfigFieldsValueTypeString
	}
}
