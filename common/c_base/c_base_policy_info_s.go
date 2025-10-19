package c_base

// SPolicyInfo 策略信息结构体
type SPolicyInfo struct {
	PolicyId               string              `json:"policyId"`               // 策略ID（如 "policy_microgrid", "policy_ess"）
	PolicyName             string              `json:"policyName"`             // 策略名称
	Description            string              `json:"description"`            // 策略描述
	ConfigFieldDefinitions []*SFieldDefinition `json:"configFieldDefinitions"` // 配置字段定义
}

