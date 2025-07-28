package internal

import (
	"common"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *CanbusProtocolProvider) Init() {
	c.once.Do(func() {
		device := common.GetDeviceById(c.deviceConfig.Id)
		if device != nil {
			c.deviceType = device.GetDriverType()
		}

		go func() {
			for {
				select {
				case <-c.ctx.Done():
					return

				case frame := <-c.receiverChan:

					task, ok := c.canTaskMap[frame.ID]
					if !ok {
						continue
					}
					// 解析数据
					c.analysisCanbus(task, frame)

					g.Log().Infof(c.ctx, "内部接收到数据：%v", frame)
				}
			}
		}()

	})

}

func (c *CanbusProtocolProvider) Close() {
	//c.cache.Clear(c.Ctx)
}
