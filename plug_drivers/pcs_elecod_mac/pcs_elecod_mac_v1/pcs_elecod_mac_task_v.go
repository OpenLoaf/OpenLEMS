package pcs_elecod_mac_v1

import (
	"canbus/p_canbus"
	"common/c_base"
)

var (
	AlarmAndFault = p_canbus.SCanbusTask{
		Name:     "AlarmAndFault",
		CanbusID: 0x1C000109,
		Metas:    []*c_base.Meta{},
		IDMatch: func(u uint32) bool {
			return true
		},
	}
)
