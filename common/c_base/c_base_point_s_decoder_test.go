package c_base

import (
	"testing"
)

// TestSModbusPointTaskDecoder 测试 SModbusPointTask.Decoder 方法
func TestSModbusPointTaskDecoder(t *testing.T) {
	// 创建测试任务：从寄存器地址100开始，读取5个寄存器
	task := &SModbusPointTask{
		Name:     "TestTask",
		Addr:     100,
		Quantity: 5,
	}

	// 模拟10字节数据 (5个寄存器 * 2字节)
	testData := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0, 0x11, 0x22}

	tests := []struct {
		name        string
		point       *SModbusPoint
		expectError bool
		errorMsg    string
	}{
		{
			name: "Valid_Register_Address",
			point: &SModbusPoint{
				SPoint: &SPoint{
					Key:  "test_reg",
					Name: "Test Register",
				},
				Address:     102, // 寄存器地址102，在任务范围内 [100-104]
				Length:      1,   // 1个寄存器
				AddressType: EPointAddressTypeByte,
				ByteEndian:  ByteEndianBig,
				WordOrder:   WordOrderHighLow,
				DataFormat:  DataFormatUInt16,
				Type:        SUseReadType,
				Factor:      1.0,
			},
			expectError: false,
		},
		{
			name: "Valid_Bit_Address",
			point: &SModbusPoint{
				SPoint: &SPoint{
					Key:  "test_bit",
					Name: "Test Bit",
				},
				Address:     1604, // 位地址1604 (寄存器100*16 + 4)，在任务范围内
				Length:      1,    // 1位
				AddressType: EPointAddressTypeBit,
				ByteEndian:  ByteEndianBig,
				WordOrder:   WordOrderHighLow,
				DataFormat:  DataFormatBitRange,
				Type:        SUseReadType,
				Factor:      1.0,
			},
			expectError: false,
		},
		{
			name: "Invalid_Register_Address_Below_Range",
			point: &SModbusPoint{
				SPoint: &SPoint{
					Key:  "test_invalid_low",
					Name: "Test Invalid Low",
				},
				Address:     99, // 寄存器地址99，低于任务起始地址100
				Length:      1,
				AddressType: EPointAddressTypeByte,
				ByteEndian:  ByteEndianBig,
				WordOrder:   WordOrderHighLow,
				DataFormat:  DataFormatUInt16,
				Type:        SUseReadType,
				Factor:      1.0,
			},
			expectError: true,
			errorMsg:    "register address range [99:99] is out of task range [100:104]",
		},
		{
			name: "Invalid_Register_Address_Above_Range",
			point: &SModbusPoint{
				SPoint: &SPoint{
					Key:  "test_invalid_high",
					Name: "Test Invalid High",
				},
				Address:     105, // 寄存器地址105，超出任务范围 [100-104]
				Length:      1,
				AddressType: EPointAddressTypeByte,
				ByteEndian:  ByteEndianBig,
				WordOrder:   WordOrderHighLow,
				DataFormat:  DataFormatUInt16,
				Type:        SUseReadType,
				Factor:      1.0,
			},
			expectError: true,
			errorMsg:    "register address range [105:105] is out of task range [100:104]",
		},
		{
			name: "Invalid_Register_Length_Exceeds_Range",
			point: &SModbusPoint{
				SPoint: &SPoint{
					Key:  "test_invalid_length",
					Name: "Test Invalid Length",
				},
				Address:     103, // 寄存器地址103
				Length:      3,   // 长度3，总范围 [103-105]，超出任务范围 [100-104]
				AddressType: EPointAddressTypeByte,
				ByteEndian:  ByteEndianBig,
				WordOrder:   WordOrderHighLow,
				DataFormat:  DataFormatUInt16,
				Type:        SUseReadType,
				Factor:      1.0,
			},
			expectError: true,
			errorMsg:    "register address range [103:105] is out of task range [100:104]",
		},
		{
			name: "Invalid_Bit_Address_Below_Range",
			point: &SModbusPoint{
				SPoint: &SPoint{
					Key:  "test_bit_invalid_low",
					Name: "Test Bit Invalid Low",
				},
				Address:     1599, // 位地址1599，低于任务起始位地址1600 (100*16)
				Length:      1,
				AddressType: EPointAddressTypeBit,
				ByteEndian:  ByteEndianBig,
				WordOrder:   WordOrderHighLow,
				DataFormat:  DataFormatBitRange,
				Type:        SUseReadType,
				Factor:      1.0,
			},
			expectError: true,
			errorMsg:    "bit address 1599 is out of task range [1600:1679]",
		},
		{
			name: "Invalid_Bit_Address_Above_Range",
			point: &SModbusPoint{
				SPoint: &SPoint{
					Key:  "test_bit_invalid_high",
					Name: "Test Bit Invalid High",
				},
				Address:     1680, // 位地址1680，超出任务范围1679 ((100+5)*16-1)
				Length:      1,
				AddressType: EPointAddressTypeBit,
				ByteEndian:  ByteEndianBig,
				WordOrder:   WordOrderHighLow,
				DataFormat:  DataFormatBitRange,
				Type:        SUseReadType,
				Factor:      1.0,
			},
			expectError: true,
			errorMsg:    "bit address 1680 is out of task range [1600:1679]",
		},
		{
			name: "Valid_Multi_Register",
			point: &SModbusPoint{
				SPoint: &SPoint{
					Key:  "test_multi_reg",
					Name: "Test Multi Register",
				},
				Address:     101, // 寄存器地址101
				Length:      2,   // 2个寄存器，范围 [101-102]，在任务范围内
				AddressType: EPointAddressTypeByte,
				ByteEndian:  ByteEndianBig,
				WordOrder:   WordOrderHighLow,
				DataFormat:  DataFormatUInt32,
				Type:        SUseReadType,
				Factor:      1.0,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := task.Decoder(testData, tt.point)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("expected error message '%s', got '%s'", tt.errorMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Errorf("expected result but got nil")
				return
			}

			// 验证结果类型
			if result.IPoint != tt.point {
				t.Errorf("expected IPoint to be the same as input point")
			}

			if result.GetValue() == nil {
				t.Errorf("expected value but got nil")
			}
		})
	}
}

