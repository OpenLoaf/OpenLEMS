package c_proto

type SModbusDeviceConfig struct {
	UnitId uint8 `json:"unitId" name:"ModbusID" min:"1" max:"255" default:"1" required:"true"  step:"1"` // 单元ID
}
