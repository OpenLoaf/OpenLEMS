package ess_pylon_checkwatt_v1

type sEssPylonCheckwattConfig struct {
	AmmeterId string `json:"ammeterId" dc:"电表ID" input_type:"ammeter_device_selector"` // 电表ID， 电表设备选择器
}
