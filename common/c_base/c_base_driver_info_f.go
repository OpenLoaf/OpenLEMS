package c_base

import (
	"common/c_enum"
	"common/c_log"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
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

func GetAllTelemetry(instance IDevice) map[string]any {

	if instance == nil || instance.GetProtocolStatus() != c_enum.EProtocolConnected { // 如果实例不是连接成功的，就不要返回数据了
		return nil
	}
	driverInfo := instance.GetConfig().DriverInfo
	telemetryMap := make(map[string]any, len(driverInfo.Telemetry))
	for _, telemetry := range driverInfo.Telemetry {
		value, err := driverInfo.getTelemetry(telemetry.Key, instance)
		if err != nil {
			// 这里有时候err也是正常的，比如系统刚启动，但是页面一直在请求
			ctx := context.WithValue(context.Background(), ConstCtxKeyDeviceId, instance.GetConfig().Id)
			c_log.Warningf(ctx, "Get telemetry %driverInfo error: %+v", telemetry.Key, err)
			//fmt.Printf("Get telemetry %driverInfo error: %+v\n", telemetry.Key, err)
			continue
		}
		if value != nil {
			telemetryMap[telemetry.Key] = value
		}
	}
	return telemetryMap
}

func GetAllTelemetryPoint(instance IDevice) []*SPointValue {
	if instance == nil { // 如果实例不是连接成功的，就不要返回数据了
		return nil
	}
	driverInfo := instance.GetConfig().DriverInfo
	var list []*SPointValue
	for _, telemetry := range driverInfo.Telemetry {
		value, err := driverInfo.getTelemetry(telemetry.Key, instance)
		if err != nil {
			// 这里有时候err也是正常的，比如系统刚启动，但是页面一直在请求
			ctx := context.WithValue(context.Background(), ConstCtxKeyDeviceId, instance.GetConfig().Id)
			c_log.Warningf(ctx, "Get telemetry %driverInfo error: %+v", telemetry.Key, err)
			//fmt.Printf("Get telemetry %driverInfo error: %+v\n", telemetry.Key, err)
			continue
		}

		list = append(list, NewPointValue(instance.GetConfig().Id, telemetry.ToPoint(ResolvingValueType(value), instance.GetConfig().Params), value))
	}
	return list
}

// ResolvingValueType TODO 提出来
func ResolvingValueType(value any) c_enum.EValueType {
	if value == nil {
		return c_enum.EString
	}

	// 使用反射获取值的实际类型
	valueType := reflect.TypeOf(value)

	// 处理指针类型
	if valueType.Kind() == reflect.Ptr {
		if reflect.ValueOf(value).IsNil() {
			return c_enum.EString
		}
		valueType = valueType.Elem()
	}

	// 根据类型返回对应的枚举值
	switch valueType.Kind() {
	case reflect.Bool:
		return c_enum.EBool
	case reflect.Int8:
		return c_enum.EInt8
	case reflect.Uint8:
		return c_enum.EUint8
	case reflect.Int16:
		return c_enum.EInt16
	case reflect.Uint16:
		return c_enum.EUint16
	case reflect.Int32:
		return c_enum.EInt32
	case reflect.Uint32:
		return c_enum.EUint32
	case reflect.Int64:
		return c_enum.EInt64
	case reflect.Uint64:
		return c_enum.EUint64
	case reflect.Float32:
		return c_enum.EFloat32
	case reflect.Float64:
		return c_enum.EFloat64
	case reflect.String:
		return c_enum.EString
	case reflect.Int:
		// int 类型根据系统架构可能是 32 位或 64 位
		// 这里统一返回 Int64，如果需要更精确的类型判断，可以根据实际值范围判断
		return c_enum.EInt64
	case reflect.Uint:
		// uint 类型根据系统架构可能是 32 位或 64 位
		// 这里统一返回 Uint64
		return c_enum.EUint64
	default:
		// 对于其他类型（如结构体、切片、映射等），默认返回字符串类型
		return c_enum.EString
	}
}

// getTelemetry 反射获取遥测信息 用于实现IDriver接口
func (s *SDriverInfo) getTelemetry(key string, instance IDevice) (any, error) {
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

	cacheKey := instance.GetConfig().Id + "_" + key
	// 先尝试读锁获取缓存
	s.reflectMethodMutex.RLock()
	if method, ok = s.reflectMethodCache[cacheKey]; ok {
		s.reflectMethodMutex.RUnlock()
	} else {
		s.reflectMethodMutex.RUnlock()

		// 缓存不存在，需要写入，使用写锁
		s.reflectMethodMutex.Lock()
		// 双重检查，防止其他goroutine已经写入
		if method, ok = s.reflectMethodCache[cacheKey]; !ok {
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
			s.reflectMethodCache[cacheKey] = method
		}
		s.reflectMethodMutex.Unlock()
	}

	// 使用defer recover处理可能的panic
	var result any
	var callErr error
	func() {
		defer func() {
			if r := recover(); r != nil {
				callErr = errors.Errorf("getTelemetry panic! key: %s, error: %+v", key, r)
				//c_log.Errorf(context.Background(), "getTelemetry panic! key: %s, error: %+v", key, r)
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
			// 检查返回值是否为 nil 指针（仅对指针类型检查）
			if values[0].Kind() == reflect.Ptr && values[0].IsNil() {
				result = nil
			} else {
				result = values[0].Interface()
			}
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
		// 检查返回值是否为 nil 指针（仅对指针类型检查）
		if values[0].Kind() == reflect.Ptr && values[0].IsNil() {
			result = nil
		} else {
			result = values[0].Interface()
		}
	}()

	if callErr != nil {
		return nil, callErr
	}

	return result, nil
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func ExecuteCustomService(functionName string, instance IDevice, params []any) error {

	// 执行自定义方法
	if instance == nil {
		c_log.Errorf(context.Background(), "ExecuteCustomService [%s] instance is nil", functionName)
		return errors.Errorf("custom service instance is nil")
	}

	ctx := context.WithValue(context.Background(), ConstCtxKeyDeviceId, instance.GetConfig().Id)

	s := instance.GetConfig().DriverInfo
	if s == nil {
		c_log.Errorf(ctx, "ExecuteCustomService [%s] instance is nil", functionName)
		return errors.Errorf("custom service %s not support", functionName)
	}

	// 判断一下是否允许这个方法调用
	var service *SDriverService
	for _, v := range s.Service {
		if v.Key == functionName {
			service = v
			break
		}
	}
	if service == nil {
		c_log.BizInfof(ctx, "执行自定义服务失败！原因：[%s]服务未定义！", functionName)
		return errors.Errorf("custom service %s not support", functionName)
	}

	// 反射前先判断缓存中是否存在
	if s.reflectMethodCache == nil {
		s.reflectMethodCache = make(map[string]reflect.Value)
	}

	var (
		method reflect.Value
		ok     bool
	)

	// 先尝试读锁获取缓存
	s.reflectMethodMutex.RLock()
	if method, ok = s.reflectMethodCache[functionName]; ok {
		s.reflectMethodMutex.RUnlock()
	} else {
		s.reflectMethodMutex.RUnlock()

		// 缓存不存在，需要写入，使用写锁
		s.reflectMethodMutex.Lock()
		// 双重检查，防止其他goroutine已经写入
		if method, ok = s.reflectMethodCache[functionName]; !ok {
			method = reflect.ValueOf(instance).MethodByName(functionName)
			if !method.IsValid() {
				s.reflectMethodMutex.Unlock()
				c_log.BizInfof(ctx, "执行自定义服务[%s]失败！原因该服务对应的方法[%s]不存在！", service.Key, service.Name)
				return errors.Errorf("service %s not found", functionName)
			}
			s.reflectMethodCache[functionName] = method
		}
		s.reflectMethodMutex.Unlock()
	}

	// 准备反射调用的参数
	var callArgs []reflect.Value
	if len(params) > 0 {
		// 获取方法类型信息以进行参数类型转换
		methodType := method.Type()
		callArgs = make([]reflect.Value, len(params))

		for i, param := range params {
			// 获取目标参数类型
			var targetType reflect.Type
			if i < methodType.NumIn() {
				targetType = methodType.In(i)
			}

			// 进行类型转换
			convertedParam, err := convertParameterType(param, targetType)
			if err != nil {
				c_log.BizInfof(ctx, "执行自定义服务[%s]失败！参数[%d]类型转换失败：%v", service.Key, i, err)
				return errors.Errorf("parameter %d type conversion failed: %v", i, err)
			}

			callArgs[i] = reflect.ValueOf(convertedParam)
		}
	}

	// 执行方法调用
	values := method.Call(callArgs)
	if len(values) == 1 {
		if err, ok := values[0].Interface().(error); ok {
			c_log.BizInfof(ctx, "执行自定义服务[%s]失败！原因：%s", service.Key, err.Error())
			return err
		}
		c_log.BizInfof(ctx, "执行自定义服务[%s]成功！", service.Key)
		return nil
	}

	//c_log.BizInfof(ctx, "执行自定义服务[%s]成功！", service.Key)

	fmt.Printf("当前函数: %s 返回的参数数据不为1 ！返回的内容为: %v", functionName, values)

	return nil
}

// convertParameterType 将参数转换为目标类型
func convertParameterType(param interface{}, targetType reflect.Type) (interface{}, error) {
	if param == nil {
		return nil, nil
	}

	// 如果目标类型为空，直接返回原参数
	if targetType == nil {
		return param, nil
	}

	// 处理 json.Number 类型
	if jsonNum, ok := param.(json.Number); ok {
		switch targetType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, err := jsonNum.Int64()
			if err != nil {
				return nil, errors.Wrapf(err, "无法将 json.Number 转换为 %s", targetType.String())
			}
			return convertIntValue(val, targetType)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := jsonNum.Int64()
			if err != nil {
				return nil, errors.Wrapf(err, "无法将 json.Number 转换为 %s", targetType.String())
			}
			if val < 0 {
				return nil, errors.Errorf("无法将负数 %d 转换为无符号类型 %s", val, targetType.String())
			}
			return convertUintValue(uint64(val), targetType)
		case reflect.Float32, reflect.Float64:
			val, err := jsonNum.Float64()
			if err != nil {
				return nil, errors.Wrapf(err, "无法将 json.Number 转换为 %s", targetType.String())
			}
			return convertFloatValue(val, targetType)
		case reflect.String:
			return string(jsonNum), nil
		}
	}

	// 处理字符串转换
	if str, ok := param.(string); ok {
		switch targetType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "无法将字符串 '%s' 转换为 %s", str, targetType.String())
			}
			return convertIntValue(val, targetType)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "无法将字符串 '%s' 转换为 %s", str, targetType.String())
			}
			return convertUintValue(val, targetType)
		case reflect.Float32, reflect.Float64:
			val, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "无法将字符串 '%s' 转换为 %s", str, targetType.String())
			}
			return convertFloatValue(val, targetType)
		case reflect.Bool:
			val, err := strconv.ParseBool(str)
			if err != nil {
				return nil, errors.Wrapf(err, "无法将字符串 '%s' 转换为 bool", str)
			}
			return val, nil
		case reflect.String:
			return str, nil
		}
	}

	// 处理布尔值
	if boolVal, ok := param.(bool); ok && targetType.Kind() == reflect.Bool {
		return boolVal, nil
	}

	// 处理数值类型的直接转换
	paramValue := reflect.ValueOf(param)
	if paramValue.Type().ConvertibleTo(targetType) {
		return paramValue.Convert(targetType).Interface(), nil
	}

	// 如果类型已经匹配，直接返回
	if paramValue.Type() == targetType {
		return param, nil
	}

	return nil, errors.Errorf("无法将类型 %s 转换为 %s", paramValue.Type().String(), targetType.String())
}

