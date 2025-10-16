package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetSummaryReq struct {
	g.Meta `path:"/system/summary" method:"get" tags:"系统相关" summary:"系统汇总信息" role:"user"`
}
type GetSummaryRes struct {
	CPU    CPUInfo    `json:"cpu" dc:"CPU 信息"`
	Memory MemoryInfo `json:"memory" dc:"内存信息"`
	Disk   DiskInfo   `json:"disk" dc:"磁盘信息(根分区)"`
	Uptime UptimeInfo `json:"uptime" dc:"系统运行时长"`
	Time   TimeInfo   `json:"time" dc:"时间与NTP信息"`
	Sys    SysInfo    `json:"sys" dc:"系统基础信息"`
}

type GetNetworkTrafficReq struct {
	g.Meta `path:"/system/net/traffic" method:"get" tags:"系统相关" summary:"网络流量信息" role:"user"`
}
type GetNetworkTrafficRes struct {
	Net NetBrief `json:"net" dc:"网络流量信息(累计字节与瞬时速率)"`
}

type GetProcessMemoryReq struct {
	g.Meta `path:"/system/process/memory" method:"get" tags:"系统相关" summary:"获取本服务进程内存信息" role:"user"`
}
type GetProcessMemoryRes struct {
	ProcessMemory ProcMemoryInfo `json:"processMemory" dc:"本服务进程内存信息"`
}

type GetTimeInfoReq struct {
	g.Meta `path:"/system/time" method:"get" tags:"系统相关" summary:"时间信息" role:"user"`
}
type GetTimeInfoRes struct {
	Time TimeInfo `json:"time" dc:"时间与NTP信息"`
}

type GetSystemInfoReq struct {
	g.Meta `path:"/system/info" method:"get" tags:"系统相关" summary:"系统信息" role:"user"`
}
type GetSystemInfoRes struct {
	Sys SysInfo `json:"sys" dc:"系统基础信息"`
}

type GetNowReq struct {
	g.Meta `path:"/system/now" method:"get" tags:"系统相关" summary:"获取系统当前时间" role:"user"`
}
type GetNowRes struct {
	Now string `json:"now" dc:"当前时间(YYYY-MM-DD HH:mm:ss)"`
}

type UpdateHostnameReq struct {
	g.Meta   `path:"/system/hostname/update" method:"post" tags:"系统相关" summary:"修改主机名" role:"admin"`
	Hostname string `json:"hostname" v:"required|regex:^[a-zA-Z0-9-]{1,63}$#主机名不能为空|主机名格式不正确" dc:"新的主机名(字母数字及短横线,<=63)"`
}
type UpdateHostnameRes struct{}

type UpdateSystemTimeReq struct {
	g.Meta   `path:"/system/time/update" method:"post" tags:"系统相关" summary:"修改系统时间" role:"admin"`
	Time     string `json:"time" v:"required#时间不能为空" dc:"新的时间, 格式: 2006-01-02 15:04:05"`
	Timezone string `json:"timezone" dc:"可选, 设置系统时区, 例如 Asia/Shanghai"`
}
type UpdateSystemTimeRes struct{}

type RebootApplyReq struct {
	g.Meta `path:"/system/reboot/apply" method:"post" tags:"系统相关" summary:"申请重启(30s内有效)" role:"admin"`
}
type RebootApplyRes struct{}

type RebootExecuteReq struct {
	g.Meta `path:"/system/reboot/execute" method:"post" tags:"系统相关" summary:"确认并执行重启(需先申请)" role:"admin"`
}
type RebootExecuteRes struct{}

type CPUInfo struct {
	Usage  float64 `json:"usage" dc:"CPU使用率(%)"`
	Load1  float64 `json:"load1" dc:"1分钟平均负载"`
	Load5  float64 `json:"load5" dc:"5分钟平均负载"`
	Load15 float64 `json:"load15" dc:"15分钟平均负载"`
}
type MemoryInfo struct {
	Usage     float64 `json:"usage" dc:"内存使用率(%)"`
	TotalGB   float64 `json:"totalGB" dc:"内存总计(GiB)"`
	UsedGB    float64 `json:"usedGB" dc:"已使用(GiB)"`
	FreeGB    float64 `json:"freeGB" dc:"可用(GiB)"`
	ProcUsage float64 `json:"procUsage" dc:"本进程内存使用率(%)"`
	ProcRSSMB float64 `json:"procRSSMB" dc:"本进程内存使用量RSS(MiB)"`
}

