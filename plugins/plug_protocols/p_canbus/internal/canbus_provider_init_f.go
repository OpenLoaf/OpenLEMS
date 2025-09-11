package internal

import (
	"common/c_enum"
	"common/c_log"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *CanbusProtocolProvider) ProtocolListen() {
	c.once.Do(func() {
		go func() {
			for {
				select {
				case <-c.ctx.Done():
					c_log.BizInfof(c.ctx, "设备关闭！")
					return

				case frame := <-c.receiverChan: // 接收canbus数据
					g.Log().Debugf(c.ctx, "收到canbus 数据: %v task长度:%v", frame, len(c.canTaskList))

					for _, task := range c.canTaskList {
						canId := task.GetCanbusID(c.deviceConfig.Params)
						//g.Log().Debugf(c.ctx, "当前的task ID: 0x%X 比对的ID:0x%X", canId, frame.ID)
						if canId == frame.ID {
							// 同一个帧ID，在一个设备下，只会响应一个Task
							if err := c.analysisCanbus(task, frame); err != nil {
								g.Log().Errorf(c.ctx, "解析 CANbus 数据失败 task:%s error:%v", task.Name, err)
							}

							now := time.Now()
							c.lastUpdateTime = &now
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
	})
	c.protocolStatus = c_enum.EProtocolConnected
}
