package c_base

import "common/c_enum"

type SStorageConfig struct {
	Enable                    bool                `json:"enable,omitempty" orm:"enable"`                                       // 是否启用
	Type                      c_enum.EStorageType `json:"type,omitempty" orm:"type"`                                           // 类型
	Url                       string              `json:"url,omitempty" orm:"url"`                                             // 地址
	SystemMetricsSurvivalDays int32               `json:"systemMetricsSurvivalDays,omitempty" orm:"systemMetricsSurvivalDays"` // 数据保存天数,0代表永久保存,-1代表不保存
	Params                    map[string]string   `json:"params,omitempty" orm:"params"`
}
