package internal

import (
	"common/c_base"
	"common/c_log"
	"common/c_proto"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *CanbusProtocolProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	if task == nil {
		return
	}

	// 将 IPointTask 转换为 *c_proto.SCanbusTask
	canbusTask, ok := task.(*c_proto.SCanbusTask)
	if !ok {
		c_log.BizErrorf(c.ctx, "[%s-%s] 任务类型转换失败！期望类型：*c_proto.SCanbusTask，实际类型：%T", c.deviceConfig.Id, task.GetName(), task)
		return
	}

	err := canbusTask.Check(c.ctx)
	if err != nil {
		c_log.BizErrorf(c.ctx, "[%s-%s] 任务注册失败！原因：%+v", c.deviceConfig.Id, canbusTask.Name, err)
		return
	}
	c.registerReadOne(canbusTask)

	// 处理额外的任务
	if len(tasks) != 0 {
		for _, t := range tasks {
			if t == nil {
				continue
			}
			canbusTask, ok := t.(*c_proto.SCanbusTask)
			if !ok {
				c_log.BizErrorf(c.ctx, "[%s-%s] 任务类型转换失败！期望类型：*c_proto.SCanbusTask，实际类型：%T", c.deviceConfig.Id, t.GetName(), t)
				continue
			}
			err := canbusTask.Check(c.ctx)
			if err != nil {
				c_log.BizErrorf(c.ctx, "[%s-%s] 任务注册失败！原因：%+v", c.deviceConfig.Id, canbusTask.Name, err)
				continue
			}
			c.registerReadOne(canbusTask)
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
