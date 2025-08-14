package p_gpio_sysfs

// SGpioChipInfo 描述一个 gpiochip 控制器
type SGpioChipInfo struct {
	Name  string // 目录名，如 gpiochip0
	Path  string // 绝对路径
	Label string // 芯片标签
	Base  int    // 起始引脚号
	Ngpio int    // 引脚数量
}

// SGpioInfo 描述一个已导出的 GPIO 引脚目录 gpioN
type SGpioInfo struct {
	Name      string // 目录名，如 gpio23
	Path      string // 绝对路径
	Number    int    // N
	Direction string // 方向(in/out) 若文件缺失则为空
	Value     string // 当前值(0/1) 若不可读则为空
}

// SGpioScanResult 扫描结果
type SGpioScanResult struct {
	Chips []*SGpioChipInfo
	Gpios []*SGpioInfo
}
