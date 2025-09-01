package entity

import (
	"common/c_status"
)

type EnergyStoreStatus struct {
	DeviceId       string `json:"key" dc:"设备ID"`
	I18nName       string `json:"name" dc:"名称"`
	LastUpdateTime string `json:"lastUpdateTime" dc:"最后更新时间"`

	Status c_status.EEnergyStoreStatus `json:"status" dc:"状态 0:未知状态 1:关机 2:待机 3:充电 4:放电 5:故障 6:同步中"`

	Soc           string `json:"soc" dc:"soc"`
	Power         string `json:"power" dc:"有功功率"`
	ApparentPower string `json:"apparentPower" dc:"视在功率"`
	ReactivePower string `json:"reactivePower" dc:"无功功率"`

	TargetPower         string `json:"targetPower" dc:"目标功率"`
	TargetReactivePower string `json:"targetReactivePower" dc:"目标无功功率"`
	TargetPowerFactor   string `json:"targetPowerFactor" dc:"目标功率因数"`

	TodayCharge      string `json:"todayCharge" dc:"今日充电量"`
	TodayDischarge   string `json:"todayDischarge" dc:"今日放电量"`
	HistoryCharge    string `json:"historyCharge" dc:"历史充电量"`
	HistoryDischarge string `json:"historyDischarge" dc:"历史放电量"`
}
