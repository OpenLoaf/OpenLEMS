package c_base

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"sync"
	"time"
)

type SMetricProtocol struct {
	metricRwMutex           sync.RWMutex // 读写锁
	metricMinuteReadCount   uint32       // 分钟读取次数
	metricMinuteFailedCount uint32       // 分钟失败次数
	metricMinuteResultSize  uint32       // 分钟读取到的数据量

	maxWaitTime     time.Duration // 最大读取时间
	maxWaitTaskName string        // 最大读取任务名称
}

func NewMetricProtocol(ctx context.Context, protocolConfig *SProtocolConfig, storage IStorage) *SMetricProtocol {

	if protocolConfig == nil {
		panic("协议配置不能为空！")
	}
	s := &SMetricProtocol{}

	gtimer.SetInterval(ctx, time.Minute, func(ctx context.Context) {
		s.metricRwMutex.RLock()
		// 保存数据到数据库
		result := map[string]any{
			"total":         s.metricMinuteReadCount,
			"success":       s.metricMinuteReadCount - s.metricMinuteFailedCount,
			"failed":        s.metricMinuteFailedCount,
			"result_size":   s.metricMinuteResultSize,
			"max_wait_ms":   s.maxWaitTime.Milliseconds(),
			"max_task_name": s.maxWaitTaskName,
		}
		err := storage.SaveProtocolMetrics(protocolConfig, result)
		if err != nil {
			g.Log().Errorf(ctx, "保存协议[%s]的统计数据失败！统计结果为：%+v", protocolConfig.Id, result)
		} else {
			g.Log().Debugf(ctx, "保存协议[%s]的统计数据成功！统计结果为：%+v", protocolConfig.Id, result)
		}

		s.metricRwMutex.RUnlock()
		s.Clear()
	})

	return s
}

func (s *SMetricProtocol) AddMinuteReadCount() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteReadCount++
}

func (s *SMetricProtocol) AddMinuteFailedCount() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteFailedCount++
}

func (s *SMetricProtocol) AddMinuteResultSize(size int) {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteResultSize += uint32(size)
}

func (s *SMetricProtocol) Clear() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteReadCount = 0
	s.metricMinuteFailedCount = 0
	s.metricMinuteResultSize = 0
	s.maxWaitTime = 0
	s.maxWaitTaskName = ""
}

// CalcReadTime 计算读取时间
func (s *SMetricProtocol) CalcReadTime(taskName string, useTime time.Duration) {
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
