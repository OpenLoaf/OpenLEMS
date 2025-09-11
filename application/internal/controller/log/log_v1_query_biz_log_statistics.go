package log

import (
	"context"

	apiv1 "application/api/log/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// QueryBizLogStatistics 查询日志统计信息
func (c *ControllerV1) QueryBizLogStatistics(ctx context.Context, req *apiv1.QueryBizLogStatisticsReq) (res *apiv1.QueryBizLogStatisticsRes, err error) {
	// 1. 按类型统计
	typeStats := struct {
		Ems      int `json:"ems" dc:"ems类型日志数量"`
		Device   int `json:"device" dc:"device类型日志数量"`
		Protocol int `json:"protocol" dc:"protocol类型日志数量"`
		Policy   int `json:"policy" dc:"policy类型日志数量"`
	}{}

	// 统计 EMS 类型日志数量
	emsCount, err := g.Model("log").Ctx(ctx).Where("type", "ems").Count()
	if err != nil {
		g.Log().Errorf(ctx, "统计EMS类型日志失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "统计EMS类型日志失败")
	}
	typeStats.Ems = emsCount

	// 统计 Device 类型日志数量
	deviceCount, err := g.Model("log").Ctx(ctx).Where("type", "device").Count()
	if err != nil {
		g.Log().Errorf(ctx, "统计Device类型日志失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "统计Device类型日志失败")
	}
	typeStats.Device = deviceCount

	// 统计 Protocol 类型日志数量
	protocolCount, err := g.Model("log").Ctx(ctx).Where("type", "protocol").Count()
	if err != nil {
		g.Log().Errorf(ctx, "统计Protocol类型日志失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "统计Protocol类型日志失败")
	}
	typeStats.Protocol = protocolCount

	// 统计 Policy 类型日志数量
	policyCount, err := g.Model("log").Ctx(ctx).Where("type", "policy").Count()
	if err != nil {
		g.Log().Errorf(ctx, "统计Policy类型日志失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "统计Policy类型日志失败")
	}
	typeStats.Policy = policyCount

	// 2. 按等级统计
	levelStats := struct {
		DEBUG int `json:"DEBUG" dc:"DEBUG等级日志数量"`
		INFO  int `json:"INFO" dc:"INFO等级日志数量"`
		WARN  int `json:"WARN" dc:"WARN等级日志数量"`
		ERROR int `json:"ERROR" dc:"ERROR等级日志数量"`
	}{}

	// 统计 DEBUG 等级日志数量
	debugCount, err := g.Model("log").Ctx(ctx).Where("level", "DEBUG").Count()
	if err != nil {
		g.Log().Errorf(ctx, "统计DEBUG等级日志失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "统计DEBUG等级日志失败")
	}
	levelStats.DEBUG = debugCount

	// 统计 INFO 等级日志数量
	infoCount, err := g.Model("log").Ctx(ctx).Where("level", "INFO").Count()
	if err != nil {
		g.Log().Errorf(ctx, "统计INFO等级日志失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "统计INFO等级日志失败")
	}
	levelStats.INFO = infoCount

	// 统计 WARN 等级日志数量
	warnCount, err := g.Model("log").Ctx(ctx).Where("level", "WARN").Count()
	if err != nil {
		g.Log().Errorf(ctx, "统计WARN等级日志失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "统计WARN等级日志失败")
	}
	levelStats.WARN = warnCount

	// 统计 ERROR 等级日志数量
	errorCount, err := g.Model("log").Ctx(ctx).Where("level", "ERROR").Count()
	if err != nil {
		g.Log().Errorf(ctx, "统计ERROR等级日志失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "统计ERROR等级日志失败")
	}
	levelStats.ERROR = errorCount

	// 3. 构建响应
	return &apiv1.QueryBizLogStatisticsRes{
		Type:  typeStats,
		Level: levelStats,
	}, nil
}
