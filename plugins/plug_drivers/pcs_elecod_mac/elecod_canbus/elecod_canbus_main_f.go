package elecod_canbus

import (
	"common/c_util"
	"fmt"
	"sync"
)

// 全局缓存相关变量
var (
	cache     = make(map[uint32]SCANFrameInfo)
	cacheLock sync.RWMutex
)

// 解析CANbus ID
func ParseCANbusID(id uint32) SCANFrameInfo {
	// 首先尝试从缓存读取
	cacheLock.RLock()
	if info, exists := cache[id]; exists {
		cacheLock.RUnlock()

		return info
	}
	cacheLock.RUnlock()

	// 缓存未命中，进行解析
	info := SCANFrameInfo{
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

func BuildScreenToMapCanbus(messageType MessageType, serviceCode uint32, params map[string]any) *uint32 {
	selfAddress, err := c_util.ToUint32(params["SelfAddress"])
	if err != nil {
		return nil
	}
	macAddress, err := c_util.ToUint32(params["MacAddress"])
	if err != nil {
		return nil
	}
	return BuildCanbusID(&SCANFrameInfo{
		TargetDeviceType: DeviceTypeMAC,
		TargetDeviceAddr: macAddress,
		SourceDeviceType: DeviceTypeScreen,
		SourceDeviceAddr: selfAddress,
		MessageType:      messageType,
		ServiceCode:      serviceCode,
	})
}

func BuildMacToScreenCanbusID(messageType MessageType, serviceCode uint32, params map[string]any) *uint32 {
	selfAddress, err := c_util.ToUint32(params["SelfAddress"])
	if err != nil {
		return nil
	}
	macAddress, err := c_util.ToUint32(params["MacAddress"])
	if err != nil {
		return nil
	}

	return BuildCanbusID(&SCANFrameInfo{
		TargetDeviceType: DeviceTypeScreen,
		TargetDeviceAddr: selfAddress,
		SourceDeviceType: DeviceTypeMAC,
		SourceDeviceAddr: macAddress,
		MessageType:      messageType,
		ServiceCode:      serviceCode,
	})
}

// 构建CANbus ID
func BuildCanbusID(info *SCANFrameInfo) *uint32 {
	if info == nil || info.TargetDeviceAddr == 0 || info.SourceDeviceAddr == 0 {
		return nil
	}

	var id uint32

	// 按照解析时的位操作进行反向操作
	id |= (uint32(info.TargetDeviceType) & 0x07) << 26 // TargetDeviceType: bits 26-28
	id |= (info.TargetDeviceAddr & 0x0F) << 22         // TargetDeviceAddr: bits 22-25
	id |= (uint32(info.SourceDeviceType) & 0x07) << 19 // SourceDeviceType: bits 19-21
	id |= (info.SourceDeviceAddr & 0x0F) << 15         // SourceDeviceAddr: bits 15-18
	id |= (uint32(info.MessageType) & 0x0F) << 11      // MessageType: bits 11-14
	id |= (info.ServiceCode & 0xFF) << 3               // ServiceCode: bits 3-10
	id |= info.Reserved & 0x07                         // Reserved: bits 0-2

	return &id
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
	info := ParseCANbusID(id)

	PrintCanFrame(id, info)
}

func PrintCanFrame(id uint32, info SCANFrameInfo) {
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
	cache = make(map[uint32]SCANFrameInfo)
	cacheLock.Unlock()
}
