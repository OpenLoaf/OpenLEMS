package internal

import (
	"common/c_base"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *CanbusProtocolProvider) ProtocolListen() {
	c.once.Do(func() {
		go func() {
			for {
				select {
				case <-c.ctx.Done():
					return

				case frame := <-c.receiverChan: // 接收canbus数据
					g.Log().Debugf(c.ctx, "收到canbus 数据: %v task长度:%v", frame, len(c.canTaskList))

					for _, task := range c.canTaskList {
						canId := task.GetCanbusID(c.deviceConfig.Params)
						if canId == nil {
							//g.Log().Debugf(c.ctx, "获取到的canID为空 %s", task.Name)
							continue
						}
						//g.Log().Debugf(c.ctx, "当前的task ID: 0x%X 比对的ID:0x%X", *canId, frame.ID)
						if *canId == frame.ID {
							// 同一个帧ID，在一个设备下，只会响应一个Task
							c.analysisCanbus(task, frame)
							break
						}
						/*if task.IDMatch != nil && task.IDMatch(frame.ID) {
							// 如果有IDMatch 并且匹配上的话，执行解析
							c.analysisCanbus(task, frame)
						} else if task.GetCanbusID() == frame.ID {
							c.analysisCanbus(task, frame)
						}*/
					}
				}
			}
		}()
		c.protocolState = c_base.EStateRunning
	})

}
