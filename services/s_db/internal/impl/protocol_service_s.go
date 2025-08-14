package impl

import (
	"context"
	"fmt"
	"s_db/s_db_interface"
	"s_db/s_db_model"
	"sync"

	"github.com/google/uuid"
)

type sProtocolServiceImpl struct {
}

var (
	protocolManageInstance s_db_interface.IProtocolService
	protocolManageOnce     sync.Once
)

func GetProtocolService() s_db_interface.IProtocolService {
	protocolManageOnce.Do(func() {
		protocolManageInstance = &sProtocolServiceImpl{}
	})
	return protocolManageInstance
}

func (s *sProtocolServiceImpl) GetProtocolList(ctx context.Context, type_ string) ([]*s_db_model.SProtocolModel, error) {
	protocols, err := (&s_db_model.SProtocolModel{}).GetByType(ctx, type_)
	if err != nil {
		return nil, err
	}
	return protocols, nil
}

func (s *sProtocolServiceImpl) UpdateProtocol(ctx context.Context, protocolId string, data map[string]interface{}) error {
	protocol := &s_db_model.SProtocolModel{}
	// 先根据ID获取协议对象
	err := protocol.GetById(ctx, protocolId)
	if err != nil {
		return err
	}

	// 从map中提取参数并更新字段
	if name, ok := data["protocolName"].(string); ok {
		protocol.Name = name
	}

	if protocolType, ok := data["protocolType"].(string); ok {
		protocol.Type = protocolType
	}

	if address, ok := data["protocolAddress"].(string); ok {
		// 处理地址和端口
		var fullAddress string
		if port, exists := data["protocolPort"]; exists {
			if portInt, ok := port.(int); ok && portInt > 0 {
				fullAddress = fmt.Sprintf("%s:%d", address, portInt)
			} else {
				fullAddress = address
			}
		} else {
			fullAddress = address
		}
		protocol.Address = fullAddress
	}

	if timeout, ok := data["protocolTimeout"].(int); ok {
		protocol.Timeout = int64(timeout)
	}

	if logLevel, ok := data["protocolLogLevel"].(string); ok {
		protocol.LogLevel = logLevel
	}

	if params, ok := data["protocolParams"].(string); ok {
		protocol.Params = params
	}

	// 执行更新操作
	return protocol.Update(ctx)
}

func (s *sProtocolServiceImpl) CreateProtocol(ctx context.Context, data map[string]interface{}) (string, error) {
	// 生成协议ID
	protocolId := uuid.NewString()

	// 处理地址和端口
	address := ""
	if addr, ok := data["protocolAddress"].(string); ok {
		address = addr
		if port, exists := data["protocolPort"]; exists {
			if portInt, ok := port.(int); ok && portInt > 0 {
				address = fmt.Sprintf("%s:%d", addr, portInt)
			}
		}
	}

	// 处理协议参数
	params := ""
	if p, ok := data["protocolParams"].(string); ok && p != "" {
		params = p
	}

	// 创建协议对象
	protocol := &s_db_model.SProtocolModel{
		Id:       protocolId,
		Name:     data["protocolName"].(string),
		Type:     data["protocolType"].(string),
		Address:  address,
		Timeout:  int64(data["protocolTimeout"].(int)),
		LogLevel: data["protocolLogLevel"].(string),
		Params:   params,
	}

	// 保存到数据库
	err := protocol.Create(ctx)
	if err != nil {
		return "", err
	}

	return protocolId, nil
}

func (s *sProtocolServiceImpl) DeleteProtocol(ctx context.Context, protocolId string) error {
	// 先检查协议是否存在
	protocol := &s_db_model.SProtocolModel{}
	err := protocol.GetById(ctx, protocolId)
	if err != nil {
		return err
	}

	// 删除协议
	return s_db_model.DeleteProtocolById(ctx, protocolId)
}
