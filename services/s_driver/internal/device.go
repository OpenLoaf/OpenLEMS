package internal

import (
	protocolCanbus "canbus"
	"common"
	"common/c_base"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/torykit/go-modbus"
	"go.einride.tech/can"
	"gpio_sysfs"
	protocolModbus "modbus"
	"os"
)

type SDeviceCmd struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewDeviceCmd(ctx context.Context) *SDeviceCmd {
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Driver")
	ctx, cancelFunc := context.WithCancel(ctx)
	return &SDeviceCmd{
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
}

func (d *SDeviceCmd) Start(activeDeviceRootId string) {
	deviceConfig := common.GetDriverConfigServ().GetDriverConfig(d.ctx, activeDeviceRootId)
	protocolConfigList := common.GetDriverConfigServ().GetProtocolsConfigList(d.ctx)

	// 捕捉panic
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(d.ctx, "启动驱动失败！原因：%v", err)
			d.Stop()
		}
	}()

	d.InitDriver(map[string]any{}, deviceConfig, protocolConfigList)
}

func (d *SDeviceCmd) Stop() {

	// 关闭所有client
	for _, driver := range common.GetDeviceAll() {
		driver.Destroy()
	}

	d.cancelFunc()
}

func (d *SDeviceCmd) InitDriver(clientCache map[string]any, config *c_base.SDriverConfig, protocolConfigList []*c_base.SProtocolConfig) c_base.IDriver {

	fmt.Print(config, "config")
	if err := config.Check(); err != nil {
		panic(err)
	}

	if config.IsEnable == false {
		g.Log().Noticef(d.ctx, "设备[%s]未启用！", config.Name)
		return nil
	}

	// 先递归初始化子设备
	if config.DeviceChildren != nil {
		for _, _device := range config.DeviceChildren {
			d.InitDriver(clientCache, _device, protocolConfigList)
		}
	}

	//if config.Id == "root" {
	//	return nil
	//}
	var protocolProvider c_base.IProtocol

	driverCtx := d.ctx
	// 设备初始化
	if config.ProtocolId != "" {
		driverCtx = context.WithValue(driverCtx, c_base.ConstCtxKeyProtocolId, config.ProtocolId)
		var protocolConfig *c_base.SProtocolConfig
		for _, _protocolConfig := range protocolConfigList {
			if _protocolConfig.Id == config.ProtocolId {
				protocolConfig = _protocolConfig
				break
			}
		}
		protocolProvider = d.getProtocolProvider(driverCtx, clientCache, config, protocolConfig)
	} else {
		driverCtx = context.WithValue(driverCtx, c_base.ConstCtxKeyProtocolId, "Virtual")
	}

	driver := d.getDriver(driverCtx, config)
	if driver == nil {
		g.Log().Errorf(driverCtx, "设备[%s]驱动加载失败！", config.Name)
		return nil
	}

	driver.Init(protocolProvider, config)

	g.Log().Noticef(driverCtx, "设备[%s]驱动加载初始化完毕！\n  设备信息: %s", config.Name, driver.GetDescription())

	if protocolProvider != nil {
		protocolProvider.Init()
	}

	common.RegisterDevice(driver)

	if config.StorageEnable {
		common.RegisterStorageDriver(config.StorageIntervalSec, driver)
	}

	return driver
}

// Block 阻塞进程
func (d *SDeviceCmd) Block() {
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		g.Log().Noticef(d.ctx, "接收到信号：%s", sig.String())
		d.Stop()
	})

	gproc.Listen()
}

func (d *SDeviceCmd) getProtocolProvider(ctx context.Context, clientCache map[string]any, deviceConfig *c_base.SDriverConfig, protocolConfig *c_base.SProtocolConfig) c_base.IProtocol {
	// 从配置中获取协议
	//protocolConfig := common.GetProtocolById(ctx, protocolId)
	if protocolConfig == nil {
		panic(gerror.Newf("协议Id: %s 配置信息不存在！", deviceConfig.ProtocolId))
	}
	if err := protocolConfig.Check(); err != nil {
		panic(err)
	}

	// 初始化协议
	switch protocolConfig.GetProtocol() {
	case c_base.EModbusRtu, c_base.EModbusTcp:
		// 从缓存中获取client，如果没有就新建后放入缓存
		var client modbus.Client
		if _client, exist := clientCache[protocolConfig.Id]; exist {
			client = _client.(modbus.Client)
		} else {
			// 创建client
			client = protocolModbus.NewModbusClient(ctx, protocolConfig)
		}

		modbusProvider, err := protocolModbus.NewModbusProvider(ctx, protocolConfig, deviceConfig, client)
		if err != nil {
			panic(err)
		}

		return modbusProvider
	case c_base.ECanbusUdp, c_base.ECanbus:
		var (
			receiverChan    <-chan can.Frame
			transmitterChan chan<- can.Frame
		)
		if _receiverChan, exist := clientCache[protocolConfig.Id+"_receiverChan"]; exist {
			receiverChan = _receiverChan.(chan can.Frame)
		}
		if _transmitterChan, exist := clientCache[protocolConfig.Id+"_transmitterChan"]; exist {
			transmitterChan = _transmitterChan.(chan<- can.Frame)
		}

		if receiverChan == nil || transmitterChan == nil {
			receiverChan, transmitterChan = protocolCanbus.NewCanbusChan(ctx, protocolConfig)
		}

		canbusProvider, err := protocolCanbus.NewCanbusProvider(ctx, protocolConfig, deviceConfig, receiverChan, transmitterChan)
		if err != nil {
			panic(err)
		}
		g.Log().Infof(ctx, "canbusProvider: %s 创建成功! Params: %v", protocolConfig.GetAddress(), protocolConfig.Params)
		return canbusProvider
	case c_base.EGpioSysfs:
		gpioSysfsProtocol, err := gpio_sysfs.NewGpioSysfsProvider(ctx, protocolConfig, deviceConfig)
		if err != nil {
			panic(err)
		}
		return gpioSysfsProtocol
	}

	panic(gerror.Newf("未知的协议类型：%s", protocolConfig.GetProtocol()))
}
