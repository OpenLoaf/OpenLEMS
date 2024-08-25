package protocol

import (
	"ems-plan/c_config"
	"plug_protocol_modbus/p_modbus"
	"time"
)

func (p *ModbusProvider) Start() error {
	// 只会执行一次监听
	p.once.Do(func() {

		go func(c chan *p_modbus.ModbusGroup) {
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

		// 打印日志
		if p.PrintCacheValue {
			go func() {

				ticker := time.NewTicker(c_config.GetSystemConfig().GetPrintCacheValueCycleDuration())
				defer ticker.Stop()
				for {
					select {
					case <-p.ctx.Done():
						p.log.Noticef(p.ctx, "关闭打印缓存值的Goroutine")
						return
					case <-ticker.C:
						p.PrintCacheValues()
					}
				}
			}()
		}

	})
}
