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
	TableDevice = "device"

	// 字段名
	FieldPid        = "pid"
	FieldProtocolId = "protocol_id"
	FieldDriver     = "driver"
	FieldLogLevel   = "log_level"
	FieldParams     = "params"
)

// 设备表结构
type SDeviceModel struct {
	g.Meta `orm:"table:device"`
	SDatabaseBasic
	Pid                string `json:"pid" orm:"pid"`
	Name               string `json:"name" orm:"name"`
	ProtocolId         string `json:"protocol_id" orm:"protocol_id"`
	Driver             string `json:"driver" orm:"driver"`
	LogLevel           string `json:"logLevel" orm:"log_level"`
	Strategy           string `json:"strategy,omitempty" orm:"strategy"`             // 	策略名称
	StorageEnable      bool   `json:"storageEnable" orm:"storage_enable"`            // 是否存储
	StorageIntervalSec int32  `json:"storageIntervalSec" orm:"storage_interval_sec"` // 存储间隔(秒),0代表默认1分钟，负数代表不存储
	Sort               int    `json:"sort" orm:"sort"`
	Enabled            bool   `json:"enabled" orm:"enabled"`
	Params             string `json:"params" orm:"params"` // 在sqlite中以json字符串形式存储设备参数
}

// GetParamsMap 获取参数的map格式
func (d *SDeviceModel) GetParamsMap() (map[string]string, error) {
	if d.Params == EmptyValue || d.Params == NullValue {
		return map[string]string{}, nil
	}

	// 先反序列化为 map[string]interface{} 来处理混合类型
	var paramsMapInterface map[string]interface{}
	err := gjson.DecodeTo(d.Params, &paramsMapInterface)
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

// SetParamsFromMap 从map设置参数
func (d *SDeviceModel) SetParamsFromMap(paramsMap g.Map) error {
	if paramsMap == nil {
		d.Params = EmptyValue
		return nil
	}

	paramsJSON, err := gjson.Encode(paramsMap)
	if err != nil {
		return err
	}

	d.Params = string(paramsJSON)
	return nil
}

// Create 创建设备记录
func (d *SDeviceModel) Create(ctx context.Context) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Insert(d)
	return err
}

// GetById 根据ID获取设备记录
func (d *SDeviceModel) GetById(ctx context.Context, id string) error {
	return g.Model(TableDevice).Ctx(ctx).Where(FieldId, id).Scan(d)
}

// GetByName 根据名称获取设备记录
func (d *SDeviceModel) GetByName(ctx context.Context, name string) error {
	return g.Model(TableDevice).Ctx(ctx).Where(FieldName, name).Scan(d)
}

// Update 更新设备记录
func (d *SDeviceModel) Update(ctx context.Context) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Where(FieldId, d.Id).Update(d)
	return err
}

// UpdateFields 更新指定字段
func (d *SDeviceModel) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Where(FieldId, d.Id).Update(data)
	return err
}

// Delete 删除设备记录
func (d *SDeviceModel) Delete(ctx context.Context) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Where(FieldId, d.Id).Delete()
	return err
}
