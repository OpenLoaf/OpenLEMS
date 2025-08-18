package modbus_checkwatt

import (
	"common/c_base"
	"common/c_log"
	"context"
	"fmt"
	"github.com/simonvetter/modbus"
	"os"
	"time"
)

func Start(ctx context.Context, handler modbus.RequestHandler, deviceConfig *c_base.SDeviceConfig) {
	// 获取配置文件
	var config = &SSessBasicConfig{}
	err := deviceConfig.ScanParams(config)
	if err != nil {
		panic(fmt.Errorf("场站储能组：%s 配置文件解析失败！", deviceConfig.Name))
	}

	if !config.ModbusServerEnable {
		c_log.Notice(ctx, "checkwatt modbus server服务未启用!")
		return
	}
	var server *modbus.ModbusServer

	fmt.Print(config, "config")
	s := &modbus.ServerConfiguration{
		// listen on localhost port 5502
		URL: config.GetModbusUrl(),
		// close idle connections after 30s of inactivity
		Timeout: config.GetTimeout() * time.Second,
		// accept 5 concurrent connections max.
		MaxClients: config.GetMaxClients(),
	}
	c_log.Info(ctx, "%v", s)
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
		c_log.Errorf(ctx, "创建checkwatt modbus server失败: %v\n", err)
		os.Exit(1)
	}

	c_log.Infof(ctx, "checkwatt modbus server配置完毕，准备启动！地址: %s\n", config.GetModbusUrl())
	// start accepting client connections
	// note that Start() returns as soon as the server is started
	err = server.Start()
	if err != nil {
		c_log.Errorf(ctx, "启动checkwatt modbus server失败: %v\n", err)
		os.Exit(1)
	}

	return
}
