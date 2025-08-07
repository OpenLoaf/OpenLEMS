package main

import (
	"fmt"
	"pcs_elecod/pcs_elecod_mac_v1"
)

func main() {
	// 测试用例
	testIDs := []uint32{
		0x1c482858,
		0x03f8a008,
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
	}

	fmt.Println("CANbus 帧解析结果:")
	fmt.Println("================================================================================")
	for _, id := range testIDs {
		pcs_elecod_mac_v1.PrintCANbusID(id)
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
