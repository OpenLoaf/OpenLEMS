// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package price

import (
	"context"

	v1 "application/api/price/v1"
)

type IPriceV1 interface {
	CreatePrice(ctx context.Context, req *v1.CreatePriceReq) (res *v1.CreatePriceRes, err error)
	UpdatePrice(ctx context.Context, req *v1.UpdatePriceReq) (res *v1.UpdatePriceRes, err error)
	DeletePrice(ctx context.Context, req *v1.DeletePriceReq) (res *v1.DeletePriceRes, err error)
	GetPriceList(ctx context.Context, req *v1.GetPriceListReq) (res *v1.GetPriceListRes, err error)
	GetPriceDetail(ctx context.Context, req *v1.GetPriceDetailReq) (res *v1.GetPriceDetailRes, err error)
	TogglePrice(ctx context.Context, req *v1.TogglePriceReq) (res *v1.TogglePriceRes, err error)
	GetCurrentPrice(ctx context.Context, req *v1.GetCurrentPriceReq) (res *v1.GetCurrentPriceRes, err error)
}
