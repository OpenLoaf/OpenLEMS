package c_base

import "time"

type SSystemConfig struct {
	ConfigPath           string `json:"configPath"`
	DriverPath           string `json:"driverPath"`
	ProtocolPath         string `json:"protocolPath"`
	PrintCacheValueCycle int    `json:"printCacheValueCycle"`
}

func (p *SSystemConfig) GetPrintCacheValueCycleDuration() time.Duration {
	if p.PrintCacheValueCycle == 0 {
		p.PrintCacheValueCycle = 3
	}
	return time.Duration(p.PrintCacheValueCycle) * time.Second
}
