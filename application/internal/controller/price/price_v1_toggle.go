package price

import (
	v1 "application/api/price/v1"
	"common/c_log"
	"context"
	"s_db"
	"s_price"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// TogglePrice 启用/停用电价
func (c *Controller) TogglePrice(ctx context.Context, req *v1.TogglePriceReq) (res *v1.TogglePriceRes, err error) {
	// 参数验证
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, err.FirstError()
	}

	// 先获取现有数据
	model, err := s_db.GetPriceService().GetPriceById(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取电价信息失败")
	}

	// 更新状态
	model.Status = req.Status

	// 调用服务层
	err = s_db.GetPriceService().UpdatePrice(ctx, model)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新电价状态失败")
	}

	// 刷新缓存
	if err := s_price.RefreshPriceCache(ctx); err != nil {
		c_log.Errorf(ctx, "刷新电价缓存失败: %+v", err)
	}

	c_log.Infof(ctx, "成功更新电价状态 - ID: %d, 状态: %s", req.Id, req.Status)
	return &v1.TogglePriceRes{}, nil
}
