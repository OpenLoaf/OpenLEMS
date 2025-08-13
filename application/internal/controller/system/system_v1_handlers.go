package system

import (
	"context"
	stdnet "net"
	"os/exec"
	"runtime"
	"strings"
	"time"

	ntv1 "application/api/network/v1"
	v1 "application/api/system/v1"
	ntctl "application/internal/controller/network"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	psnet "github.com/shirou/gopsutil/v4/net"
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
	return v1.NetBrief{UpBytes: upBytes, DownBytes: downBytes, UpBps: 0, DownBps: 0}, nil
}

func fetchIfaces(ctx context.Context) ([]v1.IfaceItem, error) {
	// 直接调用网络模块控制器的公共方法
	res, err := ntctl.NewV1().GetNetworkInterfaceList(ctx, &ntv1.GetNetworkInterfaceListReq{OnlyEthernet: false})
	if err != nil {
		return nil, err
	}
	var items []v1.IfaceItem
	for _, it := range res.Interfaces {
		items = append(items, v1.IfaceItem{Name: it.Name, IP: it.IPv4, Up: it.Connected})
	}
	return items, nil
}

func net2Interface(name string) (*stdnet.Interface, error) { return stdnet.InterfaceByName(name) }

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
	start := time.Now()
	g.Log().Infof(ctx, "[SystemSummary] start")
	t := time.Now()
	cpuI, _ := fetchCPU(ctx)
	g.Log().Infof(ctx, "[SystemSummary] cpu=%s", time.Since(t))
	t = time.Now()
	memI, _ := fetchMemory(ctx)
	g.Log().Infof(ctx, "[SystemSummary] memory=%s", time.Since(t))
	t = time.Now()
	diskI, _ := fetchDisk(ctx)
	g.Log().Infof(ctx, "[SystemSummary] disk=%s", time.Since(t))
	t = time.Now()
	upI, _ := fetchUptime(ctx)
	g.Log().Infof(ctx, "[SystemSummary] uptime=%s", time.Since(t))
	t = time.Now()
	netI, _ := fetchNet(ctx)
	g.Log().Infof(ctx, "[SystemSummary] net=%s", time.Since(t))
	t = time.Now()
	timeI, _ := fetchTime(ctx)
	g.Log().Infof(ctx, "[SystemSummary] time=%s", time.Since(t))
	t = time.Now()
	sysI, _ := fetchSys(ctx)
	g.Log().Infof(ctx, "[SystemSummary] sys=%s", time.Since(t))
	g.Log().Infof(ctx, "[SystemSummary] done cost=%s", time.Since(start))
	return &v1.GetSummaryRes{CPU: cpuI, Memory: memI, Disk: diskI, Uptime: upI, Net: netI, Time: timeI, Sys: sysI}, nil
}

func (c *ControllerV1) GetNetworkTraffic(ctx context.Context, req *v1.GetNetworkTrafficReq) (res *v1.GetNetworkTrafficRes, err error) {
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
