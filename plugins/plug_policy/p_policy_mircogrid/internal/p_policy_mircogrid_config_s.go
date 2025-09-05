package internal

type SPolicyMircogridConfig struct {
	// 时间参数
	TimeStep     float64 `json:"time_step" name:"时间间隔" desc:"时间间隔（小时）" ct:"number" vt:"float" min:"0.1" max:"24" default:"1.0" unit:"小时"`
	HorizonHours int     `json:"horizon_hours" name:"预测时域" desc:"预测时域" ct:"number" vt:"int" min:"1" max:"168" default:"24" unit:"小时"`

	// 储能参数
	BatteryCapacity     float64 `json:"battery_capacity" name:"储能容量" desc:"储能容量" ct:"number" vt:"float" min:"0" default:"232.0" unit:"kWh"`
	SocMaxRatio         float64 `json:"soc_max_ratio" name:"最大SOC比例" desc:"最大SOC比例" ct:"number" vt:"float" min:"0" max:"1" default:"0.95"`
	SocMinRatio         float64 `json:"soc_min_ratio" name:"最小SOC比例" desc:"最小SOC比例" ct:"number" vt:"float" min:"0" max:"1" default:"0.1"`
	BatteryInitSoc      float64 `json:"battery_init_soc" name:"初始SOC比例" desc:"初始SOC比例" ct:"number" vt:"float" min:"0" max:"1" default:"0.2"`
	BatteryPowerRated   float64 `json:"battery_power_rated" name:"储能额定功率" desc:"储能额定功率" ct:"number" vt:"float" min:"0" default:"100.0" unit:"kW"`
	ChargeEfficiency    float64 `json:"charge_efficiency" name:"充电效率" desc:"充电效率" ct:"number" vt:"float" min:"0" max:"1" default:"0.95"`
	DischargeEfficiency float64 `json:"discharge_efficiency" name:"放电效率" desc:"放电效率" ct:"number" vt:"float" min:"0" max:"1" default:"0.95"`

	// 需量控制参数
	EnableDemandControl bool `json:"enable_demand_control" name:"启动需量控制模式" desc:"启动需量控制模式" ct:"switch" vt:"bool" default:"false"`

	// 充电桩参数
	AllowPileReduce bool `json:"allow_pile_reduce" name:"允许充电桩降功率" desc:"不允许充电桩降功率（无充电桩）" ct:"switch" vt:"bool" default:"false"`

	// 负荷参数
	AllowLoadReduce    bool    `json:"allow_load_reduce" name:"允许负荷降功率" desc:"允许负荷降功率" ct:"switch" vt:"bool" default:"false"`
	LoadReduceMaxPower float64 `json:"load_reduce_max_power" name:"负荷允许降低的最大功率" desc:"负荷允许降低的最大功率" ct:"number" vt:"float" min:"0" default:"60.0" unit:"kW"`
	LoadReducePenalty  float64 `json:"load_reduce_penalty" name:"负荷降功率惩罚系数" desc:"负荷降功率惩罚系数" ct:"number" vt:"float" min:"0" default:"800.0" unit:"元/kWh"`

	// 变压器参数
	TransformerCapacity     float64 `json:"transformer_capacity" name:"变压器容量" desc:"变压器容量" ct:"number" vt:"float" min:"0" default:"200.0" unit:"kVA"`
	TransformerSafetyFactor float64 `json:"transformer_safety_factor" name:"变压器安全系数" desc:"变压器安全系数" ct:"number" vt:"float" min:"0" max:"1" default:"0.8"`

	// 上网售电参数
	EnableGridSell          bool    `json:"enable_grid_sell" name:"上网售电" desc:"上网售电" ct:"switch" vt:"bool" default:"true"`
	TransformerMinUseFactor float64 `json:"transformer_min_use_factor" name:"上网售电安全下限系数" desc:"上网售电安全下限系数" ct:"number" vt:"float" min:"0" max:"1" default:"0.1"`

	// 成本参数
	CurtailmentPenalty float64 `json:"curtailment_penalty" name:"弃光惩罚系数" desc:"弃光惩罚系数设为0（无光伏）" ct:"number" vt:"float" min:"0" default:"0.0" unit:"元/kWh"`
	PileReducePenalty  float64 `json:"pile_reduce_penalty" name:"充电桩降功率惩罚系数" desc:"充电桩降功率惩罚系数设为0（无充电桩）" ct:"number" vt:"float" min:"0" default:"0.0" unit:"元/kWh"`

	// 动态扩容参数
	EnableDynamicExpansion bool    `json:"enable_dynamic_expansion" name:"关闭动态扩容" desc:"关闭动态扩容" ct:"switch" vt:"bool" default:"false"`
	ExpansionCostFactor    float64 `json:"expansion_cost_factor" name:"扩容成本" desc:"扩容成本设为0（不使用）" ct:"number" vt:"float" min:"0" default:"0.0"`
	MaxExpansionRatio      float64 `json:"max_expansion_ratio" name:"最大扩容比例" desc:"最大扩容比例设为1（不扩容）" ct:"number" vt:"float" min:"1" default:"1.0"`
	LoadPriorityWeight     float64 `json:"load_priority_weight" name:"负荷优先级权重" desc:"负荷优先级权重" ct:"number" vt:"float" min:"0" default:"1000.0"`
	PilePriorityWeight     float64 `json:"pile_priority_weight" name:"充电桩优先级权重" desc:"充电桩优先级权重设为0（无充电桩）" ct:"number" vt:"float" min:"0" default:"0.0"`

	// 优化参数常量
	RegularizationWeight   float64 `json:"regularization_weight" name:"正则化权重" desc:"正则化权重" ct:"number" vt:"float" min:"0" default:"1.0e-06"`
	MutualExclusionPenalty float64 `json:"mutual_exclusion_penalty" name:"互斥操作惩罚系数" desc:"互斥操作惩罚系数" ct:"number" vt:"float" min:"0" default:"1000.0"`
}
