package c_type

import "common/c_base"

// ICoolingAc 空调
type ICoolingAc interface {
	c_base.IDevice
	GetCoolingAcStatus() (ECoolingStatus, error)
}
