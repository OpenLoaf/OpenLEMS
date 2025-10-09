package internal

import (
	"common"
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"common/c_util"
	"context"
	"time"

	"github.com/goburrow/serial"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
	"github.com/torykit/go-modbus"
)

//const notifyManage =

// NewModbusClient 创建modbus客户端，并链接
func NewModbusClient(ctx context.Context, protocolConfig *c_base.SProtocolConfig) (modbus.Client, error) {
	var client modbus.Client
	switch protocolConfig.GetProtocol() {
	case c_enum.EModbusTcp:
		var option []modbus.ClientProviderOption

		if protocolConfig.GetLogLevel() == "debug" || protocolConfig.GetLogLevel() == "DEBUG" {
			option = append(option, modbus.WithEnableLogger())
		}

		option = append(option, modbus.WithTCPTimeout(time.Duration(protocolConfig.GetTimeout())*time.Millisecond))
		tcpProvider := modbus.NewTCPClientProvider(protocolConfig.GetAddress(), option...)
		client = modbus.NewClient(tcpProvider)
	case c_enum.EModbusRtu:

		rtuConfig := &c_proto.SModbusProtocolRtuConfig{}
		err := gconv.Scan(protocolConfig.Params, rtuConfig)
		if err != nil {
			//panic(errors.Errorf("modbus rtu配置文件解析失败"))
			return nil, errors.Wrap(err, `modbus rtu配置文件解析失败`)
		}

		var option []modbus.ClientProviderOption
		if protocolConfig.GetLogLevel() == "debug" || protocolConfig.GetLogLevel() == "DEBUG" {
			option = append(option, modbus.WithEnableLogger())
		}
		option = append(option, modbus.WithSerialConfig(serial.Config{
			Address:  protocolConfig.GetAddress(),
			BaudRate: rtuConfig.BaudRate,
			DataBits: rtuConfig.DataBits,
			StopBits: rtuConfig.StopBits,
			Parity:   rtuConfig.Parity,
			Timeout:  time.Duration(protocolConfig.GetTimeout()) * time.Millisecond,
		}))

		p := modbus.NewRTUClientProvider(option...)
		client = modbus.NewClient(p)
	default:
		//panic(errors.Errorf("不支持的modbus协议！"))
		return nil, errors.Errorf("不支持的modbus协议！")
	}

	err := client.Connect()
	if err != nil {
		if protocolConfig.GetProtocol() == c_enum.EModbusRtu {
			//panic(errors.Errorf("modbus rtu 地址：[%s] 连接失败！ %v", protocolConfig.GetAddress(), err))
			return nil, errors.Wrapf(err, "modbus rtu 地址：[%s] 连接失败！", protocolConfig.GetAddress())
		}
		c_log.BizWarningf(ctx, "连接到：modbus地址 %s 失败！等待下一次连接！ %s", protocolConfig.GetAddress(), err.Error())
	} else {
		c_log.BizInfof(ctx, "modbus协议连接到：%s 成功！", protocolConfig.GetAddress())
	}

	if protocolConfig.GetProtocol() == c_enum.EModbusTcp {
		// 重连机制
		go func() {
			reconnectCount := 0
			var startTime time.Time
			var lastInterval time.Duration
			var isOffline bool

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

						// 连接正常，重置重连计数和开始时间
						if isOffline {
							c_log.BizInfof(ctx, "连接恢复正常，清除告警。")
							reconnectCount = 0
							startTime = time.Time{} // 重置开始时间
							lastInterval = 0        // 重置间隔记录
							isOffline = false
							TriggerOfflineAlarm(protocolConfig.Id, false)
						}
						// 等待一段时间再检查
						time.Sleep(3 * time.Second)
						continue
					}

					// 首次检测到掉线，记录开始时间
					if !isOffline {
						startTime = time.Now()
						isOffline = true
					}

					// todo 此处需要获取所有使用到该协议的设备，都触发连接断开警告
					TriggerOfflineAlarm(protocolConfig.Id, true)

					// 计算当前重连间隔
					elapsed := time.Since(startTime)
					var currentInterval time.Duration

					if elapsed < 5*time.Minute {
						// 前5分钟：3秒一次
						currentInterval = 3 * time.Second
					} else if elapsed < 30*time.Minute {
						// 5-30分钟：10秒一次
						currentInterval = 10 * time.Second
					} else {
						// 超过30分钟：1分钟一次
						currentInterval = 1 * time.Minute
					}

					reconnectCount++

					// 只在重连间隔切换时记录日志
					if currentInterval != lastInterval {
						if elapsed < time.Second {
							// 首次掉线且时间小于1秒，不显示掉线时长
							c_log.BizInfof(ctx, "设备掉线，重连策略切换为：%v 重连一次", currentInterval)
						} else {
							c_log.BizInfof(ctx, "设备掉线，重连策略切换为：%v 重连一次，已掉线时长：%s", currentInterval, c_util.FormatDuration(elapsed))
						}
						lastInterval = currentInterval
					}

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
							c_log.Debugf(ctx, "第%d次重连失败，等待下次重连，错误：%v", reconnectCount, err)
						} else {
							// 等待2秒，再判断是否成功
							time.Sleep(2 * time.Second)
							if client.IsConnected() {
								offlineDuration := time.Since(startTime)
								c_log.BizInfof(ctx, "modbus协议第%d次重连成功，离线时长为：%s", reconnectCount, c_util.FormatDuration(offlineDuration))
								reconnectCount = 0      // 重连成功，重置计数
								startTime = time.Time{} // 重置开始时间
								lastInterval = 0        // 重置间隔记录
								isOffline = false       // 重置离线状态
								TriggerOfflineAlarm(protocolConfig.Id, false)
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

func TriggerOfflineAlarm(protocolId string, trigger bool) {
	common.GetDeviceManager().IteratorAllDevices(func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		if device == nil || config == nil {
			return true
		}
		protocolConfig := config.ProtocolConfig
		if protocolConfig == nil || protocolConfig.Id != protocolId {
			return true
		}

		// 触发或清除离线告警
		device.UpdateAlarm(device.GetConfig().Id, &c_proto.SModbusPoint{
			SProtocolPoint: &c_base.SProtocolPoint{
				SPoint: &c_base.SPoint{
					Key:  "Offline",
					Name: "设备连接状态告警",
				},
			},
			Trigger: func(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
				trigger, err = cvt.BoolE(value)
				level = c_enum.EAlarmLevelWarn
				return
			},
			StatusExplain: func(value any) (string, error) {
				v, err := cvt.BoolE(value)
				if err != nil {
					return "", errors.Wrap(err, "状态转换失败")
				}
				if v {
					return "离线", nil
				}
				return "上线", nil
			},
		}, trigger)

		return true
	})
}
