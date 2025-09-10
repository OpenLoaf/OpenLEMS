package c_type

import (
	"common/c_base"
	"common/c_enum"
)

// ICoolingAc 空调
type ICoolingAc interface {
	c_base.IDriver
	GetCoolingAcStatus() (c_enum.ECoolingStatus, error)
}
