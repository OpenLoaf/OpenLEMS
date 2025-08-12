package main

import (
	"common/c_base"
	"context"
	"pcs_elecod/pcs_elecod_mac_v1"
	"services"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/torykit/go-modbus"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(ctx context.Context) c_base.IDriver {
	plugin := pcs_elecod_mac_v1.NewPlugin(ctx)
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

			device := cmd.InitDriver(make(map[string]modbus.Client),
				&c_base.SDriverConfig{
					Id:             "TestPcsElecodV1",
					Name:           "亿兰科PCS测试",
					ProtocolId:     "test_canbus",
					StorageEnable:  false,
					Driver:         "pcs_elecod_mac_v1",
					IsEnable:       true,
					LogLevel:       "",
					Params:         map[string]string{},
					DeviceChildren: nil,
				},
				[]*c_base.SProtocolConfig{{
					Id:       "test_canbus",
					Protocol: "canbus",
					Address:  "can0",
					Timeout:  30,
					LogLevel: "DEBUG",
					Params: map[string]string{
						"BaudRate": "250000",
					},
				}})

			gtimer.SetInterval(ctx, 1*time.Minute, func(ctx context.Context) {
				g.Log().Noticef(ctx, "设备[%s]数据:\n%v", device.GetDeviceConfig().Name, device.GetAllTelemetry(device))
			})

			<-ctx.Done()
			return err
		},
	})

	command.Run(context.Background())
}
