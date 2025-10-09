package c_base

import (
	"common/c_enum"
	"reflect"
	"sync"
)

type SDriverInfo struct {
	Name         string               `json:"name" yaml:"name" v:"required|length:6,40"`     // 驱动名称
	Type         c_enum.EDeviceType   `json:"type" yaml:"type" v:"required"`                 // 驱动类型
	ProtocolType c_enum.EProtocolType `json:"protocolType" yaml:"protocolType" v:"required"` // 协议类型
	Brand        string               `json:"brand" yaml:"brand"`                            // 品牌
	Model        string               `json:"model" yaml:"model"`                            // 型号
	Version      string               `json:"version" yaml:"version" v:"required"`           // 版本

	// 统一点位管理
	ConfigPoints []*SConfigPoint `json:"configPoints" yaml:"configPoints"`

	// 自定义服务
	Service []*SDriverService `json:"service" yaml:"service"`

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

	// 反射方法缓存
	reflectMethodCache map[string]reflect.Value // 反射方法缓存
	reflectMethodMutex sync.RWMutex             // 反射方法缓存读写锁
}

// 注意：configStructFields字段已移除，现在使用ConfigPoints字段
// 如果需要兼容旧代码，可以保留这些方法但返回空值
func (s *SDriverInfo) SetConfigStructFields(configStructFields []*SFieldDefinition) {
	// 已废弃：configStructFields字段已移除，现在使用ConfigPoints字段
}
func (s *SDriverInfo) GetConfigStructFields() []*SFieldDefinition {
	// 已废弃：configStructFields字段已移除，现在使用ConfigPoints字段
	return nil
}

// SetConfigPoints 设置配置点位列表
func (s *SDriverInfo) SetConfigPoints(configPoints []*SConfigPoint) {
	s.ConfigPoints = configPoints
}

// GetConfigPoints 获取配置点位列表
func (s *SDriverInfo) GetConfigPoints() []*SConfigPoint {
	return s.ConfigPoints
}

// AddConfigPoint 添加配置点位
func (s *SDriverInfo) AddConfigPoint(configPoint *SConfigPoint) {
	if configPoint != nil {
		s.ConfigPoints = append(s.ConfigPoints, configPoint)
	}
}

// RemoveConfigPoint 移除配置点位
func (s *SDriverInfo) RemoveConfigPoint(key string) {
	for i, point := range s.ConfigPoints {
		if point != nil && point.GetKey() == key {
			s.ConfigPoints = append(s.ConfigPoints[:i], s.ConfigPoints[i+1:]...)
			break
		}
	}
}

// GetConfigPointByKey 根据key获取配置点位
func (s *SDriverInfo) GetConfigPointByKey(key string) *SConfigPoint {
	for _, point := range s.ConfigPoints {
		if point != nil && point.GetKey() == key {
			return point
		}
	}
	return nil
}
