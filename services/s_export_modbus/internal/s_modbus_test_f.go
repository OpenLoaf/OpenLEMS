package internal

import (
	"common/c_enum"
	"testing"
)

// TestCalculateRegisterCount 测试寄存器数量计算
func TestCalculateRegisterCount(t *testing.T) {
	tests := []struct {
		valueType     c_enum.EValueType
		expectedCount uint16
	}{
		{c_enum.EBool, 1},
		{c_enum.EInt8, 1},
		{c_enum.EUint8, 1},
		{c_enum.EInt16, 1},
		{c_enum.EUint16, 1},
		{c_enum.EInt32, 2},
		{c_enum.EUint32, 2},
		{c_enum.EFloat32, 2},
		{c_enum.EInt64, 4},
		{c_enum.EUint64, 4},
		{c_enum.EFloat64, 4},
		{c_enum.EString, 0}, // 字符串类型不支持
	}

	for _, test := range tests {
		result := CalculateRegisterCount(test.valueType)
		if result != test.expectedCount {
			t.Errorf("CalculateRegisterCount(%v) = %d, expected %d",
				test.valueType, result, test.expectedCount)
		}
	}
}

// TestEncodeValueToRegisters 测试值转寄存器
func TestEncodeValueToRegisters(t *testing.T) {
	tests := []struct {
		value     interface{}
		valueType c_enum.EValueType
		expected  []uint16
	}{
		{true, c_enum.EBool, []uint16{1}},
		{false, c_enum.EBool, []uint16{0}},
		{int16(12345), c_enum.EInt16, []uint16{12345}},
		{uint16(54321), c_enum.EUint16, []uint16{54321}},
		{int32(123456789), c_enum.EInt32, []uint16{1887, 22617}},   // 大端字节序
		{float32(3.14159), c_enum.EFloat32, []uint16{16457, 4059}}, // 大端字节序
	}

	for _, test := range tests {
		result, err := EncodeValueToRegisters(test.value, test.valueType)
		if err != nil {
			t.Errorf("EncodeValueToRegisters(%v, %v) error: %v",
				test.value, test.valueType, err)
			continue
		}

		if len(result) != len(test.expected) {
			t.Errorf("EncodeValueToRegisters(%v, %v) length = %d, expected %d",
				test.value, test.valueType, len(result), len(test.expected))
			continue
		}

		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("EncodeValueToRegisters(%v, %v)[%d] = %d, expected %d",
					test.value, test.valueType, i, v, test.expected[i])
			}
		}
	}
}

// TestDecodeRegistersToValue 测试寄存器转值
func TestDecodeRegistersToValue(t *testing.T) {
	tests := []struct {
		registers []uint16
		valueType c_enum.EValueType
		expected  interface{}
	}{
		{[]uint16{1}, c_enum.EBool, true},
		{[]uint16{0}, c_enum.EBool, false},
		{[]uint16{12345}, c_enum.EInt16, int(12345)},
		{[]uint16{54321}, c_enum.EUint16, int(54321)},
		{[]uint16{1887, 22617}, c_enum.EInt32, int32(123456789)},
		{[]uint16{16457, 4059}, c_enum.EFloat32, float32(3.14159)},
	}

	for _, test := range tests {
		result, err := DecodeRegistersToValue(test.registers, test.valueType)
		if err != nil {
			t.Errorf("DecodeRegistersToValue(%v, %v) error: %v",
				test.registers, test.valueType, err)
			continue
		}

		// 对于浮点数，使用近似比较
		if test.valueType == c_enum.EFloat32 {
			expectedFloat := test.expected.(float32)
			resultFloat := result.(float32)
			if resultFloat-expectedFloat > 0.001 || resultFloat-expectedFloat < -0.001 {
				t.Errorf("DecodeRegistersToValue(%v, %v) = %v, expected %v",
					test.registers, test.valueType, result, test.expected)
			}
		} else {
			if result != test.expected {
				t.Errorf("DecodeRegistersToValue(%v, %v) = %v, expected %v",
					test.registers, test.valueType, result, test.expected)
			}
		}
	}
}

// TestCheckAddressOverlap 测试地址重叠检测
func TestCheckAddressOverlap(t *testing.T) {
	deviceMaps := map[string]*SDeviceRegisterMap{
		"device1": {
			DeviceId:       "device1",
			ModbusId:       1,
			StartAddr:      40000,
			TotalRegisters: 10,
		},
		"device2": {
			DeviceId:       "device2",
			ModbusId:       1,
			StartAddr:      40005, // 与device1重叠
			TotalRegisters: 10,
		},
		"device3": {
			DeviceId:       "device3",
			ModbusId:       2, // 不同的ModbusId，不应该冲突
			StartAddr:      40000,
			TotalRegisters: 10,
		},
		"device4": {
			DeviceId:       "device4",
			ModbusId:       1,
			StartAddr:      40020, // 不重叠
			TotalRegisters: 10,
		},
	}

	conflicts := CheckAddressOverlap(deviceMaps)

	// 应该检测到device1和device2的冲突
	expectedConflicts := []string{"device1", "device2"}
	if len(conflicts) != len(expectedConflicts) {
		t.Errorf("CheckAddressOverlap() returned %d conflicts, expected %d",
			len(conflicts), len(expectedConflicts))
	}

	// 检查是否包含预期的冲突设备
	for _, expected := range expectedConflicts {
		found := false
		for _, conflict := range conflicts {
			if conflict == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("CheckAddressOverlap() should include %s in conflicts", expected)
		}
	}
}
