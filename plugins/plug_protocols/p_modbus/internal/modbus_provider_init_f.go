package internal

import (
	"common/c_log"
	"common/c_proto"
)

func (p *ModbusProtocolProvider) ProtocolListen() {
	// 只会执行一次监听
	p.once.Do(func() {
		go func(c chan *c_proto.SModbusTask) {
			for {
				select {
				case <-p.ctx.Done():
					// d
					err := p.client.Close()
					if err != nil {
						c_log.BizErrorf(p.ctx, "关闭modbus client 失败！失败原因：%+v", err)
					}
					c_log.Noticef(p.ctx, "关闭消息查询Goroutine")
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
