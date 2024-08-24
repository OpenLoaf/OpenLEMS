package c_device

type EType string

const (
	EAmmeter     EType = "ammeter"       // 电表用来代表某个实际的分组
	ECooling     EType = "cooling_ac"    // 制冷制热
	EBms         EType = "bms"           // 电池管理系统
	EFire        EType = "fire"          // 消防
	EHumiture    EType = "humiture"      // 温湿度
	EPcs         EType = "pcs"           // 电池逆变器
	ELoad        EType = "load"          // 负载
	EPv          EType = "pv"            // 光伏
	EEnergyStore EType = "energy-store"  // 储能柜
	EChargePile  EType = "charging-pile" // 充电桩
	EGenerator   EType = "generator"     // 发电机
	EEntrance    EType = "entrance"      // 总站入口
)
