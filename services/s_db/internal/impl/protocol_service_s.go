package impl

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"s_db/s_db_interface"
	"s_db/s_db_model"
	"sync"

	"github.com/google/uuid"
)

type sProtocolServiceImpl struct {
	tableProtocol *s_db_model.SProtocolModel
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
	return s.DeleteProtocolById(ctx, protocolId)
}

// DeleteProtocolById 根据ID删除协议记录
func (s *sProtocolServiceImpl) DeleteProtocolById(ctx context.Context, id string) error {
	_, err := g.Model(s.tableProtocol).Ctx(ctx).Where(s_db_model.FieldId, id).Delete()
	return err
}

// GetAllProtocols 获取所有协议记录
func (s *sProtocolServiceImpl) GetAllProtocols(ctx context.Context) ([]*s_db_model.SProtocolModel, error) {
	var protocols []*s_db_model.SProtocolModel
	err := g.Model(s.tableProtocol).Ctx(ctx).Scan(&protocols)
	return protocols, err
}

//
//// GetByCondition 根据条件获取协议记录
//func GetProtocolsByCondition(ctx context.Context, condition g.Map) ([]*SProtocolModel, error) {
//	var protocols []*SProtocolModel
//	err := g.Model(s.tableProtocol).Ctx(ctx).Where(condition).Scan(&protocols)
//	return protocols, err
//}
//
//// GetByAddress 根据地址获取协议列表
//func GetProtocolsByAddress(ctx context.Context, address string) ([]*SProtocolModel, error) {
//	var protocols []*SProtocolModel
//	err := g.Model(s.tableProtocol).Ctx(ctx).Where(ProtocolFieldAddress, address).Scan(&protocols)
//	return protocols, err
//}
//
//// GetByLogLevel 根据日志级别获取协议列表
//func GetProtocolsByLogLevel(ctx context.Context, logLevel string) ([]*SProtocolModel, error) {
//	var protocols []*SProtocolModel
//	err := g.Model(s.tableProtocol).Ctx(ctx).Where(ProtocolFieldLogLevel, logLevel).Scan(&protocols)
//	return protocols, err
//}
//
//// Count 获取协议总数
//func CountProtocols(ctx context.Context) (int, error) {
//	count, err := g.Model(s.tableProtocol).Ctx(ctx).Count()
//	return count, err
//}
//
//// CountByCondition 根据条件获取协议数量
//func CountProtocolsByCondition(ctx context.Context, condition g.Map) (int, error) {
//	count, err := g.Model(s.tableProtocol).Ctx(ctx).Where(condition).Count()
//	return count, err
//}
//
//// Paginate 分页获取协议列表
//func PaginateProtocols(ctx context.Context, page, pageSize int) ([]*SProtocolModel, error) {
//	var protocols []*SProtocolModel
//	err := g.Model(s.tableProtocol).Ctx(ctx).Page(page, pageSize).Scan(&protocols)
//	return protocols, err
//}
//
//// Exists 检查协议是否存在
//func (p *SProtocolModel) Exists(ctx context.Context) (bool, error) {
//	count, err := g.Model(s.tableProtocol).Ctx(ctx).Where(ProtocolFieldId, p.Id).Count()
//	return count > 0, err
//}
//
//// ExistsByName 根据名称检查协议是否存在
//func ExistsProtocolByName(ctx context.Context, name string) (bool, error) {
//	count, err := g.Model(s.tableProtocol).Ctx(ctx).Where(ProtocolFieldName, name).Count()
//	return count > 0, err
//}
//
//// GetByTimeout 根据超时时间范围获取协议列表
//func GetProtocolsByTimeoutRange(ctx context.Context, minTimeout, maxTimeout int64) ([]*SProtocolModel, error) {
//	var protocols []*SProtocolModel
//	err := g.Model(s.tableProtocol).Ctx(ctx).WhereBetween(ProtocolFieldTimeout, minTimeout, maxTimeout).Scan(&protocols)
//	return protocols, err
//}
//
//// GetAllProtocolsOrderBySort 获取所有协议记录，按sort字段排序
//func GetAllProtocolsOrderBySort(ctx context.Context) ([]*SProtocolModel, error) {
//	var protocols []*SProtocolModel
//	err := g.Model(s.tableProtocol).Ctx(ctx).Order(ProtocolFieldSort).Scan(&protocols)
//	return protocols, err
//}
