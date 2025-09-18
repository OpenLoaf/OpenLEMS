package c_base

type SDriverService struct {
	Icon        string `yaml:"icon,omitempty" json:"icon" dc:"图标"`
	Key         string `json:"key" yaml:"key" dc:"执行的方法名"`
	Name        string `json:"name" yaml:"name"  dc:"服务名称，支持i18n"`
	Description string `json:"description" yaml:"description" dc:"备注"`
	//Service     func(any) error `json:"-" dc:"执行方法"`
	Params []*SConfigStructFields `json:"params" yaml:"params" dc:"参数定义"`
}
