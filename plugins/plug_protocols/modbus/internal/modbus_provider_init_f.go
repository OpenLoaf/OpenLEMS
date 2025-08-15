package internal

import (
	"common"
	"modbus/p_modbus"
)

func (p *ModbusProtocolProvider) Init() {
	// 只会执行一次监听
	p.once.Do(func() {
		device := common.GetRunningDeviceById(p.deviceConfig.Id)
		if device != nil {
			p.deviceType = device.GetDriverType()
		}

		go func(c chan *p_modbus.SModbusTask) {
			for {
				select {
				case <-p.ctx.Done():
					// d
					err := p.client.Close()
					if err != nil {
						p.log.Errorf(p.ctx, "关闭modbus client 失败！失败原因：%+v", err)
					}
					p.log.Noticef(p.ctx, "关闭消息查询Goroutine")
					return
				case query := <-c:
					/*			if !p.client.IsConnected() {
								continue
							}*/
					_, _ = p.ReadGroupSync(query, false)
				}
			}
		}(p.modbusReadChan)

	})
}

func (p *ModbusProtocolProvider) Close() {
	//p.cache.Clear(p.ctx)
}
