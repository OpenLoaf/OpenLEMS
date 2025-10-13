package c_base

import (
	"time"

	"github.com/pkg/errors"
)

const (
	ConstNewPluginFunctionName = "NewPlugin"
)

const (
	ConstStationEnergyStoreId = "station-energy-store"
)

const (
	ConstProtocol = "protocol"
	ConstSystem   = "system"
	ConstProcess  = "process"

	ConstProtocolId      = "protocol_id"
	ConstProtocolAddress = "protocol_address"
	ConstProtocolType    = "protocol_type"
	ConstDeviceId        = "device_id"
	ConstDeviceName      = "device_name"
	ConstDeviceType      = "device_type"
)

const (
	// 远程协议类型常量
	ConstRemoteModbus = "modbus"
	ConstRemoteMqtt   = "mqtt"
)

const DefaultCacheLifeTime = 10 * time.Second

var NotSupport = errors.New("function not support") // 不支持的方法
