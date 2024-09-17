// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"application/internal/model/entity"
	"common/c_base"
)

type (
	IStation interface {
		GetEnergyStoreStatus() *entity.EnergyStoreStatus
		SetEnergyStorePower(power int32) error
		SetEnergyStoreStatus(status c_base.EEnergyStoreStatus) error
	}
)

var (
	localStation IStation
)

func Station() IStation {
	if localStation == nil {
		panic("implement not found for interface IStation, forgot register?")
	}
	return localStation
}

func RegisterStation(i IStation) {
	localStation = i
}
