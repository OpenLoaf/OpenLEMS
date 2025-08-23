package internal

import (
	"common/c_base"
	"common/c_log"
	"context"
	"github.com/goburrow/serial"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/torykit/go-modbus"
	"time"
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
			// 创建一个Ticker，每3秒触发一次
			ticker := time.NewTicker(3 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					_ = client.Close()
					g.Log().Noticef(ctx, "上下文取消！连接:[%s] 取消连接重连!", protocolConfig.GetAddress())
					c_log.BizInfof(ctx, "设备关闭！")
					return
				case <-ticker.C:
					if client.IsConnected() {
						continue
					}
					c_log.BizInfof(ctx, "设备掉线，准备重连！")
					err := client.Connect()
					if err != nil {
						c_log.BizWarningf(ctx, "connect failed! watting next connect, %v", err)
					} else {
						// 等待2秒，再判断是否成功
						if client.IsConnected() {
							c_log.BizInfof(ctx, "modbus协议重连成功！")
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
