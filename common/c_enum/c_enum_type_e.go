package c_enum

type EDeviceType string

const (
	EDeviceNone          EDeviceType = ""        // 未知设备
	EDeviceAmmeter       EDeviceType = "ammeter" // 电表用来代表某个实际的分组
	EDeviceCoolingAc     EDeviceType = "ca"      // 制冷空调
	EDeviceCoolingLiquid EDeviceType = "cl"      // 液冷机组
	EDeviceBms           EDeviceType = "bms"     // 电池管理系统
	EDeviceFire          EDeviceType = "fire"    // 消防
	EDeviceHumiture      EDeviceType = "hum"     // 温湿度
	EDevicePcs           EDeviceType = "pcs"     // 电池逆变器
	EDeviceLoad          EDeviceType = "load"    // 负载
	EDevicePv            EDeviceType = "pv"      // 光伏
	EDeviceEnergyStore   EDeviceType = "ess"     // 储能柜
	EDeviceChargePile    EDeviceType = "cp"      // 充电桩
	EDeviceGenerator     EDeviceType = "gen"     // 发电机
	EDeviceGpio          EDeviceType = "gpio"    // DIY

	//EEntrance            EDeviceType = "entrance"       // 总站入口

	EStationEnergyStore EDeviceType = "sess" // 总站储能柜
)
