package log

import (
	"common/c_log"
	"context"
	"errors"
	"strings"

	apiv1 "application/api/log/v1"
	"common/c_enum"
	"s_db"
	"s_db/s_db_basic"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// DeleteBizLog 删除业务日志
func (c *ControllerV1) DeleteBizLog(ctx context.Context, req *apiv1.DeleteBizLogReq) (res *apiv1.DeleteBizLogRes, err error) {
	// 1. 参数验证
	if err := c.validateDeleteRequest(req); err != nil {
		c_log.Errorf(ctx, "删除日志参数验证失败: %+v", err)
		return nil, err
	}

	// 2. 获取日志服务
	logService := s_db.GetLogService()
	if logService == nil {
		c_log.Errorf(ctx, "日志服务未初始化")
		return nil, gerror.NewCode(gcode.CodeInternalError, "日志服务未初始化")
	}

	// 3. 标准化参数处理
	deleteParams := c.normalizeDeleteParams(req)

	// 4. 执行删除操作
	deleteCount, deleteErr := c.executeDeleteOperation(ctx, logService, deleteParams)
	if deleteErr != nil {
		c_log.Errorf(ctx, "删除业务日志失败: %+v", deleteErr)
		return nil, gerror.WrapCode(gcode.CodeInternalError, deleteErr, "删除业务日志失败")
	}

	// 5. 记录操作日志
	c_log.Infof(ctx, "删除业务日志成功，删除了 %d 条记录，条件: type=%s, level=%s",
		deleteCount, req.Type, req.Level)

	// 6. 返回成功响应
	return &apiv1.DeleteBizLogRes{
		Total: deleteCount,
	}, nil
}

// DeleteParams 删除参数结构体
type DeleteParams struct {
	Type  string
	Level string
}

// validateDeleteRequest 验证删除请求参数
func (c *ControllerV1) validateDeleteRequest(req *apiv1.DeleteBizLogReq) error {
	// 验证类型参数
	if req.Type != "" {
		// 检查是否为"all"（特殊值）
		if strings.EqualFold(req.Type, "all") {
			return nil
		}

		// 使用枚举验证类型
		validType := false
		switch c_enum.ESystemGroupType(req.Type) {
		case c_enum.ELogTypeEms, c_enum.ELogTypeDevice, c_enum.ELogTypeProtocol, c_enum.ELogTypePolicy, c_enum.ELogTypeAutomation:
			validType = true
		}

		if !validType {
			return errors.New("无效的业务类型，支持的类型: Ems, Device, Protocol, Policy, Automation, all")
		}
	}

	// 验证级别参数
	if req.Level != "" {
		// 检查是否为"ALL"（特殊值）
		if strings.EqualFold(req.Level, "ALL") {
			return nil
		}

		// 使用枚举验证级别
		validLevel := false
		switch c_enum.ELogLevel(strings.ToUpper(req.Level)) {
		case c_enum.Debug, c_enum.Info, c_enum.Warn, c_enum.Error:
			validLevel = true
		}

		if !validLevel {
			return errors.New("无效的日志级别，支持的级别: DEBUG, INFO, WARN, ERROR, ALL")
		}
	}

	return nil
}

// normalizeDeleteParams 标准化删除参数
func (c *ControllerV1) normalizeDeleteParams(req *apiv1.DeleteBizLogReq) DeleteParams {
	params := DeleteParams{
		Type:  req.Type,
		Level: req.Level,
	}

	// 处理"all"和"ALL"的特殊情况
	if strings.EqualFold(params.Type, "all") {
		params.Type = ""
	}
	if strings.EqualFold(params.Level, "ALL") {
		params.Level = ""
	}

	return params
}

// executeDeleteOperation 执行删除操作
func (c *ControllerV1) executeDeleteOperation(ctx context.Context, logService s_db_basic.ILogService, params DeleteParams) (int, error) {
	// 构建过滤条件
	filters := make(map[string]interface{})
	if params.Type != "" {
		filters["type"] = params.Type
	}
	if params.Level != "" {
		filters["level"] = params.Level
	}

	// 如果没有过滤条件，删除所有日志
	if len(filters) == 0 {
		return c.deleteAllBizLogs(ctx, logService)
	}

	// 使用统一的按条件删除方法
	return c.deleteLogsByFilters(ctx, logService, filters)
}

// deleteLogsByFilters 根据过滤条件删除日志（通用方法）
func (c *ControllerV1) deleteLogsByFilters(ctx context.Context, logService s_db_basic.ILogService, filters map[string]interface{}) (int, error) {
	c_log.Infof(ctx, "开始删除日志，过滤条件: %+v", filters)

	// 直接根据过滤条件删除日志记录
	deletedCount, err := logService.DeleteLogByFilters(ctx, filters)
	if err != nil {
		return 0, gerror.WrapCode(gcode.CodeInternalError, err, "根据条件删除日志失败")
	}

	if deletedCount == 0 {
		c_log.Infof(ctx, "没有找到符合条件的日志记录，过滤条件: %+v", filters)
	} else {
		c_log.Infof(ctx, "删除完成，共删除了 %d 条日志记录", deletedCount)
	}

	return deletedCount, nil
}

// deleteAllBizLogs 删除所有业务日志
func (c *ControllerV1) deleteAllBizLogs(ctx context.Context, logService s_db_basic.ILogService) (int, error) {
	c_log.Infof(ctx, "开始删除所有业务日志")

	// 执行删除所有日志
	err := logService.ClearAllLog(ctx)
	if err != nil {
		return 0, gerror.WrapCode(gcode.CodeInternalError, err, "删除所有日志失败")
	}

	// 由于ClearAllLog不返回删除数量，我们返回-1表示成功但数量未知
	c_log.Infof(ctx, "成功删除所有业务日志")
	return -1, nil
}
