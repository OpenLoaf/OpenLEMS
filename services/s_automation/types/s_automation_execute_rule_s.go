package types

// SAutomationExecuteRule 自动化执行规则结构体
type SAutomationExecuteRule struct {
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
	Service  string `json:"service" v:"required" dc:"服务名称"`
	Params   []any  `json:"params" dc:"参数列表"`
}

// SAutomationExecuteConfig 自动化执行配置结构体，包含多个执行规则
type SAutomationExecuteConfig struct {
	Rules []*SAutomationExecuteRule `json:"rules" v:"required|min-length:1" dc:"执行规则列表"`
}

func NewAutomationExecuteRule(deviceId, service string, params []any) *SAutomationExecuteRule {
	return &SAutomationExecuteRule{DeviceId: deviceId, Service: service, Params: params}
}

func NewAutomationExecuteConfig() *SAutomationExecuteConfig {
	return &SAutomationExecuteConfig{Rules: make([]*SAutomationExecuteRule, 0)}
}

func (c *SAutomationExecuteConfig) AddRule(rule *SAutomationExecuteRule) {
	c.Rules = append(c.Rules, rule)
}
func (c *SAutomationExecuteConfig) GetRules() []*SAutomationExecuteRule { return c.Rules }
func (c *SAutomationExecuteConfig) GetRuleCount() int                   { return len(c.Rules) }
func (c *SAutomationExecuteConfig) ClearRules()                         { c.Rules = make([]*SAutomationExecuteRule, 0) }

func (r *SAutomationExecuteRule) GetDeviceId() string         { return r.DeviceId }
func (r *SAutomationExecuteRule) SetDeviceId(deviceId string) { r.DeviceId = deviceId }
func (r *SAutomationExecuteRule) GetService() string          { return r.Service }
func (r *SAutomationExecuteRule) SetService(service string)   { r.Service = service }
func (r *SAutomationExecuteRule) GetParams() []any            { return r.Params }
func (r *SAutomationExecuteRule) SetParams(params []any)      { r.Params = params }
func (r *SAutomationExecuteRule) AddParam(param string)       { r.Params = append(r.Params, param) }
func (r *SAutomationExecuteRule) GetParamCount() int          { return len(r.Params) }
