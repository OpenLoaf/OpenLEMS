package internal

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

// GetSystemMetrics 获取系统指标
func GetSystemMetrics() map[string]any {
	result := map[string]any{}

	// 系统在线时长
	if uptime, err := host.Uptime(); err == nil {
		result[MetricUptimeMinute] = float64(uptime)
	}

	// 获取 CPU 总使用率
	if percent, err := cpu.Percent(0, false); err == nil {
		result[MetricCpu] = percent[0]
	}

	if memory, err := mem.VirtualMemory(); err == nil {
		result[MetricMemTotalMB] = float64(memory.Total / 1024 / 1024)
		result[MetricMemAvailableMB] = float64(memory.Available / 1024 / 1024)
		result[MetricMemUsedMB] = float64(memory.Used / 1024 / 1024)
		result[MetricMemUsedPercent] = memory.UsedPercent
	}

	if counters, err := net.IOCounters(false); err == nil {
		for _, counter := range counters {
			result[fmt.Sprintf("net_%s_sent_mb", counter.Name)] = float64(counter.BytesSent / 1024 / 1024)
			result[fmt.Sprintf("net_%s_recv_mb", counter.Name)] = float64(counter.BytesRecv / 1024 / 1024)
		}
	}

	if avg, err := load.Avg(); err == nil {
		result[MetricLoad1Min] = avg.Load1
		result[MetricLoad5Min] = avg.Load5
		result[MetricLoad15Min] = avg.Load15
	}

	if usageStat, err := disk.Usage("/"); err == nil {
		result[MetricDiskTotalMB] = float64(usageStat.Total / 1024 / 1024)
		result[MetricDiskFreeMB] = float64(usageStat.Free / 1024 / 1024)
		result[MetricDiskUsedMB] = float64(usageStat.Used / 1024 / 1024)
		result[MetricDiskUsedPercent] = float64(usageStat.UsedPercent)
	}

	return result
}

// GetProcessInfo 获取进程信息
func GetProcessInfo() map[string]any {
	result := map[string]any{}

	// 现有进程指标
	if p, err := process.NewProcess(int32(os.Getpid())); err == nil {
		if processCpuPercent, err := p.CPUPercent(); err == nil {
			result[MetricProcessCpuPercent] = processCpuPercent
		}
		if processMemoryPercent, err := p.MemoryPercent(); err == nil {
			result[MetricProcessMemoryPercent] = processMemoryPercent
		}
	}

	// 新增：Go runtime 指标
	result[MetricGoroutineCount] = float64(runtime.NumGoroutine())

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	result[MetricHeapAllocMB] = float64(memStats.HeapAlloc / 1024 / 1024)
	result[MetricHeapSysMB] = float64(memStats.HeapSys / 1024 / 1024)
	result[MetricGCCount] = float64(memStats.NumGC)

	return result
}

// GetSystemInfo 获取系统信息
func GetSystemInfo() map[string]string {
	result := map[string]string{}
	if info, err := host.Info(); err == nil {
		result["hostname"] = info.Hostname
		result["os"] = info.OS
		result["platform"] = info.Platform
	}

	result["pid"] = fmt.Sprintf("%d", os.Getpid())

	if bootTime, err := host.BootTime(); err == nil {
		result["boot_time"] = time.Unix(int64(bootTime), 0).Format("2006-01-02_15:04:05")
	}
	return result
}
