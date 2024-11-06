package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

type Message struct {
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

func start() {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.test.hexems.com:1883").
		SetUsername("emqx_go_client").SetPassword("public")

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	defer c.Disconnect(250)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 发布消息
	for {

		message := Message{
			Content:   "Hello World With Go Client",
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		}
		payload, err := json.Marshal(message)
		if err != nil {
			fmt.Println("json.Marshal failed")
			return
		}

		token := c.Publish("/ttt/1", 0, false, payload)
		token.Wait()

		time.Sleep(10 * time.Second)

		// 等到中断信号，断开链接
	}

	//// 取消订阅
	//if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}

	// 断开连接
	//c.Disconnect(250)
	//time.Sleep(1 * time.Second)
}
