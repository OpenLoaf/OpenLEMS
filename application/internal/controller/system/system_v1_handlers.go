package system

import (
	"context"
	"os/exec"
	"runtime"
	"strings"
	"time"

	v1 "application/api/system/v1"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	psnet "github.com/shirou/gopsutil/v4/net"
)

// 非阻塞网络速率缓存（在进程内做差）
var (
	netPrevInit bool
	netPrev     psnet.IOCountersStat
	netPrevTime time.Time
)

// 单项采集函数
func fetchCPU(ctx context.Context) (v1.CPUInfo, error) {
	// 无阻塞获取CPU百分比
	p, _ := cpu.PercentWithContext(ctx, 0, false)
	la, _ := load.AvgWithContext(ctx)
	var usage float64
	if len(p) > 0 {
		usage = p[0]
	}
	if la == nil {
		la = &load.AvgStat{}
	}
	return v1.CPUInfo{Usage: usage, Load1: la.Load1, Load5: la.Load5, Load15: la.Load15}, nil
}

func fetchMemory(ctx context.Context) (v1.MemoryInfo, error) {
	v, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return v1.MemoryInfo{}, err
	}
	toGB := func(b uint64) float64 { return float64(b) / (1024 * 1024 * 1024) }
	used := v.Total - v.Available
	return v1.MemoryInfo{Usage: v.UsedPercent, TotalGB: toGB(v.Total), UsedGB: toGB(used), FreeGB: toGB(v.Available)}, nil
}

func fetchDisk(ctx context.Context) (v1.DiskInfo, error) {
	u, err := disk.UsageWithContext(ctx, "/")
	if err != nil {
		return v1.DiskInfo{}, err
	}
	toGB := func(b uint64) float64 { return float64(b) / (1024 * 1024 * 1024) }
	return v1.DiskInfo{Usage: u.UsedPercent, TotalGB: toGB(u.Total), UsedGB: toGB(u.Used), FreeGB: toGB(u.Free)}, nil
}

func fetchUptime(ctx context.Context) (v1.UptimeInfo, error) {
	hi, err := host.InfoWithContext(ctx)
	if err != nil {
		return v1.UptimeInfo{}, err
	}
	boot := time.Unix(int64(hi.BootTime), 0)
	dur := time.Since(boot)
	days := int(dur / (24 * time.Hour))
	hours := int((dur % (24 * time.Hour)) / time.Hour)
	minutes := int((dur % time.Hour) / time.Minute)
	return v1.UptimeInfo{Day: days, Hour: hours, Minute: minutes, BootTime: boot.Format("2006-01-02 15:04:05")}, nil
}

func fetchNet(ctx context.Context) (v1.NetBrief, error) {
	// 非阻塞速率：与上次采样做差
	cur, _ := psnet.IOCountersWithContext(ctx, false)
	if len(cur) == 0 {
		return v1.NetBrief{}, nil
	}
	upBytes := cur[0].BytesSent
	downBytes := cur[0].BytesRecv
	if !netPrevInit {
		netPrevInit = true
		netPrev = cur[0]
		netPrevTime = time.Now()
		return v1.NetBrief{UpBytes: upBytes, DownBytes: downBytes, UpBps: 0, DownBps: 0}, nil
	}
	dur := time.Since(netPrevTime).Seconds()
	if dur <= 0 {
		dur = 1
	}
	upBps := float64(upBytes-netPrev.BytesSent) / dur
	downBps := float64(downBytes-netPrev.BytesRecv) / dur
	netPrev = cur[0]
	netPrevTime = time.Now()
	return v1.NetBrief{UpBytes: upBytes, DownBytes: downBytes, UpBps: upBps, DownBps: downBps}, nil
}

// 删除冗余：接口列表由网络模块独立提供

func fetchTime(ctx context.Context) (v1.TimeInfo, error) {
	now := time.Now()
	loc, _ := time.LoadLocation(now.Location().String())
	tz := loc.String()
	// NTP 状态：优先命令查询
	running := false
	server := ""
	if runtime.GOOS == "linux" {
		if out, err := exec.Command("sh", "-c", "systemctl is-active --quiet systemd-timesyncd && echo running").Output(); err == nil {
			running = strings.Contains(string(out), "running")
		}
		if out, err := exec.Command("sh", "-c", "timedatectl show -p NTP -p NTPSynchronized").Output(); err == nil {
			_ = out
		}
		// 常见配置文件
		if out, err := exec.Command("sh", "-c", "grep -Eo '^(server|pool)\\s+[^#]+' /etc/ntp.conf /etc/chrony.conf 2>/dev/null | head -1 | awk '{print $2}'").Output(); err == nil {
			server = strings.TrimSpace(string(out))
		}
	} else if runtime.GOOS == "darwin" {
		if out, err := exec.Command("sh", "-c", "systemsetup -getusingnetworktime").Output(); err == nil {
			running = strings.Contains(strings.ToLower(string(out)), "on")
		}
		if out, err := exec.Command("sh", "-c", "systemsetup -getnetworktimeserver").Output(); err == nil {
			parts := strings.Split(strings.TrimSpace(string(out)), ": ")
			if len(parts) == 2 {
				server = parts[1]
			}
		}
	}
	return v1.TimeInfo{Now: now.Format("2006-01-02 15:04:05"), Timezone: tz, NTPRunning: running, NTPServer: server}, nil
}

func fetchSys(ctx context.Context) (v1.SysInfo, error) {
	hi, err := host.InfoWithContext(ctx)
	if err != nil {
		return v1.SysInfo{}, err
	}
	kernel := hi.KernelVersion
	return v1.SysInfo{OSName: hi.Platform, OSVersion: hi.PlatformVersion, Kernel: kernel}, nil
}

// 汇总与单项接口实现
func (c *ControllerV1) GetSummary(ctx context.Context, req *v1.GetSummaryReq) (res *v1.GetSummaryRes, err error) {
	cpuI, _ := fetchCPU(ctx)
	memI, _ := fetchMemory(ctx)
	diskI, _ := fetchDisk(ctx)
	upI, _ := fetchUptime(ctx)
	timeI, _ := fetchTime(ctx)
	sysI, _ := fetchSys(ctx)
	return &v1.GetSummaryRes{CPU: cpuI, Memory: memI, Disk: diskI, Uptime: upI, Time: timeI, Sys: sysI}, nil
}

func (c *ControllerV1) GetNetworkTraffic(ctx context.Context, req *v1.GetNetworkTrafficReq) (res *v1.GetNetworkTrafficRes, err error) {
	if !netPrevInit {
		// 首次采样，等待500ms再做第二次统计，直接返回有效速率
		cur, _ := psnet.IOCountersWithContext(ctx, false)
		if len(cur) == 0 {
			return &v1.GetNetworkTrafficRes{Net: v1.NetBrief{}}, nil
		}
		netPrev = cur[0]
		netPrevTime = time.Now()
		time.Sleep(500 * time.Millisecond)
	}
	netI, _ := fetchNet(ctx)
	return &v1.GetNetworkTrafficRes{Net: netI}, nil
}

func (c *ControllerV1) GetTimeInfo(ctx context.Context, req *v1.GetTimeInfoReq) (res *v1.GetTimeInfoRes, err error) {
	timeI, _ := fetchTime(ctx)
	return &v1.GetTimeInfoRes{Time: timeI}, nil
}

func (c *ControllerV1) GetSystemInfo(ctx context.Context, req *v1.GetSystemInfoReq) (res *v1.GetSystemInfoRes, err error) {
	sysI, _ := fetchSys(ctx)
	return &v1.GetSystemInfoRes{Sys: sysI}, nil
}
