package internal

import (
	"common"
	"common/c_base"
	"common/c_log"
	"context"
	"reflect"
	"s_storage"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
)

type sMetricProtocol struct {
	metricRwMutex sync.RWMutex // 读写锁

	// 现有字段
	metricMinuteReadCount   uint32        // 分钟读取次数
	metricMinuteFailedCount uint32        // 分钟失败次数
	metricMinuteResultSize  uint32        // 分钟读取到的数据量
	maxWaitTime             time.Duration // 最大读取时间
	maxWaitTaskName         string        // 最大读取任务名称

	// 新增性能指标
	totalResponseTime int64         // 总响应时间（纳秒）
	minWaitTime       time.Duration // 最小读取时间

	// 新增可靠性指标
	consecutiveFailures uint32 // 当前连续失败次数
	timeoutCount        uint32 // 超时次数
	reconnectCount      uint32 // 重连次数

}

func newMetricProtocol(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) *sMetricProtocol {

	if protocolConfig == nil {
		panic("协议配置不能为空！")
	}
	s := &sMetricProtocol{
		minWaitTime: time.Hour * 24, // 初始值设大
	}

	gtimer.SetInterval(ctx, time.Minute, func(ctx context.Context) {
		s.metricRwMutex.Lock()
		defer s.metricRwMutex.Unlock()
		// 计算平均响应时间
		var avgResponseMs int64
		if s.metricMinuteReadCount > 0 {
			avgResponseMs = s.totalResponseTime / int64(s.metricMinuteReadCount) / int64(time.Millisecond)
		}

		// 计算成功率
		var successRate float64
		if s.metricMinuteReadCount > 0 {
			successRate = float64(s.metricMinuteReadCount-s.metricMinuteFailedCount) / float64(s.metricMinuteReadCount) * 100
		}

		// 计算平均字节数
		var avgBytesPerRequest float64
		if s.metricMinuteReadCount > 0 {
			avgBytesPerRequest = float64(s.metricMinuteResultSize) / float64(s.metricMinuteReadCount)
		}

		// 保存数据到数据库
		result := map[string]any{
			// 现有指标
			s_storage.ProtocolMetricTotal:       s.metricMinuteReadCount,
			s_storage.ProtocolMetricSuccess:     s.metricMinuteReadCount - s.metricMinuteFailedCount,
			s_storage.ProtocolMetricFailed:      s.metricMinuteFailedCount,
			s_storage.ProtocolMetricResultSize:  s.metricMinuteResultSize,
			s_storage.ProtocolMetricMaxWaitMs:   s.maxWaitTime.Milliseconds(),
			s_storage.ProtocolMetricMaxTaskName: s.maxWaitTaskName,

			// 新增性能指标
			s_storage.ProtocolMetricAvgResponseMs: avgResponseMs,
			s_storage.ProtocolMetricMinWaitMs:     s.minWaitTime.Milliseconds(),

			// 新增可靠性指标
			s_storage.ProtocolMetricSuccessRate:         successRate,
			s_storage.ProtocolMetricConsecutiveFailures: s.consecutiveFailures,
			s_storage.ProtocolMetricTimeoutCount:        s.timeoutCount,
			s_storage.ProtocolMetricReconnectCount:      s.reconnectCount,

			// 新增吞吐量指标
			s_storage.ProtocolMetricAvgBytesPerRequest: avgBytesPerRequest,
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
		s.totalResponseTime = 0
		s.minWaitTime = time.Hour * 24
		s.timeoutCount = 0
		s.reconnectCount = 0
		// 注意: consecutiveFailures 不重置，因为它是当前状态
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
	s.consecutiveFailures++
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

	// 累加总响应时间
	s.totalResponseTime += int64(useTime)

	// 更新最大响应时间
	if useTime > s.maxWaitTime {
		s.maxWaitTime = useTime
		s.maxWaitTaskName = taskName
	}

	// 更新最小响应时间
	if useTime < s.minWaitTime {
		s.minWaitTime = useTime
	}

}

// ResetConsecutiveFailures 重置连续失败计数
func (s *sMetricProtocol) ResetConsecutiveFailures() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.consecutiveFailures = 0
}

// AddTimeoutCount 增加超时次数
func (s *sMetricProtocol) AddTimeoutCount() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.timeoutCount++
}

// AddReconnectCount 增加重连次数
func (s *sMetricProtocol) AddReconnectCount() {
	s.metricRwMutex.Lock()
	defer s.metricRwMutex.Unlock()
	s.reconnectCount++
}
