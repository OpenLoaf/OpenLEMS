package internal

import (
	"context"
	"ems-plan/c_base"
	"fmt"
	"github.com/goburrow/serial"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/torykit/go-modbus"
	"time"
)

//const notifyManage =

// NewModbusClient 创建modbus客户端，并链接
func NewModbusClient(ctx context.Context, protocolConfig *c_base.SProtocolConfig) modbus.Client {
	var client modbus.Client
	switch protocolConfig.GetProtocol() {
	case c_base.EModbusTcp:
		var option []modbus.ClientProviderOption

		if protocolConfig.GetLogLevel() == "debug" || protocolConfig.GetLogLevel() == "DEBUG" {
			option = append(option, modbus.WithEnableLogger())
		}

		option = append(option, modbus.WithTCPTimeout(time.Duration(protocolConfig.GetTimeout())*time.Millisecond))
		tcpProvider := modbus.NewTCPClientProvider(protocolConfig.GetAddress(), option...)
		client = modbus.NewClient(tcpProvider)
	case c_base.EModbusRtu:

		rtuConfig := &ModbusRtuProtocolConfig{}
		err := gconv.Scan(protocolConfig.Params, rtuConfig)
		if err != nil {
			panic("modbus rtu配置文件解析失败")
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
		panic("不支持的modbus协议！")
	}

	err := client.Connect()
	if err != nil {
		if protocolConfig.GetProtocol() == c_base.EModbusRtu {
			// 如果是RTU协议，无法打开串口，那程序就没必要继续运行了
			panic(fmt.Errorf("Modbus rtu 地址：[%s] 连接失败！ %v", protocolConfig.GetAddress(), err))
		}

		g.Log().Warningf(ctx, "first connect failed! watting next connect, %v", err)

	} else {
		g.Log().Infof(ctx, "first connect success!")
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
					return
				case <-ticker.C:
					if client.IsConnected() {
						continue
					}
					err := client.Connect()
					if err != nil {
						g.Log().Warningf(ctx, "connect failed! watting next connect, %v", err)

					} else {
						g.Log().Infof(ctx, "reconnect success!")
					}
				}
			}

		}()
	} else {
		// RTU协议等ctx结束后 Close串口
		go func() {
			<-ctx.Done()
			_ = client.Close()
			g.Log().Noticef(ctx, "ModbusRUT上下文取消！连接:[%s] 取消连接!", protocolConfig.GetAddress())
		}()
	}

	return client
}
