package c_base

import "common/c_enum"

type STelemetry struct {
	Key          string               `json:"key,omitempty" yaml:"key"`                   // 遥测名称
	Name         string               `json:"name" yaml:"name"`                           // 遥测名称的国际化覆盖
	Precise      uint8                `json:"precise" yaml:"precise"`                     // 浮点数精度，默认0
	Unit         string               `json:"unit,omitempty" yaml:"unit"`                 // 单位
	Desc         string               `json:"desc,omitempty" yaml:"desc"`                 // 备注
	ValueExplain map[string]string    `json:"valueExplain,omitempty" yaml:"valueExplain"` // 值解释
	ParamExplain map[string]string    `json:"paramExplain,omitempty" yaml:"paramExplain"` // 从参数值中读取解释
	Config       *SConfigStructFields `json:"config,omitempty" yaml:"config"`             // 配置字段
}

func (s *STelemetry) ToPoint(valueType c_enum.EValueType, params map[string]any) IPoint {
	valueExplain := make(map[string]string)
	if s.ValueExplain != nil {
		valueType = c_enum.EString
		valueExplain = s.ValueExplain
	}
	if s.ParamExplain != nil {
		valueType = c_enum.EString
		for key, v := range s.ParamExplain {
			if pv, ok := params[v]; ok && pv != nil {
				valueExplain[key] = pv.(string)
			}
		}
	}

	return &SPoint{
		Key:          s.Key,
		Name:         s.Name,
		Group:        GroupTotal,
		Precise:      s.Precise,
		Desc:         s.Desc,
		Unit:         s.Unit,
		ValueType:    valueType,
		ValueExplain: valueExplain,
	}
}

type SDriverService struct {
	Key         string `json:"key" yaml:"key" dc:"执行的方法名"`
	Name        string `json:"name" yaml:"name"  dc:"服务名称，支持i18n"`
	Description string `json:"description" yaml:"description" dc:"备注"`
	//Service     func(any) error `json:"-" dc:"执行方法"`
	Params []*SConfigStructFields `json:"params" yaml:"params" dc:"参数定义"`
}
