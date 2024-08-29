package star_charge_100e

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"time"
)

func (s *StarCharge100EPcs) writeTime() {
	err := s._syncTime()
	if err != nil {

		// 创建一个每天凌晨3点执行的定时任务
		_, err := gcron.AddSingleton(s.ctx, "0 0 3 * * *", func(ctx context.Context) {
			e := s._syncTime()
			if e == nil {
				s.log.Infof(s.ctx, "_syncTime() success")
				//break
			}
		})
		if err != nil {
			panic(err)
		}

		go func() {
			// 每天凌晨3点同步一下时间

			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			for {
				select {

				case <-s.ctx.Done():
					s.log.Noticef(s.ctx, "writeTime() exit")
					return
				case <-ticker.C:
					err := s._syncTime()
					if err == nil {
						s.log.Infof(s.ctx, "_syncTime() success")
						//break
					}
				}
			}
		}()
	}
}

func (s *StarCharge100EPcs) _syncTime() error {
	//if !s.client.IsActivate() {
	//	return fmt.Errorf("modbus client is not activate")
	//}
	//now := time.Now()
	//
	//err := s.client.WriteMultipleRegisters(info.SyGroupTime, []int{now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second()})
	//if err != nil {
	//
	//	return err
	//}
	return nil
}
