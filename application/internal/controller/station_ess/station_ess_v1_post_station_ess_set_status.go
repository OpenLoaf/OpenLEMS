package station_ess

import (
	"application/internal/service"
	"context"

	"application/api/station_ess/v1"
)

func (c *ControllerV1) PostStationEssSetStatus(ctx context.Context, req *v1.PostStationEssSetStatusReq) (res *v1.PostStationEssSetStatusRes, err error) {
	err = service.Station().SetEnergyStoreStatus(req.Status)
	return &v1.PostStationEssSetStatusRes{}, err
}
