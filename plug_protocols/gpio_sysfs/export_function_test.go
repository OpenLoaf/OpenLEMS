package gpio_sysfs

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtimer"
	"testing"
	"time"
)

var Path = "/Users/zhao/Downloads/gpio"

func TestInit(t *testing.T) {
	var err error
	// 初始化测试路径
	//err := gfile.Mkdir(Path)

	// 初始化旗下文件
	err = gfile.PutContents(Path+"/direction", "in")
	if err != nil {
		t.Error(err)
	}

	err = gfile.PutContents(Path+"/value", "1")

	if err != nil {
		t.Error(err)
	}
}

func TestTimmer(t *testing.T) {
	gtimer.SetInterval(context.TODO(), 200*time.Millisecond, func(ctx context.Context) {
		fmt.Println(time.Now())
	})

	time.Sleep(30 * time.Minute)
}
