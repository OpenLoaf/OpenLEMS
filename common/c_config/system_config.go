package c_config

import "time"

type SystemConfig struct {
	ConfigPath           string `json:"configPath"`
	DriverPath           string `json:"driverPath"`
	ProtocolPath         string `json:"protocolPath"`
	PrintCacheValueCycle int    `json:"printCacheValueCycle"`
}

func (p *SystemConfig) GetPrintCacheValueCycleDuration() time.Duration {
	if p.PrintCacheValueCycle == 0 {
		p.PrintCacheValueCycle = 3
	}
	return time.Duration(p.PrintCacheValueCycle) * time.Second
}
