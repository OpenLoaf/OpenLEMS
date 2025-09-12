package c_base

import "common/c_enum"

type STelemetry struct {
	Key          string            `json:"key,omitempty" yaml:"key"`                   // 遥测名称
	Name         string            `json:"name" yaml:"name"`                           // 遥测名称的国际化覆盖
	Precise      uint8             `json:"precise" yaml:"precise"`                     // 浮点数精度，默认0
	Unit         string            `json:"unit,omitempty" yaml:"unit"`                 // 单位
	Desc         string            `json:"desc,omitempty" yaml:"desc"`                 // 备注
	ValueExplain map[string]string `json:"valueExplain,omitempty" yaml:"valueExplain"` // 值解释
}

type sTelemetryPoint struct {
	*SPoint
	dataAccess *SDataAccess
}

func (s *sTelemetryPoint) GetDataAccess() *SDataAccess {
	return s.dataAccess
}

func (s *STelemetry) ToPoint(valueType c_enum.EValueType) IPoint {
	if s.ValueExplain != nil {
		valueType = c_enum.EString
	}

	dataAccess := &SDataAccess{
		ValueType: valueType,
	}
	return &sTelemetryPoint{
		dataAccess: dataAccess,
		SPoint: &SPoint{
			Key:  s.Key,
			Name: s.Name,
			Group: &SPointGroup{
				GroupKey:  "Total",
				GroupName: "汇总",
				GroupSort: -1,
			},
			Precise: s.Precise,
			Desc:    s.Desc,
			Unit:    s.Unit,
		},
	}
}

type SDriverService struct {
	Name        string `json:"name" yaml:"name"  dc:"服务名称，支持i18n"`
	DisplayName string `json:"displayName" yaml:"displayName" dc:"执行的方法名"`
	Description string `json:"description" yaml:"description" dc:"备注"`
	//Service     func(any) error `json:"-" dc:"执行方法"`
}
