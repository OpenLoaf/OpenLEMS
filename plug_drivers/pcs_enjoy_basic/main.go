package main

import (
	"fmt"
)

// 设备类型定义
var deviceTypes = map[uint32]string{
	0b000: "广播地址",
	0b001: "MAC (主控制器)",
	0b010: "MDC (运动控制器)",
	0b011: "STS (状态控制器)",
	0b111: "屏 (显示屏)",
}

// 信息类型定义
var messageTypes = map[uint32]string{
	0x1: "配置信息",
	0x2: "控制信息",
	0x3: "告警信息",
	0x4: "状态信息",
	0x5: "硬件信息",
}

// CANbus帧信息结构体
type CANFrameInfo struct {
	TargetDeviceType string
	TargetDeviceAddr uint32
	SourceDeviceType string
	SourceDeviceAddr uint32
	MessageType      string
	ServiceCode      uint32
	Reserved         uint32
	FrameFormat      string
	FrameType        string
}

// 解析CANbus ID
func parseCANbusID(id uint32) CANFrameInfo {
	var info CANFrameInfo

	// 提取各个字段 (基于29位扩展帧格式)
	// 目标设备类型 [31:29] -> [28:26] (因为实际使用29位)
	targetDeviceType := (id >> 26) & 0x07
	// 目标设备地址 [28:25] -> [25:22]
	targetDeviceAddr := (id >> 22) & 0x0F
	// 源设备类型 [24:22] -> [21:19]
	sourceDeviceType := (id >> 19) & 0x07
	// 源设备地址 [21:18] -> [18:15]
	sourceDeviceAddr := (id >> 15) & 0x0F
	// 信息类型 [17:14] -> [14:11]
	messageType := (id >> 11) & 0x0F
	// 服务码 [13:6] -> [10:3]
	serviceCode := (id >> 3) & 0xFF
	// 预留 [5:3] -> [2:0]
	reserved := id & 0x07

	// 填充结构体
	if deviceName, exists := deviceTypes[targetDeviceType]; exists {
		info.TargetDeviceType = deviceName
	} else {
		info.TargetDeviceType = fmt.Sprintf("未知类型(0x%X)", targetDeviceType)
	}
	info.TargetDeviceAddr = targetDeviceAddr

	if deviceName, exists := deviceTypes[sourceDeviceType]; exists {
		info.SourceDeviceType = deviceName
	} else {
		info.SourceDeviceType = fmt.Sprintf("未知类型(0x%X)", sourceDeviceType)
	}
	info.SourceDeviceAddr = sourceDeviceAddr

	if msgType, exists := messageTypes[messageType]; exists {
		info.MessageType = msgType
	} else {
		info.MessageType = fmt.Sprintf("未知信息类型(0x%X)", messageType)
	}

	info.ServiceCode = serviceCode
	info.Reserved = reserved

	// 通讯配置字段在实际CAN帧中不在ID中，这里暂时设为默认值
	info.FrameFormat = "扩展帧"
	info.FrameType = "数据帧"

	return info
}

// 打印解析结果
func printCANbusID(id uint32) {
	info := parseCANbusID(id)

	// 获取信息类型的数值
	messageTypeNum := (id >> 11) & 0x0F

	fmt.Printf("ID: 0x%X (%029b)\t源=%s(%d) -> 目标=%s(%d)\t信息类型=%s(%d)\t服务码=0x%X\t通讯配置: %s|%s\n",
		id,
		id,
		info.SourceDeviceType,
		info.SourceDeviceAddr,
		info.TargetDeviceType,
		info.TargetDeviceAddr,
		info.MessageType,
		messageTypeNum,
		info.ServiceCode,
		info.FrameFormat,
		info.FrameType)
}

func main() {
	// 测试用例
	testIDs := []uint32{
		0x08789010, // 表为对设备号为2的MDC下发控制信息
		0x1C508808, // 表为设备号为2的MDC上传信息
		0x1C501808, // 该表为1号设备上传告警信息
		0x1C50A808, // 该表为2号设备
	}

	fmt.Println("CANbus ID 解析结果:")
	fmt.Println("================================================================================")
	for _, id := range testIDs {
		printCANbusID(id)
	}
}
