package internal

import (
	"canbus/p_canbus"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *CanbusProtocolProvider) RegisterRead(ctx context.Context, group *p_canbus.SCanbusTask, gs ...*p_canbus.SCanbusTask) {
	if group == nil {
		return
	}
	group.Check()
	c.registerReadOne(ctx, group)
	if len(gs) != 0 {
		for _, q := range gs {
			c.registerReadOne(ctx, q)
		}
	}
}

func (c *CanbusProtocolProvider) registerReadOne(ctx context.Context, group *p_canbus.SCanbusTask) {
	if group.Name == "" {
		panic(gerror.Newf("[%v-%v] 参数错误！modbusQuery的name为空！%+v", c.deviceConfig.Id, group.Name, group))
	}
	c.canTaskMap[group.CanbusID] = group
}
