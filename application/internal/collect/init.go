package collect

import (
	"context"
	"ems-plan/c_base"
	"plug_protocol_modbus"
	"plug_protocol_modbus/p_modbus"

	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

func Create(ctx context.Context, clientConfigs []*c_base.SProtocolConfig) error {
	var cabinetIdSet = &gset.Set{}

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
				deviceConfig := &p_modbus.SModbusDeviceConfig{}
				err := gconv.Scan(_device, deviceConfig)
				if err != nil {
					return err
				}
				if deviceConfig.Id == "" {
					panic(fmt.Sprintf("设备Id不能为空！"))
				}

				deviceCtx := context.WithValue(newCtx, "DeviceName", fmt.Sprintf("%s:%s", strings.ToUpper(string(deviceConfig.Group)), deviceConfig.Id))
				if !deviceConfig.Enable {
					g.Log().Warningf(deviceCtx, "设备%s Enable为fasle, 设备不启用！", deviceConfig.Name)
					continue
				}
				// 加到cabinetIds中
				if deviceConfig.CabinetId != 0 {
					cabinetIdSet.AddIfNotExist(deviceConfig.CabinetId)
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

				err = dv.Init(modbusProvider, deviceConfig)
				if err != nil {
					return err
				}

				// 柜子的电表，加到缓存中
				//if deviceConfig.DeviceConfig.Group == config.Ammeter && deviceConfig.Location == config.Cabinet {
				//	cabinetIdAmmeterMap[deviceConfig.CabinetId] = dv.(driver.IAmmeter)
				//}

				// 开始监听(此处防止device未调用Listen方法而执行)
				modbusProvider.Init(dv.GetType())

				// 设置告警
				//protocol.SetupProtocol(modbusProvider)

				//device.RegisterInstance(dv)
				g.Log().Noticef(deviceCtx, "设备%s加载完成！Id为：%s", deviceConfig.Name, dv.GetId())
			}

		case "canbus":
		case "canbus_tcp":
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
