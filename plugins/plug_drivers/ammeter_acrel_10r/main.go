//go:generate ../build.sh
package main

import (
	"ammeter_acrel_10r_v1/ammeter_acrel_10r_v1"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
	"s_driver"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(ctx context.Context) c_base.IDevice {
	plugin := ammeter_acrel_10r_v1.NewPlugin(ctx)
	plugin.GetDriverDescription().BuildTime = buildTime
	plugin.GetDriverDescription().CommitHash = commitHash
	return plugin
}

func main() {
	command := c_base.PluginDriverCommand(func() c_base.IDevice {
		return NewPlugin(context.Background())
	})

	// 此处可添加自定义命令
	_ = command.AddCommand(&gcmd.Command{
		Name:  "test",
		Usage: "test",
		Brief: "测试启动",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			_ = s_driver.GetDriverManager(ctx)

			//	cmd.InitDriver(make(map[string]modbus.Client), &c_base.SDeviceConfig{
			//		Id:         "TestAmmeterAcrel10r",
			//		Name:       "安科瑞10R电表",
			//		ProtocolId: "192.168.0.100:5000",
			//		Driver:     "ammeter_acrel_10r_v1",
			//		Enabled:   true,
			//		LogLevel:   "",
			//		Params: map[string]string{
			//			"unitId": "1",
			//		},
			//		DeviceChildren: nil,
			//	}, []*c_base.SProtocolConfig{
			//		{Id: "192.168.0.100:5000", Type: c_base.EModbusTcp, Address: "192.168.0.100:5000", Timeout: 30, LogLevel: "DEBUG", Params: nil},
			//	})
			//
			//	cmd.Block()
			//	return err
			//},
			return err
		},
	})
	command.Run(context.Background())
}
