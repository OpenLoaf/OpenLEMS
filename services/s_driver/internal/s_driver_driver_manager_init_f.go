package internal

import (
	"common/c_base"
	"common/c_enum"
	"context"
	"p_canbus"
	"p_gpiod"
	"p_modbus"
	"runtime"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
	"github.com/torykit/go-modbus"
	"go.einride.tech/can"
)

func (m *SDeviceManager) IsProtocolActive(protocolId string) bool {
	return m.protocolClientCache.Contains(protocolId)
}

func (m *SDeviceManager) Shutdown() {
	// 关闭所有client
	//m.deviceConfigTree.IteratorDesc(func(key, value any) bool {
	//	deviceWrapper := value.(*SDeviceWrapper)
	//	if deviceWrapper.deviceState == c_enum.EStateRunning {
	//		deviceWrapper.Shutdown()
	//	}
	//	return true
	//})
	m.cancelFunc()
	m.state = c_enum.EStateStopped
	return
}

func (m *SDeviceManager) Cleanup() error {
	// 定时清理无用连接
	return nil
}

func (m *SDeviceManager) Status() c_enum.EServerState {
	return m.state
}

func (m *SDeviceManager) GetChildDeviceInstance(pid string) []c_base.IDevice {
	var deviceInstances = make([]c_base.IDevice, 0)
	flatList := m.GetFlatList()
	for _, deviceConfig := range flatList {
		if deviceConfig != nil && deviceConfig.Pid == pid {
			if instance := m.deviceInstanceMap.Get(deviceConfig.Id); instance != nil {
				if dev, ok := instance.(c_base.IDevice); ok {
					deviceInstances = append(deviceInstances, dev)
				}
			}
		}
	}
	return deviceInstances
}

// getProtocolProvider 不管连接是否能创建，都会返回provider。但是如果系统不支持的时候，就不会返回
func (m *SDeviceManager) getProtocolProvider(deviceCtx context.Context, deviceConfig *c_base.SDeviceConfig) (c_base.IProtocol, error) {
	// 从配置中获取协议
	protocolConfig := deviceConfig.ProtocolConfig
	if protocolConfig == nil {
		return nil, errors.Errorf("协议Id: %s 配置信息不存在！", deviceConfig.ProtocolId)
	}
	if err := protocolConfig.Check(); err != nil {
		return nil, errors.Wrapf(err, "检查协议配置失败，格式异常！")
	}

	// 初始化协议
	switch protocolConfig.GetProtocol() {
	case c_enum.EModbusRtu, c_enum.EModbusTcp:
		// 从缓存中获取client，如果没有就新建后放入缓存
		var client modbus.Client
		if _client := m.protocolClientCache.Get(protocolConfig.Id); _client != nil {
			client = _client.(modbus.Client)
		} else {
			// 创建client
			c, err := p_modbus.NewModbusClient(m.ctx, protocolConfig)
			if err != nil {
				return nil, err
			}
			client = c
			m.protocolClientCache.Set(protocolConfig.Id, client)
		}
		modbusProvider, err := p_modbus.NewModbusProvider(deviceCtx, protocolConfig, deviceConfig, client)
		if err != nil {
			return nil, err
		}

		return modbusProvider, nil
	case c_enum.ECanbusUdp, c_enum.ECanbus:
		// 验证系统支持：只有Linux系统才支持canbus
		if runtime.GOOS != "linux" {
			return nil, errors.Errorf("canbus协议仅在Linux系统上支持，当前系统: %s", runtime.GOOS)
		}

		var (
			receiverChan    <-chan can.Frame
			transmitterChan chan<- can.Frame
		)
		if _receiverChan := m.protocolClientCache.Get(protocolConfig.Id + "_receiverChan"); _receiverChan != nil {
			receiverChan = _receiverChan.(chan can.Frame)
		}
		if _transmitterChan := m.protocolClientCache.Get(protocolConfig.Id + "_transmitterChan"); _transmitterChan != nil {
			transmitterChan = _transmitterChan.(chan<- can.Frame)
		}

		if receiverChan == nil || transmitterChan == nil {
			r, t, err := p_canbus.NewCanbusChan(deviceCtx, protocolConfig)
			if err != nil {
				return nil, err
			}
			receiverChan = r
			transmitterChan = t
			m.protocolClientCache.Set(protocolConfig.Id+"_receiverChan", receiverChan)
			m.protocolClientCache.Set(protocolConfig.Id+"_transmitterChan", transmitterChan)
		}

		canbusProvider, err := p_canbus.NewCanbusProvider(deviceCtx, protocolConfig, deviceConfig, receiverChan, transmitterChan)
		if err != nil {
			return nil, err
		}
		g.Log().Infof(deviceCtx, "canbusProvider: %s 创建成功! Params: %v", protocolConfig.GetAddress(), protocolConfig.Params)
		return canbusProvider, nil
	case c_enum.EGpiod:
		// 验证系统支持：只有Linux系统才支持gpiod
		//if runtime.GOOS != "linux" {
		//	return nil, errors.Errorf("gpiod协议仅在Linux系统上支持，当前系统: %s", runtime.GOOS)
		//}
		if m.protocolClientCache.Contains(protocolConfig.Id) {
			return nil, errors.Errorf("[%s]已被其他设备占用！", protocolConfig.Id)
		}

		gpioProtocol, err := p_gpiod.NewGpiodProvider(deviceCtx, protocolConfig, deviceConfig)
		m.protocolClientCache.Set(protocolConfig.Id, gpioProtocol)

		if err != nil {
			return nil, err
		}
		return gpioProtocol, nil
	case c_enum.EGpioSfs:
		// 未来完善
	}

	return nil, errors.Errorf("未知的协议类型：%s", protocolConfig.GetProtocol())
}
