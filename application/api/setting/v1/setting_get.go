package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetSettingReq struct {
	g.Meta `path:"/setting" method:"get" tags:"设置相关" summary:"获取公开设置信息"`
}
type GetSettingRes struct {
	Settings []SettingItem `json:"settings" dc:"公开设置信息列表"`
}

// SettingItem 设置项结构体
type SettingItem struct {
	Id        string `json:"id" dc:"设置ID"`
	Value     string `json:"value" dc:"设置值"`
	GroupName string `json:"groupName" dc:"分组名称"`
	Remark    string `json:"remark" dc:"备注"`
}

// GetSettingDetailReq 获取设置详情请求
type GetSettingDetailReq struct {
	g.Meta `path:"/setting/detail" method:"post" tags:"设置相关" summary:"获取设置详情"`
	Id     string `json:"id" v:"required" dc:"设置ID"`
}

// GetSettingDetailRes 获取设置详情响应
type GetSettingDetailRes struct {
	Id        string     `json:"id" dc:"设置ID"`
	Value     string     `json:"value" dc:"设置值"`
	IsPublic  bool       `json:"isPublic" dc:"是否公开"`
	Enabled   bool       `json:"enabled" dc:"是否启用"`
	Remark    string     `json:"remark" dc:"备注"`
	Sort      int        `json:"sort" dc:"排序"`
	Group     string     `json:"group" dc:"分组"`
	CreatedAt *time.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" dc:"更新时间"`
}

// UpdateSettingReq 更新设置请求
type UpdateSettingReq struct {
	g.Meta   `path:"/setting/update" method:"post" tags:"设置相关" summary:"更新设置信息"`
	Id       string `json:"id" v:"required" dc:"设置ID"`
	Value    string `json:"value" dc:"设置值"`
	Group    string `json:"group" v:"required" dc:"分组"`
	IsPublic *bool  `json:"isPublic" dc:"是否公开，默认为false"`
	Enabled  *bool  `json:"enabled" dc:"是否启用，默认为true"`
	Remark   string `json:"remark" dc:"备注"`
	Sort     int    `json:"sort" dc:"排序"`
}

// UpdateSettingRes 更新设置响应
type UpdateSettingRes struct {
}
