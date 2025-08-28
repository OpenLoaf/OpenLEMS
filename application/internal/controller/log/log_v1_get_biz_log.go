package log

import (
	apiv1 "application/api/log/v1"
	"common/c_log"

	"github.com/gogf/gf/v2/frame/g"
)

// ControllerV1 由 log_new.go 的 NewV1 返回
// func New() *ControllerV1 { return &ControllerV1{} }

func (c *ControllerV1) GetBizLog(ctx g.Ctx, req *apiv1.GetBizLogReq) (res *apiv1.GetBizLogRes, err error) {
	// 构建查询参数
	params := c_log.LogQueryParams{
		Type:     req.Type,
		Id:       req.Id,
		Date:     req.Date,
		Page:     req.Page,
		PageSize: req.PageSize,
		Level:    req.Level,
	}

	// 使用业务日志查询接口
	result, err := c_log.BizQueryLogs(ctx, params)
	if err != nil {
		return nil, err
	}

	// 转换为API响应格式
	lines := make([]apiv1.LogLine, 0, len(result.Lines))
	for _, line := range result.Lines {
		lines = append(lines, apiv1.LogLine{
			Id:        line.Id,
			Type:      line.Type,
			Level:     line.Level,
			Content:   line.Content,
			CreatedAt: line.CreatedAt,
		})
	}

	return &apiv1.GetBizLogRes{
		Total: result.Total,
		Lines: lines,
	}, nil
}
