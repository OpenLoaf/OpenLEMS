package c_proto

type SCanbusConfig struct {
	BaudRate uint32 `json:"baudRate" name:"波特率" required:"true" ct:"singleSelect" vt:"int" valueExplain:"10000:10k,20000:20k,50000:50k,100000:100k,125000:125k,250000:250k,500000:500k,1000000:1000k" default:"250000" unit:"bps" dc:"CAN总线通信波特率，标准值：10k, 20k, 50k, 100k, 125k, 250k, 500k, 1000k"`
}
