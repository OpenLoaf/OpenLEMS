package plug_protocol_gpio_sysfs

import (
	"context"
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtimer"
	"plug_protocol_gpio_sysfs/internal"
	"plug_protocol_gpio_sysfs/p_gpio_sysfs"
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

func TestFile(t *testing.T) {
	provider, err := internal.NewGpioSysfsProvider(context.TODO(), &c_base.SProtocolConfig{
		Name:           "",
		Protocol:       "",
		Address:        "",
		Timeout:        0,
		LogLevel:       "INFO",
		Config:         nil,
		Enable:         true,
		DeviceChildren: nil,
	}, &p_gpio_sysfs.SGpioSysfsDeviceConfig{
		SDriverConfig: c_base.SDriverConfig{},
		Direction:     p_gpio_sysfs.EGpioDirectionIn,
		Path:          Path,
		ExportPath:    "",
		ExportPort:    0,
	})
	if err != nil {
		t.Error(err)
	}

	provider.Init(c_base.EGpio)

	time.Sleep(30 * time.Second)
}

func TestTimmer(t *testing.T) {
	gtimer.SetInterval(context.TODO(), 200*time.Millisecond, func(ctx context.Context) {
		fmt.Println(time.Now())
	})

	time.Sleep(30 * time.Minute)
}
