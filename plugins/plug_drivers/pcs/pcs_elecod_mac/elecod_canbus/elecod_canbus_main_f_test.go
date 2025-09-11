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
	source := DeviceTypeScreen
	sourceAddress := ScreenAddress
	target := DeviceTypeMAC
	targetAddress := MacAddress

	cmd := &SCANFrameInfo{
		TargetDeviceType: target,
		TargetDeviceAddr: targetAddress,
		SourceDeviceType: source,
		SourceDeviceAddr: sourceAddress,
		MessageType:      MessageTypeControl,
		ServiceCode:      0x01,
	}

	id := BuildCanbusID(cmd)
	fmt.Printf("待机: 0x%X\n", *id)
	PrintCANbusID(*id)

	cmd.ServiceCode = 0x02
	id = BuildCanbusID(cmd)
	fmt.Printf("启动: 0x%X\n", *id)
	PrintCANbusID(*id)

	cmd.ServiceCode = 0x03
	id = BuildCanbusID(cmd)
	fmt.Printf("停止: 0x%X\n", *id)
	PrintCANbusID(*id)
}

func TestSync(t *testing.T) {
	// 同步帧
	cmd := &SCANFrameInfo{
		TargetDeviceType: DeviceTypeBroadcast,
		TargetDeviceAddr: 0b1111,
		SourceDeviceType: DeviceTypeScreen,
		SourceDeviceAddr: ScreenAddress,
		MessageType:      MessageTypeStatus,
		ServiceCode:      0x01,
	}

	id := BuildCanbusID(cmd)
	fmt.Printf("待机: 0x%X\n", *id)
	PrintCANbusID(*id)
}

func TestPrintCanbusID(t *testing.T) {
	// 测试用例
	testIDs := []uint32{
		//0x1C50A808, // 模拟量1
		0x1C512810,
		0x1c482808,
		0x1c482810,
		0x1c482818,
		0x1c482820,
		0x1c482828,
		0x1c482830,
		0x1c482838,
		0x1c482840,
		0x1c482848,
		0x1c482850,
		//
		//0x0BF88808, // 配置电池
		//0x08789010, // 开机
		//0x08789018, // 关机
		0x4389008, //待机
		0x4389010, // 开机
		0x4389018, // 关机

	}
	// 目标地址（屏幕）
	// 111
	fmt.Printf("待机：0x%X\n", 0b00100001110001001000000001000)
	fmt.Printf("开机：0x%X\n", 0b00100001110001001000000010000)
	fmt.Printf("关机：0x%X\n", 0b00100001110001001000000011000)
	fmt.Println("CANbus 帧解析结果:")
	fmt.Println("================================================================================")
	for _, id := range testIDs {
		PrintCANbusID(id)
	}

}

func TestAA(t *testing.T) {
	m := make(map[string]any)
	m["selfAddress"] = 1
	m["macAddress"] = 0
	id := BuildMacToScreenCanbusID(MessageTypeAnalog, 0x02, m)
	fmt.Printf("ID: 0x%X\n", *id)

	if *id == 0x1C512810 {
		fmt.Println("EQ")
	}
}
