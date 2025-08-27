package internal

import (
	"common/c_base"
	"common/c_log"
	"context"
	"time"

	"github.com/goburrow/serial"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/torykit/go-modbus"
)

//const notifyManage =

// NewModbusClient 创建modbus客户端，并链接
func NewModbusClient(ctx context.Context, protocolConfig *c_base.SProtocolConfig) (modbus.Client, error) {
	var client modbus.Client
	switch protocolConfig.GetProtocol() {
	case c_base.EModbusTcp:
		var option []modbus.ClientProviderOption

		if protocolConfig.GetLogLevel() == "debug" || protocolConfig.GetLogLevel() == "DEBUG" {
			option = append(option, modbus.WithEnableLogger())
		}

		option = append(option, modbus.WithTCPTimeout(time.Duration(protocolConfig.GetTimeout())*time.Second))
		tcpProvider := modbus.NewTCPClientProvider(protocolConfig.GetAddress(), option...)
		client = modbus.NewClient(tcpProvider)
	case c_base.EModbusRtu:

		rtuConfig := &ModbusRtuProtocolConfig{}
		err := gconv.Scan(protocolConfig.Params, rtuConfig)
		if err != nil {
			//panic(gerror.New("modbus rtu配置文件解析失败"))
			return nil, gerror.Wrap(err, `modbus rtu配置文件解析失败`)
		}

		var option []modbus.ClientProviderOption
		if protocolConfig.GetLogLevel() == "debug" || protocolConfig.GetLogLevel() == "DEBUG" {
			option = append(option, modbus.WithEnableLogger())
		}
		option = append(option, modbus.WithSerialConfig(serial.Config{
			Address:  protocolConfig.GetAddress(),
			BaudRate: rtuConfig.GetBaudRate(),
			DataBits: rtuConfig.GetDataBits(),
			StopBits: rtuConfig.GetStopBits(),
			Parity:   rtuConfig.GetParity(),
			Timeout:  time.Duration(protocolConfig.GetTimeout()) * time.Millisecond,
		}))

		p := modbus.NewRTUClientProvider(option...)
		client = modbus.NewClient(p)
	default:
		//panic(gerror.New("不支持的modbus协议！"))
		return nil, gerror.New("不支持的modbus协议！")
	}

	err := client.Connect()
	if err != nil {
		if protocolConfig.GetProtocol() == c_base.EModbusRtu {
			//panic(gerror.Newf("modbus rtu 地址：[%s] 连接失败！ %v", protocolConfig.GetAddress(), err))
			return nil, gerror.Wrapf(err, "modbus rtu 地址：[%s] 连接失败！ %v", protocolConfig.GetAddress())
		}
		c_log.BizWarningf(ctx, "连接到：modbus地址 %s 失败！等待下一次连接！ %v", protocolConfig.GetAddress(), err)
	} else {
		c_log.BizInfof(ctx, "modbus协议连接到：%s 成功！", protocolConfig.GetAddress())
	}

	if protocolConfig.GetProtocol() == c_base.EModbusTcp {
		// 重连机制
		go func() {
			// 重连间隔配置：第一次3秒、第二次10秒、第三次30秒、第四次60秒、之后每5分钟
			reconnectIntervals := []time.Duration{
				3 * time.Second,  // 第1次
				10 * time.Second, // 第2次
				30 * time.Second, // 第3次
				60 * time.Second, // 第4次
				5 * time.Minute,  // 第5次及之后
				10 * time.Minute, // 第6次及之后
			}

			reconnectCount := 0
			var currentInterval time.Duration

			for {
				select {
				case <-ctx.Done():
					_ = client.Close()
					g.Log().Noticef(ctx, "上下文取消！连接:[%s] 取消连接重连!", protocolConfig.GetAddress())
					c_log.BizInfof(ctx, "设备关闭！")
					return
				default:
					// 检查连接状态
					if client.IsConnected() {

						// 连接正常，重置重连计数
						if reconnectCount > 0 {
							c_log.BizInfof(ctx, "连接恢复正常，重置重连计数")
							reconnectCount = 0
						}
						// 等待一段时间再检查
						time.Sleep(3 * time.Second)
						continue
					}

					// 计算当前重连间隔
					if reconnectCount < len(reconnectIntervals) {
						currentInterval = reconnectIntervals[reconnectCount]
					} else {
						currentInterval = reconnectIntervals[len(reconnectIntervals)-1] // 使用最后一个间隔（5分钟）
					}

					reconnectCount++
					c_log.BizInfof(ctx, "设备掉线，准备第%d次重连，等待间隔：%v", reconnectCount, currentInterval)

					// 等待重连间隔
					select {
					case <-ctx.Done():
						_ = client.Close()
						g.Log().Noticef(ctx, "上下文取消！连接:[%s] 取消连接重连!", protocolConfig.GetAddress())
						c_log.BizInfof(ctx, "设备关闭！")
						return
					case <-time.After(currentInterval):
						// 执行重连
						err := client.Connect()
						if err != nil {
							c_log.BizWarningf(ctx, "第%d次重连失败，等待下次重连，错误：%v", reconnectCount, err)
						} else {
							// 等待2秒，再判断是否成功
							time.Sleep(2 * time.Second)
							if client.IsConnected() {
								c_log.BizInfof(ctx, "modbus协议第%d次重连成功！", reconnectCount)
								reconnectCount = 0 // 重连成功，重置计数
							} else {
								c_log.BizWarningf(ctx, "第%d次重连连接建立但状态检查失败", reconnectCount)
							}
						}
					}
				}
			}
		}()
	} else {
		// RTU协议等ctx结束后 Close串口
		go func() {
			<-ctx.Done()
			_ = client.Close()
			c_log.BizInfof(ctx, "连接:[%s] 关闭!", protocolConfig.GetAddress())
		}()
	}

	return client, nil
}
