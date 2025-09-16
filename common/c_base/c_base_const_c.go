package c_base

import (
	"time"

	"github.com/pkg/errors"
)

const (
	ConstCtxKeyGroupName    = "GroupName"
	ConstCtxKeyProtocolId   = "ProtocolId"
	ConstCtxKeyPolicyId     = "PolicyId"
	ConstCtxKeyDeviceId     = "DeviceId"
	ConstCtxKeyDeviceName   = "DeviceName"
	ConstCtxKeyDeviceDetail = "DeviceDetail"

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

const DefaultCacheLifeTime = 10 * time.Second

var NotSupport = errors.New("function not support") // 不支持的方法
