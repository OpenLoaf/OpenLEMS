package internal

// SAutomationExecuteRule 自动化执行规则结构体
type SAutomationExecuteRule struct {
	DeviceId string   `json:"deviceId"` // 设备ID
	Service  string   `json:"service"`  // 服务名称
	Params   []string `json:"params"`   // 参数列表
}

// SAutomationExecuteConfig 自动化执行配置结构体，包含多个执行规则
type SAutomationExecuteConfig struct {
	Rules []*SAutomationExecuteRule `json:"rules"` // 执行规则列表
}

// NewAutomationExecuteRule 创建新的执行规则
func NewAutomationExecuteRule(deviceId, service string, params []string) *SAutomationExecuteRule {
	return &SAutomationExecuteRule{
		DeviceId: deviceId,
		Service:  service,
		Params:   params,
	}
}

// NewAutomationExecuteConfig 创建新的执行配置
func NewAutomationExecuteConfig() *SAutomationExecuteConfig {
	return &SAutomationExecuteConfig{
		Rules: make([]*SAutomationExecuteRule, 0),
	}
}

// AddRule 添加执行规则
func (c *SAutomationExecuteConfig) AddRule(rule *SAutomationExecuteRule) {
	c.Rules = append(c.Rules, rule)
}

// GetRules 获取所有执行规则
func (c *SAutomationExecuteConfig) GetRules() []*SAutomationExecuteRule {
	return c.Rules
}

// GetRuleCount 获取执行规则数量
func (c *SAutomationExecuteConfig) GetRuleCount() int {
	return len(c.Rules)
}

// ClearRules 清空所有执行规则
func (c *SAutomationExecuteConfig) ClearRules() {
	c.Rules = make([]*SAutomationExecuteRule, 0)
}

// GetDeviceId 获取设备ID
func (r *SAutomationExecuteRule) GetDeviceId() string {
	return r.DeviceId
}

// SetDeviceId 设置设备ID
func (r *SAutomationExecuteRule) SetDeviceId(deviceId string) {
	r.DeviceId = deviceId
}

// GetService 获取服务名称
func (r *SAutomationExecuteRule) GetService() string {
	return r.Service
}

// SetService 设置服务名称
func (r *SAutomationExecuteRule) SetService(service string) {
	r.Service = service
}

// GetParams 获取参数列表
func (r *SAutomationExecuteRule) GetParams() []string {
	return r.Params
}

// SetParams 设置参数列表
func (r *SAutomationExecuteRule) SetParams(params []string) {
	r.Params = params
}

// AddParam 添加参数
func (r *SAutomationExecuteRule) AddParam(param string) {
	r.Params = append(r.Params, param)
}

// GetParamCount 获取参数数量
func (r *SAutomationExecuteRule) GetParamCount() int {
	return len(r.Params)
}
