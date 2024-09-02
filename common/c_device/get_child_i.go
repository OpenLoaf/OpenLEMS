package c_device

import "ems-plan/c_base"

type IGetChildren interface {
	GetChildren() []c_base.IDriver
}
