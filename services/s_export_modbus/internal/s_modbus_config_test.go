package internal

import (
	"encoding/json"
	"testing"
)

func TestSModbusConfig(t *testing.T) {
	// 测试启用状态的配置
	enabledConfig := SModbusConfig{
		Enabled:    true,
		ListenPort: 502,
		DeviceIds:  []string{"device1", "device2"},
	}

	// 测试禁用状态的配置
	disabledConfig := SModbusConfig{
		Enabled:    false,
		ListenPort: 502,
		DeviceIds:  []string{"device1", "device2"},
	}

	// 测试JSON序列化
	enabledJSON, err := json.Marshal(enabledConfig)
	if err != nil {
		t.Fatalf("序列化启用配置失败: %v", err)
	}

	disabledJSON, err := json.Marshal(disabledConfig)
	if err != nil {
		t.Fatalf("序列化禁用配置失败: %v", err)
	}

	// 验证序列化结果
	expectedEnabled := `{"enabled":true,"listenPort":502,"deviceIds":["device1","device2"]}`
	expectedDisabled := `{"enabled":false,"listenPort":502,"deviceIds":["device1","device2"]}`

	if string(enabledJSON) != expectedEnabled {
		t.Errorf("启用配置序列化结果不匹配，期望: %s, 实际: %s", expectedEnabled, string(enabledJSON))
	}

	if string(disabledJSON) != expectedDisabled {
		t.Errorf("禁用配置序列化结果不匹配，期望: %s, 实际: %s", expectedDisabled, string(disabledJSON))
	}

	// 测试JSON反序列化
	var parsedEnabled SModbusConfig
	err = json.Unmarshal(enabledJSON, &parsedEnabled)
	if err != nil {
		t.Fatalf("反序列化启用配置失败: %v", err)
	}

	var parsedDisabled SModbusConfig
	err = json.Unmarshal(disabledJSON, &parsedDisabled)
	if err != nil {
		t.Fatalf("反序列化禁用配置失败: %v", err)
	}

	// 验证反序列化结果
	if parsedEnabled.Enabled != true {
		t.Errorf("启用配置反序列化失败，期望: true, 实际: %v", parsedEnabled.Enabled)
	}

	if parsedDisabled.Enabled != false {
		t.Errorf("禁用配置反序列化失败，期望: false, 实际: %v", parsedDisabled.Enabled)
	}
}
