package internal

import (
	"common/c_log"
	"common/c_proto"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *CanbusProtocolProvider) RegisterTask(task *c_proto.SCanbusTask, tasks ...*c_proto.SCanbusTask) {
	if task == nil {
		return
	}

	err := task.Check(c.ctx)
	if err != nil {
		c_log.BizErrorf(c.ctx, "[%s-%s] 任务注册失败！原因：%+v", c.deviceConfig.Id, task.Name, err)
		return
	}
	c.registerReadOne(task)

	// 处理额外的任务
	if len(tasks) != 0 {
		for _, t := range tasks {
			if t == nil {
				continue
			}
			err := t.Check(c.ctx)
			if err != nil {
				c_log.BizErrorf(c.ctx, "[%s-%s] 任务注册失败！原因：%+v", c.deviceConfig.Id, t.Name, err)
				continue
			}
			c.registerReadOne(t)
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
