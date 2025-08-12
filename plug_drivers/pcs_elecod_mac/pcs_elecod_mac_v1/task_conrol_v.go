package pcs_elecod_mac_v1

import (
	"canbus/p_canbus"
)

var (
	sandBy = &p_canbus.SCanbusTask{
		Name:        "待机",
		GetCanbusID: func(params map[string]string) uint32 { return 0x438900D },
		IsRemote:    true,
		IsExtended:  true,
	}
)
