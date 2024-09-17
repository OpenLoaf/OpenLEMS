package main

import (
	"ammeter_acrel_10r_v1/ammeter_acrel_10r_v1"
	"common/c_base"
	"context"
	"driver"
	"github.com/gogf/gf/v2/os/gcmd"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(ctx context.Context) c_base.IDriver {
	plugin := ammeter_acrel_10r_v1.NewPlugin(ctx)
	plugin.GetDescription().BuildTime = buildTime
	plugin.GetDescription().CommitHash = commitHash
	return plugin
}

func main() {
	command := c_base.PluginDriverCommand(func() c_base.IDriver {
		return NewPlugin(context.Background())
	})

	// 此处可添加自定义命令
	_ = command.AddCommand(&gcmd.Command{
		Name:  "test",
		Usage: "test",
		Brief: "测试启动",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			cmd := driver.NewDeviceCmd(ctx)

			cmd.InitDriver(&c_base.SDriverConfig{
				Id:         "TestAmmeterAcrel10r",
				Name:       "安科瑞10R电表",
				ProtocolId: "192.168.0.100:5000",
				Driver:     "ammeter_acrel_10r_v1",
				Enable:     true,
				LogLevel:   "",
				Params: map[string]string{
					"unitId": "1",
				},
				DeviceChildren: nil,
			}, []*c_base.SProtocolConfig{
				{Id: "192.168.0.100:5000", Protocol: c_base.EModbusTcp, Address: "192.168.0.100:5000", Timeout: 30, LogLevel: "DEBUG", Params: nil},
			})

			cmd.Block()
			return err
		},
	})
	command.Run(context.Background())
}
