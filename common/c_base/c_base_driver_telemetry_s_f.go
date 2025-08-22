package c_base

import (
	"common/c_log"
	"context"
	"fmt"
	"reflect"
	"strings"
)

type STelemetry struct {
	Name        string `json:"name,omitempty" yaml:"name"`     // 遥测名称
	DisplayName string `json:"displayName" yaml:"displayName"` // 遥测名称的国际化覆盖
	Unit        string `json:"unit,omitempty" yaml:"unit"`     // 单位
	Remark      string `json:"remark,omitempty" yaml:"remark"` // 备注
}

func (s *STelemetry) String() string {
	i18nKey := "-"
	if s.DisplayName != "" {
		i18nKey = s.DisplayName
	}
	return fmt.Sprintf("%s\t%s\t%s\t%s\t", s.Name, i18nKey, s.Unit, s.Remark)
}

// GetTelemetry 反射获取遥测信息 用于实现IDriver接口
func (s *SDriverDescription) GetTelemetry(key string, instance any) (any, error) {

	// 反射前先判断缓存中是否存在
	if s.reflectMethodCache == nil {
		s.reflectMethodCache = make(map[string]reflect.Value)
	}

	var (
		method reflect.Value
		ok     bool
	)

	// 如果缓冲中不存在，就反射新增
	if method, ok = s.reflectMethodCache[key]; !ok {
		functionName := fmt.Sprintf("Get%s", capitalizeFirstLetter(key))
		method = reflect.ValueOf(instance).MethodByName(functionName)
		if !method.IsValid() {
			return nil, fmt.Errorf("method %s not found", key)
		}
		s.reflectMethodCache[key] = method
	}

	defer func() {
		if r := recover(); r != nil {
			c_log.Errorf(context.Background(), "GetTelemetry Painc! key: %s Error: %v\n", s, r)
		}
	}()

	// 空参数调用
	value := method.Call(nil)
	if len(value) == 1 {
		return value[0].Interface(), nil
	}

	if len(value) != 2 {
		return nil, fmt.Errorf("function %s return value length is not 2", key)
	}
	if value[1].Interface() != nil {
		return nil, value[1].Interface().(error)
	}
	return value[0].Interface(), nil
}

func (s *SDriverDescription) GetAllTelemetry(instance any) map[string]any {
	telemetryMap := make(map[string]any, len(s.Telemetry))
	for _, telemetry := range s.Telemetry {
		value, err := s.GetTelemetry(telemetry.Name, instance)
		if err != nil {
			continue
		}
		telemetryMap[telemetry.Name] = value
	}
	return telemetryMap
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
