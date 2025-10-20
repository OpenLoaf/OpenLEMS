package price

import (
	v1 "application/api/price/v1"
	"context"
	"s_db"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetPriceDetail 获取电价详情
func (c *Controller) GetPriceDetail(ctx context.Context, req *v1.GetPriceDetailReq) (res *v1.GetPriceDetailRes, err error) {
	// 参数验证
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, err.FirstError()
	}

	// 调用服务层
	model, err := s_db.GetPriceService().GetPriceById(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取电价详情失败")
	}

	// 使用 gconv.Scan 直接转换
	var dto v1.Price
	if err := gconv.Scan(model, &dto); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "转换电价DTO失败")
	}

	return &v1.GetPriceDetailRes{
		Price: &dto,
	}, nil
}
