package internal

import (
	"common"
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
				}
			}
		}()

	})

}

func (c *CanbusProtocolProvider) Close() {
	//c.cache.Clear(c.Ctx)
}
