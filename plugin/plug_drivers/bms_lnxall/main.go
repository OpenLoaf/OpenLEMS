package main

import (
	"bms_lnxall/bms_lnxall_v1"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtimer"
	"services"
	"time"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(ctx context.Context) c_base.IDriver {
	plugin := bms_lnxall_v1.NewPlugin(ctx)
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
			cmd := services.NewDeviceCmd(ctx)

			device := cmd.InitDriver(make(map[string]any), &c_base.SDriverConfig{
				Id:            "TestAmmeterAcrel10r",
				Name:          "协能BMS",
				ProtocolId:    "test_modbus",
				StorageEnable: false,
				Driver:        "bms_lnxall_v1",
				IsEnable:      true,
				LogLevel:      "",
				Params: map[string]string{
					"unitId": "34",
				},
				DeviceChildren: nil,
			}, []*c_base.SProtocolConfig{{
				Id:       "test_modbus",
				Protocol: "modbus_tcp",
				Address:  "127.0.0.1:2509",
				Timeout:  30,
				LogLevel: "DEBUG",
				Params:   nil,
			}})

			gtimer.SetInterval(ctx, 1*time.Second, func(ctx context.Context) {
				g.Log().Noticef(ctx, "设备[%s]数据:\n%v", device.GetDeviceConfig().Name, device.GetAllTelemetry(device))
			})

			<-ctx.Done()
			return err
		},
	})
	command.Run(context.Background())
}
