package c_base

import (
	"fmt"
	"reflect"
)

type SCustomService struct {
	Name        string `json:"name" yaml:"name"  dc:"服务名称，支持i18n"`
	DisplayName string `json:"displayName" yaml:"displayName" dc:"执行的方法名"`
	Description string `json:"description" yaml:"description" dc:"备注"`
}

func (s *SDriverDescription) ExecuteCustomService(functionName string, instance any, params any) error {
	// 执行自定义方法
	if instance == nil {
		return fmt.Errorf("custom service instance is nil")
	}

	// 反射前先判断缓存中是否存在
	if s.reflectMethodCache == nil {
		s.reflectMethodCache = make(map[string]reflect.Value)
	}

	var (
		method reflect.Value
		ok     bool
	)

	// 如果缓冲中不存在，就反射新增
	if method, ok = s.reflectMethodCache[functionName]; !ok {
		method = reflect.ValueOf(instance).MethodByName(functionName)
		if !method.IsValid() {
			return fmt.Errorf("service %s not found", functionName)
		}
		s.reflectMethodCache[functionName] = method
	}

	// 空参数调用
	values := method.Call(nil)
	if len(values) == 1 {
		if err, ok := values[0].Interface().(error); ok {
			return err
		}
		return nil
	}

	fmt.Printf("当前函数: %s 返回的参数数据不为1 ！返回的内容为: %v", functionName, values)

	return nil
}
