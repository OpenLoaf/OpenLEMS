package c_none

import "fmt"

var (
	NoneErr                = fmt.Errorf("this device is not real")
	NoneProtocol           = &sNoneProtocol{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NonePcs                = &sNonePcs{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneBms                = &sNoneBms{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneAmmeter            = &sNoneAmmeter{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NonePv                 = &sNonePv{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneLoad               = &sNoneLoad{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneCharge             = &sNoneCharge{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneGenerator          = &sNoneGenerator{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneCoolingAc          = &sNoneCoolingAc{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneCoolingLiquid      = &sNoneCoolingLiquid{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneFire               = &sNoneFire{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneHumiture           = &sNoneHumiture{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneGpio               = &sNoneGpio{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneEnergyStore        = &sNoneEnergyStore{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneStationEnergyStore = &sNoneStationEnergyStore{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneStationEntrance    = &sNoneStationEntrance{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
)
