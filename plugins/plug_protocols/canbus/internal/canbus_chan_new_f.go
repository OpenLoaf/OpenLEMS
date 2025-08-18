package internal

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go.einride.tech/can"
	"go.einride.tech/can/pkg/candevice"
	"go.einride.tech/can/pkg/socketcan"
	"net"
)

func NewCanbusConnect(ctx context.Context, protocolConfig *c_base.SProtocolConfig) (<-chan can.Frame, chan<- can.Frame, error) {

	canbusConfig := GetCanbusConfig(protocolConfig)

	// 启动接口
	canInterface, err := candevice.New(protocolConfig.GetAddress())
	if err != nil {
		return nil, nil, gerror.Wrapf(err, "canbus初始化失败！%s")
	}
	isUp, _ := canInterface.IsUp()
	if !isUp {
		err = canInterface.SetBitrate(canbusConfig.BaudRate)
		if err != nil {
			return nil, nil, gerror.Wrapf(err, "canbus[%v] 设置波特率为 %v 失败！%s", protocolConfig.GetAddress(), canbusConfig.BaudRate)
		}
		err = canInterface.SetUp()
		if err != nil {
			return nil, nil, err
		}
	}

	//defer conn.SetDown()

	go func() {
		<-ctx.Done()
		_ = canInterface.SetDown()
		g.Log().Noticef(ctx, "Canbus上下文取消! canbus协议[%s]已关闭！", protocolConfig.GetAddress())
	}()

	// 创建接收器
	var conn net.Conn
	switch protocolConfig.GetProtocol() {
	case c_base.ECanbus:
		conn, err = socketcan.DialContext(context.Background(), "can", protocolConfig.GetAddress())
		g.Log().Infof(ctx, "canbus协议连接[%s]初始化成功！", protocolConfig.GetAddress())
	case c_base.ECanbusUdp:
		conn, err = socketcan.DialContext(context.Background(), "udp", protocolConfig.GetAddress())
		g.Log().Infof(ctx, "canbus udp协议连接[%s]初始化成功！", protocolConfig.GetAddress())
	default:
		return nil, nil, gerror.Newf("错误的参数传递！%s 进入了canbus协议连接初始化！", protocolConfig.GetProtocol())
	}
	if err != nil {
		return nil, nil, gerror.Wrapf(err, "canbus协议连接[%s]初始化失败！%s", protocolConfig.GetAddress())
	}

	// 初始化 can 接收器
	receiver := socketcan.NewReceiver(conn)
	receiverChan := make(chan can.Frame, 100) // 用于接收 CAN 帧, 100个缓存帧

	// 初始化 can 发送器
	transmitter := socketcan.NewTransmitter(conn)
	transmitterChan := make(chan can.Frame, 100) // 用于发送 CAN 帧, 100个缓存帧

	// 启动接收 goroutine
	go func() {
		defer close(receiverChan) // 当 goroutine 退出时关闭接收通道
		for receiver.Receive() {
			select {
			case <-ctx.Done(): // 监听上下文取消信号
				g.Log().Noticef(ctx, "接收goroutine: 上下文取消，停止接收CAN帧。")
				return
			default:
				// 读取 CAN 帧
				frame := receiver.Frame()
				//g.Log().Debugf(ctx, "接口 %s 接收到 CAN 帧: %s", protocolConfig.GetAddress(), frame.String()) // 可以打印调试信息
				// 成功接收到帧，发送到通道
				receiverChan <- frame
			}
		}
	}()

	// 启动发送 goroutine
	go func() {
		for {
			select {
			case <-ctx.Done(): // 监听上下文取消信号
				g.Log().Noticef(ctx, "发送goroutine: 上下文取消，停止发送CAN帧。")
				return
			case frame := <-transmitterChan: // 从发送通道接收要发送的帧
				err := transmitter.TransmitFrame(ctx, frame)
				if err != nil {
					g.Log().Errorf(ctx, "发送CAN帧失败: %s", err.Error())
					// 这里可以根据错误类型选择是否继续发送
				} else {
					g.Log().Debugf(ctx, "成功发送 CAN 帧: %s", frame.String()) // 可以打印调试信息
				}
				// 可以添加一个小的延迟以避免CPU过度占用，如果发送频率很高的话
				// time.Sleep(1 * time.Millisecond)
			}
		}
	}()

	return receiverChan, transmitterChan, nil
}
