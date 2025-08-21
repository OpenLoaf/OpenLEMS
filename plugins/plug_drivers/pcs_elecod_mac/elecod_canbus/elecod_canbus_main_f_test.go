package elecod_canbus

import (
	"fmt"
	"testing"
)

const (
	MacAddress    uint32 = 0x00
	ScreenAddress uint32 = 0x01
)

// Test_buildCANbusID 生成CANbusID
func Test_buildCANbusID(t *testing.T) {

	standby := &SCANFrameInfo{
		TargetDeviceType: DeviceTypeMAC,
		TargetDeviceAddr: MacAddress,
		SourceDeviceType: DeviceTypeScreen,
		SourceDeviceAddr: ScreenAddress,
		MessageType:      MessageTypeControl,
		ServiceCode:      0x01,
	}

	fmt.Printf("待机: 0x%X\n", BuildCanbusID(standby))

	start := &SCANFrameInfo{
		TargetDeviceType: DeviceTypeMAC,
		TargetDeviceAddr: MacAddress,
		SourceDeviceType: DeviceTypeScreen,
		SourceDeviceAddr: ScreenAddress,
		MessageType:      MessageTypeControl,
		ServiceCode:      0x02,
	}
	fmt.Printf("启动: 0x%X\n", BuildCanbusID(start))

	stop := &SCANFrameInfo{
		TargetDeviceType: DeviceTypeMAC,
		TargetDeviceAddr: MacAddress,
		SourceDeviceType: DeviceTypeScreen,
		SourceDeviceAddr: ScreenAddress,
		MessageType:      MessageTypeControl,
		ServiceCode:      0x03,
	}
	fmt.Printf("停止: 0x%X\n", BuildCanbusID(stop))
}
