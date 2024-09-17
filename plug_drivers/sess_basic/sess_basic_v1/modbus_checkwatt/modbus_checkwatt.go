package modbus_checkwatt

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/simonvetter/modbus"
	"os"
	"time"
)

func Start(ctx context.Context, handler modbus.RequestHandler, deviceConfig *c_base.SDriverConfig) {
	// 获取配置文件
	var config = &SSessBasicConfig{}
	err := gconv.Scan(deviceConfig.Params, config)
	if err != nil {
		panic(gerror.Newf("场站储能组：%s 配置文件解析失败！", deviceConfig.Name))
	}

	if !config.ModbusServerEnable {
		g.Log().Notice(ctx, "checkwatt modbus server服务未启用!")
		return
	}
	var server *modbus.ModbusServer

	// create the server object
	server, err = modbus.NewServer(&modbus.ServerConfiguration{
		// listen on localhost port 5502
		URL: config.GetModbusUrl(),
		// close idle connections after 30s of inactivity
		Timeout: config.GetTimeout() * time.Second,
		// accept 5 concurrent connections max.
		MaxClients: config.GetMaxClients(),
	}, handler)
	if err != nil {
		g.Log().Errorf(ctx, "创建checkwatt modbus server失败: %v\n", err)
		os.Exit(1)
	}

	g.Log().Infof(ctx, "checkwatt modbus server配置完毕，准备启动！地址: %s\n", config.GetModbusUrl())
	// start accepting client connections
	// note that Start() returns as soon as the server is started
	err = server.Start()
	if err != nil {
		g.Log().Errorf(ctx, "启动checkwatt modbus server失败: %v\n", err)
		os.Exit(1)
	}

	return
}
