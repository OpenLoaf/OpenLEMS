package model

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据库相关常量
const (
	// 表名
	TableProtocol = "protocol"

	// 字段名
	ProtocolFieldId        = "id"
	ProtocolFieldName      = "name"
	ProtocolFieldType      = "type"
	ProtocolFieldAddress   = "address"
	ProtocolFieldTimeout   = "timeout"
	ProtocolFieldLogLevel  = "log_level"
	ProtocolFieldParams    = "params"
	ProtocolFieldSort      = "sort"
	ProtocolFieldCreatedAt = "created_at"
	ProtocolFieldUpdatedAt = "updated_at"

	// 特殊值
	ProtocolNullValue  = "null"
	ProtocolEmptyValue = ""
)

// 协议表结构
type Protocol struct {
	g.Meta   `orm:"table:protocol"`
	Id       string `json:"id" orm:"id,primary"`
	Name     string `json:"name" orm:"name"`
	Type     string `json:"type" orm:"type"`
	Address  string `json:"address" orm:"address"`
	Timeout  int64  `json:"timeout" orm:"timeout"`
	LogLevel string `json:"log_level" orm:"log_level"`
	// 在sqlite中以json字符串形式存储设备参数
	Params    string `json:"params" orm:"params"`
	Sort      int    `json:"sort" orm:"sort"`
	CreatedAt string `json:"created_at" orm:"created_at"`
	UpdatedAt string `json:"updated_at" orm:"updated_at"`
}

func (p *Protocol) GetParamsMap() (map[string]string, error) {
	if p.Params == ProtocolEmptyValue || p.Params == ProtocolNullValue {
		return map[string]string{}, nil
	}

	// 先反序列化为 map[string]interface{} 来处理混合类型
	var paramsMapInterface map[string]interface{}
	err := gjson.DecodeTo(p.Params, &paramsMapInterface)
	if err != nil {
		return nil, err
	}

	// 转换为 map[string]string
	paramsMap := make(map[string]string)
	for key, value := range paramsMapInterface {
		paramsMap[key] = fmt.Sprintf("%v", value)
	}

	return paramsMap, nil
}

// Create 创建协议记录
func (p *Protocol) Create(ctx context.Context) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Insert(p)
	return err
}

// GetById 根据ID获取协议记录
func (p *Protocol) GetById(ctx context.Context, id string) error {
	return g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldId, id).Scan(p)
}

// GetByName 根据名称获取协议记录
func (p *Protocol) GetByName(ctx context.Context, name string) error {
	return g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldName, name).Scan(p)
}

// GetByType 根据类型获取协议记录
func (p *Protocol) GetByType(ctx context.Context, type_ string) ([]*Protocol, error) {
	var protocols []*Protocol
	err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldType, type_).Scan(&protocols)
	return protocols, err
}

// Update 更新协议记录
func (p *Protocol) Update(ctx context.Context) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldId, p.Id).Update(p)
	return err
}

// UpdateFields 更新指定字段
func (p *Protocol) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldId, p.Id).Update(data)
	return err
}

// Delete 删除协议记录
func (p *Protocol) Delete(ctx context.Context) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldId, p.Id).Delete()
	return err
}

// DeleteById 根据ID删除协议记录
func DeleteProtocolById(ctx context.Context, id string) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldId, id).Delete()
	return err
}

// GetAll 获取所有协议记录
func GetAllProtocols(ctx context.Context) ([]*Protocol, error) {
	var protocols []*Protocol
	err := g.Model(TableProtocol).Ctx(ctx).Scan(&protocols)
	return protocols, err
}

// GetByCondition 根据条件获取协议记录
func GetProtocolsByCondition(ctx context.Context, condition g.Map) ([]*Protocol, error) {
	var protocols []*Protocol
	err := g.Model(TableProtocol).Ctx(ctx).Where(condition).Scan(&protocols)
	return protocols, err
}

// GetByAddress 根据地址获取协议列表
func GetProtocolsByAddress(ctx context.Context, address string) ([]*Protocol, error) {
	var protocols []*Protocol
	err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldAddress, address).Scan(&protocols)
	return protocols, err
}

// GetByLogLevel 根据日志级别获取协议列表
func GetProtocolsByLogLevel(ctx context.Context, logLevel string) ([]*Protocol, error) {
	var protocols []*Protocol
	err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldLogLevel, logLevel).Scan(&protocols)
	return protocols, err
}

// Count 获取协议总数
func CountProtocols(ctx context.Context) (int, error) {
	count, err := g.Model(TableProtocol).Ctx(ctx).Count()
	return count, err
}

// CountByCondition 根据条件获取协议数量
func CountProtocolsByCondition(ctx context.Context, condition g.Map) (int, error) {
	count, err := g.Model(TableProtocol).Ctx(ctx).Where(condition).Count()
	return count, err
}

// Paginate 分页获取协议列表
func PaginateProtocols(ctx context.Context, page, pageSize int) ([]*Protocol, error) {
	var protocols []*Protocol
	err := g.Model(TableProtocol).Ctx(ctx).Page(page, pageSize).Scan(&protocols)
	return protocols, err
}

// Exists 检查协议是否存在
func (p *Protocol) Exists(ctx context.Context) (bool, error) {
	count, err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldId, p.Id).Count()
	return count > 0, err
}

// ExistsByName 根据名称检查协议是否存在
func ExistsProtocolByName(ctx context.Context, name string) (bool, error) {
	count, err := g.Model(TableProtocol).Ctx(ctx).Where(ProtocolFieldName, name).Count()
	return count > 0, err
}

// GetByTimeout 根据超时时间范围获取协议列表
func GetProtocolsByTimeoutRange(ctx context.Context, minTimeout, maxTimeout int64) ([]*Protocol, error) {
	var protocols []*Protocol
	err := g.Model(TableProtocol).Ctx(ctx).WhereBetween(ProtocolFieldTimeout, minTimeout, maxTimeout).Scan(&protocols)
	return protocols, err
}

// GetAllProtocolsOrderBySort 获取所有协议记录，按sort字段排序
func GetAllProtocolsOrderBySort(ctx context.Context) ([]*Protocol, error) {
	var protocols []*Protocol
	err := g.Model(TableProtocol).Ctx(ctx).Order(ProtocolFieldSort).Scan(&protocols)
	return protocols, err
}
