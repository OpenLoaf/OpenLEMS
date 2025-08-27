package internal

import (
	"common/c_base"
	"common/c_log"
	"common/c_proto"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

func (p *ModbusProtocolProvider) RegisterTask(task c_base.ITask, tasks ...c_base.ITask) {
	if task == nil {
		return
	}
	if modbusTask, ok := task.(*c_proto.SModbusTask); ok {
		err := modbusTask.Check(p.ctx)
		if err == nil {
			p.registerReadOne(modbusTask)
		}
	} else {
		c_log.BizErrorf(p.ctx, "modbus task type(%T) not support", task.GetName())
	}
	if len(tasks) != 0 {
		for _, t := range tasks {
			if modbusTask, ok := t.(*c_proto.SModbusTask); ok {
				err := modbusTask.Check(p.ctx)
				if err == nil {
					p.registerReadOne(modbusTask)
				}
			} else {
				c_log.BizErrorf(p.ctx, "modbus task type(%T) not support", task.GetName())
			}
		}
	}
	return
}

func (p *ModbusProtocolProvider) registerReadOne(group *c_proto.SModbusTask) {
	if group.Name == "" {
		panic(gerror.Newf("[%v-%v] 参数错误！modbusQuery的name为空！%+v", p.deviceId, group.Name, group))
	}

	var (
		isPermanent = !group.Transitory // 永久的查询
		name        = group.Name
	)
	ctx := context.WithValue(p.ctx, c_base.ConstCtxKeyDeviceDetail, group.Name)

	// 预处理一下数据
	setDefaultValue(group)

	if _, ok := p.preQuery[name]; ok {
		// 如果存在就不再创建
		c_log.BizInfof(ctx, "[%s] 查询任务已存在！不再创建！", name)
		return
	}
	p.preQuery[name] = isPermanent
	// 如果没有设置，默认为1秒一次循环

	var cycle time.Duration
	if group.CycleMill == 0 {
		cycle = time.Second
	} else {
		cycle = time.Duration(group.CycleMill * int64(time.Millisecond))
	}

	if isPermanent {
		// 永久的查询
		go func() {
			tk := time.NewTicker(cycle)
			defer tk.Stop()
			for {
				select {
				case <-ctx.Done():
					delete(p.preQuery, name)
					c_log.Debugf(ctx, "关闭永久触发查询指令的Goroutine")
					return
				case <-tk.C:
					// 如果没有连接，就延迟3秒后再执行下个周期
					if p.GetStatus() != c_base.EProtocolConnected {
						time.Sleep(3 * time.Second)
						continue
					}
					//p.modbusReadChan <- group
					_, _ = p.ReadGroupSync(group, false)
				}
			}
		}()
	} else {
		go func() {
			// 会超时的查询
			lifetime := time.After(group.TransitoryTime)
			tk := time.NewTicker(cycle)
			defer tk.Stop()
			for {
				select {
				case <-ctx.Done():
					delete(p.preQuery, name)
					c_log.Debugf(ctx, "ctx取消,关闭超时的Goroutine")
					return
				case <-lifetime:
					delete(p.preQuery, name)
					c_log.Debugf(ctx, "预读自动过期,关闭Goroutine")
					return
				case <-tk.C:
					// 这里等待一个周期(也就是跳过一个周期)exit
					//time.Sleep(cycle)
					// 如果没有连接，就延迟3秒后再执行下个周期
					if p.GetStatus() != c_base.EProtocolConnected {
						time.Sleep(3 * time.Second)
						continue
					}
					//p.modbusReadChan <- group
					_, _ = p.ReadGroupSync(group, false)
				}
			}
		}()
	}

	c_log.BizInfof(ctx, "启动Modbus定时读取任务成功！任务名称：[%s] 查询周期: %0.3fs", name, cycle.Seconds())

}

// setDefaultValue 检查点位是否规范
func setDefaultValue(group *c_proto.SModbusTask) {
	for _, meta := range group.Metas {
		// 如果倍率没有设置，就默认为1
		if meta.Factor == 0 {
			meta.Factor = 1
		}
	}
}
