package c_base

import (
	"time"
)

type SDeviceDetail struct {
	Id   string `json:"id,omitempty" `   // 设备ID
	Pid  string `json:"pid,omitempty" `  // 父设备Id
	Name string `json:"name,omitempty" ` // 设备名称

	LogLevel           string         `json:"logLevel,omitempty" ` // 日志等级
	Strategy           string         `json:"strategy,omitempty" ` // 	策略名称
	StorageEnable      bool           `json:"StorageEnable" `      // 是否存储
	StorageIntervalSec int32          `json:"storageIntervalSec" ` // 存储间隔(秒),0代表默认1分钟，负数代表不存储
	Sort               int            `json:"sort" `
	Enabled            bool           `json:"enabled" `          // 是否启用
	Params             map[string]any `json:"params,omitempty" ` // 额外参数

	Driver            string              `json:"driver,omitempty" ` // 驱动名
	DriverType        EDeviceType         `json:"type"`              // 驱动类型
	DriverDescription *SDriverDescription `json:"description"`       // 驱动描述

	ProtocolId      string         `json:"protocolId,omitempty" `   // 协议配置ID,如果配置了肯定是实体设备
	ProtocolType    EProtocolType  `json:"protocolType,omitempty" ` // 协议
	ProtocolName    string         `json:"protocolName,omitempty" `
	ProtocolAddress string         `json:"address,omitempty" ` // 地址
	ProtocolParams  map[string]any `json:"protocolParams,omitempty" `
	ProtocolEnable  bool           `json:"protocolEnable" ` // 协议是否启动

	/* 二次计算后的值 */
	IsPhysics         bool       `json:"isPhysics,omitempty" ` // 是否是物理设备
	HasChildDevice    bool       `json:"hasChildDevice" dc:"是否有子设备"`
	AlarmLevel        string     `json:"alarmLevel" dc:"告警级别"`
	LastUpdateTime    *time.Time `json:"lastUpdateTime" dc:"最后更新时间"`
	DeviceServerState string     `json:"deviceServerState,omitempty"` // 设备服务状态
}

func NewDeviceDetailNoInstance(deviceConfig *SDeviceConfig, driverInfo *SDriverInfo, protocolConfig *SProtocolConfig) *SDeviceDetail {
	return NewDeviceDetail(deviceConfig, driverInfo, protocolConfig, nil, EStateInit)
}

func NewDeviceDetail(deviceConfig *SDeviceConfig, driverInfo *SDriverInfo, protocolConfig *SProtocolConfig, instance IDevice, deviceServerState EServerState) *SDeviceDetail {
	if deviceConfig == nil {
		panic("deviceConfig is nil")
	}
	deviceDetail := &SDeviceDetail{
		Id:                 deviceConfig.Id,
		Pid:                deviceConfig.Pid,
		Name:               deviceConfig.Name,
		ProtocolId:         deviceConfig.ProtocolId,
		Driver:             deviceConfig.Driver,
		LogLevel:           deviceConfig.LogLevel,
		Strategy:           deviceConfig.Strategy,
		StorageEnable:      deviceConfig.StorageEnable,
		StorageIntervalSec: deviceConfig.StorageIntervalSec,
		Sort:               deviceConfig.Sort,
		Enabled:            deviceConfig.Enabled,
		Params:             deviceConfig.Params,
	}

	if driverInfo != nil {
		deviceDetail.DriverType = driverInfo.Type
		deviceDetail.DriverDescription = driverInfo.Description
	}
	if protocolConfig != nil {
		deviceDetail.ProtocolName = protocolConfig.Name
		deviceDetail.ProtocolAddress = protocolConfig.Address
		deviceDetail.ProtocolType = protocolConfig.Type
		deviceDetail.ProtocolParams = protocolConfig.Params
		deviceDetail.ProtocolEnable = protocolConfig.Enabled
	} else {
		deviceDetail.ProtocolEnable = false
	}

	if instance != nil {
		if deviceServerState != EStateError && deviceServerState != EStateInit {
			deviceDetail.AlarmLevel = instance.GetAlarmLevel().String()
			deviceDetail.LastUpdateTime = instance.GetLastUpdateTime()
			deviceDetail.IsPhysics = instance.IsPhysical()
		}
		deviceDetail.DeviceServerState = deviceServerState.String()
	}

	return deviceDetail
}
