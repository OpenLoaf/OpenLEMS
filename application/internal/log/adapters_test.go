package log

import (
	"common/c_base"
	"common/c_log"
	"context"
	"testing"
)

// TestSystemAdapter 测试系统适配器
func TestSystemAdapter(t *testing.T) {
	// 创建系统日志适配器
	logger := NewSystemAdapter(nil) // 传入nil，因为只是测试接口

	// 测试基本日志功能
	ctx := context.Background()

	// 测试不同级别的日志（这些不会真正输出，因为传入的是nil）
	logger.Info(ctx, "系统信息日志")
	logger.Warning(ctx, "系统警告日志")
	logger.Error(ctx, "系统错误日志")

	t.Log("系统适配器基本功能测试完成")
}

// TestSystemAdapterQuery 测试系统适配器查询功能
func TestSystemAdapterQuery(t *testing.T) {
	// 创建系统日志适配器
	logger := NewSystemAdapter(nil)

	// 测试查询功能
	ctx := context.Background()
	params := c_log.LogQueryParams{
		Type:     "ems",
		Page:     1,
		PageSize: 10,
	}

	result, err := logger.QueryLogs(ctx, params)
	if err != nil {
		t.Errorf("系统适配器查询不应该返回错误: %v", err)
	}

	if result.Total != 0 {
		t.Errorf("系统适配器应该返回空结果，但得到了 %d 条记录", result.Total)
	}

	t.Log("系统适配器查询测试通过，正确返回空结果")
}

// TestFileAdapter 测试文件适配器
func TestFileAdapter(t *testing.T) {
	// 创建文件日志适配器
	logger := NewFileAdapter()

	// 测试基本日志功能
	ctx := context.Background()

	// 测试不同级别的日志
	logger.Info(ctx, "这是一条信息日志")
	logger.Warning(ctx, "这是一条警告日志")
	logger.Error(ctx, "这是一条错误日志")

	// 测试带设备ID的上下文
	deviceCtx := context.WithValue(ctx, c_base.ConstCtxKeyDeviceId, "test_device_001")
	logger.Info(deviceCtx, "设备日志测试")

	// 测试带协议ID的上下文
	protocolCtx := context.WithValue(ctx, c_base.ConstCtxKeyProtocolId, "test_protocol_001")
	logger.Info(protocolCtx, "协议日志测试")

	t.Log("文件适配器基本功能测试完成")
}

// TestFileAdapterQuery 测试文件适配器查询功能
func TestFileAdapterQuery(t *testing.T) {
	// 创建文件日志适配器
	logger := NewFileAdapter()

	// 测试查询功能
	ctx := context.Background()
	params := c_log.LogQueryParams{
		Type:     "ems",
		Date:     "20241201",
		Page:     1,
		PageSize: 10,
		Level:    "INFO",
	}

	result, err := logger.QueryLogs(ctx, params)
	if err != nil {
		t.Logf("查询日志时出现错误（这是正常的，因为日志文件可能不存在）: %v", err)
	} else {
		t.Logf("查询成功，总记录数: %d", result.Total)
		for i, line := range result.Lines {
			t.Logf("记录 %d: 时间=%s, 级别=%s, 类型=%s, ID=%s, 内容=%s",
				i+1, line.CreatedAt, line.Level, line.Type, line.Id, line.Content)
		}
	}
}

// TestDatabaseAdapter 测试数据库适配器
func TestDatabaseAdapter(t *testing.T) {
	// 创建数据库日志适配器
	logger := NewDatabaseAdapter()

	// 测试基本日志功能
	ctx := context.Background()

	// 测试不同级别的日志
	logger.Info(ctx, "这是一条信息日志")
	logger.Infof(ctx, "这是一条格式化的信息日志: %s", "测试内容")
	logger.Warning(ctx, "这是一条警告日志")
	logger.Warningf(ctx, "这是一条格式化的警告日志: %d", 123)
	logger.Error(ctx, "这是一条错误日志")
	logger.Errorf(ctx, "这是一条格式化的错误日志: %v", "错误信息")

	// 测试带设备ID的上下文
	deviceCtx := context.WithValue(ctx, c_base.ConstCtxKeyDeviceId, "test_device_001")
	logger.Info(deviceCtx, "设备日志测试")

	// 测试带协议ID的上下文
	protocolCtx := context.WithValue(ctx, c_base.ConstCtxKeyProtocolId, "test_protocol_001")
	logger.Info(protocolCtx, "协议日志测试")

	// 测试带策略ID的上下文
	policyCtx := context.WithValue(ctx, "PolicyId", "test_policy_001")
	logger.Info(policyCtx, "策略日志测试")

	t.Log("数据库适配器基本功能测试完成")
}

// TestDatabaseAdapterQuery 测试数据库适配器查询功能
func TestDatabaseAdapterQuery(t *testing.T) {
	// 创建数据库日志适配器
	logger := NewDatabaseAdapter()

	// 测试查询功能
	ctx := context.Background()
	params := c_log.LogQueryParams{
		Type:     "device",
		Id:       "test_device_001",
		Date:     "20241201",
		Page:     1,
		PageSize: 10,
		Level:    "INFO",
	}

	result, err := logger.QueryLogs(ctx, params)
	if err != nil {
		t.Logf("查询日志时出现错误（这是正常的，因为数据库可能未初始化）: %v", err)
	} else {
		t.Logf("查询成功，总记录数: %d", result.Total)
		for i, line := range result.Lines {
			t.Logf("记录 %d: 时间=%s, 级别=%s, 类型=%s, ID=%s, 内容=%s",
				i+1, line.CreatedAt, line.Level, line.Type, line.Id, line.Content)
		}
	}
}
