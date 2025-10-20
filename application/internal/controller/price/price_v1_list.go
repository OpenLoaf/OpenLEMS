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

// GetPriceList 获取电价列表
func (c *Controller) GetPriceList(ctx context.Context, req *v1.GetPriceListReq) (res *v1.GetPriceListRes, err error) {
	// 参数验证
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, err.FirstError()
	}

	// 构造过滤条件
	filters := gconv.Map(req)
	delete(filters, "page")
	delete(filters, "pageSize")

	// 移除 nil 值
	for key, value := range filters {
		if value == nil {
			delete(filters, key)
		}
	}

	// 调用服务层
	prices, total, err := s_db.GetPriceService().GetPricePage(ctx, req.Page, req.PageSize, filters)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取电价列表失败")
	}

	// 使用 gconv.Scan 直接转换列表
	var priceList []*v1.Price
	if err := gconv.Scan(prices, &priceList); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "转换电价列表失败")
	}

	return &v1.GetPriceListRes{
		List:  priceList,
		Total: total,
	}, nil
}
