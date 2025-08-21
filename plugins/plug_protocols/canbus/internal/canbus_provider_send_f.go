package internal

import (
	"c_protocol"
	"canbus/p_canbus"
	"github.com/gogf/gf/v2/errors/gerror"
	"go.einride.tech/can"
)

func (c *CanbusProtocolProvider) SendMessage(task *p_canbus.SCanbusTask, values []int64) error {

	//valueBytes := meta.ReadType.Encoder(values[i], meta.Factor, meta.Offset, meta.Endianness)

	bytes := make([]byte, can.MaxDataLength)

	// 需要验证一下meta的顺序是否正确
	metaIndex := uint16(0)
	for i, meta := range task.Metas {
		if metaIndex == 0 {
			metaIndex = meta.Addr
		} else {
			if meta.Addr != (metaIndex + c_protocol.ReadTypeRegisterSize(meta.ReadType)) {
				return gerror.Newf("点位的顺序不正确！点位：%s, 地址：%d，实际地址应该为: %d", meta.Name, meta.Addr, metaIndex+c_protocol.ReadTypeRegisterSize(meta.ReadType))
			}
			metaIndex = meta.Addr
		}
		valueBytes := c_protocol.ReadTypeEncoder(meta.ReadType, values[i], meta.Factor, meta.Offset, meta.Endianness)

		copy(bytes[i*2:], valueBytes)
	}

	frame := can.Frame{
		ID:         *task.GetCanbusID(c.deviceConfig.Params),
		Length:     uint8(len(values)),
		Data:       can.Data(bytes),
		IsRemote:   task.IsRemote,
		IsExtended: task.IsExtended,
	}

	//g.Log().Debugf(c.ctx, "发送的内容为: %v", frame)

	c.transmitterChan <- frame

	return nil
}
