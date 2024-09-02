package station_ess

import (
	"application/internal/service"
	"context"

	"application/api/station_ess/v1"
)

func (c *ControllerV1) PostStationEssSetPower(ctx context.Context, req *v1.PostStationEssSetPowerReq) (res *v1.PostStationEssSetPowerRes, err error) {
	err = service.Station().SetEnergyStorePower(req.Power)
	return &v1.PostStationEssSetPowerRes{}, err
}
