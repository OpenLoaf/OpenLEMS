package internal_meta

import (
	"common/c_base"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	hexData := "010007001E0000E1"

	// 将十六进制字符串转换为字节切片
	canbusRawData, err := hex.DecodeString(hexData)
	assert.NoError(t, err, "Failed to decode hex string")
	assert.Len(t, canbusRawData, 8, "Decoded data should be 8 bytes long")

	t.Logf("Testing with CANbus Raw Data: %x", canbusRawData)

	// 根据图片定义 "探测器故障" 的 Meta
	// 图片中 "探测器故障" 在 "Byte/bit 3.1"
	// 转换为 0-indexed 字节索引: Byte 3 -> index 2
	// 转换为 0-indexed 比特索引: Bit 1 -> bit 0
	detectorFaultMeta := &c_base.Meta{
		Name:       "DetectorFault",
		Cn:         "探测器故障",
		Addr:       2,                 // 字节索引 2 (对应 Byte 3)
		ReadType:   c_base.RBit0,      // 比特索引 0 (对应 Bit 1)
		BitLength:  1,                 // 读取 1 位
		Endianness: c_base.EBigEndian, // 大端序
		Desc:       "位 1 -故障 位 0 -正常",
	}

	Level2Alarm := &c_base.Meta{Name: "Level2Alarm", Cn: "2级报警", Addr: 1, ReadType: c_base.RBit6, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -报警 位 0 -正常", Trigger: c_base.IsNotZero}
	GasCapsuleHardwareFault := &c_base.Meta{Name: "GasCapsuleHardwareFault", Cn: "气溶胶硬件故障", Addr: 2, ReadType: c_base.RBit1, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}
	MainCircuitVoltageFault := &c_base.Meta{Name: "MainCircuitVoltageFault", Cn: "主电欠压故障", Addr: 2, ReadType: c_base.RBit2, SystemType: c_base.SBool, Unit: "", Desc: "位 1 -故障 位 0 -正常", Trigger: c_base.IsNotZero}

	s, err := ParseCanbusData(canbusRawData, Level2Alarm)
	fmt.Println(s)
	s, err = ParseCanbusData(canbusRawData, GasCapsuleHardwareFault)
	fmt.Println(s)
	s, err = ParseCanbusData(canbusRawData, MainCircuitVoltageFault)
	fmt.Println(s)

	// 预期结果分析：
	// canbusRawData 的第三个字节 (索引 2) 是 0x07。
	// 0x07 的二进制表示是 0000 0111。
	// Bit 0 (最右边，0-indexed) 的值是 1。
	// 因此，预期 "探测器故障" 的结果是 true。

	// 调用解析函数
	parsedValue, err := ParseCanbusData(canbusRawData, detectorFaultMeta)

	// 验证没有错误
	assert.NoError(t, err, "ParseCanbusData should not return an error")

	// 验证返回的类型是否是 bool
	boolValue, ok := parsedValue.(bool)
	assert.True(t, ok, "Parsed value should be a boolean")

	// 验证解析结果是否符合预期
	assert.True(t, boolValue, "DetectorFault should be true (fault)")

	t.Logf("Parsed '探测器故障': %v (Expected: true)", boolValue)
}

// TestEndiannessImportance 测试字节序重排的重要性
func TestEndiannessImportance(t *testing.T) {
	// 构造测试数据：4字节数据 [0x12, 0x34, 0x56, 0x78]
	testData := []byte{0x12, 0x34, 0x56, 0x78}

	// 测试读取第2个字节（索引1，0x34）的第4位
	// 0x34 = 00110100，第4位（从右数起，0-indexed）是 0

	t.Run("BigEndian", func(t *testing.T) {
		meta := &c_base.Meta{
			Addr:       1,                 // 第2个字节 (0x34)
			ReadType:   c_base.RBit4,      // 第4位
			BitLength:  1,                 // 读取1位
			Endianness: c_base.EBigEndian, // 大端序
		}

		result, err := ParseCanbusData(testData, meta)
		assert.NoError(t, err)

		boolValue, ok := result.(bool)
		assert.True(t, ok, "Result should be boolean")

		// 0x34 = 00110100，第4位（从右数起，0-indexed）是1
		assert.True(t, boolValue, "Bit 4 of 0x34 should be true")
		t.Logf("BigEndian - Byte 0x34, Bit 4: %v", boolValue)
	})

	t.Run("LittleEndian", func(t *testing.T) {
		meta := &c_base.Meta{
			Addr:       1,                    // 在小端序下，索引1应该对应重排后的字节
			ReadType:   c_base.RBit4,         // 第4位
			BitLength:  1,                    // 读取1位
			Endianness: c_base.ELittleEndian, // 小端序
		}

		result, err := ParseCanbusData(testData, meta)
		assert.NoError(t, err)

		boolValue, ok := result.(bool)
		assert.True(t, ok, "Result should be boolean")

		// 在小端序下，字节会被重排，但对于单字节读取，结果应该相同
		// 因为我们读取的仍然是同一个字节位置的相同位
		t.Logf("LittleEndian - Bit 4: %v", boolValue)
	})

	t.Run("CrossByteRead", func(t *testing.T) {
		// 测试跨字节的多位读取，这里字节序的影响会更明显
		// 读取从第1个字节开始的12位数据
		meta := &c_base.Meta{
			Addr:       1,                 // 从第2个字节开始
			ReadType:   c_base.RBit0,      // 从第0位开始
			BitLength:  12,                // 读取12位（跨越两个字节）
			Endianness: c_base.EBigEndian, // 大端序
		}

		result, err := ParseCanbusData(testData, meta)
		assert.NoError(t, err)

		uint16Value, ok := result.(uint16)
		assert.True(t, ok, "Result should be uint16 for 12-bit read")

		// 在大端序下：[0x34, 0x56] 的12位应该是 0x345
		// 0x34 = 00110100, 0x56 = 01010110
		// 合并的12位：001101000101 = 0x345 = 837
		t.Logf("CrossByte BigEndian - 12 bits: 0x%x (%d)", uint16Value, uint16Value)
	})
}

// TestBitIndexConversion 测试位索引转换的正确性
func TestBitIndexConversion(t *testing.T) {
	// 测试单个字节的所有位
	testByte := byte(0xA5) // 10100101
	testData := []byte{testByte}

	expectedBits := []bool{
		true,  // bit 0 (LSB): 1
		false, // bit 1: 0
		true,  // bit 2: 1
		false, // bit 3: 0
		false, // bit 4: 0
		true,  // bit 5: 1
		false, // bit 6: 0
		true,  // bit 7 (MSB): 1
	}

	for bitIndex := 0; bitIndex < 8; bitIndex++ {
		t.Run(fmt.Sprintf("Bit%d", bitIndex), func(t *testing.T) {
			meta := &c_base.Meta{
				Addr:       0,
				ReadType:   c_base.EReadType(bitIndex), // RBit0 到 RBit7
				BitLength:  1,
				Endianness: c_base.EBigEndian,
			}

			result, err := ParseCanbusData(testData, meta)
			assert.NoError(t, err)

			boolValue, ok := result.(bool)
			assert.True(t, ok, "Result should be boolean")

			assert.Equal(t, expectedBits[bitIndex], boolValue,
				"Bit %d of 0x%02X should be %v", bitIndex, testByte, expectedBits[bitIndex])

			t.Logf("Bit %d of 0x%02X: %v (expected: %v)",
				bitIndex, testByte, boolValue, expectedBits[bitIndex])
		})
	}
}
