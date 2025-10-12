package internal

import (
	"common/c_base"
	"common/c_enum"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"time"
)

// SRegisterMapping 寄存器映射结构体
type SRegisterMapping struct {
	Point         c_base.IPoint // 点位信息
	StartOffset   uint16        // 相对起始地址的偏移
	RegisterCount uint16        // 占用寄存器数量
}

// SDeviceRegisterMap 设备寄存器映射结构体
type SDeviceRegisterMap struct {
	DeviceId       string              // 设备ID
	ModbusId       uint8               // Modbus从站ID
	StartAddr      uint16              // 起始地址
	Mappings       []*SRegisterMapping // 点位映射列表
	TotalRegisters uint16              // 总寄存器数量
	IsOnline       bool                // 设备是否在线
	LastUpdateTime time.Time           // 最后更新时间
}

// SModbusDeviceStatus 设备状态结构体
type SModbusDeviceStatus struct {
	DeviceId       string    // 设备ID
	ModbusId       uint8     // Modbus从站ID
	StartAddr      uint16    // 起始地址
	IsOnline       bool      // 设备是否在线
	LastUpdateTime time.Time // 最后更新时间
	Error          string    // 错误信息（如果有）
}

// CalculateRegisterCount 计算数据类型占用寄存器数
func CalculateRegisterCount(valueType c_enum.EValueType) uint16 {
	switch valueType {
	case c_enum.EBool, c_enum.EInt8, c_enum.EUint8:
		return 1
	case c_enum.EInt16, c_enum.EUint16:
		return 1
	case c_enum.EInt32, c_enum.EUint32, c_enum.EFloat32:
		return 2
	case c_enum.EInt64, c_enum.EUint64, c_enum.EFloat64:
		return 4
	case c_enum.EString:
		// 字符串类型不支持，返回0表示跳过
		return 0
	default:
		return 1
	}
}

// BuildDeviceRegisterMap 构建设备寄存器映射
func BuildDeviceRegisterMap(deviceId string, device c_base.IDevice, externalParam *c_base.SExternalParam) (*SDeviceRegisterMap, error) {
	if device == nil {
		return nil, errors.New("设备实例为空")
	}

	if externalParam == nil {
		return nil, errors.New("设备外部参数为空")
	}

	// 获取设备点位列表
	exportPoints := device.GetExportModbusPoints()
	if len(exportPoints) == 0 {
		return nil, errors.New("设备没有可导出的Modbus点位")
	}

	// 构建寄存器映射
	mappings := make([]*SRegisterMapping, 0)
	currentOffset := uint16(0)

	// 添加固定点位：设备在线状态（1个寄存器）
	onlineMapping := &SRegisterMapping{
		Point:         nil, // 系统固定点位，不需要IPoint
		StartOffset:   currentOffset,
		RegisterCount: 1,
	}
	mappings = append(mappings, onlineMapping)
	currentOffset += 1

	// 添加固定点位：通讯时间戳（2个寄存器）
	timestampMapping := &SRegisterMapping{
		Point:         nil, // 系统固定点位，不需要IPoint
		StartOffset:   currentOffset,
		RegisterCount: 2,
	}
	mappings = append(mappings, timestampMapping)
	currentOffset += 2

	// 添加设备实际点位
	for _, point := range exportPoints {
		if point == nil {
			continue
		}

		// 跳过字符串类型
		if point.GetValueType() == c_enum.EString {
			continue
		}

		registerCount := CalculateRegisterCount(point.GetValueType())
		if registerCount == 0 {
			continue // 跳过不支持的类型
		}

		// 检查是否超过100个寄存器的限制
		if currentOffset+registerCount > 100 {
			return nil, fmt.Errorf("设备 %s 的寄存器数量超过100个限制", deviceId)
		}

		mapping := &SRegisterMapping{
			Point:         point,
			StartOffset:   currentOffset,
			RegisterCount: registerCount,
		}
		mappings = append(mappings, mapping)
		currentOffset += registerCount
	}

	// 检查总寄存器数量
	if currentOffset > 100 {
		return nil, fmt.Errorf("设备 %s 的总寄存器数量 %d 超过100个限制", deviceId, currentOffset)
	}

	return &SDeviceRegisterMap{
		DeviceId:       deviceId,
		ModbusId:       externalParam.ModbusId,
		StartAddr:      uint16(externalParam.ModbusRegisterAddr),
		Mappings:       mappings,
		TotalRegisters: currentOffset,
		IsOnline:       device.GetProtocolStatus() == c_enum.EProtocolConnected,
		LastUpdateTime: time.Now(),
	}, nil
}

