package internal

import (
	"common/c_modbus"
)

func (p *ModbusProtocolProvider) ProtocolListen() {
	// 只会执行一次监听
	p.once.Do(func() {
		go func(c chan *c_modbus.SModbusTask) {
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
