package internal

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	ctx    context.Context
	config *c_base.SMqttConfig
	client mqtt.Client
}

func NewMqttClient(ctx context.Context, config *c_base.SMqttConfig) *MqttClient {
	return &MqttClient{
		ctx:    ctx,
		config: config,
	}
}

func (c *MqttClient) start() {
	if c.config == nil {
		panic("mqtt config is nil")
	}
	if c.config.Timeout == 0 {
		c.config.Timeout = 5000
	}

	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(c.config.Url).SetUsername(c.config.Username).SetPassword(c.config.Password)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	//opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(60 * time.Second)

	c.client = mqtt.NewClient(opts)

	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (c *MqttClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) {
	if token := c.client.Subscribe(topic, qos, callback); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (c *MqttClient) Unsubscribe(topics ...string) {
	if token := c.client.Unsubscribe(topics...); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (c *MqttClient) SendMap(topic string, payload map[string]any) {
	c.SendStruct(topic, payload)
}

func (c *MqttClient) SendStruct(topic string, payload interface{}) {
	j, err := gjson.Encode(payload)
	if err != nil {
		panic(err)
	}
	c.SendJson(topic, string(j))
}

func (c *MqttClient) SendJson(topic string, payload string) {
	token := c.client.Publish(topic, 0, false, payload)
	token.WaitTimeout(time.Second * time.Duration(c.config.Timeout))
	if token.Error() != nil {
		panic(token.Error())
	}
}

func (c *MqttClient) Disconnect() {
	c.client.Disconnect(250)
}
