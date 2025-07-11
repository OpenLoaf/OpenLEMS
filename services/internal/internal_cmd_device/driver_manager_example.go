package internal_cmd_device

import (
	"common/c_base"
	"context"
	"fmt"
	"log"
)

// 使用示例：演示如何使用通用驱动管理器
func ExampleDriverManager() {
	ctx := context.Background()

	// 创建驱动管理器
	manager := NewDriverManager()

	// 获取所有驱动名称
	fmt.Println("=== 所有驱动名称 ===")
	driverNames := manager.GetAllDriverNames()
	for _, name := range driverNames {
		fmt.Printf("- %s\n", name)
	}

	// 获取所有驱动详细信息
	fmt.Println("\n=== 所有驱动详细信息 ===")
	driversInfo := manager.GetAllDriversInfo(ctx)
	for _, info := range driversInfo {
		fmt.Printf("驱动: %s, 类型: %s, 可用: %t\n", info.Name, info.Type, info.Available)
		if info.Description != nil {
			fmt.Printf("  品牌: %s, 型号: %s, 版本: %s\n",
				info.Description.Brand, info.Description.Model, info.Description.Version)
		}
	}

	// 获取特定类型的驱动
	fmt.Println("\n=== BMS类型的驱动 ===")
	bmsDrivers := manager.GetDriversByType(ctx, c_base.EDeviceBms)
	for _, driver := range bmsDrivers {
		fmt.Printf("- %s (可用: %t)\n", driver.Name, driver.Available)
	}

	// 检查驱动是否可用
	fmt.Println("\n=== 检查驱动可用性 ===")
	if len(driverNames) > 0 {
		testDriver := driverNames[0]
		available := manager.IsDriverAvailable(ctx, testDriver)
		fmt.Printf("驱动 %s 是否可用: %t\n", testDriver, available)
	}

	// 获取支持的设备类型
	fmt.Println("\n=== 支持的设备类型 ===")
	deviceTypes := manager.GetSupportedDeviceTypes(ctx)
	for _, deviceType := range deviceTypes {
		fmt.Printf("- %s\n", deviceType)
	}
}

// 使用示例：驱动管理器辅助工具
func ExampleDriverManagerHelper() {
	ctx := context.Background()

	// 创建驱动管理器辅助工具
	helper := NewDriverManagerHelper()

	// 获取驱动统计信息
	fmt.Println("=== 驱动统计信息 ===")
	stats := helper.GetDriverStats(ctx)
	fmt.Printf("总数: %d\n", stats["total"])
	fmt.Printf("可用: %d\n", stats["available"])
	fmt.Printf("不可用: %d\n", stats["unavailable"])

	if typeCount, ok := stats["byType"].(map[c_base.EDeviceType]int); ok {
		fmt.Println("按类型分组:")
		for deviceType, count := range typeCount {
			fmt.Printf("  %s: %d\n", deviceType, count)
		}
	}

	// 搜索驱动
	fmt.Println("\n=== 搜索驱动 ===")
	searchResults := helper.SearchDrivers(ctx, "bms")
	fmt.Printf("搜索 'bms' 的结果:\n")
	for _, result := range searchResults {
		fmt.Printf("- %s\n", result.Name)
	}

	// 验证驱动配置
	fmt.Println("\n=== 验证驱动配置 ===")
	testConfig := &c_base.SDriverConfig{
		Id:     "test-device",
		Name:   "测试设备",
		Driver: "bms_lnxall_v1.0.0",
	}

	if err := helper.ValidateDriverConfig(ctx, testConfig); err != nil {
		fmt.Printf("配置验证失败: %v\n", err)
	} else {
		fmt.Println("配置验证成功")
	}

	// 获取驱动遥测信息
	fmt.Println("\n=== 获取驱动遥测信息 ===")
	telemetryInfo, err := helper.GetDriverTelemetryInfo(ctx, "bms_lnxall_v1.0.0")
	if err != nil {
		fmt.Printf("获取遥测信息失败: %v\n", err)
	} else if telemetryInfo != nil {
		fmt.Printf("遥测点位数量: %d\n", len(telemetryInfo))
		for i, tel := range telemetryInfo {
			if i < 3 { // 只显示前3个
				fmt.Printf("  - %s: %s\n", tel.Name, tel.NationalizationName)
			}
		}
		if len(telemetryInfo) > 3 {
			fmt.Printf("  ... 还有 %d 个遥测点位\n", len(telemetryInfo)-3)
		}
	}
}

// 使用示例：创建驱动实例
func ExampleCreateDriver() {
	ctx := context.Background()

	// 创建驱动管理器
	manager := NewDriverManager()

	// 配置一个设备
	deviceConfig := &c_base.SDriverConfig{
		Id:       "test-bms-001",
		Name:     "测试BMS设备",
		Driver:   "bms_lnxall_v1.0.0",
		IsEnable: true,
	}

	// 创建驱动实例
	driver, err := manager.CreateDriver(ctx, deviceConfig)
	if err != nil {
		log.Printf("创建驱动失败: %v", err)
		return
	}

	fmt.Printf("成功创建驱动: %s\n", driver.GetDriverType())

	// 获取驱动描述
	description := driver.GetDescription()
	if description != nil {
		fmt.Printf("驱动描述: %s %s v%s\n",
			description.Brand, description.Model, description.Version)
	}

	// 获取设备配置
	config := driver.GetDeviceConfig()
	if config != nil {
		fmt.Printf("设备配置: %s (%s)\n", config.Name, config.Id)
	}
}

// 通用驱动管理器使用指南
func PrintUsageGuide() {
	fmt.Println(`
=== 通用驱动管理器使用指南 ===

1. 基本用法:
   manager := NewDriverManager()
   
2. 获取驱动信息:
   - GetAllDriverNames() - 获取所有驱动名称
   - GetAllDriversInfo(ctx) - 获取所有驱动详细信息
   - GetDriverInfo(ctx, driverName) - 获取指定驱动信息
   - GetDriversByType(ctx, deviceType) - 按类型获取驱动
   
3. 驱动操作:
   - CreateDriver(ctx, config) - 创建驱动实例
   - IsDriverAvailable(ctx, driverName) - 检查驱动可用性
   - GetDriverDescription(ctx, driverName) - 获取驱动描述
   - GetSupportedDeviceTypes(ctx) - 获取支持的设备类型
   
4. 辅助工具:
   helper := NewDriverManagerHelper()
   - GetDriverStats(ctx) - 获取统计信息
   - ValidateDriverConfig(ctx, config) - 验证配置
   - SearchDrivers(ctx, keyword) - 搜索驱动
   - GetDriverTelemetryInfo(ctx, driverName) - 获取遥测信息
   
5. 特性:
   - 支持开发环境和生产环境
   - 自动处理驱动加载失败
   - 提供详细的错误信息
   - 支持按类型筛选驱动
   - 内置搜索功能
`)
}
