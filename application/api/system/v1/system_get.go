package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSummaryReq struct {
	g.Meta `path:"/system/summary" method:"get" tags:"系统" summary:"系统汇总信息"`
}
type GetSummaryRes struct {
	CPU    CPUInfo    `json:"cpu"`
	Memory MemoryInfo `json:"memory"`
	Disk   DiskInfo   `json:"disk"`
	Uptime UptimeInfo `json:"uptime"`
	Net    NetBrief   `json:"net"`
	Time   TimeInfo   `json:"time"`
	Sys    SysInfo    `json:"sys"`
}

type GetNetworkTrafficReq struct {
	g.Meta `path:"/system/net/traffic" method:"get" tags:"系统" summary:"网络流量信息"`
}
type GetNetworkTrafficRes struct {
	Net NetBrief `json:"net"`
}

type GetTimeInfoReq struct {
	g.Meta `path:"/system/time" method:"get" tags:"系统" summary:"时间信息"`
}
type GetTimeInfoRes struct {
	Time TimeInfo `json:"time"`
}

type GetSystemInfoReq struct {
	g.Meta `path:"/system/info" method:"get" tags:"系统" summary:"系统信息"`
}
type GetSystemInfoRes struct {
	Sys SysInfo `json:"sys"`
}

type CPUInfo struct {
	Usage  float64 `json:"usage"`
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}
type MemoryInfo struct {
	Usage   float64 `json:"usage"`
	TotalGB float64 `json:"totalGB"`
	UsedGB  float64 `json:"usedGB"`
	FreeGB  float64 `json:"freeGB"`
}
type DiskInfo struct {
	Usage   float64 `json:"usage"`
	TotalGB float64 `json:"totalGB"`
	UsedGB  float64 `json:"usedGB"`
	FreeGB  float64 `json:"freeGB"`
}
type UptimeInfo struct {
	Day      int    `json:"day"`
	Hour     int    `json:"hour"`
	Minute   int    `json:"minute"`
	BootTime string `json:"bootTime"`
}
type NetBrief struct {
	UpBytes   uint64  `json:"upBytes"`
	DownBytes uint64  `json:"downBytes"`
	UpBps     float64 `json:"upBps"`
	DownBps   float64 `json:"downBps"`
}
type IfaceItem struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Up   bool   `json:"up"`
}
type TimeInfo struct {
	Now        string `json:"now"`
	Timezone   string `json:"timezone"`
	NTPRunning bool   `json:"ntpRunning"`
	NTPServer  string `json:"ntpServer"`
}
type SysInfo struct {
	OSName    string `json:"osName"`
	OSVersion string `json:"osVersion"`
	Kernel    string `json:"kernel"`
}
