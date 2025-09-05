package c_base

import (
	"common/c_log"
	"common/c_status"
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func BuildDescriptionFromYaml(yamlData []byte, IsVirtualDevice ...bool) *SDriverInfo {
	info := &SDriverInfo{}
	err := yaml.Unmarshal(yamlData, info)
	if err != nil {
		panic(errors.Errorf("解析版本信息失败！请检查version.yaml文件!%+v", err))
	}

	if len(IsVirtualDevice) > 0 && IsVirtualDevice[0] {
		info.IsVirtualDevice = true
	}

	return info
}

// GetTelemetry 反射获取遥测信息 用于实现IDriver接口
func (s *SDriverInfo) GetTelemetry(key string, instance any) (any, error) {
	// 输入参数验证
	if key == "" {
		return nil, errors.New("telemetry key cannot be empty")
	}
	if instance == nil {
		return nil, errors.New("instance cannot be nil")
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
	if method, ok = s.reflectMethodCache[key]; !ok {
		functionName := fmt.Sprintf("Get%s", capitalizeFirstLetter(key))

		// 获取实例的反射值
		instanceValue := reflect.ValueOf(instance)
		if !instanceValue.IsValid() {
			return nil, errors.Errorf("invalid instance for telemetry key: [%s]", key)
		}

		// 获取方法
		method = instanceValue.MethodByName(functionName)
		if !method.IsValid() {
			return nil, errors.Errorf("TelemetryKey: [%s] method [%s] not found", key, functionName)
		}

		// 检查方法是否可以被调用
		if !method.CanInterface() {
			return nil, errors.Errorf("TelemetryKey: [%s] method [%s] cannot be called", key, functionName)
		}

		// 缓存方法
		s.reflectMethodCache[key] = method
	}

	// 使用defer recover处理可能的panic
	var result any
	var callErr error
	func() {
		defer func() {
			if r := recover(); r != nil {
				callErr = errors.Errorf("GetTelemetry panic! key: %s, error: %+v", key, r)
				c_log.Errorf(context.Background(), "GetTelemetry panic! key: %s, error: %+v", key, r)
			}
		}()

		// 空参数调用
		values := method.Call(nil)
		if len(values) == 0 {
			callErr = errors.Errorf("function %s returned no values", key)
			return
		}

		if len(values) == 1 {
			// 只有返回值，没有error
			result = values[0].Interface()
			return
		}

		if len(values) != 2 {
			callErr = errors.Errorf("function %s return value length is not 1 or 2, got %d", key, len(values))
			return
		}

		// 检查第二个返回值是否为error
		if values[1].Interface() != nil {
			// 安全地转换为error类型
			if err, ok := values[1].Interface().(error); ok {
				callErr = err
			} else {
				callErr = errors.Errorf("function %s second return value is not error type", key)
			}
			return
		}

		// 返回成功结果
		result = values[0].Interface()
	}()

	if callErr != nil {
		return nil, callErr
	}

	return result, nil
}

func (s *SDriverInfo) GetAllTelemetry(instance IDevice) map[string]any {
	if instance == nil || instance.GetProtocolStatus() != c_status.EProtocolConnected { // 如果实例不是连接成功的，就不要返回数据了
		return nil
	}

	telemetryMap := make(map[string]any, len(s.Telemetry))
	for _, telemetry := range s.Telemetry {
		value, err := s.GetTelemetry(telemetry.Name, instance)
		if err != nil {
			// 这里有时候err也是正常的，比如系统刚启动，但是页面一直在请求
			c_log.Debugf(context.Background(), "Get telemetry %s error: %+v", telemetry.Name, err)
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

func (s *SDriverInfo) ExecuteCustomService(functionName string, instance IDevice, params any) error {

	// 执行自定义方法
	if instance == nil {
		c_log.Errorf(context.Background(), "ExecuteCustomService [%s] instance is nil", functionName)
		return fmt.Errorf("custom service instance is nil")
	}

	ctx := context.WithValue(context.Background(), ConstCtxKeyDeviceId, instance.GetConfig().Id)

	// 判断一下是否允许这个方法调用
	var service *SDriverService
	for _, v := range s.Service {
		if v.Name == functionName {
			service = v
			break
		}
	}
	if service == nil {
		c_log.BizInfof(ctx, "执行自定义服务失败！原因：[%s]服务未定义！", functionName)
		return fmt.Errorf("custom service %s not support", functionName)
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
			c_log.BizInfof(ctx, "执行自定义服务[%s]失败！原因该服务对应的方法[%s]不存在！", service.DisplayName, service.Name)
			return fmt.Errorf("service %s not found", functionName)
		}
		s.reflectMethodCache[functionName] = method
	}

	// 空参数调用
	values := method.Call(nil)
	if len(values) == 1 {
		if err, ok := values[0].Interface().(error); ok {
			c_log.BizInfof(ctx, "执行自定义服务[%s]失败！原因：%s", service.DisplayName, err.Error())
			return err
		}
		c_log.BizInfof(ctx, "执行自定义服务[%s]成功！", service.DisplayName)
		return nil
	}

	c_log.BizInfof(ctx, "执行自定义服务[%s]成功！", service.DisplayName)

	fmt.Printf("当前函数: %s 返回的参数数据不为1 ！返回的内容为: %v", functionName, values)

	return nil
}
