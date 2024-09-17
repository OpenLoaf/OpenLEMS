package c_device

import "common/c_base"

type IGetChildren interface {
	GetChildren() []c_base.IDriver
}
