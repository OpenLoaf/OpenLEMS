package c_base

import (
	"common/c_enum"
	"reflect"
	"regexp"
	"strconv"
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
		f, err := BuildConfigPoints(deviceConfig[0])
		if err != nil {
			panic(errors.Errorf("配置对象中的Fileds解析失败!%+v", err))
		}
		info.ConfigPoints = f
	}
	return info
}

// initConfigFieldsTagCache 初始化SConfigStructFields的标签映射缓存
func initConfigFieldsTagCache() {
	if configFieldsTagCache != nil {
		return
	}

	configFieldsTagCache = make(map[string]*TagMapping)
	structType := reflect.TypeOf(SFieldDefinition{})

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
//   - key/key/key: JSON序列化字段名，必需标签
//   - group/group/group: 字段分组
//   - valueType/value_type/vt: 值类型，可选值：string、int、float、bool
//   - componentType/component_type/ct: 组件类型，如text、number、switch等
//   - step/step/step: 数值步长
//   - required/required/req: 是否必填
//   - unit/unit/unit: 单位信息
//   - min/min/min: 最小值限制
//   - max/max/max: 最大值限制
//   - default/default/def: 默认值
//   - valueExplain/valueExplain/ve: 值解释配置，格式：key1:value1,key2:value2
//   - paramExplain/paramExplain/pe: 参数解释配置，格式：key1:value1,key2:value2
//   - regex/regex/regex: 正则表达式验证规则
//   - regexFailedMessage/regex_failed_message/rfm: 正则验证失败提示
//   - description/description/desc: 字段描述信息
//
// 示例：
//
//	type BaseConfig struct {
//	    UnitId uint8 `key:"unitId" name:"ModbusID" min:"1" max:"255" def:"1"`
//	}
//
//	type GpioDeviceConfig struct {
//	    BaseConfig                    // 匿名嵌入，字段会被平铺展开
//	    Key string `key:"name" name:"设备名称" desc:"设备的显示名称" ct:"text" vt:"string" regex:"^[a-zA-Z0-9_-]+$" rfm:"只能包含字母、数字、下划线和连字符"`
//	    Port int    `key:"port" name:"端口号" desc:"设备通信端口" ct:"number" vt:"int" min:"1" max:"65535" def:"8080"`
//	    Enabled bool `key:"enable" name:"启用状态" desc:"是否启用设备" ct:"switch" vt:"bool" def:"true"`
//	}
func BuildConfigStructFields(config any) ([]*SFieldDefinition, error) {
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

		// 按优先级检查标签：short > json > yaml > v
		if mapping.ShortTag != "" {
			if value := field.Tag.Get(mapping.ShortTag); value != "" {
				return parseTagValue(value, mapping.FieldName)
			}
		}
		if mapping.JsonTag != "" {
			if value := field.Tag.Get(mapping.JsonTag); value != "" {
				return parseTagValue(value, mapping.FieldName)
			}
		}
		if mapping.YamlTag != "" {
			if value := field.Tag.Get(mapping.YamlTag); value != "" {
				return parseTagValue(value, mapping.FieldName)
			}
		}
		// 支持GoFrame验证规则v标签
		if value := field.Tag.Get("v"); value != "" {
			return parseGoFrameValidationTag(value, mapping.FieldName)
		}
		break
	}

	return ""
}

