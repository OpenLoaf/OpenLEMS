package c_base

import (
	"fmt"
)

type STelemetry struct {
	Name                string `json:"name,omitempty" yaml:"name"`                     // 遥测名称
	NationalizationName string `json:"nationalizationName" yaml:"nationalizationName"` // 遥测名称的国际化覆盖
	Unit                string `json:"unit,omitempty" yaml:"unit"`                     // 单位
	Remark              string `json:"remark,omitempty" yaml:"remark"`                 // 备注
}

func (s *STelemetry) String() string {
	i18nKey := "-"
	if s.NationalizationName != "" {
		i18nKey = s.NationalizationName
	}
	return fmt.Sprintf("%s\t%s\t%s\t%s\t", s.Name, i18nKey, s.Unit, s.Remark)
}
