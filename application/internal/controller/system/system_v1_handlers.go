package system

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"s_db"
	"s_db/s_db_basic"
	"strings"
	"time"

	"github.com/pkg/errors"

	v1 "application/api/system/v1"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	psnet "github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

var rebootApplyExpireAt time.Time

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
	pm, _ := fetchProcMemory(ctx)
	return v1.MemoryInfo{Usage: v.UsedPercent, TotalGB: toGB(v.Total), UsedGB: toGB(used), FreeGB: toGB(v.Available), ProcUsage: pm.Usage, ProcRSSMB: pm.RSSMB}, nil
}

// fetchProcMemory 返回本服务进程内存信息
func fetchProcMemory(ctx context.Context) (v1.ProcMemoryInfo, error) {
	p, err := process.NewProcessWithContext(ctx, int32(os.Getpid()))
	if err != nil {
		return v1.ProcMemoryInfo{}, err
	}
	pm, err := p.MemoryInfoWithContext(ctx)
	if err != nil {
		return v1.ProcMemoryInfo{}, err
	}
	vm, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil || vm.Total == 0 {
		return v1.ProcMemoryInfo{
			Usage: 0,
			RSSMB: float64(pm.RSS) / (1024 * 1024),
			VMSMB: float64(pm.VMS) / (1024 * 1024),
		}, nil
	}
	usage := (float64(pm.RSS) / float64(vm.Total)) * 100
	return v1.ProcMemoryInfo{
		Usage: usage,
		RSSMB: float64(pm.RSS) / (1024 * 1024),
		VMSMB: float64(pm.VMS) / (1024 * 1024),
	}, nil
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
	procUptimeSeconds := 0
	procStartTime := ""
	if p, err := process.NewProcessWithContext(ctx, int32(os.Getpid())); err == nil {
		if ctms, err := p.CreateTimeWithContext(ctx); err == nil {
			start := time.Unix(0, ctms*int64(time.Millisecond))
			if start.Before(time.Now()) {
				procUptimeSeconds = int(time.Since(start).Seconds())
				procStartTime = start.Format("2006-01-02 15:04:05")
			}
		}
	}
	return v1.UptimeInfo{Day: days, Hour: hours, Minute: minutes, BootTime: boot.Format("2006-01-02 15:04:05"), ProcUptimeSeconds: procUptimeSeconds, ProcStartTime: procStartTime}, nil
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
	return v1.SysInfo{Hostname: hi.Hostname, OSName: hi.Platform, OSVersion: hi.PlatformVersion, Kernel: kernel}, nil
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
		// 首次采样，等待100ms再做第二次统计，直接返回有效速率
		cur, _ := psnet.IOCountersWithContext(ctx, false)
		if len(cur) == 0 {
			return &v1.GetNetworkTrafficRes{Net: v1.NetBrief{}}, nil
		}
		netPrev = cur[0]
		netPrevTime = time.Now()
		time.Sleep(100 * time.Millisecond)
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

func (c *ControllerV1) GetProcessMemory(ctx context.Context, req *v1.GetProcessMemoryReq) (res *v1.GetProcessMemoryRes, err error) {
	pm, _ := fetchProcMemory(ctx)
	return &v1.GetProcessMemoryRes{ProcessMemory: pm}, nil
}

// GetNow 返回当前系统时间
func (c *ControllerV1) GetNow(ctx context.Context, req *v1.GetNowReq) (res *v1.GetNowRes, err error) {
	return &v1.GetNowRes{Now: time.Now().Format("2006-01-02 15:04:05")}, nil
}

// UpdateHostname 修改主机名
func (c *ControllerV1) UpdateHostname(ctx context.Context, req *v1.UpdateHostnameReq) (res *v1.UpdateHostnameRes, err error) {
	switch runtime.GOOS {
	case "linux":
		if out, err := exec.Command("hostnamectl", "set-hostname", req.Hostname).CombinedOutput(); err != nil {
			return nil, errors.Errorf("set-hostname failed: %v, %s", err, string(out))
		}
	case "darwin":
		args := []string{"-c", fmt.Sprintf("scutil --set HostName '%s'; scutil --set LocalHostName '%s'; scutil --set ComputerName '%s'", req.Hostname, req.Hostname, req.Hostname)}
		if out, err := exec.Command("sh", args...).CombinedOutput(); err != nil {
			return nil, errors.Errorf("set hostname failed: %v, %s", err, string(out))
		}
	default:
		return nil, errors.Errorf("unsupported OS")
	}
	return &v1.UpdateHostnameRes{}, nil
}

// UpdateSystemTime 修改系统时间/时区
func (c *ControllerV1) UpdateSystemTime(ctx context.Context, req *v1.UpdateSystemTimeReq) (res *v1.UpdateSystemTimeRes, err error) {
	if req.Timezone != "" {
		switch runtime.GOOS {
		case "linux":
			if out, err := exec.Command("timedatectl", "set-timezone", req.Timezone).CombinedOutput(); err != nil {
				return nil, errors.Errorf("set-timezone failed: %v, %s", err, string(out))
			}
		case "darwin":
			if out, err := exec.Command("systemsetup", "-settimezone", req.Timezone).CombinedOutput(); err != nil {
				return nil, errors.Errorf("settimezone failed: %v, %s", err, string(out))
			}
		}
	}

	if req.Time != "" {
		switch runtime.GOOS {
		case "linux":
			if out, err := exec.Command("date", "-s", req.Time).CombinedOutput(); err != nil {
				return nil, errors.Errorf("set time failed: %v, %s", err, string(out))
			}
			// 同步到硬件时钟
			exec.Command("hwclock", "-w").Run()
		case "darwin":
			// macOS: 使用 systemsetup -setdate/-settime, 需要拆分
			// 尝试使用 date 命令: [[MMDDHHmm[[CC]YY][.ss]]]
			// 将 2006-01-02 15:04:05 转换为 010215042006.05
			// 简化：调用 shell 用 date -f 解析
			script := fmt.Sprintf("date -j -f '%s' '%s' +%s >/dev/null && sudo sntp -sS 2>/dev/null || true", "%Y-%m-%d %H:%M:%S", req.Time, "%Y%m%d%H%M.%S")
			if out, err := exec.Command("sh", "-c", script).CombinedOutput(); err != nil {
				_ = out // 某些系统无 sntp
			}
		}
	}
	return &v1.UpdateSystemTimeRes{}, nil
}

// RebootApply 设置30s内有效的重启令牌
func (c *ControllerV1) RebootApply(ctx context.Context, req *v1.RebootApplyReq) (res *v1.RebootApplyRes, err error) {
	rebootApplyExpireAt = time.Now().Add(30 * time.Second)
	return &v1.RebootApplyRes{}, nil
}

// RebootExecute 若令牌有效则执行重启
func (c *ControllerV1) RebootExecute(ctx context.Context, req *v1.RebootExecuteReq) (res *v1.RebootExecuteRes, err error) {
	if time.Now().After(rebootApplyExpireAt) {
		return nil, errors.Errorf("reboot token expired or not applied")
	}
	switch runtime.GOOS {
	case "linux":
		// 以非阻塞方式触发重启
		go exec.Command("sh", "-c", "(sleep 1; systemctl reboot || reboot) >/dev/null 2>&1 &").Start()
	default:
		// 其他系统模拟成功
	}
	// 清空令牌
	rebootApplyExpireAt = time.Time{}
	return &v1.RebootExecuteRes{}, nil
}

func (c *ControllerV1) GetSetting(ctx context.Context, req *v1.GetSettingReq) (res *v1.GetSettingRes, err error) {
	// 获取所有设置信息
	settings, err := s_db.GetSettingService().GetAllSettings(ctx)
	if err != nil {
		return nil, err
	}

	// 构建设置映射
	settingsMap := make(map[string]string)
	for _, setting := range settings {
		// 只返回启用的设置
		if setting.Enabled {
			settingsMap[setting.Id] = setting.Value
		}
	}

	res = &v1.GetSettingRes{
		Settings: settingsMap,
	}
	return
}

func (c *ControllerV1) UpdateStorageTime(ctx context.Context, req *v1.UpdateStorageTimeReq) (res *v1.UpdateStorageTimeRes, err error) {
	// 更新设备数据保留天数
	err = s_db.GetSettingService().SetSettingValueByName(ctx, s_db_basic.SettingDeviceRetentionDays, fmt.Sprintf("%d", req.DeviceRetentionDays))
	if err != nil {
		return nil, errors.Errorf("更新设备数据保留天数失败: %+v", err)
	}

	// 更新系统数据保留天数
	err = s_db.GetSettingService().SetSettingValueByName(ctx, s_db_basic.SettingSystemRetentionDays, fmt.Sprintf("%d", req.SystemRetentionDays))
	if err != nil {
		return nil, errors.Errorf("更新系统数据保留天数失败: %+v", err)
	}

	// 更新日志数据保留天数
	err = s_db.GetSettingService().SetSettingValueByName(ctx, s_db_basic.SettingLogRetentionDays, fmt.Sprintf("%d", req.LogRetentionDays))
	if err != nil {
		return nil, errors.Errorf("更新日志数据保留天数失败: %+v", err)
	}

	return &v1.UpdateStorageTimeRes{}, nil
}
