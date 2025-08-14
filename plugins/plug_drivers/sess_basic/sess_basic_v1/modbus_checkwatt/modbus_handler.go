package modbus_checkwatt

import (
	"github.com/simonvetter/modbus"
)

type BaseModbusHandler struct {
	Cabinets map[uint8]*EssHandler
}

func (b *BaseModbusHandler) HandleCoils(req *modbus.CoilsRequest) (res []bool, err error) {
	err = modbus.ErrIllegalFunction
	return
}

func (b *BaseModbusHandler) HandleDiscreteInputs(req *modbus.DiscreteInputsRequest) (res []bool, err error) {
	err = modbus.ErrIllegalFunction
	return
}

func (b *BaseModbusHandler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	if ok := b.Cabinets[req.UnitId]; ok != nil {
		return b.Cabinets[req.UnitId].HandleHoldingRegisters(req)
	} else {
		err = modbus.ErrIllegalFunction
		return
	}
}

func (b *BaseModbusHandler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {
	if ok := b.Cabinets[req.UnitId]; ok != nil {
		return b.Cabinets[req.UnitId].HandleInputRegisters(req)
	} else {
		err = modbus.ErrIllegalFunction
		return
	}
}
