package internal

import (
	"canbus/p_canbus"
	"common"
	"common/c_base"
	"github.com/gogf/gf/v2/frame/g"
	"go.einride.tech/can"
	"time"
)

func (c *CanbusProtocolProvider) analysisCanbus(task *p_canbus.SCanbusTask, frame can.Frame) {
	if task.CanbusID != frame.ID {
		return
	}

	g.Log().Infof(c.ctx, "===> 收到数据：%v", frame)

	data := &frame.Data
	for i := 0; i < len(task.Metas); i++ {
		meta := task.Metas[i]
		if meta == nil {
			continue
		}
		c.analysisSingleCanbusMeta(meta, data, task.Lifetime)
	}

}

func (c *CanbusProtocolProvider) analysisSingleCanbusMeta(meta *c_base.Meta, frameData *can.Data, lifeTime time.Duration) {
	//frameData.

	//value, err := meta.ReadType.ReadValue(result[index:], meta.BitLength, meta.Endianness)
	v, err := common.MetaTransformCanbus(c.ctx, c.deviceConfig.Id, c.deviceType, c, meta, frameData[:], c.cache, lifeTime)

	if err != nil {
		g.Log().Errorf(c.ctx, "解析数据失败：%v", err)
		return
	}
	g.Log().Infof(c.ctx, "解析数据成功：%s : %v", meta.Cn, v)
}
