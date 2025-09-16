package internal

import (
	"common/c_base"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/glist"
)

// BuildTree 从平面列表构建树形结构
func (m *SDeviceManager) BuildTree(devices []*c_base.SDeviceConfig) []*c_base.SDeviceConfig {
	// 创建设备ID到设备的映射
	deviceMap := make(map[string]*c_base.SDeviceConfig)
	var rootDevices []*c_base.SDeviceConfig

	// 首先将所有设备放入映射中
	for _, device := range devices {
		deviceMap[device.Id] = device
		// 清空子设备列表，避免重复
		device.ChildDeviceConfig = nil
	}

	// 构建树形结构
	for _, device := range devices {
		if device.Pid == "" || device.Pid == device.Id {
			// 根设备
			rootDevices = append(rootDevices, device)
		} else {
			// 子设备，添加到父设备下
			if parent, exists := deviceMap[device.Pid]; exists {
				parent.ChildDeviceConfig = append(parent.ChildDeviceConfig, device)
			} else {
				// 父设备不存在，作为根设备处理
				rootDevices = append(rootDevices, device)
			}
		}
	}

	return rootDevices
}

// FindDevice 查找设备
func (m *SDeviceManager) FindDevice(deviceId string) *c_base.SDeviceConfig {
	// 将线程安全的列表转换为切片
	var devices []*c_base.SDeviceConfig
	m.deviceConfigTree.Iterator(func(e *glist.Element) bool {
		if config, ok := e.Value.(*c_base.SDeviceConfig); ok {
			devices = append(devices, config)
		}
		return true
	})
	return m.findDeviceRecursive(devices, deviceId)
}

// findDeviceRecursive 递归查找设备
func (m *SDeviceManager) findDeviceRecursive(devices []*c_base.SDeviceConfig, deviceId string) *c_base.SDeviceConfig {
	for _, device := range devices {
		if device.Id == deviceId {
			return device
		}
		// 递归查找子设备
		if found := m.findDeviceRecursive(device.ChildDeviceConfig, deviceId); found != nil {
			return found
		}
	}
	return nil
}

// GetFlatList 获取平面列表
func (m *SDeviceManager) GetFlatList() []*c_base.SDeviceConfig {
	var flatList []*c_base.SDeviceConfig
	// 将线程安全的列表转换为切片
	var devices []*c_base.SDeviceConfig
	m.deviceConfigTree.Iterator(func(e *glist.Element) bool {
		if config, ok := e.Value.(*c_base.SDeviceConfig); ok {
			devices = append(devices, config)
		}
		return true
	})
	m.flattenDevices(devices, &flatList)
	return flatList
}

// flattenDevices 递归扁平化设备列表
func (m *SDeviceManager) flattenDevices(devices []*c_base.SDeviceConfig, result *[]*c_base.SDeviceConfig) {
	for _, device := range devices {
		*result = append(*result, device)
		// 递归处理子设备
		m.flattenDevices(device.ChildDeviceConfig, result)
	}
}

// PrintTree 打印树形结构（用于调试）
func (m *SDeviceManager) PrintTree() {
	// 将线程安全的列表转换为切片
	var devices []*c_base.SDeviceConfig
	m.deviceConfigTree.Iterator(func(e *glist.Element) bool {
		if config, ok := e.Value.(*c_base.SDeviceConfig); ok {
			devices = append(devices, config)
		}
		return true
	})
	m.printTreeRecursive(devices, 0)
}

// printTreeRecursive 递归打印树形结构
func (m *SDeviceManager) printTreeRecursive(devices []*c_base.SDeviceConfig, level int) {
	indent := strings.Repeat("  ", level)
	for _, device := range devices {
		fmt.Printf("%s├─ %s (ID: %s, PID: %s)\n", indent, device.Name, device.Id, device.Pid)
		if len(device.ChildDeviceConfig) > 0 {
			m.printTreeRecursive(device.ChildDeviceConfig, level+1)
		}
	}
}

// ExecuteFromBottom 从最底部的节点开始执行任务
func (m *SDeviceManager) ExecuteFromBottom(executor func(deviceConfig *c_base.SDeviceConfig)) {
	// 将线程安全的列表转换为切片
	var devices []*c_base.SDeviceConfig
	m.deviceConfigTree.Iterator(func(e *glist.Element) bool {
		if config, ok := e.Value.(*c_base.SDeviceConfig); ok {
			devices = append(devices, config)
		}
		return true
	})
	m.executeFromBottomRecursive(devices, executor)
}

// executeFromBottomRecursive 递归从底部执行任务
func (m *SDeviceManager) executeFromBottomRecursive(devices []*c_base.SDeviceConfig, executor func(deviceConfig *c_base.SDeviceConfig)) {
	for _, device := range devices {
		// 先递归处理子设备
		m.executeFromBottomRecursive(device.ChildDeviceConfig, executor)
		// 然后执行当前设备
		executor(device)
	}
}
