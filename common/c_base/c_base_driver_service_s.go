package c_base

type STelemetry struct {
	Name        string `json:"name,omitempty" yaml:"name"`     // 遥测名称
	DisplayName string `json:"displayName" yaml:"displayName"` // 遥测名称的国际化覆盖
	Unit        string `json:"unit,omitempty" yaml:"unit"`     // 单位
	Remark      string `json:"remark,omitempty" yaml:"remark"` // 备注
}

type SDriverService struct {
	Name        string `json:"name" yaml:"name"  dc:"服务名称，支持i18n"`
	DisplayName string `json:"displayName" yaml:"displayName" dc:"执行的方法名"`
	Description string `json:"description" yaml:"description" dc:"备注"`
	//Service     func(any) error `json:"-" dc:"执行方法"`
}
