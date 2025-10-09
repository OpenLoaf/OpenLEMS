package c_base

import (
	"common/c_enum"
	"common/c_log"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

func GetAllTelemetry(instance IDevice) map[string]any {
	if instance == nil || instance.GetProtocolStatus() != c_enum.EProtocolConnected {
		return nil
	}

	result := make(map[string]any)

	// 获取所有点位数据
	allPoints := instance.GetPoints()
	pointValueList := instance.GetPointValueList()

	// 创建点位值映射，便于查找
	pointValueMap := make(map[string]*SPointValue)
	for _, pv := range pointValueList {
		pointValueMap[pv.GetKey()] = pv
	}

	for _, point := range allPoints {
		// 检查是否是SReflectPoint类型（遥测点位）
		if reflectPoint, ok := point.(*SReflectPoint); ok {
			// 通过反射获取遥测数据
			value, err := getTelemetryByReflect(reflectPoint.MethodName, instance)
			if err != nil {
				c_log.Warningf(context.Background(), "=>获取遥测数据失败: 点位=%s, 方法=%s, 错误=%v",
					point.GetKey(), reflectPoint.MethodName, err)
				continue
			}
			result[point.GetKey()] = value
		} else {
			// 协议点位，从缓存中获取数据
			if pv, exists := pointValueMap[point.GetKey()]; exists && pv != nil {
				result[point.GetKey()] = pv.GetValue()
			}
		}
	}

	return result
}

func GetAllTelemetryPoint(instance IDevice) []*SPointValue {
	if instance == nil {
		return nil
	}

	var result []*SPointValue

	// 获取所有点位数据
	allPoints := instance.GetPoints()
	pointValueList := instance.GetPointValueList()

	// 创建点位值映射，便于查找
	pointValueMap := make(map[string]*SPointValue)
	for _, pv := range pointValueList {
		pointValueMap[pv.GetKey()] = pv
	}

	for _, point := range allPoints {
		// 检查是否是SReflectPoint类型（遥测点位）
		if reflectPoint, ok := point.(*SReflectPoint); ok {
			// 通过反射获取遥测数据
			value, err := getTelemetryByReflect(reflectPoint.MethodName, instance)
			if err != nil {
				c_log.Warningf(context.Background(), "获取遥测数据失败: 点位=%s, 方法=%s, 错误=%v",
					point.GetKey(), reflectPoint.MethodName, err)
				continue
			}

			// 创建SPointValue
			pointValue := &SPointValue{
				IPoint: point,
				value:  value,
			}
			result = append(result, pointValue)
		} else {
			// 协议点位，从缓存中获取数据
			if pv, exists := pointValueMap[point.GetKey()]; exists && pv != nil {
				result = append(result, pv)
			}
		}
	}

	return result
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
		default:
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
		default:
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

// getTelemetryByReflect 通过反射获取遥测数据
func getTelemetryByReflect(methodName string, instance IDevice) (any, error) {
	if instance == nil || methodName == "" {
		return nil, errors.New("instance or methodName is nil")
	}

	// 获取实例的反射值
	instanceValue := reflect.ValueOf(instance)
	//if instanceValue.Kind() == reflect.Ptr {
	//	instanceValue = instanceValue.Elem()
	//}

	// 查找方法
	method := instanceValue.MethodByName(methodName)
	if !method.IsValid() {
		return nil, errors.Errorf("method %s not found", methodName)
	}

	// 调用方法
	results := method.Call(nil)
	if len(results) == 0 {
		return nil, errors.Errorf("method %s returned no values", methodName)
	}

	// 返回第一个结果
	return results[0].Interface(), nil
}
