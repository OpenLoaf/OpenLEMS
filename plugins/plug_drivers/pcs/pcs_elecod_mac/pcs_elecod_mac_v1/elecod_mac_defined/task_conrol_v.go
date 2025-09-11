package elecod_mac_defined

import (
	"canbus/p_canbus"

	"github.com/shockerli/cvt"
	"pcs_elecod/elecod_canbus"
)

var (
	CmdStandby = &p_canbus.SCanbusTask{
		Name: "待机",
		GetCanbusID: func(params map[string]any) *uint32 {
			macAddress, err := cvt.Uint32E(params["macAddress"])
			if err != nil {
				return nil
			}
			return buildControlCanbusId(elecod_canbus.DeviceTypeMAC, macAddress, elecod_canbus.MessageTypeControl, 0x01, params)
		},
		IsRemote:   true,
		IsExtended: true,
	}

	CmdStart = &p_canbus.SCanbusTask{
		Name: "开机",
		GetCanbusID: func(params map[string]any) *uint32 {
			macAddress, err := cvt.Uint32E(params["macAddress"])
			if err != nil {
				return nil
			}
			return buildControlCanbusId(elecod_canbus.DeviceTypeMAC, macAddress, elecod_canbus.MessageTypeControl, 0x02, params)
		},
		IsRemote:   true,
		IsExtended: true,
	}

	CmdShutdown = &p_canbus.SCanbusTask{
		Name: "关机",
		GetCanbusID: func(params map[string]any) *uint32 {
			macAddress, err := cvt.Uint32E(params["macAddress"])
			if err != nil {
				return nil
			}
			return buildControlCanbusId(elecod_canbus.DeviceTypeMAC, macAddress, elecod_canbus.MessageTypeControl, 0x03, params)
		},
		IsRemote:   true,
		IsExtended: true,
	}

	CmdHealth = &p_canbus.SCanbusTask{
		Name: "心跳",
		GetCanbusID: func(params map[string]any) *uint32 {
			return buildControlCanbusId(elecod_canbus.DeviceTypeBroadcast, 0b1111, elecod_canbus.MessageTypeStatus, 0x01, params)
		},
		IsRemote:   true,
		IsExtended: true,
	}
)

func buildControlCanbusId(targetType elecod_canbus.DeviceType, targetAddr uint32, messageType elecod_canbus.MessageType, serviceCode uint32, params map[string]any) *uint32 {
	selfAddress, err := cvt.Uint32E(params["selfAddress"])
	if err != nil {
		return nil
	}

	info := &elecod_canbus.SCANFrameInfo{
		TargetDeviceType: targetType,
		TargetDeviceAddr: targetAddr,
		SourceDeviceType: elecod_canbus.DeviceTypeScreen,
		SourceDeviceAddr: selfAddress,
		MessageType:      messageType,
		ServiceCode:      serviceCode,
	}
	return elecod_canbus.BuildCanbusID(info)

}
