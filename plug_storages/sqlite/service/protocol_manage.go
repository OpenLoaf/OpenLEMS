package service

import (
	"context"
	"sqlite/model"
)

type IProtocolManage interface {
	GetProtocolList(ctx context.Context, type_ string) ([]*model.Protocol, error)
}

type sProtocolManage struct {
}

func NewProtocolManage(ctx context.Context) IProtocolManage {
	return &sProtocolManage{
		// TODO: implement GetProtocolList method
	}
}

func (s *sProtocolManage) GetProtocolList(ctx context.Context, type_ string) ([]*model.Protocol, error) {
	protocols, err := (&model.Protocol{}).GetByType(ctx, type_)
	if err != nil {
		return nil, err
	}
	return protocols, nil
}
