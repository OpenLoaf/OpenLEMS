package internal

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
	"time"
)

type PvMessage struct {
}

func Test_start(t *testing.T) {
	config := &c_base.SMqttConfig{
		Enable:   true,
		Url:      "tcp://mqtt.test.hexems.com:1883",
		Username: "goclient",
		Password: "123456",
		Timeout:  0,
		Params:   nil,
	}

	ctx := context.Background()
	client := NewMqttClient(ctx, config)

	client.start()

	count := 1

	for {
		client.SendMap("testtopic/2", map[string]any{
			"test":  "test",
			"count": count,
		})
		count++
		g.Log().Infof(ctx, "send message %d", count)
		time.Sleep(1 * time.Second)
	}

}