type ProcMemoryInfo struct {
	Usage float64 `json:"usage" dc:"本进程占系统内存百分比(%)"`
	RSSMB float64 `json:"rssMB" dc:"常驻集(物理内存)MiB"`
	VMSMB float64 `json:"vmsMB" dc:"虚拟内存MiB"`
}

type DiskInfo struct {
	Usage   float64 `json:"usage" dc:"磁盘使用率(%)"`
	TotalGB float64 `json:"totalGB" dc:"磁盘总计(GiB)"`
	UsedGB  float64 `json:"usedGB" dc:"已使用(GiB)"`
	FreeGB  float64 `json:"freeGB" dc:"可用(GiB)"`
}

type UptimeInfo struct {
	Day               int    `json:"day" dc:"运行天数"`
	Hour              int    `json:"hour" dc:"运行小时"`
	Minute            int    `json:"minute" dc:"运行分钟"`
	BootTime          string `json:"bootTime" dc:"开机时间(YYYY-MM-DD HH:mm:ss)"`
	ProcUptimeSeconds int    `json:"procUptimeSeconds" dc:"当前程序运行秒数"`
	ProcStartTime     string `json:"procStartTime" dc:"程序启动时间(YYYY-MM-DD HH:mm:ss)"`
}

type NetBrief struct {
	UpBytes   uint64  `json:"upBytes" dc:"累计上传字节数"`
	DownBytes uint64  `json:"downBytes" dc:"累计下载字节数"`
	UpBps     float64 `json:"upBps" dc:"上传速度(B/s)"`
	DownBps   float64 `json:"downBps" dc:"下载速度(B/s)"`
}

type IfaceItem struct {
	Name string `json:"name" dc:"接口名称"`
	IP   string `json:"ip" dc:"IPv4地址"`
	Up   bool   `json:"up" dc:"是否在线"`
}

type TimeInfo struct {
	Now        string `json:"now" dc:"系统当前时间"`
	Timezone   string `json:"timezone" dc:"系统时区"`
	NTPRunning bool   `json:"ntpRunning" dc:"NTP服务是否启用"`
	NTPServer  string `json:"ntpServer" dc:"NTP服务器地址"`
}

type SysInfo struct {
	Hostname  string `json:"hostname" dc:"主机名"`
	OSName    string `json:"osName" dc:"系统名称/发行版"`
	OSVersion string `json:"osVersion" dc:"系统版本号"`
	Kernel    string `json:"kernel" dc:"内核版本号"`
}

type UpdateStorageTimeReq struct {
	g.Meta              `path:"/system/storage-time" method:"post" tags:"系统相关" summary:"更新存储时间参数" role:"admin"`
	DeviceRetentionDays int `json:"deviceRetentionDays" v:"min:1#设备数据保留天数必须大于0" dc:"设备数据保留天数"`
	SystemRetentionDays int `json:"systemRetentionDays" v:"min:1#系统数据保留天数必须大于0" dc:"系统数据保留天数"`
	LogRetentionDays    int `json:"logRetentionDays" v:"min:1#日志数据保留天数必须大于0" dc:"日志数据保留天数"`
}
type UpdateStorageTimeRes struct{}

type GetStorageStatsReq struct {
	g.Meta `path:"/system/storage/stats" method:"get" tags:"系统相关" summary:"获取存储统计信息" role:"user"`
}
type GetStorageStatsRes struct {
	Stats StorageStatsInfo `json:"stats" dc:"存储统计信息"`
}

// StorageStatsInfo 存储统计信息结构体
type StorageStatsInfo struct {
	TotalSeries      int64     `json:"total_series" dc:"总时间序列数量"`
	TotalSamples     int64     `json:"total_samples" dc:"总样本数量"`
	StorageSize      int64     `json:"storage_size" dc:"存储大小（字节）"`
	StorageSizeMB    float64   `json:"storage_size_mb" dc:"存储大小（MB）"`
	OldestTimestamp  time.Time `json:"oldest_timestamp" dc:"最老数据时间戳"`
	NewestTimestamp  time.Time `json:"newest_timestamp" dc:"最新数据时间戳"`
	RetentionTime    int64     `json:"retention_time" dc:"数据保留时间（秒）"`
	RetentionHours   float64   `json:"retention_hours" dc:"数据保留时间（小时）"`
	AvgSeriesSize    float64   `json:"avg_series_size" dc:"平均每个序列占用数据大小（字节）"`
	SamplesPerSecond float64   `json:"samples_per_second" dc:"每秒存储样本数"`
}
