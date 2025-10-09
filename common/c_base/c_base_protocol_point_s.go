package c_base

// SProtocolPoint 协议点位基础结构
type SProtocolPoint struct {
	*SPoint                 // 嵌套基础点位信息
	DataAccess *SDataAccess `json:"dataAccess" v:"required" dc:"数据访问配置"` // 数据访问配置
}

// GetDataAccess 获取数据访问配置
func (s *SProtocolPoint) GetDataAccess() *SDataAccess {
	return s.DataAccess
}

// 注意：不需要重复实现IPoint接口方法
// 通过结构体嵌套自动继承SPoint的方法实现
// SPoint字段将在启动时验证是否设置
