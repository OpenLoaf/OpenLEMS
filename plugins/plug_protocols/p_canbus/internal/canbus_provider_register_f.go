package internal

import (
	"common/c_log"
	"common/c_proto"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *CanbusProtocolProvider) RegisterCanbusTask(group *c_proto.SCanbusTask, gs ...*c_proto.SCanbusTask) {
	if group == nil {
		return
	}
	err := group.Check(c.ctx)
	if err != nil {
		c_log.BizErrorf(c.ctx, "[%s-%s] 任务注册失败！原因：%+v", c.deviceConfig.Id, group.Name, err)
		return
	}
	c.registerReadOne(group)
	if len(gs) != 0 {
		for _, q := range gs {
			c.registerReadOne(q)
		}
	}
}

func (c *CanbusProtocolProvider) registerReadOne(group *c_proto.SCanbusTask) {
	if group.Name == "" {
		panic(gerror.Newf("[%s-%s] 参数错误！modbusQuery的name为空！%+v", c.deviceConfig.Id, group.Name, group))
	}
	//c.canTaskMap[group.GetCanbusID] = group
	c.canTaskList = append(c.canTaskList, group)
}
