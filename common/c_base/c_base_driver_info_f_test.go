package c_base

import (
	"common/c_enum"
	"testing"
)

func TestResolvingValueType(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected c_enum.EValueType
	}{
		// 基本类型测试
		{"nil value", nil, c_enum.EString},
		{"bool true", true, c_enum.EBool},
		{"bool false", false, c_enum.EBool},
		{"int8", int8(42), c_enum.EInt8},
		{"uint8", uint8(42), c_enum.EUint8},
		{"int16", int16(42), c_enum.EInt16},
		{"uint16", uint16(42), c_enum.EUint16},
		{"int32", int32(42), c_enum.EInt32},
		{"uint32", uint32(42), c_enum.EUint32},
		{"int64", int64(42), c_enum.EInt64},
		{"uint64", uint64(42), c_enum.EUint64},
		{"float32", float32(3.14), c_enum.EFloat32},
		{"float64", float64(3.14), c_enum.EFloat64},
		{"string", "hello", c_enum.EString},
		{"int", 42, c_enum.EInt64},
		{"uint", uint(42), c_enum.EUint64},

		// 指针类型测试
		{"*bool", boolPtr(true), c_enum.EBool},
		{"*int32", int32Ptr(42), c_enum.EInt32},
		{"*string", stringPtr("hello"), c_enum.EString},
		{"nil pointer", (*int)(nil), c_enum.EString},

		// 复杂类型测试
		{"slice", []int{1, 2, 3}, c_enum.EString},
		{"map", map[string]int{"a": 1}, c_enum.EString},
		{"struct", struct{}{}, c_enum.EString},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ResolvingValueType(tt.value)
			if result != tt.expected {
				t.Errorf("ResolvingValueType(%v) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

// 辅助函数用于创建指针
func boolPtr(b bool) *bool {
	return &b
}

func int32Ptr(i int32) *int32 {
	return &i
}

func stringPtr(s string) *string {
	return &s
}
