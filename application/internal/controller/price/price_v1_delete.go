package price

import (
	v1 "application/api/price/v1"
	"context"
	"s_db"
	"s_price"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// DeletePrice 删除电价
func (c *Controller) DeletePrice(ctx context.Context, req *v1.DeletePriceReq) (res *v1.DeletePriceRes, err error) {
	// 参数验证
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, err.FirstError()
	}

	// 调用服务层
	err = s_db.GetPriceService().DeletePrice(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "删除电价失败")
	}

	// 刷新缓存
	if err := s_price.RefreshPriceCache(ctx); err != nil {
		g.Log().Errorf(ctx, "刷新电价缓存失败: %+v", err)
	}

	g.Log().Infof(ctx, "成功删除电价 - ID: %d", req.Id)
	return &v1.DeletePriceRes{}, nil
}
