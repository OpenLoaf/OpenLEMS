package internal

import (
	"common/c_base"
	"context"
	"p_modbus"

	"github.com/pkg/errors"
	"github.com/torykit/go-modbus"
)

func (m *SDeviceManager) IsProtocolActive(protocolId string) bool {
	return m.protocolClientCache[protocolId] != nil
}

func (m *SDeviceManager) Shutdown() {
	// 关闭所有client
	//m.deviceConfigTree.IteratorDesc(func(key, value any) bool {
	//	deviceWrapper := value.(*SDeviceWrapper)
	//	if deviceWrapper.deviceState == c_base.EStateRunning {
	//		deviceWrapper.Shutdown()
	//	}
	//	return true
	//})
	m.cancelFunc()
	m.state = c_base.EStateStopped
	return
}

func (m *SDeviceManager) Cleanup() error {
	// 定时清理无用连接
	return nil
}

func (m *SDeviceManager) Status() c_base.EServerState {
	return m.state
}

func (m *SDeviceManager) GetChildDeviceInstance(pid string) []c_base.IDevice {
	var deviceInstances = make([]c_base.IDevice, 0)
	flatList := m.GetFlatList()
	for _, deviceConfig := range flatList {
		if deviceConfig != nil && deviceConfig.Pid == pid {
			if instance, exist := m.deviceInstanceMap[deviceConfig.Id]; exist {
				deviceInstances = append(deviceInstances, instance)
			}
		}
	}
	return deviceInstances
}

//func (d *SDeviceManager) getModbusProvider() (c_proto.IModbusProtocol,error) {
//
//}

func (m *SDeviceManager) getProtocolProvider(ctx context.Context, deviceConfig *c_base.SDeviceConfig) (c_base.IProtocol, error) {
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
	case c_base.EModbusRtu, c_base.EModbusTcp:
		// 从缓存中获取client，如果没有就新建后放入缓存
		var client modbus.Client
		if _client, exist := m.protocolClientCache[protocolConfig.Id]; exist {
			client = _client.(modbus.Client)
		} else {
			// 创建client
			c, err := p_modbus.NewModbusClient(ctx, protocolConfig)
			if err != nil {
				return nil, err
			}
			client = c
			m.protocolClientCache[protocolConfig.Id] = client
		}
		modbusProvider, err := p_modbus.NewModbusProvider(ctx, deviceConfig.GetType(), protocolConfig, deviceConfig, client)
		if err != nil {
			return nil, err
		}

		return modbusProvider, nil
	case c_base.ECanbusUdp, c_base.ECanbus:
		//var (
		//	receiverChan    <-chan can.Frame
		//	transmitterChan chan<- can.Frame
		//)
		//if _receiverChan, exist := m.protocolClientCache[protocolConfig.Id+"_receiverChan"]; exist {
		//	receiverChan = _receiverChan.(chan can.Frame)
		//}
		//if _transmitterChan, exist := m.protocolClientCache[protocolConfig.Id+"_transmitterChan"]; exist {
		//	transmitterChan = _transmitterChan.(chan<- can.Frame)
		//}
		//
		//if receiverChan == nil || transmitterChan == nil {
		//	r, t, err := protocolCanbus.NewCanbusChan(ctx, protocolConfig)
		//	if err != nil {
		//		return nil, err
		//	}
		//	receiverChan = r
		//	transmitterChan = t
		//	m.protocolClientCache[protocolConfig.Id+"_receiverChan"] = receiverChan
		//	m.protocolClientCache[protocolConfig.Id+"_transmitterChan"] = transmitterChan
		//}
		//
		//canbusProvider, err := protocolCanbus.NewCanbusProvider(ctx, deviceType, protocolConfig, deviceConfig, receiverChan, transmitterChan)
		//if err != nil {
		//	return nil, err
		//}
		//g.Log().Infof(ctx, "canbusProvider: %s 创建成功! Params: %v", protocolConfig.GetAddress(), protocolConfig.Params)
		//return canbusProvider, nil
	case c_base.EGpioSysfs:
		//gpioSysfsProtocol, err := gpio_sysfs.NewGpioSysfsProvider(ctx, protocolConfig, deviceConfig)
		//if err != nil {
		//	return nil, err
		//}
		//return gpioSysfsProtocol, nil
	}

	return nil, errors.Errorf("未知的协议类型：%s", protocolConfig.GetProtocol())
}
