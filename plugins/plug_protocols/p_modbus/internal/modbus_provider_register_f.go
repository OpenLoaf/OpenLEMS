package internal

import (
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"context"
	"p_base"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

// validateModbusPoint 验证Modbus点位配置是否正确
func validateModbusPoint(point *c_proto.SModbusPoint) error {
	if point == nil {
		return errors.New("point is nil")
	}

	if point.DataAccess == nil {
		return errors.Errorf("point %s: DataAccess is nil", point.GetKey())
	}
	// 使用 GoFrame 验证器验证 DataAccess 结构体
	if err := g.Validator().Data(point.DataAccess).Run(context.Background()); err != nil {
		return errors.Wrapf(err, "point %s: DataAccess validation failed", point.GetKey())
	}

	// 验证模式配置
	isBitMode := p_base.IsBitMode(point.DataAccess.BitLength, point.DataAccess.DataFormat)
	isByteMode := p_base.IsByteMode(point.DataAccess.ByteLength, point.DataAccess.DataFormat)

	if !isBitMode && !isByteMode {
		return errors.Errorf("point %s: neither bit mode nor byte mode is valid", point.GetKey())
	}

	return nil
}

func (p *ModbusProtocolProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	if task == nil {
		return
	}
	if modbusTask, ok := task.(*c_proto.SModbusPointTask); ok {
		// 验证任务配置
		err := modbusTask.Check(p.ctx)
		if err != nil {
			c_log.BizErrorf(p.ctx, "modbus task %s check failed: %v", modbusTask.GetName(), err)
			return
		}

		// 验证任务中的所有点位配置
		for _, point := range modbusTask.Points {
			if err := validateModbusPoint(point); err != nil {
				c_log.BizErrorf(p.ctx, "modbus point validation failed: %v", err)
				return
			}
		}

		p.registerReadOne(modbusTask)
	} else {
		c_log.BizErrorf(p.ctx, "modbus task type(%T) not support", task.GetName())
	}
	if len(tasks) != 0 {
		for _, t := range tasks {
			if modbusTask, ok := t.(*c_proto.SModbusPointTask); ok {
				// 验证任务配置
				err := modbusTask.Check(p.ctx)
				if err != nil {
					c_log.BizErrorf(p.ctx, "modbus task %s check failed: %v", modbusTask.GetName(), err)
					continue
				}

				// 验证任务中的所有点位配置
				hasError := false
				for _, point := range modbusTask.Points {
					if err := validateModbusPoint(point); err != nil {
						c_log.BizErrorf(p.ctx, "modbus point validation failed: %v", err)
						hasError = true
						break
					}
				}

				if !hasError {
					p.registerReadOne(modbusTask)
				}
			} else {
				c_log.BizErrorf(p.ctx, "modbus task type(%T) not support", t.GetName())
			}
		}
	}
	return
}

func (p *ModbusProtocolProvider) registerReadOne(group *c_proto.SModbusPointTask) {
	if group.Name == "" {
		panic(errors.Errorf("[%v-%v] 参数错误！modbusQuery的name为空！%+v", p.deviceId, group.Name, group))
	}

	var (
		isPermanent = !group.Transitory // 永久的查询
		name        = group.Name
	)
	ctx := context.WithValue(p.ctx, c_base.ConstCtxKeyDeviceDetail, group.Name)

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
					if p.GetProtocolStatus() != c_enum.EProtocolConnected {
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
					if p.GetProtocolStatus() != c_enum.EProtocolConnected {
						time.Sleep(3 * time.Second)
						continue
					}
					//p.modbusReadChan <- group
					_, _ = p.ReadGroupSync(group, false)
				}
			}
		}()
	}
	displayName := group.GetName()
	if displayName == "" {
		displayName = group.Name
	}

	c_log.BizInfof(ctx, "启动Modbus定时读取任务成功！任务名称：[%s] 查询周期: %0.3fs", displayName, cycle.Seconds())

}
