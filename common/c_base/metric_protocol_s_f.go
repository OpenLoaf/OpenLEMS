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
}

func NewMetricProtocol(ctx context.Context, protocolConfig *SProtocolConfig, storage IStorage) *SMetricProtocol {
	s := &SMetricProtocol{}

	if storage == nil || protocolConfig == nil {
		return s
	}

	gtimer.SetInterval(ctx, time.Minute, func(ctx context.Context) {
		s.metricRwMutex.RLock()
		// 保存数据到数据库
		result := map[string]any{
			"read_count":    s.metricMinuteReadCount,
			"success_count": s.metricMinuteReadCount - s.metricMinuteFailedCount,
			"failed_count":  s.metricMinuteFailedCount,
			"result_size":   s.metricMinuteResultSize,
		}
		_ = storage.SaveProtocolMetrics(protocolConfig, result)
		g.Log().Debugf(ctx, "保存协议[%s]的统计数据成功！统计结果为：%+v", protocolConfig.Id, result)
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

func (s *SMetricProtocol) AddMinuteResultSize(size uint32) {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteResultSize += size
}

func (s *SMetricProtocol) Clear() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.metricMinuteReadCount = 0
	s.metricMinuteFailedCount = 0
	s.metricMinuteResultSize = 0
}