// TestSModbusPointTaskDecoderWithCustomDecoder 测试自定义解码器
func TestSModbusPointTaskDecoderWithCustomDecoder(t *testing.T) {
	task := &SModbusPointTask{
		Name:     "CustomDecoderTask",
		Addr:     200,
		Quantity: 2,
		CustomDecoder: func(bytes []byte, task *SModbusPointTask, point IPoint) (any, error) {
			// 简单的自定义解码器：返回字节数组的长度
			return len(bytes), nil
		},
	}

	point := &SModbusPoint{
		SPoint: &SPoint{
			Key:  "custom_test",
			Name: "Custom Test",
		},
		Address:     200,
		Length:      1,
		AddressType: EPointAddressTypeByte,
		DataFormat:  DataFormatUInt16,
		Type:        SUseReadType,
		Factor:      1.0,
	}

	testData := []byte{0x12, 0x34, 0x56, 0x78}

	result, err := task.Decoder(testData, point)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if result == nil {
		t.Errorf("expected result but got nil")
		return
	}

	// 验证自定义解码器被调用
	if result.GetValue() != 4 { // 字节数组长度
		t.Errorf("expected value 4, got %v", result.GetValue())
	}
}

// TestSModbusPointTaskDecoderAddressCalculation 测试地址计算的正确性
func TestSModbusPointTaskDecoderAddressCalculation(t *testing.T) {
	// 创建任务：从寄存器地址10开始，读取3个寄存器
	task := &SModbusPointTask{
		Name:     "AddressCalcTask",
		Addr:     10,
		Quantity: 3,
	}

	// 6字节数据 (3个寄存器 * 2字节)
	testData := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66}

	t.Run("Register_Address_Calculation", func(t *testing.T) {
		// 测试寄存器地址11（第二个寄存器）
		point := &SModbusPoint{
			SPoint: &SPoint{
				Key:  "reg_11",
				Name: "Register 11",
			},
			Address:     11,
			Length:      1,
			AddressType: EPointAddressTypeByte,
			ByteEndian:  ByteEndianBig,
			WordOrder:   WordOrderHighLow,
			DataFormat:  DataFormatUInt16,
			Type:        SUseReadType,
			Factor:      1.0,
		}

		result, err := task.Decoder(testData, point)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		// 寄存器11对应testData[2:4] = 0x3344
		expected := uint16(0x3344)
		if result.GetValue() != expected {
			t.Errorf("expected value %d, got %v", expected, result.GetValue())
		}
	})

	t.Run("Bit_Address_Calculation", func(t *testing.T) {
		// 测试位地址164（寄存器10的第4位）
		point := &SModbusPoint{
			SPoint: &SPoint{
				Key:  "bit_164",
				Name: "Bit 164",
			},
			Address:     164, // 寄存器10*16 + 4 = 164
			Length:      1,
			AddressType: EPointAddressTypeBit,
			ByteEndian:  ByteEndianBig,
			WordOrder:   WordOrderHighLow,
			DataFormat:  DataFormatBitRange,
			Type:        SUseReadType,
			Factor:      1.0,
		}

		result, err := task.Decoder(testData, point)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		// 验证结果不为nil（具体值取决于位提取逻辑）
		if result.GetValue() == nil {
			t.Errorf("expected bit value but got nil")
		}
	})
}
