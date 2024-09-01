package collect

import (
	"context"
	common "ems-plan"
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"plug_protocol_gpio_sysfs"
	"plug_protocol_gpio_sysfs/p_gpio_sysfs"
	"plug_protocol_modbus"
	"strings"
)

func Create(ctx context.Context, clientConfigs []*c_base.SProtocolConfig) error {

	for _, protocolConfig := range clientConfigs {

		newCtx := context.WithValue(ctx, "I18nName", protocolConfig.Name)

		//newCtx.Value()

		if !protocolConfig.Enable {
			g.Log().Noticef(newCtx, "协议%s 连接地址：%s Enable为fasle, 协议不启用！", protocolConfig.GetProtocol(), protocolConfig.GetAddress())
			continue
		}

		g.Log().Infof(newCtx, "搜索到协议：%s 连接地址：%s 超时时间：%s毫秒 日志等级：%v", protocolConfig.GetProtocol(), protocolConfig.GetAddress(), protocolConfig.GetTimeout(), protocolConfig.GetLogLevel())

		switch protocolConfig.GetProtocol() {
		case c_base.EModbusRtu, c_base.EModbusTcp:
			// 创建client
			client := plug_protocol_modbus.NewModbusClient(newCtx, protocolConfig)

			// 把所有的设备配置文件转换为列表
			for _, _device := range protocolConfig.DeviceChildren {
				deviceConfig := &c_base.SDriverConfig{}
				err := gconv.Scan(_device, deviceConfig)
				if err != nil {
					return err
				}
				if deviceConfig.Id == "" {
					panic(fmt.Sprintf("设备Id不能为空！"))
				}

				deviceCtx := context.WithValue(newCtx, "DeviceName", fmt.Sprintf("%s:%s", strings.ToUpper(string(deviceConfig.Type)), deviceConfig.Id))
				if !deviceConfig.Enable {
					g.Log().Warningf(deviceCtx, "设备%s Enable为fasle, 设备不启用！", deviceConfig.Name)
					continue
				}

				/*	// 通过加载插件的方式来调用
					symbol, err := util.OpenPlugin(ctx, "/Users/zhao/Documents/01.Code/Zhao/ems/application/manifest/protocol/modbus_basic_v1.protocol")
					if err != nil {
						return err
					}
					fmt.Printf("%T\n", symbol)
					provider := symbol.(protocol.NewModbusProvider)
					modbusProvider, err := provider(ctx, protocolConfig, deviceConfig, client)*/

				dv := loadDriver(deviceCtx, deviceConfig)

				modbusProvider, err := plug_protocol_modbus.NewModbusProvider(deviceCtx, protocolConfig, deviceConfig, client) // 进行直接掉用
				if err != nil {
					return err
				}

				dv.Init(modbusProvider, deviceConfig)

				// 柜子的电表，加到缓存中
				//if deviceConfig.DeviceConfig.StationType == config.Ammeter && deviceConfig.Location == config.Cabinet {
				//	cabinetIdAmmeterMap[deviceConfig.CabinetId] = dv.(driver.IAmmeter)
				//}

				// 开始监听(此处防止device未调用Listen方法而执行)
				modbusProvider.Init()

				// 设置告警
				//protocol.SetupProtocol(modbusProvider)

				//device.RegisterInstance(dv)
				g.Log().Noticef(deviceCtx, "设备%s加载完成！Id为：%s", deviceConfig.Name, dv.GetDeviceConfig().Id)

				common.DeviceInstance.RegisterInstance(dv)
			}

		case c_base.ECanbusTcp:
		case c_base.ECanbus:
		case c_base.EGpioSysfs:

			for _, _device := range protocolConfig.DeviceChildren {
				deviceConfig := &c_base.SDriverConfig{}
				err := gconv.Scan(_device, deviceConfig)
				if err != nil {
					return err
				}
				if deviceConfig.Id == "" {
					panic(fmt.Sprintf("设备Id不能为空！"))
				}

				deviceCtx := context.WithValue(newCtx, "DeviceName", fmt.Sprintf("%s:%s", strings.ToUpper(string(deviceConfig.Type)), deviceConfig.Id))
				if !deviceConfig.Enable {
					g.Log().Warningf(deviceCtx, "设备%s Enable为fasle, 设备不启用！", deviceConfig.Name)
					continue
				}

				impl := &p_gpio_sysfs.SDriverGpioImpl{}

				gpioSysfsProtocol, err := plug_protocol_gpio_sysfs.NewGpioSysfsProvider(deviceCtx, protocolConfig, deviceConfig)
				if err != nil {
					return err
				}
				impl.Init(gpioSysfsProtocol, deviceConfig)

				gpioSysfsProtocol.Init()

				common.DeviceInstance.RegisterInstance(impl)
			}

		}
	}
	// 初始化所有的entity
	_tempInstanceCache.Init(ctx)

	g.Log().Infof(ctx, "所有设备加载完成！")

	//modbus.Init(&modbus.BaseModbusHandler{
	//	Cabinets: essHandlerMap,
	//})

	return nil
}
