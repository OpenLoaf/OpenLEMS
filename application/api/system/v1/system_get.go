package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSummaryReq struct {
	g.Meta `path:"/system/summary" method:"get" tags:"系统相关" summary:"系统汇总信息"`
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
	g.Meta `path:"/system/net/traffic" method:"get" tags:"系统相关" summary:"网络流量信息"`
}
type GetNetworkTrafficRes struct {
	Net NetBrief `json:"net" dc:"网络流量信息(累计字节与瞬时速率)"`
}

type GetTimeInfoReq struct {
	g.Meta `path:"/system/time" method:"get" tags:"系统相关" summary:"时间信息"`
}
type GetTimeInfoRes struct {
	Time TimeInfo `json:"time" dc:"时间与NTP信息"`
}

type GetSystemInfoReq struct {
	g.Meta `path:"/system/info" method:"get" tags:"系统相关" summary:"系统信息"`
}
type GetSystemInfoRes struct {
	Sys SysInfo `json:"sys" dc:"系统基础信息"`
}

type CPUInfo struct {
	Usage  float64 `json:"usage" dc:"CPU使用率(%)"`
	Load1  float64 `json:"load1" dc:"1分钟平均负载"`
	Load5  float64 `json:"load5" dc:"5分钟平均负载"`
	Load15 float64 `json:"load15" dc:"15分钟平均负载"`
}
type MemoryInfo struct {
	Usage   float64 `json:"usage" dc:"内存使用率(%)"`
	TotalGB float64 `json:"totalGB" dc:"内存总计(GiB)"`
	UsedGB  float64 `json:"usedGB" dc:"已使用(GiB)"`
	FreeGB  float64 `json:"freeGB" dc:"可用(GiB)"`
}
type DiskInfo struct {
	Usage   float64 `json:"usage" dc:"磁盘使用率(%)"`
	TotalGB float64 `json:"totalGB" dc:"磁盘总计(GiB)"`
	UsedGB  float64 `json:"usedGB" dc:"已使用(GiB)"`
	FreeGB  float64 `json:"freeGB" dc:"可用(GiB)"`
}
type UptimeInfo struct {
	Day      int    `json:"day" dc:"运行天数"`
	Hour     int    `json:"hour" dc:"运行小时"`
	Minute   int    `json:"minute" dc:"运行分钟"`
	BootTime string `json:"bootTime" dc:"开机时间(YYYY-MM-DD HH:mm:ss)"`
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
	OSName    string `json:"osName" dc:"系统名称/发行版"`
	OSVersion string `json:"osVersion" dc:"系统版本号"`
	Kernel    string `json:"kernel" dc:"内核版本号"`
}
