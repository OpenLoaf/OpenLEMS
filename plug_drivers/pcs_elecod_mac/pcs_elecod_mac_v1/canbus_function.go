package pcs_elecod_mac_v1

import (
	"fmt"
	"sync"
)

// 全局缓存相关变量
var (
	cache     = make(map[uint32]CANFrameInfo)
	cacheLock sync.RWMutex
)

// 解析CANbus ID
func parseCANbusID(id uint32) CANFrameInfo {
	// 首先尝试从缓存读取
	cacheLock.RLock()
	if info, exists := cache[id]; exists {
		cacheLock.RUnlock()

		return info
	}
	cacheLock.RUnlock()

	// 缓存未命中，进行解析
	info := CANFrameInfo{
		TargetDeviceType: DeviceType((id >> 26) & 0x07),
		TargetDeviceAddr: (id >> 22) & 0x0F,
		SourceDeviceType: DeviceType((id >> 19) & 0x07),
		SourceDeviceAddr: (id >> 15) & 0x0F,
		MessageType:      MessageType((id >> 11) & 0x0F),
		ServiceCode:      (id >> 3) & 0xFF,
		Reserved:         id & 0x07,
	}

	// 将结果存入缓存
	cacheLock.Lock()
	cache[id] = info
	cacheLock.Unlock()

	return info
}

// 解析数据帧
func parseDataFrame(data [8]byte) {
	// 按照DATA1(LB:HB)DATA2(LB:HB)DATA3(LB:HB)DATA4(LB:HB)格式解析
	// 每个DATA占用2个字节，LB在前，HB在后
	data1 := uint16(data[0]) | (uint16(data[1]) << 8) // LB + HB
	data2 := uint16(data[2]) | (uint16(data[3]) << 8)
	data3 := uint16(data[4]) | (uint16(data[5]) << 8)
	data4 := uint16(data[6]) | (uint16(data[7]) << 8)

	fmt.Printf("数据帧: DATA1=0x%04X DATA2=0x%04X DATA3=0x%04X DATA4=0x%04X\n",
		data1, data2, data3, data4)
	fmt.Printf("数据帧: DATA1=%d DATA2=%d DATA3=%d DATA4=%d\n",
		data1, data2, data3, data4)
}

// 打印解析结果
func PrintCANbusID(id uint32) {
	info := parseCANbusID(id)

	PrintCanFrame(id, info)
}

func PrintCanFrame(id uint32, info CANFrameInfo) {
	fmt.Printf("ID: 0x%X (%029b)\t源=%s(%d) -> 目标=%s(%d)\t信息类型=%s(%d)\t服务码=0x%X\t\n",
		id,
		id,
		info.SourceDeviceType,
		info.SourceDeviceAddr,
		info.TargetDeviceType,
		info.TargetDeviceAddr,
		info.MessageType,
		uint32(info.MessageType),
		info.ServiceCode)
}

// PrintCANbusFrame 打印完整的CANbus帧信息（ID + 数据）
func PrintCANbusFrame(id uint32, data [8]byte) {
	PrintCANbusID(id)
	parseDataFrame(data)
	fmt.Println()
}

// ClearCache 清空缓存（可选，用于测试或内存管理）
func ClearCache() {
	cacheLock.Lock()
	cache = make(map[uint32]CANFrameInfo)
	cacheLock.Unlock()
}
