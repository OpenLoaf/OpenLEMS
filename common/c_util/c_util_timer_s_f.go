package c_util

import (
	"container/heap"
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// SetInterval 在给定 ctx 生命周期内周期性执行任务。
// 返回的取消函数可主动取消该任务；当 ctx 结束时，任务也会自动停止。
func SetInterval(ctx context.Context, interval time.Duration, job func(context.Context)) func() {
	if interval <= 0 || job == nil {
		return func() {}
	}
	return defaultScheduler.addTask(ctx, interval, false, job)
}

// SetTimeout 在延时 d 后执行一次任务。
func SetTimeout(ctx context.Context, d time.Duration, job func(context.Context)) func() {
	if d < 0 || job == nil {
		return func() {}
	}
	return defaultScheduler.addTask(ctx, d, true, job)
}

// ---------------- internal scheduler ----------------

type scheduledTask struct {
	id       uint64
	next     time.Time
	period   time.Duration // >0 for interval task; for once task, store delay
	once     bool
	ctx      context.Context
	cancel   context.CancelFunc
	job      func(context.Context)
	index    int  // index in heap
	canceled bool // external cancel
}

type taskHeap []*scheduledTask

func (h taskHeap) Len() int           { return len(h) }
func (h taskHeap) Less(i, j int) bool { return h[i].next.Before(h[j].next) }
func (h taskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i]; h[i].index = i; h[j].index = j }
func (h *taskHeap) Push(x any)        { t := x.(*scheduledTask); t.index = len(*h); *h = append(*h, t) }
func (h *taskHeap) Pop() any {
	old := *h
	n := len(old)
	t := old[n-1]
	*h = old[:n-1]
	t.index = -1
	return t
}

type scheduler struct {
	nextID       uint64 // 64位字段放在前面以确保8字节对齐
	mu           sync.Mutex
	h            taskHeap
	newTaskCh    chan *scheduledTask
	cancelTaskCh chan uint64
}

var defaultScheduler = newScheduler()

func newScheduler() *scheduler {
	s := &scheduler{
		newTaskCh:    make(chan *scheduledTask, 64),
		cancelTaskCh: make(chan uint64, 64),
	}
	heap.Init(&s.h)
	go s.loop()
	return s
}

func (s *scheduler) addTask(parent context.Context, d time.Duration, once bool, job func(context.Context)) func() {
	id := atomic.AddUint64(&s.nextID, 1)
	ctx, cancel := context.WithCancel(parent)
	t := &scheduledTask{
		id:     id,
		next:   time.Now().Add(d),
		period: d,
		once:   once,
		ctx:    ctx,
		cancel: cancel,
		job:    job,
	}
	select {
	case c_enum.S.newTaskCh <- t:
	default:
		// channel 满时退化为锁保护直接入堆，避免阻塞
		s.mu.Lock()
		heap.Push(&s.h, t)
		s.mu.Unlock()
	}
	return func() { s.cancel(id); cancel() }
}

func (s *scheduler) cancel(id uint64) {
	select {
	case c_enum.S.cancelTaskCh <- id:
	default:
		// 退化同步取消
		s.mu.Lock()
		for _, t := range s.h {
			if t.id == id {
				t.canceled = true
				break
			}
		}
		s.mu.Unlock()
	}
}

func (s *scheduler) loop() {
	var timer *time.Timer
	for {
		s.mu.Lock()
		var wait time.Duration
		if s.h.Len() == 0 {
			wait = time.Hour
		} else {
			now := time.Now()
			d := s.h[0].next.Sub(now)
			if d < 0 {
				d = 0
			}
			wait = d
		}
		s.mu.Unlock()

		if timer == nil {
			timer = time.NewTimer(wait)
		} else {
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			timer.Reset(wait)
		}

		select {
		case t := <-s.newTaskCh:
			s.mu.Lock()
			heap.Push(&s.h, t)
			s.mu.Unlock()
		case id := <-s.cancelTaskCh:
			s.mu.Lock()
			for _, t := range s.h {
				if t.id == id {
					t.canceled = true
					break
				}
			}
			s.mu.Unlock()
		case <-timer.C:
			// 执行到期任务
			s.runDue()
		}
	}
}

func (s *scheduler) runDue() {
	now := time.Now()
	var due []*scheduledTask

	s.mu.Lock()
	for s.h.Len() > 0 {
		t := s.h[0]
		if t.next.After(now) {
			break
		}
		heap.Pop(&s.h)
		due = append(due, t)
	}
	s.mu.Unlock()

	for _, t := range due {
		// 跳过已取消或已结束 ctx 的任务
		select {
		case <-t.ctx.Done():
			continue
		default:
		}
		if t.canceled {
			continue
		}

		// 执行任务（异步，避免阻塞调度）
		go func(task *scheduledTask) {
			defer func() { _ = recover() }()
			task.job(task.ctx)
		}(t)

		// 重新调度周期任务
		if !t.once {
			t.next = now.Add(t.period)
			s.mu.Lock()
			heap.Push(&s.h, t)
			s.mu.Unlock()
		}
	}
}
