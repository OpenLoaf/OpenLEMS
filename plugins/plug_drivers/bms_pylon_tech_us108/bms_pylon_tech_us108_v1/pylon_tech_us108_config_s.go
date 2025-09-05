package bms_pylon_tech_us108_v1

import (
	"common/c_proto"
)

type sPylonTechUs108BmsConfig struct {
	c_proto.SModbusDeviceConfig
	SyncTime       bool    `json:"syncTime" name:"是否同步时间" ct:"switch" default:"false"`                             // 是否同步时间
	RatedPower     *uint32 `json:"ratedPower input_type:" name:"额定功率" min:"0" max:"1000" default:"100"  unit:"kW"` // 额定功率
	Capacity       *uint32 `json:"capacity" min:"0" name:"容量" max:"2000" default:"232"  unit:"kWh"`                // 容量
	MaxInputPower  *uint32 `json:"maxInputPower" name:"最大输入功率" min:"0" unit:"kW"`                                  // 最大输入功率
	MaxOutputPower *uint32 `json:"maxOutputPower" name:"最大输出功率" min:"0" unit:"kW"`                                 // 最大输出功率
}

func (s *sPylonTechUs108BmsConfig) getMaxInputPower(power float32) *float32 {
	if s.MaxInputPower == nil {
		return nil
	}
	maxPower := float32(*s.MaxInputPower)
	if maxPower != 0 && power < maxPower {
		return &maxPower
	}
	return nil
}

func (s *sPylonTechUs108BmsConfig) getMaxOutputPower(power float32) *float32 {
	if s.MaxOutputPower == nil {
		return nil
	}
	maxPower := float32(*s.MaxOutputPower)
	if maxPower != 0 && power < maxPower {
		return &maxPower
	}
	return nil
}
