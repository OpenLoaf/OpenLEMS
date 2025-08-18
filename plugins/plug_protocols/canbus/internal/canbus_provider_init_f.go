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
					g.Log().Debugf(c.ctx, "收到canbus 数据: %v", frame)

					for _, task := range c.canTaskList {
						if task.GetCanbusID == nil {
							continue
						}
						c.analysisCanbus(task, frame)

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
