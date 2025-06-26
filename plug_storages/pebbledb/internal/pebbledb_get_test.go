package internal

import (
	"common/c_base"
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// TestPebbledbRead 测试读取功能的示例
func TestPebbledbRead(t *testing.T) {
	ctx := context.Background()

	// 创建pebbledb实例
	pdb := NewPebbledb(ctx).(*Pebbledb)
	defer pdb.Close()

	// 先写入一些测试数据
	setupTestData(pdb)

	// 演示各种读取功能
	fmt.Println("=== PebbleDB 读取功能测试 ===")

	// 1. 获取数据库统计信息
	fmt.Println("\n1. 数据库统计信息:")
	stats := pdb.GetStats()
	printJSON(stats)

	// 2. 获取所有键名
	fmt.Println("\n2. 获取所有键名 (前100个):")
	keys, err := pdb.GetAllKeys("", 100)
	if err != nil {
		t.Errorf("获取键名失败: %v", err)
		return
	}
	for i, key := range keys {
		fmt.Printf("  [%d] %s\n", i+1, key)
	}

	// 3. 根据设备ID获取数据
	fmt.Println("\n3. 根据设备ID获取数据:")
	deviceData, err := pdb.GetDeviceData("test_device_001", 5)
	if err != nil {
		t.Errorf("获取设备数据失败: %v", err)
		return
	}
	for i, data := range deviceData {
		fmt.Printf("  [%d] 设备ID: %s, 类型: %s, 时间: %s\n",
			i+1, data.DeviceID, data.DeviceType,
			time.Unix(data.Timestamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("       字段: %v\n", data.Fields)
	}

	// 4. 根据设备类型获取数据
	fmt.Println("\n4. 根据设备类型获取数据:")
	typeData, err := pdb.GetDeviceDataByType(c_base.EDevicePcs, 3)
	if err != nil {
		t.Errorf("获取设备类型数据失败: %v", err)
		return
	}
	for i, data := range typeData {
		fmt.Printf("  [%d] 设备ID: %s, 时间: %s\n",
			i+1, data.DeviceID,
			time.Unix(data.Timestamp, 0).Format("2006-01-02 15:04:05"))
	}

	// 5. 获取时间范围内的数据
	fmt.Println("\n5. 获取时间范围内的数据:")
	startTime := time.Now().Add(-1 * time.Hour)
	endTime := time.Now()
	rangeData, err := pdb.GetDeviceDataByTimeRange("test_device_001", startTime, endTime, 5)
	if err != nil {
		t.Errorf("获取时间范围数据失败: %v", err)
		return
	}
	for i, data := range rangeData {
		fmt.Printf("  [%d] 时间: %s, 字段数量: %d\n",
			i+1, time.Unix(data.Timestamp, 0).Format("2006-01-02 15:04:05"),
			len(data.Fields))
	}

	// 6. 获取协议指标数据
	fmt.Println("\n6. 获取协议指标数据:")
	protocolData, err := pdb.GetProtocolMetrics("test_device_001", 3)
	if err != nil {
		t.Errorf("获取协议指标数据失败: %v", err)
		return
	}
	for i, data := range protocolData {
		fmt.Printf("  [%d] 协议ID: %s, 地址: %s, 指标数量: %d\n",
			i+1, data.ProtocolID, data.ProtocolAddress, len(data.Metrics))
	}

	// 7. 获取系统指标数据
	fmt.Println("\n7. 获取系统指标数据:")
	systemData, err := pdb.GetSystemMetrics("system", 3)
	if err != nil {
		t.Errorf("获取系统指标数据失败: %v", err)
		return
	}
	for i, data := range systemData {
		fmt.Printf("  [%d] 测量: %s, 时间: %s, 标签: %v\n",
			i+1, data.Measurement,
			time.Unix(data.Timestamp, 0).Format("2006-01-02 15:04:05"),
			data.Tags)
	}

	// 8. 获取最新设备数据
	fmt.Println("\n8. 获取最新设备数据:")
	latestData, err := pdb.GetLatestDeviceData("test_device_001")
	if err != nil {
		t.Errorf("获取最新设备数据失败: %v", err)
		return
	}
	fmt.Printf("   最新数据时间: %s\n",
		time.Unix(latestData.Timestamp, 0).Format("2006-01-02 15:04:05"))
	printJSON(latestData.Fields)

	// 9. 根据键名获取数据
	fmt.Println("\n9. 根据键名获取数据:")
	if len(keys) > 0 {
		keyData, err := pdb.GetDataByKey(keys[0])
		if err != nil {
			t.Errorf("根据键名获取数据失败: %v", err)
			return
		}
		fmt.Printf("   键名: %s\n", keys[0])
		printJSON(keyData)
	}

	fmt.Println("\n=== 测试完成 ===")
}

// setupTestData 设置测试数据
func setupTestData(pdb *Pebbledb) {
	// 写入设备数据
	deviceFields := map[string]any{
		"voltage":     220.5,
		"current":     10.2,
		"power":       2251.5,
		"frequency":   50.0,
		"temperature": 25.3,
	}

	for i := 0; i < 5; i++ {
		deviceID := fmt.Sprintf("test_device_%03d", i+1)
		deviceFields["power"] = 2000.0 + float64(i*100)
		pdb.SaveDevices(deviceID, c_base.EDevicePcs, deviceFields)
		time.Sleep(10 * time.Millisecond) // 确保时间戳不同
	}

	// 写入协议指标数据
	protocolConfig := &c_base.SProtocolConfig{
		Id:       "protocol_001",
		Address:  "192.168.1.100:502",
		Protocol: c_base.EModbusTcp,
	}

	deviceConfig := &c_base.SDriverConfig{
		Id:   "test_device_001",
		Name: "测试设备001",
	}

	protocolMetrics := map[string]any{
		"response_time": 15.5,
		"error_count":   0,
		"success_rate":  100.0,
	}

	for i := 0; i < 3; i++ {
		protocolMetrics["response_time"] = 10.0 + float64(i*5)
		pdb.SaveProtocolMetrics(protocolConfig, deviceConfig, protocolMetrics)
		time.Sleep(10 * time.Millisecond)
	}

	// 写入系统指标数据
	systemTags := map[string]string{
		"host": "localhost",
		"env":  "test",
	}

	systemMetrics := map[string]any{
		"cpu_usage":    45.2,
		"memory_usage": 67.8,
		"disk_usage":   23.4,
	}

	for i := 0; i < 3; i++ {
		systemMetrics["cpu_usage"] = 40.0 + float64(i*5)
		pdb.SaveSystemMetrics("system", systemTags, systemMetrics)
		time.Sleep(10 * time.Millisecond)
	}
}

// printJSON 格式化打印JSON
func printJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "   ", "  ")
	if err != nil {
		fmt.Printf("   JSON序列化失败: %v\n", err)
		return
	}
	fmt.Printf("   %s\n", string(jsonData))
}

// DemoPebbledbUsage 演示PebbleDB使用方法的示例函数
func DemoPebbledbUsage() {
	fmt.Println("=== PebbleDB 使用演示 ===")

	ctx := context.Background()

	// 1. 创建PebbleDB实例
	pdb := NewPebbledb(ctx).(*Pebbledb)
	defer pdb.Close()

	g.Log().Infof(ctx, "PebbleDB实例创建成功")

	// 2. 查看数据库统计信息
	stats := pdb.GetStats()
	g.Log().Infof(ctx, "数据库统计信息: %+v", stats)

	// 3. 如果没有数据，先创建一些测试数据
	if stats["total_count"].(int) == 0 {
		g.Log().Infof(ctx, "数据库为空，创建测试数据...")
		setupTestData(pdb)

		// 重新获取统计信息
		stats = pdb.GetStats()
		g.Log().Infof(ctx, "创建测试数据后统计信息: %+v", stats)
	}

	// 4. 查询设备数据示例
	deviceData, err := pdb.GetDeviceData("test_device_001", 5)
	if err != nil {
		g.Log().Errorf(ctx, "查询设备数据失败: %v", err)
	} else {
		g.Log().Infof(ctx, "查询到 %d 条设备数据", len(deviceData))
		if len(deviceData) > 0 {
			latest := deviceData[0]
			g.Log().Infof(ctx, "最新数据: DeviceID=%s, Type=%s, Timestamp=%d",
				latest.DeviceID, latest.DeviceType, latest.Timestamp)
		}
	}

	// 5. 查询系统指标示例
	systemData, err := pdb.GetSystemMetrics("system", 3)
	if err != nil {
		g.Log().Errorf(ctx, "查询系统指标失败: %v", err)
	} else {
		g.Log().Infof(ctx, "查询到 %d 条系统指标数据", len(systemData))
	}

	// 6. 获取所有键名示例
	keys, err := pdb.GetAllKeys("devices/", 10)
	if err != nil {
		g.Log().Errorf(ctx, "获取键名失败: %v", err)
	} else {
		g.Log().Infof(ctx, "设备数据键名示例:")
		for i, key := range keys {
			if i < 3 { // 只显示前3个
				g.Log().Infof(ctx, "  %s", key)
			}
		}
	}

	g.Log().Infof(ctx, "PebbleDB使用演示完成")
}
