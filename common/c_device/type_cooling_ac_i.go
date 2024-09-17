package c_device

import "common/c_base"

// ICoolingAc 空调
type ICoolingAc interface {
	c_base.IDriver
	GetCoolingAcStatus() (ECoolingStatus, error)
}
