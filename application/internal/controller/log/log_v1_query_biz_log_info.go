package log

import (
	"common/c_log"
	"context"

	apiv1 "application/api/log/v1"
	"s_db/s_db_model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// QueryBizLogInfo 查询日志信息
func (c *ControllerV1) QueryBizLogInfo(ctx context.Context, req *apiv1.QueryBizLogInfoReq) (res *apiv1.QueryBizLogInfoRes, err error) {
	// 1. 获取日志总数
	var logModel s_db_model.SLogModel
	total, err := logModel.GetCount(ctx)
	if err != nil {
		c_log.Errorf(ctx, "获取日志总数失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取日志总数失败")
	}

	// 2. 获取第一条日志时间
	var firstLogTime string
	if total > 0 {
		var firstLog s_db_model.SLogModel
		err = g.Model("log").Ctx(ctx).Order("created_at ASC").Limit(1).Scan(&firstLog)
		if err != nil {
			c_log.Errorf(ctx, "获取第一条日志时间失败: %+v", err)
			return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取第一条日志时间失败")
		}
		if firstLog.CreatedAt != nil {
			firstLogTime = firstLog.CreatedAt.Format("2006-01-02 15:04:05")
		}
	}

	// 3. 获取最新日志时间
	var latestLogTime string
	if total > 0 {
		var latestLog s_db_model.SLogModel
		err = g.Model("log").Ctx(ctx).Order("created_at DESC").Limit(1).Scan(&latestLog)
		if err != nil {
			c_log.Errorf(ctx, "获取最新日志时间失败: %+v", err)
			return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取最新日志时间失败")
		}
		if latestLog.CreatedAt != nil {
			latestLogTime = latestLog.CreatedAt.Format("2006-01-02 15:04:05")
		}
	}

	// 4. 构建响应
	return &apiv1.QueryBizLogInfoRes{
		Total:         total,
		FirstLogTime:  firstLogTime,
		LatestLogTime: latestLogTime,
	}, nil
}
