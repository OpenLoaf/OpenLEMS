package util

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"os"
	"runtime/pprof"
	"time"
)

func GetSystemMetrics() map[string]any {
	result := map[string]any{}
	// 系统在线时长
	if uptime, err := host.Uptime(); err == nil {
		result["uptime_minute"] = time.Duration(uptime) * time.Minute
	}

	// 获取 CPU 总使用率
	if percent, err := cpu.Percent(0, false); err == nil {
		result["cpu"] = percent[0]
	}
	if percent, err := cpu.Percent(0, true); err == nil {
		for i, f := range percent {
			result[fmt.Sprintf("cpu_%d", i)] = f
		}
	}

	if memory, err := mem.VirtualMemory(); err == nil {
		result["mem_total_mb"] = memory.Total / 1024 / 1024
		result["mem_available_mb"] = memory.Available / 1024 / 1024
		result["mem_used_mb"] = memory.Used / 1024 / 1024
		result["mem_used_percent"] = memory.UsedPercent
	}

	if counters, err := net.IOCounters(false); err == nil {
		for _, counter := range counters {
			result[fmt.Sprintf("net_%s_sent_mb", counter.Name)] = counter.BytesSent / 1024 / 1024
			result[fmt.Sprintf("net_%s_recv_mb", counter.Name)] = counter.BytesRecv / 1024 / 1024
		}
	}

	if avg, err := load.Avg(); err == nil {
		result["load_1min"] = avg.Load1
		result["load_5min"] = avg.Load5
		result["load_15min"] = avg.Load15
	}

	if usageStat, err := disk.Usage("/"); err == nil {
		result["disk_total_mb"] = usageStat.Total / 1024 / 1024
		result["disk_free_mb"] = usageStat.Free / 1024 / 1024
		result["disk_used_mb"] = usageStat.Used / 1024 / 1024
		result["disk_used_percent"] = usageStat.UsedPercent
	}

	return result
}

func GetProcessInfo() map[string]any {
	result := map[string]any{}

	if p, err := process.NewProcess(int32(os.Getpid())); err == nil {
		if processCpuPercent, err := p.CPUPercent(); err == nil {
			result["cpu_percent"] = processCpuPercent
		}
		if processMemoryPercent, err := p.MemoryPercent(); err == nil {
			result["memory_percent"] = processMemoryPercent
		}
		if processMemoryInfo, err := p.MemoryInfo(); err == nil {
			result["memory_rss_mb"] = processMemoryInfo.RSS / 1024 / 1024 // 物理内存
		}
		if threads, err := p.NumThreads(); err == nil {
			result["threads"] = threads
		}
	}
	// 获取pprof是否启动
	isPprofEnabled := g.Config().MustGet(context.Background(), "server.pprofEnabled").Bool()

	if isPprofEnabled {
		// 获取堆使用情况
		heapStats := pprof.Lookup("heap")
		if heapStats != nil {
			result["heap_alloc"] = heapStats.Count()
		}

		// 获取goroutine数量
		goroutineStats := pprof.Lookup("goroutine")
		if goroutineStats != nil {
			result["goroutine_count"] = goroutineStats.Count()
		}

		// 获取线程创建情况
		threadCreateStats := pprof.Lookup("threadcreate")
		if threadCreateStats != nil {
			result["thread_create_count"] = threadCreateStats.Count()
		}

		// 获取阻塞分析
		blockStats := pprof.Lookup("block")
		if blockStats != nil {
			result["block_count"] = blockStats.Count()
		}
	}

	return result
}

func GetSystemInfo() map[string]string {
	result := map[string]string{}
	if info, err := host.Info(); err == nil {
		result["hostname"] = info.Hostname
		result["os"] = info.OS
		result["platform"] = info.Platform
		result["platform_family"] = info.PlatformFamily
		result["platform_version"] = info.PlatformVersion
		result["kernel_version"] = info.KernelVersion
		//result["virtualization_system"] = info.VirtualizationSystem
		//result["virtualization_role"] = info.VirtualizationRole
	}

	result["pid"] = fmt.Sprintf("%d", os.Getpid())

	if bootTime, err := host.BootTime(); err == nil {
		result["boot_time"] = time.Unix(int64(bootTime), 0).Format("2006-01-02 15:04:05")
	}
	return result
}
