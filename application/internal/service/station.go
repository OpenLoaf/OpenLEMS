// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"application/internal/model/entity"
)

type (
	IStation interface {
		GetEssStatus() *entity.EssStatus
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
