package internal

import (
	"common/c_proto"
	"p_base"
	"time"

	"github.com/pkg/errors"
	"go.einride.tech/can"
)

// SendMessage 发送 CANbus 消息
//
// 此函数根据 task 中的 point 描述和 values 来生成 CANbus 帧数据，支持以下功能：
// 1. 优先使用自定义编码器（CustomEncoder），如果存在的话
// 2. 否则使用默认编码器，根据每个 point 的数据访问配置（DataAccess）进行数据编码
// 3. 支持多种数据格式：整数、浮点数、BCD码、字符串、位数据等
// 4. 支持字节序和字序处理
// 5. 支持系数和偏移量转换
// 6. 按顺序连续放置数据，确保无空缺
//
// 参数说明：
//   - task: CANbus 任务配置，包含点位信息和 CAN ID 配置
//   - values: 要发送的值数组，长度必须与 task.Points 长度一致
//
// 返回值：
//   - error: 发送过程中的错误，nil 表示成功
//
// 使用示例：
//
//	// 使用默认编码器
//	task := &c_proto.SCanbusTask{
//		Points: []*c_proto.SCanbusPoint{...},
//		GetCanbusID: func(params map[string]any) uint32 { ... },
//	}
//	values := []int64{1234, 5678}
//
//	// 使用自定义编码器
//	taskWithCustom := &c_proto.SCanbusTask{
//		Points: []*c_proto.SCanbusPoint{...},
//		GetCanbusID: func(params map[string]any) uint32 { ... },
//		CustomEncoder: func(task *c_proto.SCanbusTask, values []any) ([]byte, error) {
//			// 自定义编码逻辑
//			return []byte{0x01, 0x02, 0x03, 0x04}, nil
//		},
//	}
//
//	// 发送消息
//	err := provider.SendMessage(task, values)
//	if err != nil {
//		// 处理错误
//	}
func (c *CanbusProtocolProvider) SendMessage(task *c_proto.SCanbusTask, values []float64) error {
	// 验证输入参数
	if task == nil {
		return errors.New("task cannot be nil")
	}
	if len(values) == 0 {
		return errors.New("values cannot be empty")
	}
	if len(task.Points) == 0 {
		return errors.New("task points cannot be empty")
	}
	if len(values) != len(task.Points) {
		return errors.Errorf("values length (%d) must match points length (%d)", len(values), len(task.Points))
	}

	var canData []byte
	var err error

	// 检查是否有自定义编码器
	if task.CustomEncoder != nil {
		// 使用自定义编码器
		canData, err = task.CustomEncoder(task, convertFloat64ToAny(values))
		if err != nil {
			return errors.Wrap(err, "custom encoder failed")
		}
	} else {
		// 使用默认编码器：遍历每个点位，根据其配置编码数据
		canData, err = c.encodeWithDefaultEncoder(task, values)
		if err != nil {
			return err
		}
	}

	// 验证数据长度
	if len(canData) > can.MaxDataLength {
		return errors.Errorf("encoded data length (%d) exceeds CAN frame max length (%d)",
			len(canData), can.MaxDataLength)
	}

	// 创建并发送 CAN 帧
	return c.sendCanFrame(task, canData)
}

// encodeWithDefaultEncoder 使用默认编码器编码数据
func (c *CanbusProtocolProvider) encodeWithDefaultEncoder(task *c_proto.SCanbusTask, values []float64) ([]byte, error) {
	canData := make([]byte, can.MaxDataLength)
	currentByteIndex := 0

	for i, point := range task.Points {
		if point.DataAccess == nil {
			return nil, errors.Errorf("point %s has no data access configuration", point.GetName())
		}

		// 获取当前值
		value := values[i]

		// 使用编码器将值编码为字节
		encodedBytes, err := p_base.EncoderBytes(
			value,
			point.DataAccess.DataFormat,
			point.DataAccess.ByteEndian,
			point.DataAccess.WordOrder,
			point.DataAccess.Offset,
			point.DataAccess.Factor,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to encode value for point %s", point.GetName())
		}

		// 验证数据长度
		if len(encodedBytes) > len(canData)-currentByteIndex {
			return nil, errors.Errorf("encoded data length (%d) exceeds remaining CAN frame space for point %s",
				len(encodedBytes), point.GetName())
		}

		// 将编码后的数据复制到 CAN 数据缓冲区
		copy(canData[currentByteIndex:currentByteIndex+len(encodedBytes)], encodedBytes)
		currentByteIndex += len(encodedBytes)
	}

	return canData[:currentByteIndex], nil
}

// sendCanFrame 创建并发送 CAN 帧，支持重试机制
func (c *CanbusProtocolProvider) sendCanFrame(task *c_proto.SCanbusTask, canData []byte) error {
	// 创建 CAN 帧
	canID := task.GetCanbusID(c.deviceConfig.Params)

	frame := can.Frame{
		ID:         canID,
		Length:     uint8(len(canData)),
		Data:       can.Data(canData),
		IsRemote:   task.IsRemote,
		IsExtended: task.IsExtended,
	}

	// 重试发送 CAN 帧，每次间隔1秒
	for attempt := 1; attempt <= task.SendMaxRetries; attempt++ {
		select {
		case c.transmitterChan <- frame:
			// 发送成功
			return nil
		default:
			// channel 满了，需要重试
			if attempt < task.SendMaxRetries {
				// 等待1秒后重试
				time.Sleep(time.Second)
				continue
			}
			// 最后一次尝试失败，返回错误
			return errors.Errorf("failed to send CAN frame after %d attempts: transmitter channel is full", task.SendMaxRetries)
		}
	}

	return errors.New("unexpected error in sendCanFrame")
}

// convertFloat64ToAny 将 int64 数组转换为 any 数组
func convertFloat64ToAny(values []float64) []any {
	result := make([]any, len(values))
	for i, v := range values {
		result[i] = v
	}
	return result
}
