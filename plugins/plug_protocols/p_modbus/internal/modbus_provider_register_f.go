package internal

import (
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

// validateTask 验证任务配置和点位配置
func (p *ModbusProtocolProvider) validateTask(task *c_proto.SModbusPointTask) error {
	if task == nil {
		return errors.New("task is nil")
	}

	// 验证任务配置
	if err := task.Check(p.ctx); err != nil {
		return errors.Wrapf(err, "modbus task %s check failed", task.GetName())
	}

	// 验证任务中的所有点位配置
	for _, point := range task.Points {
		if err := validateModbusPoint(point); err != nil {
			return errors.Wrapf(err, "modbus point validation failed")
		}
	}

	return nil
}

func (p *ModbusProtocolProvider) RegisterTask(task *c_proto.SModbusPointTask, tasks ...*c_proto.SModbusPointTask) {
	// 处理第一个任务
	if task != nil {
		if err := p.validateTask(task); err != nil {
			c_log.BizError(p.ctx, "task validation failed", err)
			return
		}
		p.registerReadOne(task)
	}

	// 处理额外的任务
	for _, t := range tasks {
		if err := p.validateTask(t); err != nil {
			c_log.BizError(p.ctx, "task validation failed", err)
			continue
		}
		p.registerReadOne(t)
	}
}

func (p *ModbusProtocolProvider) registerReadOne(group *c_proto.SModbusPointTask) {
	if group.Name == "" {
		panic(errors.Errorf("[%v-%v] 参数错误！modbusQuery的name为空！%+v", p.deviceId, group.Name, group))
	}

	var (
		isPermanent = !group.Transitory // 永久的查询
		name        = group.Name
	)
	ctx := context.WithValue(p.ctx, c_enum.ELogTypeEms, group.Name)

	if p.preQuery.Contains(name) {
		// 如果存在就不再创建
		c_log.BizInfof(ctx, "[%s] 查询任务已存在！不再创建！", name)
		return
	}
	p.preQuery.Set(name, isPermanent)
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
					p.preQuery.Remove(name)
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
					p.preQuery.Remove(name)
					c_log.Debugf(ctx, "ctx取消,关闭超时的Goroutine")
					return
				case <-lifetime:
					p.preQuery.Remove(name)
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
