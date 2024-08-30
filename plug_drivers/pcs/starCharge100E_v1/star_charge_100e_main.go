//go:generate ../build.sh
package main

import (
	"common_station/c_station"
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"plug_protocol_modbus"
	"plug_protocol_modbus/p_modbus"
	"starCharge100E_v1/star_charge_100e"
	"time"
)

// NewPlugin 必须的方法，不能取消
func NewPlugin(ctx context.Context) (c_device.IPcs, error) {
	return &star_charge_100e.StarCharge100EPcs{
		Ctx: ctx,
	}, nil
}

func main() {

	var (
		protocolConfig = &c_base.SProtocolConfig{
			Protocol: c_base.EModbusTcp,
			//Address:  "10.211.55.3:502",
			Address: "127.0.0.1:1503",
			Timeout: 1000,
		}
		deviceConfig = &p_modbus.SModbusDeviceConfig{
			SDriverConfig: c_base.SDriverConfig{
				Id:              "测试Pcs-1",
				Name:            "测试Pcs-1",
				Driver:          "",
				Group:           c_station.EGroupNan,
				CabinetId:       1,
				IsMaster:        false,
				Enable:          true,
				LogLevel:        "INFO",
				PrintCacheValue: false,
			},
			UnitId: 1,
		}
	)

	ctx := gctx.New()
	// 加载modbus插件
	//baseProtocol, err := util.OpenPlugin(ctx, "/Users/zhao/Documents/01.Code/Zhao/ems/application/manifest/protocol/modbus_basic_v1.protocol")
	//if err != nil {
	//	return
	//}
	// 调用modbus插件的NewModbusProvider方法，创建一个modbus provider
	//newModbusProviderPlugin := baseProtocol.(protocol.NewModbusProvider)

	client := plug_protocol_modbus.NewModbusClient(gctx.New(), protocolConfig)
	provider, err := plug_protocol_modbus.NewModbusProvider(ctx, protocolConfig, deviceConfig, client)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建一个Device对象，并初始化
	device := &star_charge_100e.StarCharge100EPcs{
		Ctx: context.TODO(),
	}
	err = device.Init(provider, deviceConfig)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 每2秒打印一次缓存值
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			provider.PrintCacheValues()
		}
	}()

	for {
		time.Sleep(time.Hour)
	}
}
