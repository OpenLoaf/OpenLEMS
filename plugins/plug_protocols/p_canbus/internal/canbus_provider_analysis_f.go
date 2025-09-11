package internal

import (
	"common/c_base"
	"common/c_proto"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"go.einride.tech/can"
)

func (c *CanbusProtocolProvider) analysisCanbus(task *c_proto.SCanbusTask, frame can.Frame) {
	//if task.GetCanbusID != frame.ID {
	//	return
	//}

	g.Log().Debugf(c.ctx, "===> 收到匹配的task数据：taskName: %s  数据：%v", task.Name, frame)

	data := &frame.Data
	for i := 0; i < len(task.Points); i++ {
		meta := task.Points[i]
		if meta == nil {
			continue
		}
		c.analysisSingleCanbusMeta(meta, data, task.Lifetime)
	}

}

func (c *CanbusProtocolProvider) analysisSingleCanbusMeta(meta *c_proto.SCanbusPoint, frameData *can.Data, lifeTime time.Duration) {
	//frameData.

	//value, err := meta.ReadType.ReadValue(result[index:], meta.BitLength, meta.Endianness)
	v, err := c_protocol.MetaTransformCanbus(c.ctx, c.deviceConfig.Id, c.deviceType, c, meta, frameData[:], c.cache, lifeTime)

	if err != nil {
		g.Log().Errorf(c.ctx, "解析CAN: %v 数据失败 meta:%s ：%s", frameData, meta.Name, err.Error())
		return
	}
	g.Log().Debugf(c.ctx, "解析数据成功：%s : %v", meta.Cn, v)
}
