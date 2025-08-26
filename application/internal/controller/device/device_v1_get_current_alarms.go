package device

import (
	v1 "application/api/device/v1"
	"context"
)

func (c *ControllerV1) GetCurrentAlarms(ctx context.Context, req *v1.GetCurrentAlarmsReq) (res *v1.GetCurrentAlarmsRes, err error) {

	return &v1.GetCurrentAlarmsRes{CurrentTotal: 0, Items: []v1.AlarmItem{}}, nil
}
