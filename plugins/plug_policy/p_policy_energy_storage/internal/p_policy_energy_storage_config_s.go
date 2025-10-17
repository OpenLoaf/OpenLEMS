package internal

// SPolicyEnergyStorageConfig 储能站策略配置
type SPolicyEnergyStorageConfig struct {
	// 储能设备列表
	EssDevices []string `json:"essDevices" name:"储能设备列表" desc:"储能设备ID列表" ct:"multiSelect"`

	// 策略模式
	StrategyMode string `json:"strategyMode" name:"策略模式" desc:"local:本地模式, cloud:云端模式" ct:"singleSelect" opts:"local:本地模式,cloud:云端模式" default:"local"`

	// 策略类型
	StrategyType string `json:"strategyType" name:"策略类型" desc:"策略类型标识" ct:"text" default:"policy_ess"`
}
