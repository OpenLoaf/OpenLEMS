package main

import (
	"fmt"
	"pcs_elecod/elecod_canbus"
)

func main() {
	// 测试用例
	testIDs := []uint32{
		//0x1C50A808, // 模拟量1
		0x1C502810,
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
		elecod_canbus.PrintCANbusID(id)
	}

	//// 演示单独解析数据帧的功能
	//fmt.Println("单独数据帧解析示例:")
	//fmt.Println("--------------------------------------------------------------------------------")
	//sampleData := [8]byte{0xAB, 0xCD, 0xEF, 0x12, 0x34, 0x56, 0x78, 0x9A}
	//fmt.Printf("原始数据: %02X %02X %02X %02X %02X %02X %02X %02X\n",
	//	sampleData[0], sampleData[1], sampleData[2], sampleData[3],
	//	sampleData[4], sampleData[5], sampleData[6], sampleData[7])
	//parseDataFrame(sampleData)
}
