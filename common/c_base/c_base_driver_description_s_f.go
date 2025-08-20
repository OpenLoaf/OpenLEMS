package c_base

import (
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"

	"gopkg.in/yaml.v3"
)

// SDriverDescription 驱动的描述内容
type SDriverDescription struct {
	Brand         string            `json:"brand" yaml:"brand"`                 // 品牌
	Model         string            `json:"model" yaml:"model"`                 // 型号
	Version       string            `json:"version" yaml:"version"`             // 版本
	Create        string            `json:"create" yaml:"create"`               // 创建时间
	Image         string            `json:"image" yaml:"image"`                 // 图片
	BuildTime     string            `json:"buildTime" yaml:"buildTime"`         // 编译时间
	CommitHash    string            `json:"commitHash" yaml:"commitHash"`       // 提交哈希
	Author        string            `json:"author" yaml:"author"`               // 作者
	Remark        string            `json:"remark" yaml:"remark"`               // 备注
	ProtocolType  string            `json:"protocolType" yaml:"protocolType"`   // 协议类型
	Telemetry     []*STelemetry     `json:"telemetry" yaml:"telemetry"`         // 遥测
	CustomService []*SCustomService `json:"customService" yaml:"customService"` // 自定义服务

	reflectMethodCache map[string]reflect.Value // 反射方法缓存
}

func BuildDescriptionFromYaml(yamlData []byte) *SDriverDescription {
	description := &SDriverDescription{}
	err := yaml.Unmarshal(yamlData, description)
	if err != nil {
		panic(fmt.Errorf("解析版本信息失败！请检查version.yaml文件!%v", err))
	}

	return description
}

func (s *SDriverDescription) String() string {

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

// GetDriverDescription 获取描述信息 用于实现IDriver接口
func (s *SDriverDescription) GetDriverDescription() *SDriverDescription {
	return s
}
