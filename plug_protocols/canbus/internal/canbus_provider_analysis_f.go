package internal

import (
	"canbus/p_canbus"
	"common/c_base"
	"go.einride.tech/can"
)

func (c *CanbusProtocolProvider) analysisCanbus(task *p_canbus.SCanbusTask, frame can.Frame) {
	if task.CanbusID != frame.ID {
		return
	}

	data := &frame.Data
	for i := 0; i < len(task.Metas); i++ {
		meta := task.Metas[i]
		if meta == nil {
			continue
		}
		c.analysisSingleCanbusMeta(meta, data)
	}

}

func (c *CanbusProtocolProvider) analysisSingleCanbusMeta(meta *c_base.Meta, frameData *can.Data) {
	//frameData.
}
