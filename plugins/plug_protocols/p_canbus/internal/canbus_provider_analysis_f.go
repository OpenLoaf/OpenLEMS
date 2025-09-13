package internal

import (
	"common/c_base"
	"common/c_proto"
	"p_base"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go.einride.tech/can"
)

func (c *CanbusProtocolProvider) analysisCanbus(task *c_proto.SCanbusTask, frame can.Frame) error {
	// 检查是否有自定义解码器
	if task.CustomDecoder != nil {
		// 使用自定义解码器处理整个帧
		var lastErr error
		successCount := 0

		for _, point := range task.Points {
			if point == nil {
				continue
			}

			value, err := task.CustomDecoder(task, frame.Data[:], point)
			if err != nil {
				lastErr = gerror.Wrapf(err, "自定义解码器解析失败 task:%s point:%s", task.Name, point.GetName())
				g.Log().Errorf(c.ctx, "自定义解码器解析失败 task:%s point:%s error:%v",
					task.Name, point.GetName(), err)
				continue
			}
			//point.TriggerAlarm(value)

			// 处理解析成功后的步骤
			if err := c.handleDecodedValue(task, point, value, "自定义解码器"); err != nil {
				lastErr = err
				continue
			}

			successCount++
		}

		// 如果所有点位都解析失败，返回错误
		if successCount == 0 && lastErr != nil {
			return gerror.Wrapf(lastErr, "自定义解码器解析所有点位失败 task:%s", task.Name)
		}

		// 如果有部分成功，记录警告但不返回错误
		if successCount > 0 && lastErr != nil {
			g.Log().Warningf(c.ctx, "自定义解码器部分解析失败 task:%s 成功:%d 失败:%d",
				task.Name, successCount, len(task.Points)-successCount)
		}

		return nil
	}

	g.Log().Debugf(c.ctx, "===> 收到匹配的task数据：taskName: %s  数据：%v", task.Name, frame)

	// 使用默认解码器：根据每个点位的 ByteIndex 解析
	frameData := frame.Data[:]
	var lastErr error
	successCount := 0

	for i, point := range task.Points {
		if point == nil || point.DataAccess == nil {
			lastErr = gerror.Newf("点位配置为空 task:%s pointIndex:%d", task.Name, i)
			g.Log().Errorf(c.ctx, "点位配置为空 task:%s pointIndex:%d", task.Name, i)
			continue
		}

		// 解析单个点位，使用 point 的 ByteIndex
		value, err := c.analysisSingleCanbusMeta(point, frameData, int(point.DataAccess.ByteIndex), task.Lifetime)
		if err != nil {
			lastErr = gerror.Wrapf(err, "解析点位失败 task:%s point:%s", task.Name, point.GetName())
			g.Log().Errorf(c.ctx, "解析点位失败 task:%s point:%s error:%v",
				task.Name, point.GetName(), err)
			continue
		}

		// 处理解析成功后的步骤
		if err := c.handleDecodedValue(task, point, value, "默认解码器"); err != nil {
			lastErr = err
			continue
		}

		successCount++
	}

	// 如果所有点位都解析失败，返回错误
	if successCount == 0 && lastErr != nil {
		return gerror.Wrapf(lastErr, "默认解码器解析所有点位失败 task:%s", task.Name)
	}

	// 如果有部分成功，记录警告但不返回错误
	if successCount > 0 && lastErr != nil {
		g.Log().Warningf(c.ctx, "默认解码器部分解析失败 task:%s 成功:%d 失败:%d",
			task.Name, successCount, len(task.Points)-successCount)
	}

	return nil
}

// handleDecodedValue 处理解码成功后的公共步骤
// 包括：缓存结果、更新告警、记录日志
func (c *CanbusProtocolProvider) handleDecodedValue(task *c_proto.SCanbusTask, point *c_proto.SCanbusPoint, value any, decoderType string) error {
	// 缓存解析结果
	pointValue := c_base.NewPointValue(c.deviceId, point, value)
	if cacheErr := c.IProtocolCacheValue.CacheValue(pointValue, task.Lifetime); cacheErr != nil {
		err := gerror.Wrapf(cacheErr, "缓存解析结果失败 task:%s point:%s", task.Name, point.GetName())
		g.Log().Errorf(c.ctx, "缓存解析结果失败 task:%s point:%s error:%v",
			task.Name, point.GetName(), cacheErr)
		return err
	}

	// 更新告警
	c.UpdateAlarm(c.deviceId, point, value)

	// 记录成功日志
	g.Log().Debugf(c.ctx, "%s解析成功 task:%s point:%s value:%v",
		decoderType, task.Name, point.GetName(), value)

	return nil
}

// analysisSingleCanbusMeta 解析单个 CANbus 点位数据
// 使用 DecoderBytes 函数根据点位配置解析数据
func (c *CanbusProtocolProvider) analysisSingleCanbusMeta(point *c_proto.SCanbusPoint, frameData []byte, currentByteIndex int, lifeTime time.Duration) (any, error) {
	if point.DataAccess == nil {
		g.Log().Errorf(c.ctx, "点位 %s 没有数据访问配置", point.GetName())
		return nil, nil
	}

	// 使用 DecoderBytes 解析数据
	value, err := p_base.DecoderBytes(
		frameData,                   // 原始字节数据
		uint16(currentByteIndex),    // 字节起始索引
		point.DataAccess.ByteLength, // 字节长度
		point.DataAccess.BitIndex,   // 位起始索引
		point.DataAccess.BitLength,  // 位长度
		point.DataAccess.ByteEndian, // 字节序
		point.DataAccess.WordOrder,  // 字序
		point.DataAccess.DataFormat, // 数据格式
		point.ValueType,             // 返回格式类型
		point.DataAccess.Offset,     // 偏移量
		point.DataAccess.Factor,     // 系数
	)

	if err != nil {
		g.Log().Errorf(c.ctx, "解析点位 %s 失败: %v", point.GetName(), err)
		return nil, err
	}

	return value, nil
}
