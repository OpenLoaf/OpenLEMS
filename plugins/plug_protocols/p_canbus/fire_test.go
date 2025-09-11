package p_canbus

import (
	"context"
	"fmt"
	"testing"

	"go.einride.tech/can/pkg/candevice"
	"go.einride.tech/can/pkg/socketcan"
)

func TestFireCan0(t *testing.T) {
	id := "can0"

	d, _ := candevice.New(id)
	_ = d.SetBitrate(250000)
	_ = d.SetUp()
	defer func() {
		d.SetDown()
		println("close!!")
	}()

	// Error handling omitted to keep example simple
	conn, err := socketcan.DialContext(context.Background(), "can", id)
	if err != nil {
		panic(err)
	}

	println("listening...")

	recv := socketcan.NewReceiver(conn)
	for recv.Receive() {
		frame := recv.Frame()

		fmt.Println(frame.String())

		if frame.ID == 0x1C000109 {
			// 确保数据长度足够，以避免运行时错误
			// can.Data 是一个 [8]byte 数组，所以 len(frame.Data) 总是 8
			// 但为了安全起见，检查一下是否有足够的字节来访问
			if len(frame.Data) >= 8 {
				// 解析数据
				// 直接访问字节数据，例如 Byte 1 是 frame.Data[0]
				detectorNumber := frame.Data[0] // 探测器编号 (Byte 1)

				// 使用 Data.Bit(globalBitIndex) 方法来获取单个位的值
				// 计算全局位索引：(ByteNumber - 1) * 8 + (BitInByteNumber - 1)

				// Byte 2 的位
				tempAlarm := frame.Data.Bit(8)    // 温度报警 (Byte 2, Bit 1 -> global bit 8)
				smokeAlarm := frame.Data.Bit(9)   // 烟雾报警 (Byte 2, Bit 2 -> global bit 9)
				coAlarm := frame.Data.Bit(10)     // CO 报警 (Byte 2, Bit 3 -> global bit 10)
				h2Alarm := frame.Data.Bit(11)     // H2 报警 (Byte 2, Bit 4 -> global bit 11)
				vocAlarm := frame.Data.Bit(12)    // VOC 报警 (Byte 2, Bit 5 -> global bit 12)
				level1Alarm := frame.Data.Bit(13) // 1级报警 (Byte 2, Bit 6 -> global bit 13)
				level2Alarm := frame.Data.Bit(14) // 2级报警 (Byte 2, Bit 7 -> global bit 14)

				// Byte 3 的位
				detectorFault := frame.Data.Bit(16)              // 探测器故障 (Byte 3, Bit 1 -> global bit 16)
				aerosolHardwareFault := frame.Data.Bit(17)       // 气溶胶硬件故障 (Byte 3, Bit 2 -> global bit 17)
				mainPowerUndervoltageFault := frame.Data.Bit(18) // 主电欠压故障 (Byte 3, Bit 3 -> global bit 18)

				// 报文编号 (Byte 8)
				messageNumber := frame.Data[7]

				fmt.Printf("       探测器编号: %d\n", detectorNumber)
				fmt.Printf("       温度报警: %t (true-报警, false-正常)\n", tempAlarm)
				fmt.Printf("       烟雾报警: %t (true-报警, false-正常)\n", smokeAlarm)
				fmt.Printf("       CO 报警: %t (true-报警, false-正常)\n", coAlarm)
				fmt.Printf("       H2 报警: %t (true-报警, false-正常)\n", h2Alarm)
				fmt.Printf("       VOC 报警: %t (true-报警, false-正常)\n", vocAlarm)
				fmt.Printf("       1级报警: %t (true-报警, false-正常)\n", level1Alarm)
				fmt.Printf("       2级报警: %t (true-报警, false-正常)\n", level2Alarm)
				fmt.Printf("       探测器故障: %t (true-故障, false-正常)\n", detectorFault)
				fmt.Printf("       气溶胶硬件故障: %t (true-故障, false-正常)\n", aerosolHardwareFault)
				fmt.Printf("       主电欠压故障: %t (true-故障, false-正常)\n", mainPowerUndervoltageFault)
				fmt.Printf("       报文编号: %d\n", messageNumber)
			} else {
				// 理论上不会发生，因为 can.Data 总是 8 字节长
				fmt.Printf("       Received frame with ID 0x1C000109, but data length is less than 8 bytes.\n")
			}
		}
	}
	println("done!")
}
