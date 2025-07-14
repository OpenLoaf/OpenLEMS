package c_base

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

type SDescription struct {
	Brand        string        `json:"brand" yaml:"brand"`               // 品牌
	Model        string        `json:"model" yaml:"model"`               // 型号
	Version      string        `json:"version" yaml:"version"`           // 版本
	Create       string        `json:"create" yaml:"create"`             // 创建时间
	BuildTime    string        `json:"buildTime" yaml:"buildTime"`       // 编译时间
	CommitHash   string        `json:"commitHash" yaml:"commitHash"`     // 提交哈希
	Author       string        `json:"author" yaml:"author"`             // 作者
	Remark       string        `json:"remark" yaml:"remark"`             // 备注
	ProtocolType string        `json:"protocolType" yaml:"protocolType"` // 协议类型
	Telemetry    []*STelemetry `json:"telemetry" yaml:"telemetry"`       // 遥测

	reflectMethodCache map[string]reflect.Value // 反射方法缓存
}

func BuildDescriptionFromYaml(ctx context.Context, yaml []byte) *SDescription {
	// g.Log().Debugf(ctx, "BuildDescriptionFromYaml: %s", string(yaml))
	description := &SDescription{}
	err := gyaml.DecodeTo(yaml, &description)
	if err != nil {
		panic(gerror.Newf("解析版本信息失败！请检查version.yaml文件!%v", err))
	}

	return description
}

func (s *SDescription) String() string {

	// 创建一个 strings.Builder 来构建表格内容
	var builder strings.Builder

	// 创建一个新的 tabwriter，写入 strings.Builder
	writer := tabwriter.NewWriter(&builder, 0, 0, 3, ' ', 0)

	_, _ = writer.Write([]byte(fmt.Sprintf("Brand\t:\t%s\t\n", s.Brand)))
	_, _ = writer.Write([]byte(fmt.Sprintf("Model\t:\t%s\t\n", s.Model)))
	_, _ = writer.Write([]byte(fmt.Sprintf("Version\t:\t%s\t\n", s.Version)))
	_, _ = writer.Write([]byte(fmt.Sprintf("Author\t:\t%s\t\n", s.Author)))
	_, _ = writer.Write([]byte(fmt.Sprintf("CreateTime\t:\t%s\t\n", s.Create)))
	_, _ = writer.Write([]byte(fmt.Sprintf("ProtocolType\t:\t%s\t\n", s.ProtocolType)))

	if s.BuildTime != "" {
		_, _ = writer.Write([]byte(fmt.Sprintf("BuildTime\t:\t%s\t\n", s.BuildTime)))
	}
	if s.CommitHash != "" {
		_, _ = writer.Write([]byte(fmt.Sprintf("CommitHash\t:\t%s\t\n", s.CommitHash)))
	}
	_, _ = writer.Write([]byte(fmt.Sprintf("Remark\t:\t%s\t\n", s.Remark)))

	if len(s.Telemetry) != 0 {
		_, _ = writer.Write([]byte("\nTelemetry Information:\t\n"))
		_, _ = writer.Write([]byte("Name\tNationalization\tUnit\tRemark\t"))

		for _, telemetry := range s.Telemetry {
			_, _ = writer.Write([]byte("\n" + telemetry.String()))
		}

	}

	_ = writer.Flush()
	return builder.String()
}

// GetDescription 获取描述信息 用于实现IDriver接口
func (s *SDescription) GetDescription() *SDescription {
	return s
}

// GetTelemetry 反射获取遥测信息 用于实现IDriver接口
func (s *SDescription) GetTelemetry(key string, instance any) (any, error) {

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
		functionName := fmt.Sprintf("Get%s", gstr.UcFirst(key))
		method = reflect.ValueOf(instance).MethodByName(functionName)
		if !method.IsValid() {
			return 0, gerror.Newf("method %s not found", key)
		}
		s.reflectMethodCache[key] = method
	}

	// 空参数调用
	value := method.Call(nil)
	if len(value) == 1 {
		return value[0].Interface(), nil
	}

	if len(value) != 2 {
		return 0, gerror.Newf("function %s return value length is not 2", key)
	}
	if value[1].Interface() != nil {
		return 0, value[1].Interface().(error)
	}
	return value[0].Interface(), nil
}

func (s *SDescription) GetAllTelemetry(instance any) map[string]any {
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
