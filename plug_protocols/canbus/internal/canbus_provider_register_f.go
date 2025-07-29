package internal

import (
	"canbus/p_canbus"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *CanbusProtocolProvider) RegisterRead(group *p_canbus.SCanbusTask, gs ...*p_canbus.SCanbusTask) {
	if group == nil {
		return
	}
	group.Check()
	c.registerReadOne(group)
	if len(gs) != 0 {
		for _, q := range gs {
			c.registerReadOne(q)
		}
	}
}

func (c *CanbusProtocolProvider) registerReadOne(group *p_canbus.SCanbusTask) {
	if group.Name == "" {
		panic(gerror.Newf("[%v-%v] 参数错误！modbusQuery的name为空！%+v", c.deviceConfig.Id, group.Name, group))
	}
	c.canTaskMap[group.CanbusID] = group

}
