package c_base

import "github.com/torykit/go-modbus"

type IService interface {
	Start()

	Stop()

	InitDriver(clientCache map[string]modbus.Client, config *SDriverConfig, protocolConfigList []*SProtocolConfig) IDriver
}