// parseGoFrameValidationTag 解析GoFrame验证规则v标签
// 支持常见的验证规则：required、min、max、length等
func parseGoFrameValidationTag(vTag, fieldName string) string {
	if vTag == "" {
		return ""
	}

	// 解析验证规则
	rules := strings.Split(vTag, "|")

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)

		switch {
		case rule == "required":
			// required规则表示必填
			if fieldName == "Required" {
				return "true"
			}
		case strings.HasPrefix(rule, "min:"):
			// min规则，如 min:1
			if fieldName == "Min" {
				value := strings.TrimPrefix(rule, "min:")
				if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
					return strconv.FormatInt(intVal, 10)
				}
			}
		case strings.HasPrefix(rule, "max:"):
			// max规则，如 max:255
			if fieldName == "Max" {
				value := strings.TrimPrefix(rule, "max:")
				if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
					return strconv.FormatInt(intVal, 10)
				}
			}
		case strings.HasPrefix(rule, "length:"):
			// length规则，如 length:6,40
			if fieldName == "Min" || fieldName == "Max" {
				value := strings.TrimPrefix(rule, "length:")
				parts := strings.Split(value, ",")
				if len(parts) == 2 {
					if fieldName == "Min" {
						if intVal, err := strconv.ParseInt(parts[0], 10, 64); err == nil {
							return strconv.FormatInt(intVal, 10)
						}
					} else {
						if intVal, err := strconv.ParseInt(parts[1], 10, 64); err == nil {
							return strconv.FormatInt(intVal, 10)
						}
					}
				}
			}
		case strings.HasPrefix(rule, "regex:"):
			// regex规则，如 regex:^[a-zA-Z0-9_-]+$
			if fieldName == "Regex" {
				value := strings.TrimPrefix(rule, "regex:")
				return value
			}
		}
	}

	return ""
}

// parseTagValue 解析标签值，对于某些字段类型不进行逗号分割
func parseTagValue(value, fieldName string) string {
	// 对于正则表达式和正则表达式失败消息字段，不进行逗号分割
	if fieldName == "Regex" || fieldName == "RegexFailedMessage" {
		return value
	}
	// 对于其他字段，按逗号分割并取第一部分（移除omitempty等选项）
	return strings.Split(value, ",")[0]
}

