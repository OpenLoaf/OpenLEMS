package c_base

import (
	"reflect"
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

	configStructFields []*SConfigStructFields   // 不能从yaml中倒入
	reflectMethodCache map[string]reflect.Value // 反射方法缓存
}

func (s *SDriverInfo) SetConfigStructFields(configStructFields []*SConfigStructFields) {
	s.configStructFields = configStructFields
}
func (s *SDriverInfo) GetConfigStructFields() []*SConfigStructFields {
	return s.configStructFields
}
