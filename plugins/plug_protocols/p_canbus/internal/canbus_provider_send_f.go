package internal

import (
	"common/c_proto"

	"github.com/gogf/gf/v2/errors/gerror"
	"go.einride.tech/can"
)

func (c *CanbusProtocolProvider) SendMessage(task *c_proto.SCanbusTask, values []int64) error {

	//valueBytes := point.ReadType.Encoder(values[i], point.Factor, point.Offset, point.Endianness)

	bytes := make([]byte, can.MaxDataLength)

	// 需要验证一下meta的顺序是否正确
	metaIndex := uint16(0)
	for i, point := range task.Points {
		if metaIndex == 0 {
			metaIndex = point.Addr
		} else {
			if point.Addr != (metaIndex + c_protocol.ReadTypeRegisterSize(point.ReadType)) {
				return gerror.Newf("点位的顺序不正确！点位：%s, 地址：%d，实际地址应该为: %d", point.Name, point.Addr, metaIndex+c_protocol.ReadTypeRegisterSize(point.ReadType))
			}
			metaIndex = point.Addr
		}
		valueBytes := c_protocol.ReadTypeEncoder(point.ReadType, values[i], point.Factor, point.Offset, point.Endianness)

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
