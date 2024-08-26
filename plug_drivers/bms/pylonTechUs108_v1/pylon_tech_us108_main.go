//go:generate ../build.sh
package main

import (
	"common_station/c_station"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"plug_protocol_modbus"
	"plug_protocol_modbus/p_modbus"
	"pylonTechUs108_v1/pylon_tech_us108"
	"time"
)

// NewPlugin 必须的方法，不能取消
func NewPlugin() (c_device.IBms, error) {
	return &pylon_tech_us108.PylonTechUs108Bms{}, nil
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
				Id:              "测试BMS-1",
				Name:            "测试BMS-1",
				Driver:          "",
				Group:           c_station.EGroupNan,
				CabinetId:       1,
				IsMaster:        false,
				Enable:          true,
				LogLevel:        "INFO",
				PrintCacheValue: false,
			},
			UnitId: 2,
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
		glog.Fatal(ctx, err)
	}

	// 创建一个Device对象，并初始化
	device := &pylon_tech_us108.PylonTechUs108Bms{}
	err = device.Init(ctx, provider, gconv.MapStrStr(deviceConfig))
	if err != nil {
		glog.Fatal(ctx, err)
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
