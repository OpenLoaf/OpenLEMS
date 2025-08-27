package c_base

import (
	"common/c_log"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"reflect"
	"strings"
)

type SDriverInfo struct {
	Name         string            `json:"name" yaml:"name" v:"required|length:6,40"`     // 驱动名称
	Type         EDeviceType       `json:"type" yaml:"type" v:"required"`                 // 驱动类型
	ProtocolType EProtocolType     `json:"protocolType" yaml:"protocolType" v:"required"` // 协议类型
	Brand        string            `json:"brand" yaml:"brand"`                            // 品牌
	Model        string            `json:"model" yaml:"model"`                            // 型号
	Version      string            `json:"version" yaml:"version" v:"required"`           // 版本
	Telemetry    []*STelemetry     `json:"telemetry" yaml:"telemetry"`                    // 遥测
	Service      []*SDriverService `json:"service" yaml:"service"`                        // 自定义服务

	Enabled      bool   `json:"enabled" yaml:"enabled" v:"required"`
	Path         string `json:"path"`         // 路径
	HashCode     string `json:"hashCode"`     // 哈希码
	FileSizeByte int64  `json:"fileSizeByte"` // 文件大小

	Create     string `json:"create" yaml:"create"`         // 创建时间
	Image      string `json:"image" yaml:"image"`           // 图片
	BuildTime  string `json:"buildTime" yaml:"buildTime"`   // 编译时间
	CommitHash string `json:"commitHash" yaml:"commitHash"` // 提交哈希
	Author     string `json:"author" yaml:"author"`         // 作者
	Remark     string `json:"remark" yaml:"remark"`         // 备注

	reflectMethodCache map[string]reflect.Value // 反射方法缓存
}

func BuildDescriptionFromYaml(yamlData []byte) *SDriverInfo {
	description := &SDriverInfo{}
	err := yaml.Unmarshal(yamlData, description)
	if err != nil {
		panic(errors.Errorf("解析版本信息失败！请检查version.yaml文件!%+v", err))
	}

	return description
}

// GetTelemetry 反射获取遥测信息 用于实现IDriver接口
func (s *SDriverInfo) GetTelemetry(key string, instance any) (any, error) {

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
			return nil, errors.Errorf("method %s not found", key)
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
		return nil, errors.Errorf("function %s return value length is not 2", key)
	}
	if value[1].Interface() != nil {
		return nil, value[1].Interface().(error)
	}
	return value[0].Interface(), nil
}

func (s *SDriverInfo) GetAllTelemetry(instance any) map[string]any {
	telemetryMap := make(map[string]any, len(s.Telemetry))
	for _, telemetry := range s.Telemetry {
		value, err := s.GetTelemetry(telemetry.Name, instance)
		if err != nil {
			c_log.Errorf(context.Background(), "Get telemetry %s error: %+v", telemetry.Name, err)
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
