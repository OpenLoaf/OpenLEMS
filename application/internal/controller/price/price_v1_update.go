package price

import (
	v1 "application/api/price/v1"
	"context"
	"s_db"
	"s_db/s_db_model"
	"s_price"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// UpdatePrice 更新电价
func (c *Controller) UpdatePrice(ctx context.Context, req *v1.UpdatePriceReq) (res *v1.UpdatePriceRes, err error) {
	// 参数验证
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, err.FirstError()
	}

	// 构造Model
	model := &s_db_model.SPriceModel{
		Id:            req.Id,
		Description:   req.Description,
		Priority:      req.Priority,
		Status:        req.Status,
		DateRange:     gconv.String(req.DateRange),
		TimeRange:     gconv.String(req.TimeRange),
		PriceSegments: gconv.String(req.PriceSegments),
		RemoteId:      gconv.String(req.RemoteId),
	}

	// 调用服务层
	err = s_db.GetPriceService().UpdatePrice(ctx, model)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新电价失败")
	}

	// 刷新缓存
	if err := s_price.RefreshPriceCache(ctx); err != nil {
		g.Log().Errorf(ctx, "刷新电价缓存失败: %+v", err)
	}

	g.Log().Infof(ctx, "成功更新电价 - ID: %d", req.Id)
	return &v1.UpdatePriceRes{}, nil
}
