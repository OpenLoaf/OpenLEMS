package public

import "core/c_window_counter/internal"

// IWindowCounter 滑动窗口计数器接口
// 用于统计指定时间窗口内的请求数量，计算QPS等指标
type IWindowCounter = internal.IWindowCounter

// SWindowCounterStats 计数器统计信息
type SWindowCounterStats = internal.SWindowCounterStats