// EncodeValueToRegisters 值转寄存器
func EncodeValueToRegisters(value any, valueType c_enum.EValueType) ([]uint16, error) {
	registerCount := CalculateRegisterCount(valueType)
	if registerCount == 0 {
		return nil, errors.New("不支持的数据类型")
	}

	registers := make([]uint16, registerCount)

	switch valueType {
	case c_enum.EBool:
		if boolVal, ok := value.(bool); ok {
			if boolVal {
				registers[0] = 1
			} else {
				registers[0] = 0
			}
		} else {
			return nil, errors.New("布尔值类型转换失败")
		}

	case c_enum.EInt8, c_enum.EUint8:
		if intVal, ok := value.(int); ok {
			registers[0] = uint16(intVal & 0xFF)
		} else {
			return nil, errors.New("8位整数类型转换失败")
		}

	case c_enum.EInt16, c_enum.EUint16:
		switch v := value.(type) {
		case int16:
			registers[0] = uint16(v)
		case uint16:
			registers[0] = v
		case int:
			registers[0] = uint16(v & 0xFFFF)
		case uint:
			registers[0] = uint16(v & 0xFFFF)
		default:
			return nil, errors.New("16位整数类型转换失败")
		}

	case c_enum.EInt32, c_enum.EUint32, c_enum.EFloat32:
		var bytes [4]byte
		switch v := value.(type) {
		case int32:
			binary.BigEndian.PutUint32(bytes[:], uint32(v))
		case uint32:
			binary.BigEndian.PutUint32(bytes[:], v)
		case float32:
			// 对于浮点数，需要正确的位操作
			bits := math.Float32bits(v)
			binary.BigEndian.PutUint32(bytes[:], bits)
		case int:
			binary.BigEndian.PutUint32(bytes[:], uint32(v))
		case uint:
			binary.BigEndian.PutUint32(bytes[:], uint32(v))
		default:
			return nil, errors.New("32位数据类型转换失败")
		}
		registers[0] = binary.BigEndian.Uint16(bytes[0:2])
		registers[1] = binary.BigEndian.Uint16(bytes[2:4])

	case c_enum.EInt64, c_enum.EUint64, c_enum.EFloat64:
		var bytes [8]byte
		switch v := value.(type) {
		case int64:
			binary.BigEndian.PutUint64(bytes[:], uint64(v))
		case uint64:
			binary.BigEndian.PutUint64(bytes[:], v)
		case float64:
			binary.BigEndian.PutUint64(bytes[:], uint64(v))
		default:
			return nil, errors.New("64位数据类型转换失败")
		}
		registers[0] = binary.BigEndian.Uint16(bytes[0:2])
		registers[1] = binary.BigEndian.Uint16(bytes[2:4])
		registers[2] = binary.BigEndian.Uint16(bytes[4:6])
		registers[3] = binary.BigEndian.Uint16(bytes[6:8])

	default:
		return nil, errors.New("不支持的数据类型")
	}

	return registers, nil
}

// DecodeRegistersToValue 寄存器转值
func DecodeRegistersToValue(registers []uint16, valueType c_enum.EValueType) (any, error) {
	registerCount := CalculateRegisterCount(valueType)
	if registerCount == 0 {
		return nil, errors.New("不支持的数据类型")
	}

	if len(registers) < int(registerCount) {
		return nil, errors.New("寄存器数量不足")
	}

	switch valueType {
	case c_enum.EBool:
		return registers[0] != 0, nil

	case c_enum.EInt8, c_enum.EUint8:
		return int(registers[0] & 0xFF), nil

	case c_enum.EInt16, c_enum.EUint16:
		return int(registers[0]), nil

	case c_enum.EInt32, c_enum.EUint32, c_enum.EFloat32:
		var bytes [4]byte
		binary.BigEndian.PutUint16(bytes[0:2], registers[0])
		binary.BigEndian.PutUint16(bytes[2:4], registers[1])

		switch valueType {
		case c_enum.EInt32:
			return int32(binary.BigEndian.Uint32(bytes[:])), nil
		case c_enum.EUint32:
			return binary.BigEndian.Uint32(bytes[:]), nil
		case c_enum.EFloat32:
			// 对于浮点数，需要正确的位操作
			bits := binary.BigEndian.Uint32(bytes[:])
			return math.Float32frombits(bits), nil
		}

	case c_enum.EInt64, c_enum.EUint64, c_enum.EFloat64:
		var bytes [8]byte
		binary.BigEndian.PutUint16(bytes[0:2], registers[0])
		binary.BigEndian.PutUint16(bytes[2:4], registers[1])
		binary.BigEndian.PutUint16(bytes[4:6], registers[2])
		binary.BigEndian.PutUint16(bytes[6:8], registers[3])

		switch valueType {
		case c_enum.EInt64:
			return int64(binary.BigEndian.Uint64(bytes[:])), nil
		case c_enum.EUint64:
			return binary.BigEndian.Uint64(bytes[:]), nil
		case c_enum.EFloat64:
			return float64(binary.BigEndian.Uint64(bytes[:])), nil
		}
	}

	return nil, errors.New("不支持的数据类型")
}

// CheckAddressOverlap 检查地址重叠
func CheckAddressOverlap(deviceMaps map[string]*SDeviceRegisterMap) []string {
	var conflictDevices []string

	// 按ModbusId分组
	modbusGroups := make(map[uint8][]*SDeviceRegisterMap)
	for _, deviceMap := range deviceMaps {
		modbusGroups[deviceMap.ModbusId] = append(modbusGroups[deviceMap.ModbusId], deviceMap)
	}

	// 检查每个ModbusId组内的地址重叠
	for _, devices := range modbusGroups {
		if len(devices) <= 1 {
			continue
		}

		// 检查地址重叠
		for i := 0; i < len(devices); i++ {
			for j := i + 1; j < len(devices); j++ {
				device1 := devices[i]
				device2 := devices[j]

				// 检查地址范围是否重叠
				device1Start := device1.StartAddr
				device1End := device1Start + device1.TotalRegisters - 1
				device2Start := device2.StartAddr
				device2End := device2Start + device2.TotalRegisters - 1

				// 检查是否重叠
				if !(device1End < device2Start || device2End < device1Start) {
					// 地址重叠，两个设备都标记为冲突
					conflictDevices = append(conflictDevices, device1.DeviceId, device2.DeviceId)
				}
			}
		}
	}

	return conflictDevices
}
