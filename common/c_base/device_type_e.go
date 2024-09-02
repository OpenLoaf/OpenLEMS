package c_base

type EDeviceType string

const (
	EDeviceNone          EDeviceType = ""               // 未知设备
	EDeviceAmmeter       EDeviceType = "ammeter"        // 电表用来代表某个实际的分组
	EDeviceCoolingAc     EDeviceType = "cooling_ac"     // 制冷空调
	EDeviceCoolingLiquid EDeviceType = "cooling_liquid" // 液冷机组
	EDeviceBms           EDeviceType = "bms"            // 电池管理系统
	EDeviceFire          EDeviceType = "fire"           // 消防
	EDeviceHumiture      EDeviceType = "humiture"       // 温湿度
	EDevicePcs           EDeviceType = "pcs"            // 电池逆变器
	EDeviceLoad          EDeviceType = "load"           // 负载
	EDevicePv            EDeviceType = "pv"             // 光伏
	EDeviceEnergyStore   EDeviceType = "energy-store"   // 储能柜
	EDeviceChargePile    EDeviceType = "charging-pile"  // 充电桩
	EDeviceGenerator     EDeviceType = "generator"      // 发电机
	EDeviceGpio          EDeviceType = "gpio"           // DIY

	//EEntrance            EDeviceType = "entrance"       // 总站入口

	EStationEnergyStore EDeviceType = "station-energy-store" // 总站储能柜
)
