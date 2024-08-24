package c_device

type ICoolingAcBasic interface {
	ICoolingBasic
	GetCoolingAcStatus() (ECoolingStatus, error)
}

// IAcCooling 空调
type IAcCooling interface {
	IInfo
	ICoolingAcBasic
}
