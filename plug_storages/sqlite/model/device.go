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
	FieldId            = "id"
	FieldPid           = "pid"
	FieldGid           = "gid"
	FieldProtocolId    = "protocol_id"
	FieldName          = "name"
	FieldDriver        = "driver"
	FieldLogLevel      = "log_level"
	FieldEnable        = "enable"
	FieldParams        = "params"
	FieldRetentionDays = "retention_days"

	// 特殊值
	NullValue  = "null"
	EmptyValue = ""
)

// 设备表结构
type Device struct {
	g.Meta     `orm:"table:device"`
	Id         string `json:"id" orm:"id,primary"`
	Pid        string `json:"pid" orm:"pid"`
	Gid        uint   `json:"gid" orm:"gid"`
	ProtocolId string `json:"protocol_id" orm:"protocol_id"`
	Name       string `json:"name" orm:"name"`
	Driver     string `json:"driver" orm:"driver"`
	LogLevel   string `json:"log_level" orm:"log_level"`
	Enable     bool   `json:"enable" orm:"enable"`
	// 在sqlite中以json字符串形式存储设备参数
	Params        string `json:"params" orm:"params"`
	RetentionDays int    `json:"retention_days" orm:"retention_days"`
}

// GetParamsMap 获取参数的map格式
func (d *Device) GetParamsMap() (map[string]string, error) {
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
func (d *Device) SetParamsFromMap(paramsMap g.Map) error {
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
func (d *Device) Create(ctx context.Context) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Insert(d)
	return err
}

// GetById 根据ID获取设备记录
func (d *Device) GetById(ctx context.Context, id string) error {
	return g.Model(TableDevice).Ctx(ctx).Where(FieldId, id).Scan(d)
}

// GetByName 根据名称获取设备记录
func (d *Device) GetByName(ctx context.Context, name string) error {
	return g.Model(TableDevice).Ctx(ctx).Where(FieldName, name).Scan(d)
}

// Update 更新设备记录
func (d *Device) Update(ctx context.Context) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Where(FieldId, d.Id).Update(d)
	return err
}

// UpdateFields 更新指定字段
func (d *Device) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Where(FieldId, d.Id).Update(data)
	return err
}

// Delete 删除设备记录
func (d *Device) Delete(ctx context.Context) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Where(FieldId, d.Id).Delete()
	return err
}

// DeleteById 根据ID删除设备记录
func DeleteDeviceById(ctx context.Context, id string) error {
	_, err := g.Model(TableDevice).Ctx(ctx).Where(FieldId, id).Delete()
	return err
}

// GetAll 获取所有设备记录
func GetAllDevices(ctx context.Context) ([]*Device, error) {
	var devices []*Device
	err := g.Model(TableDevice).Ctx(ctx).Scan(&devices)
	return devices, err
}

// GetByCondition 根据条件获取设备记录
func GetDevicesByCondition(ctx context.Context, condition g.Map) ([]*Device, error) {
	var devices []*Device
	err := g.Model(TableDevice).Ctx(ctx).Where(condition).Scan(&devices)
	return devices, err
}

// GetByPid 根据父设备ID获取子设备列表
func GetDevicesByPid(ctx context.Context, pid string) ([]*Device, error) {
	var devices []*Device
	err := g.Model(TableDevice).Ctx(ctx).Where(FieldPid, pid).Scan(&devices)
	return devices, err
}

// GetByGid 根据组ID获取设备列表
func GetDevicesByGid(ctx context.Context, gid uint) ([]*Device, error) {
	var devices []*Device
	err := g.Model(TableDevice).Ctx(ctx).Where(FieldGid, gid).Scan(&devices)
	return devices, err
}

// GetByProtocolId 根据协议ID获取设备列表
func GetDevicesByProtocolId(ctx context.Context, protocolId string) ([]*Device, error) {
	var devices []*Device
	err := g.Model(TableDevice).Ctx(ctx).Where(FieldProtocolId, protocolId).Scan(&devices)
	return devices, err
}

// GetEnabledDevices 获取所有启用的设备
func GetEnabledDevices(ctx context.Context) ([]*Device, error) {
	var devices []*Device
	err := g.Model(TableDevice).Ctx(ctx).Where(FieldEnable, true).Scan(&devices)
	return devices, err
}

// Count 获取设备总数
func CountDevices(ctx context.Context) (int, error) {
	count, err := g.Model(TableDevice).Ctx(ctx).Count()
	return count, err
}

// CountByCondition 根据条件获取设备数量
func CountDevicesByCondition(ctx context.Context, condition g.Map) (int, error) {
	count, err := g.Model(TableDevice).Ctx(ctx).Where(condition).Count()
	return count, err
}

// Paginate 分页获取设备列表
func PaginateDevices(ctx context.Context, page, pageSize int) ([]*Device, error) {
	var devices []*Device
	err := g.Model(TableDevice).Ctx(ctx).Page(page, pageSize).Scan(&devices)
	return devices, err
}
