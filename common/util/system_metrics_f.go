package util

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"time"
)

func GetSystemMetrics() map[string]any {
	result := map[string]any{}
	// 开始时间
	if uptime, err := host.Uptime(); err == nil {
		result["uptime_minute"] = time.Duration(uptime) * time.Minute
	}

	// 获取 CPU 总使用率
	if percent, err := cpu.Percent(0, false); err != nil {
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
	return result
}
