package internal

import (
	"common/c_base"
	"common/c_enum"
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
		if _client, exist := m.protocolClientCache[protocolConfig.Id]; exist {
			client = _client.(modbus.Client)
		} else {
			// 创建client
			c, err := p_modbus.NewModbusClient(m.ctx, protocolConfig)
			if err != nil {
				return nil, err
			}
			client = c
			m.protocolClientCache[protocolConfig.Id] = client
		}
		modbusProvider, err := p_modbus.NewModbusProvider(deviceCtx, deviceConfig.GetType(), protocolConfig, deviceConfig, client)
		if err != nil {
			return nil, err
		}

		return modbusProvider, nil
	case c_enum.ECanbusUdp, c_enum.ECanbus:
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
		//	r, t, err := protocolCanbus.NewCanbusChan(deviceCtx, protocolConfig)
		//	if err != nil {
		//		return nil, err
		//	}
		//	receiverChan = r
		//	transmitterChan = t
		//	m.protocolClientCache[protocolConfig.Id+"_receiverChan"] = receiverChan
		//	m.protocolClientCache[protocolConfig.Id+"_transmitterChan"] = transmitterChan
		//}
		//
		//canbusProvider, err := protocolCanbus.NewCanbusProvider(deviceCtx, deviceType, protocolConfig, deviceConfigTree, receiverChan, transmitterChan)
		//if err != nil {
		//	return nil, err
		//}
		//g.Log().Infof(deviceCtx, "canbusProvider: %s 创建成功! Params: %v", protocolConfig.GetAddress(), protocolConfig.Params)
		//return canbusProvider, nil
	case c_enum.EGpioSysfs:
		//gpioSysfsProtocol, err := gpio_sysfs.NewGpioSysfsProvider(deviceCtx, protocolConfig, deviceConfigTree)
		//if err != nil {
		//	return nil, err
		//}
		//return gpioSysfsProtocol, nil
	}

	return nil, errors.Errorf("未知的协议类型：%s", protocolConfig.GetProtocol())
}
