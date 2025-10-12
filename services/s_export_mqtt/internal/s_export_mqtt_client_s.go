package internal

import (
	"context"
	"fmt"
	"strings"
	"time"

	"common"
	"common/c_base"
	"common/c_log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
)

// SMqttClient MQTT客户端封装
type SMqttClient struct {
	config       *SMqttExportConfig // MQTT配置
	client       mqtt.Client        // MQTT客户端
	ticker       *time.Ticker       // 定时器
	ctx          context.Context    // 上下文
	cancel       context.CancelFunc // 取消函数
	formatter    IDataFormatter     // 数据格式化器
	systemNumber string             // 系统序列号
}

// NewMqttClient 创建新的MQTT客户端
func NewMqttClient(config *SMqttExportConfig, formatter IDataFormatter, systemNumber string) *SMqttClient {
	return &SMqttClient{
		config:       config,
		formatter:    formatter,
		systemNumber: systemNumber,
	}
}

// Start 启动MQTT客户端
func (c *SMqttClient) Start(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)

	// 创建MQTT客户端
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", c.config.ServerAddress, c.config.ServerPort))
	opts.SetClientID(fmt.Sprintf("ems_mqtt_export_%d", time.Now().Unix()))
	opts.SetAutoReconnect(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetMaxReconnectInterval(30 * time.Second)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(10 * time.Second)

	// 设置连接丢失处理
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		c_log.Warningf(c.ctx, "MQTT连接丢失: %v", err)
	})

	// 设置连接成功处理
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		c_log.Infof(c.ctx, "MQTT连接成功: %s:%d", c.config.ServerAddress, c.config.ServerPort)
	})

	c.client = mqtt.NewClient(opts)

	// 连接到MQTT服务器
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		return errors.Wrapf(token.Error(), "连接MQTT服务器失败: %s:%d", c.config.ServerAddress, c.config.ServerPort)
	}

	// 启动定时推送
	c.ticker = time.NewTicker(time.Duration(c.config.UploadPeriod) * time.Second)
	go c.publishLoop()

	c_log.Infof(c.ctx, "MQTT客户端启动成功: %s:%d, 推送周期: %d秒", c.config.ServerAddress, c.config.ServerPort, c.config.UploadPeriod)
	return nil
}

// Stop 停止MQTT客户端
func (c *SMqttClient) Stop() error {
	if c.cancel != nil {
		c.cancel()
	}

	if c.ticker != nil {
		c.ticker.Stop()
	}

	if c.client != nil && c.client.IsConnected() {
		c.client.Disconnect(250)
	}

	c_log.Infof(c.ctx, "MQTT客户端已停止: %s:%d", c.config.ServerAddress, c.config.ServerPort)
	return nil
}

// publishLoop 定时推送循环
func (c *SMqttClient) publishLoop() {
	for {
		select {
		case <-c.ctx.Done():
			c_log.Debugf(c.ctx, "MQTT推送循环已停止")
			return
		case <-c.ticker.C:
			if err := c.publishDeviceData(); err != nil {
				c_log.Errorf(c.ctx, "推送设备数据失败: %v", err)
			}
		}
	}
}

// publishDeviceData 推送设备数据
func (c *SMqttClient) publishDeviceData() error {
	if !c.client.IsConnected() {
		return errors.New("MQTT客户端未连接")
	}

	// 遍历配置的设备ID列表
	for _, deviceId := range c.config.DeviceIds {
		// 获取设备实例
		device := common.GetDeviceManager().GetDeviceById(deviceId)
		if device == nil {
			c_log.Warningf(c.ctx, "设备不存在: %s", deviceId)
			continue
		}

		// 检查设备连接状态
		if device.GetProtocolStatus() != common.GetDeviceManager().GetDeviceById(deviceId).GetProtocolStatus() {
			c_log.Debugf(c.ctx, "设备未连接，跳过推送: %s", deviceId)
			continue
		}

		// 获取设备所有点位数据
		deviceData := c_base.GetAllTelemetry(device)
		if len(deviceData) == 0 {
			c_log.Debugf(c.ctx, "设备无数据，跳过推送: %s", deviceId)
			continue
		}

		// 格式化数据
		jsonData, err := c.formatter.Format(c.ctx, deviceId, deviceData, c.systemNumber)
		if err != nil {
			c_log.Errorf(c.ctx, "格式化设备数据失败: 设备ID=%s, 错误=%v", deviceId, err)
			continue
		}

		// 构建topic
		topic := c.buildTopic()
		if topic == "" {
			c_log.Errorf(c.ctx, "构建topic失败: 设备ID=%s", deviceId)
			continue
		}

		// 发布MQTT消息
		if token := c.client.Publish(topic, 0, false, jsonData); token.Wait() && token.Error() != nil {
			c_log.Errorf(c.ctx, "发布MQTT消息失败: 设备ID=%s, Topic=%s, 错误=%v", deviceId, topic, token.Error())
			continue
		}

		c_log.Debugf(c.ctx, "成功推送设备数据: 设备ID=%s, Topic=%s, 数据长度=%d", deviceId, topic, len(jsonData))
	}

	return nil
}

// buildTopic 构建MQTT topic
func (c *SMqttClient) buildTopic() string {
	// 如果配置了自定义pushChannel且rewriteChannel=true，使用自定义
	if c.config.RewriteChannel && c.config.PushChannel != "" {
		return c.config.PushChannel
	}

	// 使用formatter的模板，替换{system_number}为实际值
	topicTemplate := c.formatter.GetTopicTemplate()
	return strings.Replace(topicTemplate, "{system_number}", c.systemNumber, -1)
}
