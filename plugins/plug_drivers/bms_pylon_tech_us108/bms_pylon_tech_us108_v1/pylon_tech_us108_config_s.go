package bms_pylon_tech_us108_v1

import (
	"common/c_proto"
)

type sPylonTechUs108BmsConfig struct {
	c_proto.SModbusDeviceConfig
	SyncTime       bool    `json:"syncTime" input_type:"bool"`                            // 是否同步时间
	RatedPower     *uint32 `json:"ratedPower input_type:"int" min:0 max:1000 default:100` // 额定功率
	Capacity       *uint32 `json:"capacity" input_type:"int"`                             // 容量
	MaxInputPower  *uint32 `json:"maxInputPower"`                                         // 最大输入功率
	MaxOutputPower *uint32 `json:"maxOutputPower"`                                        // 最大输出功率
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
