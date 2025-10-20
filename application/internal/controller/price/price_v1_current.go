package price

import (
	v1 "application/api/price/v1"
	"context"
	"s_price"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetCurrentPrice 获取当前激活电价
func (c *Controller) GetCurrentPrice(ctx context.Context, req *v1.GetCurrentPriceReq) (res *v1.GetCurrentPriceRes, err error) {
	// 获取当前激活的电价
	activePrice, err := s_price.GetCurrentActivePrice(ctx)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取当前激活电价失败")
	}

	if activePrice == nil {
		return &v1.GetCurrentPriceRes{Price: nil}, nil
	}

	// 使用 gconv.Scan 直接转换
	var dto v1.Price
	if err := gconv.Scan(activePrice, &dto); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "转换电价DTO失败")
	}

	return &v1.GetCurrentPriceRes{
		Price: &dto,
	}, nil
}
