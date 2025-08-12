package internal

import (
	"canbus/p_canbus"
	"common"
	"common/c_base"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"go.einride.tech/can"
)

func (c *CanbusProtocolProvider) analysisCanbus(task *p_canbus.SCanbusTask, frame can.Frame) {
	//if task.GetCanbusID != frame.ID {
	//	return
	//}

	g.Log().Debugf(c.ctx, "===> 收到匹配的task数据：taskName: %s  数据：%v", task.Name, frame)

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
	g.Log().Debugf(c.ctx, "解析数据成功：%s : %v", meta.Cn, v)
}
