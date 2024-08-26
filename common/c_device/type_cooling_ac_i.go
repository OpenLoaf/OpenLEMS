package c_device

import "ems-plan/c_base"

type ICoolingAcBasic interface {
	ICoolingBasic
	GetCoolingAcStatus() (ECoolingStatus, error)
}

// ICoolingAc 空调
type ICoolingAc interface {
	c_base.IDriver
	ICoolingAcBasic
}
