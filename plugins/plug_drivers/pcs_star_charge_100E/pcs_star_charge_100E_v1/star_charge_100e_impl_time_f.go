package pcs_star_charge_100E_v1

import (
	"common/c_log"
	"common/c_timer"
	"context"
	"time"
)

func (s *sPcsStarCharge100E) writeTime() {
	err := s._syncTime()
	if err != nil {
		// 每天凌晨3点执行一次
		scheduleDaily3AM := func() {}
		scheduleDaily3AM = func() {
			d := next3AMDuration(time.Now())
			c_timer.SetTimeout(s.DeviceCtx, d, func(ctx context.Context) {
				if e := s._syncTime(); e == nil {
					c_log.Infof(s.DeviceCtx, "_syncTime() success (daily)")
				}
				// 继续安排下一次
				scheduleDaily3AM()
			})
		}
		scheduleDaily3AM()

		// 失败后每5秒重试，成功后取消
		var cancelRetry func()
		cancelRetry = c_timer.SetInterval(s.DeviceCtx, 5*time.Second, func(ctx context.Context) {
			if e := s._syncTime(); e == nil {
				c_log.Infof(s.DeviceCtx, "_syncTime() success (retry)")
				cancelRetry()
			}
		})
	}
}

// 计算从当前时间到下一次凌晨3点的等待时长
func next3AMDuration(now time.Time) time.Duration {
	loc := now.Location()
	next := time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, loc)
	if !next.After(now) {
		next = next.Add(24 * time.Hour)
	}
	return next.Sub(now)
}

func (s *sPcsStarCharge100E) _syncTime() error {
	//if !s.client.IsActivate() {
	//	return gerror.Newf("modbus client is not activate")
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