// convertIntValue 将 int64 值转换为指定的整数类型
func convertIntValue(val int64, targetType reflect.Type) (interface{}, error) {
	switch targetType.Kind() {
	case reflect.Int:
		return int(val), nil
	case reflect.Int8:
		if val < -128 || val > 127 {
			return nil, errors.Errorf("值 %d 超出 int8 范围", val)
		}
		return int8(val), nil
	case reflect.Int16:
		if val < -32768 || val > 32767 {
			return nil, errors.Errorf("值 %d 超出 int16 范围", val)
		}
		return int16(val), nil
	case reflect.Int32:
		if val < -2147483648 || val > 2147483647 {
			return nil, errors.Errorf("值 %d 超出 int32 范围", val)
		}
		return int32(val), nil
	case reflect.Int64:
		return val, nil
	default:
		return nil, errors.Errorf("不支持的整数类型: %s", targetType.String())
	}
}

// convertUintValue 将 uint64 值转换为指定的无符号整数类型
func convertUintValue(val uint64, targetType reflect.Type) (interface{}, error) {
	switch targetType.Kind() {
	case reflect.Uint:
		return uint(val), nil
	case reflect.Uint8:
		if val > 255 {
			return nil, errors.Errorf("值 %d 超出 uint8 范围", val)
		}
		return uint8(val), nil
	case reflect.Uint16:
		if val > 65535 {
			return nil, errors.Errorf("值 %d 超出 uint16 范围", val)
		}
		return uint16(val), nil
	case reflect.Uint32:
		if val > 4294967295 {
			return nil, errors.Errorf("值 %d 超出 uint32 范围", val)
		}
		return uint32(val), nil
	case reflect.Uint64:
		return val, nil
	default:
		return nil, errors.Errorf("不支持的无符号整数类型: %s", targetType.String())
	}
}

// convertFloatValue 将 float64 值转换为指定的浮点数类型
func convertFloatValue(val float64, targetType reflect.Type) (interface{}, error) {
	switch targetType.Kind() {
	case reflect.Float32:
		return float32(val), nil
	case reflect.Float64:
		return val, nil
	default:
		return nil, errors.Errorf("不支持的浮点数类型: %s", targetType.String())
	}
}
