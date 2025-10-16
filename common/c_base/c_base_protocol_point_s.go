package c_base

// no imports required

// SProtocolPoint 协议点位基础结构
type SProtocolPoint struct {
	*SPoint                       // 嵌套基础点位信息
	DataAccess   *SDataAccess     `json:"dataAccess" v:"required" dc:"数据访问配置"`        // 数据访问配置
	ValueExplain []*SFieldExplain `json:"valueExplain,omitempty" yaml:"valueExplain"` // 值解释
}

// GetDataAccess 获取数据访问配置
func (s *SProtocolPoint) GetDataAccess() *SDataAccess {
	return s.DataAccess
}

func (s *SProtocolPoint) GetValueExplain() []*SFieldExplain {
	if len(s.ValueExplain) == 0 {
		return s.SPoint.GetValueExplain()
	}
	return s.ValueExplain
}

// GetValueExplainByValue 获取值解释，优先使用自身的 ValueExplain，其次回退到嵌入的 SPoint 逻辑
func (s *SProtocolPoint) GetValueExplainByValue(value any) (string, error) {
	// 统一用 SPoint 的公共实现，优先自身 explains（GetValueExplain 已做回退）
	return s.SPoint.explainByValueCommon(value, s.GetValueExplain(), s.SPoint.Precise)
}

// AsProtocolPoint 转换为协议点位，SProtocolPoint 本身就是协议点位，返回自身
func (s *SProtocolPoint) AsProtocolPoint() *SProtocolPoint {
	return s
}
