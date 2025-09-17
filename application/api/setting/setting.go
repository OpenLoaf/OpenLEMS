// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package setting

import (
	"context"

	v1 "application/api/setting/v1"
)

type ISettingV1 interface {
	GetSetting(ctx context.Context, req *v1.GetSettingReq) (res *v1.GetSettingRes, err error)
	GetSettingDetail(ctx context.Context, req *v1.GetSettingDetailReq) (res *v1.GetSettingDetailRes, err error)
	UpdateSetting(ctx context.Context, req *v1.UpdateSettingReq) (res *v1.UpdateSettingRes, err error)
}
