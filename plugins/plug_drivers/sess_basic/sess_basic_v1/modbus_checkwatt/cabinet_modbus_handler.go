package modbus_checkwatt

import (
	"common/c_base"
	"common/c_log"
	"common/c_type"
	"context"
	"encoding/binary"
	"github.com/simonvetter/modbus"
	"sync"
)

type EssHandler struct {
	Ctx context.Context
	c_type.IStationEnergyStore
}

func (c *EssHandler) HandleCoils(req *modbus.CoilsRequest) (res []bool, err error) {
	return nil, modbus.ErrIllegalFunction
}

func (c *EssHandler) HandleDiscreteInputs(req *modbus.DiscreteInputsRequest) (res []bool, err error) {
	return nil, modbus.ErrIllegalFunction
}

func (c *EssHandler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	if req.IsWrite {
		res = append(res, req.Args...)
		return
	}

	for i := 0; i < int(req.Quantity); i++ {
		res = append(res, 0xFFFF)
	}
	return
}

func (c *EssHandler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {
	res = make([]uint16, req.Quantity)
	// 全部初始化为0xFFFF
	for i := 0; i < int(req.Quantity); i++ {
		res[i] = 0xFFFF
	}
	var wg sync.WaitGroup
	for regAddr := req.Addr; regAddr < req.Addr+req.Quantity; regAddr++ {
		index := regAddr - req.Addr
		switch regAddr {
		case 0xF000: // 0xF000 SOC
			updateValue(index, res, &wg, 10, 0, c.GetSoc)
		case 0xF001: // 0xF001 Active Power
			updateValue(index, res, &wg, 10, 0, c.GetPower)
		case 0xF002: // 0xF002 Reactive Power
			updateValue(index, res, &wg, 10, 0, c.GetReactivePower)
		case 0xF003: // 0xF003 Energy storage status
			updateValue(index, res, &wg, 1, 0, func() (int, error) {
				status, err := c.GetStatus()
				if err != nil {
					return 0xFF, err
				}
				switch status {
				case c_base.EPcsStatusOff, c_base.EPcsStatusSync:
					return 0, nil
				case c_base.EPcsStatusUnknown, c_base.EPcsStatusFault:
					return 0xFF, nil
				case c_base.EPcsStatusCharge:
					return 0x02, nil
				case c_base.EPcsStatusDischarge:
					return 0x03, nil
				case c_base.EPcsStatusStandby:
					return 0x01, nil
				default:
				}
				return 0xFF, err
			})
		case 0xF004: // 0xF004 Maximum allowable charging power
			updateValue(index, res, &wg, 10, 0, c.GetMaxInputPower)
		case 0xF005: // 0xF005 Maximum allowable discharge power
			updateValue(index, res, &wg, 10, 0, c.GetMaxOutputPower)
		case 0xF006: // 0xF006 AC Voltage
			res[index] = 0xFFFF
			//updateValue(index, res, &wg, 10, 0, c.GetA)
		case 0xF007: // 0xF007 AC Frequency
			//updateValue(index, res, &wg, 10, 0, c.GetF)
			res[index] = 0xFFFF
		case 0xF008: // 0xF008 Dc Power
			updateValue(index, res, &wg, 10, 0, c.GetDcPower)
		case 0xF009: // 0xF009 Dc Voltage
			//updateValue(index, res, &wg, 10, 0, c.Get)
			res[index] = 0xFFFF
		case 0xF00A: // 0xF00A Dc Current
			//updateValue(index, res, &wg, 10, 0, c.GetDcCurrent)
			res[index] = 0xFFFF
		case 0xF00B: // 0xF00B IGBT Average Temperature
			res[index] = 0xFFFF

		case 0xF00C: // 0xF00C Battery Average Temperature
			res[index] = 0xFFFF
			updateValue(index, res, &wg, 10, 0, c.GetCellAvgTemp)
		case 0xF00D: // 0xF00D Battery Average Voltage
			res[index] = 0xFFFF
			updateValue(index, res, &wg, 10, 0, c.GetCellAvgVoltage)
		case 0xF00E: // 0xF00E Rated Power
			updateValue(index, res, &wg, 1, int32(0), func() (int32, error) {
				return c.GetRatedPower(), nil
			})
		case 0xF00F: // 0xF00F Rated RatedPower
			updateValue(index, res, &wg, 1, 0, c.GetCapacity)
		case 0xF010: // 0xF010 Alarm Level
			res[index] = 0xFFFF
		case 0xF011: // 0xF011 Fire System Status
			res[index] = 0xFFFF
		case 0xF012: // 0xF012 Smoke Detection Alarm
			res[index] = 0xFFFF
		case 0xF013: // 0xF013 Carbonmonoxide Concentration
			res[index] = 0xFFFF
		case 0xF014: // 0xF014 Charge Today
			updateValue(index, res, &wg, 10, 0, c.GetTodayIncomingQuantity)
		case 0xF015: // 0xF015 Discharge Today
			updateValue(index, res, &wg, 10, 0, c.GetTodayOutgoingQuantity)
		case 0xF016: // 0xF016 Charge History
			updateDoubleValue(index, res, &wg, 10, 0, c.GetHistoryIncomingQuantity)
			regAddr++
		case 0xF018: // 0xF018 Discharge History
			updateDoubleValue(index, res, &wg, 10, 0, c.GetHistoryOutgoingQuantity)
			regAddr++

		default:
			return nil, modbus.ErrIllegalFunction
		}
	}
	//res = append(res, 0x1230)

	wg.Wait()
	c_log.Debugf(c.Ctx, "Req: SourceAddress:%v,点位:%v,数量:%v Res: %v", req.ClientAddr, req.Addr, req.Quantity, res)

	return
}
func updateValue[T c_base.Number](index uint16, result []uint16, wg *sync.WaitGroup, factor, offset T, fc func() (T, error)) {
	updateUint16(index, result, wg, func() uint16 {
		value, err := fc()
		if err != nil {
			return 0xFFFF
		}
		return uint16(value*factor + offset)
	})
}

func updateDoubleValue[T c_base.Number](index uint16, result []uint16, wg *sync.WaitGroup, factor, offset T, fc func() (T, error)) {
	updateUint32(index, result, wg, func() uint32 {
		value, err := fc()
		if err != nil {
			return 0xFFFF
		}
		return uint32(value*factor + offset)
	})
}

func updateUint16(index uint16, result []uint16, wg *sync.WaitGroup, fc func() uint16) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		value := fc()
		result[index] = value
	}()
}

func updateUint32(index uint16, result []uint16, wg *sync.WaitGroup, fc func() uint32) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		value := fc()
		bt := make([]byte, 4)
		binary.BigEndian.PutUint32(bt, value)
		result[index] = binary.BigEndian.Uint16(bt[0:2])
		result[index+1] = binary.BigEndian.Uint16(bt[2:4])
	}()
}
