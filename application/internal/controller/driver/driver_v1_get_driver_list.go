package driver

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "application/api/driver/v1"

	"services"
)

func (c *ControllerV1) GetDriverList(ctx context.Context, req *v1.GetDriverListReq) (res *v1.GetDriverListRes, err error) {
	driverManager := services.GetDriverManager()
	driverList := driverManager.GetAllDriverNames()

	for _, driverName := range driverList {
		driverInfo, err := driverManager.GetDriverInfo(ctx, driverName)
		if err != nil {
			continue
		}
		fmt.Println(driverInfo)
	}
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