// populateFieldConfigFromTags 动态填充字段配置
// 通过反射SConfigStructFields结构体，动态解析所有支持的标签
func populateFieldConfigFromTags(field reflect.StructField, fieldConfig *SFieldDefinition) {
	structType := reflect.TypeOf(SFieldDefinition{})
	fieldConfigValue := reflect.ValueOf(fieldConfig).Elem()

	// 遍历SConfigStructFields的所有字段
	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)
		if !structField.IsExported() {
			continue
		}

		// 跳过Code字段，因为它已经在外部处理了
		if structField.Name == "Key" {
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
			// 处理 ValueExplain 和 ParamExplain 字段
			if structField.Name == "ValueExplain" || structField.Name == "ParamExplain" {
				// 解析 Explain 字符串格式为 []*SFieldExplain 对象数组
				explains := ParseExplainString(tagValue)
				targetField.Set(reflect.ValueOf(explains))
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
func buildConfigStructFieldsRecursive(structType reflect.Type, prefix string) ([]*SFieldDefinition, error) {
	var fields []*SFieldDefinition

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
		jsonName := parseFieldTagValue(field, "Key") // Code字段对应json标签
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

		fieldConfig := &SFieldDefinition{
			Key:           jsonName,
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

// ValidateConfigStructFields 验证配置结构体字段定义是否正确
// 用于验证给前端的数据结构是否正常
// 参数：
//   - config: 配置结构体实例
//
// 返回值：
//   - bool: 验证是否通过
//   - error: 验证失败的错误信息
func ValidateConfigStructFields(config any) (bool, error) {
	if config == nil {
		return false, errors.New("config is nil")
	}

	// 构建字段定义
	fields, err := BuildConfigStructFields(config)
	if err != nil {
		return false, errors.Wrap(err, "failed to build config struct fields")
	}

	// 验证每个字段定义
	for _, field := range fields {
		if err := validateFieldDefinition(field); err != nil {
			return false, errors.Wrapf(err, "field '%s' validation failed", field.Key)
		}
	}

	return true, nil
}

// ValidateConfigData 验证前端传递的数据是否符合字段定义要求
// 用于验证前端传递过来的数据是否正常
// 参数：
//   - config: 配置结构体实例（用于获取字段定义）
//   - data: 前端传递的数据（map格式）
//
// 返回值：
//   - bool: 验证是否通过
//   - error: 验证失败的错误信息
func ValidateConfigData(config any, data map[string]any) (bool, error) {
	if config == nil {
		return false, errors.New("config is nil")
	}
	if data == nil {
		return false, errors.New("data is nil")
	}

	// 构建字段定义
	fields, err := BuildConfigStructFields(config)
	if err != nil {
		return false, errors.Wrap(err, "failed to build config struct fields")
	}

	// 创建字段映射，便于快速查找
	fieldMap := make(map[string]*SFieldDefinition)
	for _, field := range fields {
		fieldMap[field.Key] = field
	}

	// 验证每个数据字段
	for key, value := range data {
		field, exists := fieldMap[key]
		if !exists {
			// 如果字段不在定义中，跳过验证（允许额外字段）
			continue
		}

		if err := validateFieldValue(field, value); err != nil {
			return false, errors.Wrapf(err, "field '%s' value validation failed", key)
		}
	}

	// 检查必填字段
	for _, field := range fields {
		if field.Required {
			if _, exists := data[field.Key]; !exists {
				return false, errors.Errorf("required field '%s' is missing", field.Key)
			}
		}
	}

	return true, nil
}

// validateFieldDefinition 验证字段定义是否正确
func validateFieldDefinition(field *SFieldDefinition) error {
	if field == nil {
		return errors.New("field definition is nil")
	}

	// 检查必需字段
	if field.Key == "" {
		return errors.New("field key is required")
	}
	if field.ComponentType == "" {
		return errors.New("component type is required")
	}
	if field.ValueType == "" {
		return errors.New("value type is required")
	}

	// 检查组件类型和值类型的匹配性
	if err := validateComponentValueTypeMatch(field.ComponentType, field.ValueType); err != nil {
		return errors.Wrap(err, "component type and value type mismatch")
	}

	// 检查数值范围
	if field.Min != nil && field.Max != nil && *field.Min > *field.Max {
		return errors.New("min value cannot be greater than max value")
	}

	// 检查正则表达式
	if field.Regex != nil && *field.Regex != "" {
		if _, err := regexp.Compile(*field.Regex); err != nil {
			return errors.Errorf("invalid regex pattern: %v", err)
		}
	}

	// 检查选择组件的选项配置
	if err := validateSelectOptions(field); err != nil {
		return errors.Wrap(err, "select options validation failed")
	}

	return nil
}

// validateFieldValue 验证字段值是否符合定义要求
func validateFieldValue(field *SFieldDefinition, value any) error {
	if field == nil {
		return errors.New("field definition is nil")
	}

	// 处理nil值
	if value == nil {
		if field.Required {
			return errors.New("required field cannot be nil")
		}
		return nil // 非必填字段允许nil值
	}

	// 类型转换和验证
	switch field.ValueType {
	case c_enum.EConfigFieldsValueTypeString:
		return validateStringValue(field, value)
	case c_enum.EConfigFieldsValueTypeInt:
		return validateIntValue(field, value)
	case c_enum.EConfigFieldsValueTypeFloat:
		return validateFloatValue(field, value)
	case c_enum.EConfigFieldsValueTypeBoolean:
		return validateBoolValue(field, value)
	default:
		return errors.Errorf("unsupported value type: %s", field.ValueType)
	}
}

// validateComponentValueTypeMatch 验证组件类型和值类型是否匹配
func validateComponentValueTypeMatch(componentType c_enum.EConfigFieldsComponentType, valueType c_enum.EConfigFieldsValueType) error {
	switch componentType {
	case c_enum.EConfigFieldsComponentTypeText:
		if valueType != c_enum.EConfigFieldsValueTypeString {
			return errors.Errorf("text component requires string value type, got %s", valueType)
		}
	case c_enum.EConfigFieldsComponentTypeNumber:
		if valueType != c_enum.EConfigFieldsValueTypeInt && valueType != c_enum.EConfigFieldsValueTypeFloat {
			return errors.Errorf("number component requires int or float value type, got %s", valueType)
		}
	case c_enum.EConfigFieldsComponentTypeSwitch:
		if valueType != c_enum.EConfigFieldsValueTypeBoolean {
			return errors.Errorf("switch component requires bool value type, got %s", valueType)
		}
	case c_enum.EConfigFieldsComponentTypeSingleSelect, c_enum.EConfigFieldsComponentTypeMultiSelect:
		if valueType != c_enum.EConfigFieldsValueTypeString {
			return errors.Errorf("select component requires string value type, got %s", valueType)
		}
	case c_enum.EConfigFieldsComponentTypeDate, c_enum.EConfigFieldsComponentTypeTime, c_enum.EConfigFieldsComponentTypeDateTime:
		if valueType != c_enum.EConfigFieldsValueTypeString {
			return errors.Errorf("date/time component requires string value type, got %s", valueType)
		}
	}
	return nil
}

// validateSelectOptions 验证选择组件的选项配置
func validateSelectOptions(field *SFieldDefinition) error {
	if field.ComponentType != c_enum.EConfigFieldsComponentTypeSingleSelect &&
		field.ComponentType != c_enum.EConfigFieldsComponentTypeMultiSelect {
		return nil // 非选择组件不需要验证选项
	}

	// 检查是否有选项配置（通过ValueExplain或ParamExplain）
	hasOptions := len(field.ValueExplain) > 0 || len(field.ParamExplain) > 0
	if !hasOptions {
		return errors.New("select component must have options configured via ValueExplain or ParamExplain")
	}

	// 验证ValueExplain格式
	for _, explain := range field.ValueExplain {
		if explain.Key == "" || explain.Value == "" {
			return errors.New("ValueExplain entries must have both key and value")
		}
	}

	// 验证ParamExplain格式
	for _, explain := range field.ParamExplain {
		if explain.Key == "" || explain.Value == "" {
			return errors.New("ParamExplain entries must have both key and value")
		}
	}

	return nil
}

// validateStringValue 验证字符串值
func validateStringValue(field *SFieldDefinition, value any) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.Errorf("expected string value, got %T", value)
	}

	// 长度验证（使用min/max作为长度限制）
	if field.Min != nil && int64(len(strValue)) < *field.Min {
		return errors.Errorf("string length %d is below minimum %d", len(strValue), *field.Min)
	}
	if field.Max != nil && int64(len(strValue)) > *field.Max {
		return errors.Errorf("string length %d is above maximum %d", len(strValue), *field.Max)
	}

	// 正则表达式验证
	if field.Regex != nil && *field.Regex != "" {
		matched, err := regexp.MatchString(*field.Regex, strValue)
		if err != nil {
			return errors.Errorf("regex validation error: %v", err)
		}
		if !matched {
			if field.RegexFailedMessage != nil && *field.RegexFailedMessage != "" {
				return errors.New(*field.RegexFailedMessage)
			}
			return errors.Errorf("value '%s' does not match regex pattern", strValue)
		}
	}

	// 选择组件选项验证
	if field.ComponentType == c_enum.EConfigFieldsComponentTypeSingleSelect {
		return validateSingleSelectValue(field, strValue)
	}
	if field.ComponentType == c_enum.EConfigFieldsComponentTypeMultiSelect {
		return validateMultiSelectValue(field, strValue)
	}

	return nil
}

// validateIntValue 验证整数值
func validateIntValue(field *SFieldDefinition, value any) error {
	intValue := cvt.Int64(value)

	// 范围验证
	if field.Min != nil && intValue < *field.Min {
		return errors.Errorf("value %d is below minimum %d", intValue, *field.Min)
	}
	if field.Max != nil && intValue > *field.Max {
		return errors.Errorf("value %d is above maximum %d", intValue, *field.Max)
	}

	return nil
}

// validateFloatValue 验证浮点数值
func validateFloatValue(field *SFieldDefinition, value any) error {
	floatValue := cvt.Float64(value)

	// 范围验证
	if field.Min != nil && floatValue < float64(*field.Min) {
		return errors.Errorf("value %f is below minimum %d", floatValue, *field.Min)
	}
	if field.Max != nil && floatValue > float64(*field.Max) {
		return errors.Errorf("value %f is above maximum %d", floatValue, *field.Max)
	}

	return nil
}

// validateBoolValue 验证布尔值
func validateBoolValue(field *SFieldDefinition, value any) error {
	_, ok := value.(bool)
	if !ok {
		return errors.Errorf("expected bool value, got %T", value)
	}
	return nil
}

// validateSingleSelectValue 验证单选值
func validateSingleSelectValue(field *SFieldDefinition, value string) error {
	// 检查值是否在ValueExplain中
	for _, explain := range field.ValueExplain {
		if explain.Key == value {
			return nil
		}
	}
	return errors.Errorf("value '%s' is not a valid option", value)
}

// validateMultiSelectValue 验证多选值
func validateMultiSelectValue(field *SFieldDefinition, value string) error {
	// 多选值通常是逗号分隔的字符串
	values := strings.Split(value, ",")
	validKeys := make(map[string]bool)

	// 构建有效选项映射
	for _, explain := range field.ValueExplain {
		validKeys[explain.Key] = true
	}

	// 验证每个值
	for _, v := range values {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if !validKeys[v] {
			return errors.Errorf("value '%s' is not a valid option", v)
		}
	}

	return nil
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

// BuildConfigPoints 将配置结构体转换为SConfigPoint列表
// 这是新的统一点位管理系统的一部分，用于替代原有的BuildConfigStructFields方法
func BuildConfigPoints(config any) ([]*SConfigPoint, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	// 先获取字段定义
	fields, err := BuildConfigStructFields(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build config struct fields")
	}

	// 转换为SConfigPoint
	var configPoints []*SConfigPoint
	for _, field := range fields {
		configPoint := convertFieldDefinitionToConfigPoint(field)
		configPoints = append(configPoints, configPoint)
	}

	return configPoints, nil
}

// convertFieldDefinitionToConfigPoint 将SFieldDefinition转换为SConfigPoint
func convertFieldDefinitionToConfigPoint(field *SFieldDefinition) *SConfigPoint {
	// 创建基础SPoint
	basePoint := &SPoint{
		Key:          field.Key,
		Name:         field.Name,
		ValueType:    convertConfigValueTypeToEValueType(field.ValueType),
		Unit:         getStringValue(field.Unit),
		Desc:         field.Description,
		Min:          getInt64Value(field.Min),
		Max:          getInt64Value(field.Max),
		Precise:      getPrecisionFromStep(field.Step),
		ValueExplain: field.ValueExplain,
	}

	// 创建SConfigPoint
	configPoint := &SConfigPoint{
		SPoint:             basePoint,
		Required:           field.Required,
		Default:            field.Default,
		Regex:              field.Regex,
		RegexFailedMessage: field.RegexFailedMessage,
		Step:               field.Step,
	}

	return configPoint
}

// convertConfigValueTypeToEValueType 将EConfigFieldsValueType转换为EValueType
func convertConfigValueTypeToEValueType(valueType c_enum.EConfigFieldsValueType) c_enum.EValueType {
	switch valueType {
	case c_enum.EConfigFieldsValueTypeString:
		return c_enum.EString
	case c_enum.EConfigFieldsValueTypeInt:
		return c_enum.EInt32
	case c_enum.EConfigFieldsValueTypeFloat:
		return c_enum.EFloat32
	case c_enum.EConfigFieldsValueTypeBoolean:
		return c_enum.EBool
	default:
		return c_enum.EString
	}
}

// getPrecisionFromStep 从步长中获取精度
func getPrecisionFromStep(step *float32) uint8 {
	if step == nil {
		return 0
	}

	// 根据步长计算精度
	stepVal := *step
	if stepVal >= 1 {
		return 0
	} else if stepVal >= 0.1 {
		return 1
	} else if stepVal >= 0.01 {
		return 2
	} else if stepVal >= 0.001 {
		return 3
	} else {
		return 4
	}
}

// getStringValue 获取字符串值，处理指针类型
func getStringValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

// getInt64Value 获取int64值，处理指针类型
func getInt64Value(ptr *int64) int64 {
	if ptr != nil {
		return *ptr
	}
	return 0
}
