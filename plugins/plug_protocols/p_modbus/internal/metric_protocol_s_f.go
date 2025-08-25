package internal

import (
	"common"
	"common/c_base"
	"common/c_log"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"reflect"
	"sync"
	"time"
)

type sMetricProtocol struct {
	metricRwMutex           sync.RWMutex // 读写锁
	metricMinuteReadCount   uint32       // 分钟读取次数
	metricMinuteFailedCount uint32       // 分钟失败次数
	metricMinuteResultSize  uint32       // 分钟读取到的数据量

	maxWaitTime     time.Duration // 最大读取时间
	maxWaitTaskName string        // 最大读取任务名称
}

func newMetricProtocol(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) *sMetricProtocol {

	if protocolConfig == nil {
		panic("协议配置不能为空！")
	}
	s := &sMetricProtocol{}

	gtimer.SetInterval(ctx, time.Minute, func(ctx context.Context) {
		s.metricRwMutex.Lock()
		defer s.metricRwMutex.Unlock()
		// 保存数据到数据库
		result := map[string]any{
			"total":         s.metricMinuteReadCount,
			"success":       s.metricMinuteReadCount - s.metricMinuteFailedCount,
			"failed":        s.metricMinuteFailedCount,
			"result_size":   s.metricMinuteResultSize,
			"max_wait_ms":   s.maxWaitTime.Milliseconds(),
			"max_task_name": s.maxWaitTaskName,
		}
		g.Log().Debugf(ctx, "保存协议[%s]的统计数据，统计结果为：%+v", protocolConfig.Id, result)
		storage := common.GetStorageInstance()
		if storage == nil || reflect.ValueOf(storage).IsNil() {
			g.Log().Debugf(ctx, "没有找到存储实例，无法保存协议[%s]的统计数据，跳过本次存储！", protocolConfig.Id)
			return
		}

		err := storage.SaveProtocolMetrics(protocolConfig, deviceConfig, result)
		if err != nil {
			c_log.BizErrorf(ctx, "保存协议[%s]的统计数据失败！统计结果为：%+v 异常原因：%v", protocolConfig.Id, result, err)
		} else {
			g.Log().Debugf(ctx, "保存协议[%s]的统计数据成功！统计结果为：%+v", protocolConfig.Id, result)
		}

		s.metricMinuteReadCount = 0
		s.metricMinuteFailedCount = 0
		s.metricMinuteResultSize = 0
		s.maxWaitTime = 0
		s.maxWaitTaskName = ""
	})

	return s
}

func (s *sMetricProtocol) AddMinuteReadCount() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteReadCount++
}

func (s *sMetricProtocol) AddMinuteFailedCount() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteFailedCount++
}

func (s *sMetricProtocol) AddMinuteResultSize(size int) {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteResultSize += uint32(size)
}

// CalcReadTime 计算读取时间
func (s *sMetricProtocol) CalcReadTime(taskName string, useTime time.Duration) {
	if useTime < time.Millisecond {
		// 小于1毫秒的不统计
		return
	}
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	if useTime > s.maxWaitTime {
		s.maxWaitTime = useTime
		s.maxWaitTaskName = taskName
	}
}
