package c_proto

type SModbusDeviceConfig struct {
	UnitId uint8 `json:"unitId" min:"1" max:"256"` // 单元ID
}
