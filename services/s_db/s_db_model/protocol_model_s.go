package s_db_model

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
	ProtocolFieldAddress  = "address"
	ProtocolFieldTimeout  = "timeout"
	ProtocolFieldLogLevel = "log_level"
	ProtocolFieldParams   = "params"
)

// 协议表结构
type SProtocolModel struct {
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
	Enable    bool   `json:"enable" orm:"enable"`
	CreatedAt string `json:"created_at" orm:"created_at"`
	UpdatedAt string `json:"updated_at" orm:"updated_at"`
}

func (p *SProtocolModel) GetParamsMap() (map[string]string, error) {
	if p.Params == EmptyValue || p.Params == NullValue {
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
func (p *SProtocolModel) Create(ctx context.Context) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Insert(p)
	return err
}

// GetById 根据ID获取协议记录
func (p *SProtocolModel) GetById(ctx context.Context, id string) error {
	return g.Model(TableProtocol).Ctx(ctx).Where(FieldId, id).Scan(p)
}

// GetByName 根据名称获取协议记录
func (p *SProtocolModel) GetByName(ctx context.Context, name string) error {
	return g.Model(TableProtocol).Ctx(ctx).Where(FieldName, name).Scan(p)
}

// GetByType 根据类型获取协议记录
func (p *SProtocolModel) GetByType(ctx context.Context, type_ string) ([]*SProtocolModel, error) {
	var protocols []*SProtocolModel
	err := g.Model(TableProtocol).Ctx(ctx).Where(FieldType, type_).Scan(&protocols)
	return protocols, err
}

// Update 更新协议记录
func (p *SProtocolModel) Update(ctx context.Context) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Where(FieldId, p.Id).Update(p)
	return err
}

// UpdateFields 更新指定字段
func (p *SProtocolModel) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Where(FieldId, p.Id).Update(data)
	return err
}

// Delete 删除协议记录
func (p *SProtocolModel) Delete(ctx context.Context) error {
	_, err := g.Model(TableProtocol).Ctx(ctx).Where(FieldId, p.Id).Delete()
	return err
}
