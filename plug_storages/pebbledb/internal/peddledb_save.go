package internal

import (
	"common/c_base"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sPebbledbConfig struct {
	Path string `json:"path,omitempty" dc:"数据库文件路径"`
}

type Pebbledb struct {
	db           *pebble.DB
	ctx          context.Context
	pebbleConfig *sPebbledbConfig
}

func NewPebbledb(ctx context.Context) c_base.IStorage {

	var pebbleConfig = &sPebbledbConfig{
		Path: "./out/pebbledb",
	}

	// 创建PebbleDB实例
	db, err := pebble.Open(pebbleConfig.Path, &pebble.Options{})
	if err != nil {
		panic(gerror.Newf("创建Pebbledb实例失败！%v", err))
	}

	d := &Pebbledb{
		db:           db,
		ctx:          ctx,
		pebbleConfig: pebbleConfig,
	}

	g.Log().Infof(ctx, "PebbleDB 初始化成功，数据路径: %s", pebbleConfig.Path)
	return d
}

func (p *Pebbledb) SaveDevices(deviceId string, deviceType c_base.EDeviceType, fields map[string]any) error {
	if len(fields) == 0 {
		return nil
	}

	// 创建设备数据结构
	deviceData := map[string]any{
		"device_id":   deviceId,
		"device_type": string(deviceType),
		"timestamp":   time.Now().Unix(),
		"fields":      fields,
	}

	// 序列化数据
	data, err := json.Marshal(deviceData)
	if err != nil {
		return gerror.Wrapf(err, "序列化设备数据失败, deviceId: %s", deviceId)
	}

	// 生成键名：devices/{deviceType}/{deviceId}/{timestamp}
	key := fmt.Sprintf("devices/%s/%s/%d", string(deviceType), deviceId, time.Now().UnixNano())

	// 写入数据
	err = p.db.Set([]byte(key), data, pebble.Sync)
	if err != nil {
		return gerror.Wrapf(err, "写入设备数据失败, deviceId: %s", deviceId)
	}

	g.Log().Debugf(p.ctx, "设备数据写入成功, deviceId: %s, key: %s", deviceId, key)
	return nil
}

func (p *Pebbledb) SaveProtocolMetrics(protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDriverConfig, metrics map[string]any) error {
	if len(metrics) == 0 {
		return nil
	}

	// 创建协议指标数据结构
	protocolData := map[string]any{
		"device_id":        deviceConfig.Id,
		"device_name":      deviceConfig.Name,
		"protocol_id":      protocolConfig.Id,
		"protocol_address": protocolConfig.Address,
		"protocol_type":    string(protocolConfig.Protocol),
		"timestamp":        time.Now().Unix(),
		"metrics":          metrics,
	}

	// 序列化数据
	data, err := json.Marshal(protocolData)
	if err != nil {
		return gerror.Wrapf(err, "序列化协议指标数据失败, deviceId: %s, protocolId: %s", deviceConfig.Id, protocolConfig.Id)
	}

	// 生成键名：protocol_metrics/{deviceId}/{protocolId}/{timestamp}
	key := fmt.Sprintf("protocol_metrics/%s/%s/%d", deviceConfig.Id, protocolConfig.Id, time.Now().UnixNano())

	// 写入数据
	err = p.db.Set([]byte(key), data, pebble.Sync)
	if err != nil {
		return gerror.Wrapf(err, "写入协议指标数据失败, deviceId: %s, protocolId: %s", deviceConfig.Id, protocolConfig.Id)
	}

	g.Log().Debugf(p.ctx, "协议指标数据写入成功, deviceId: %s, protocolId: %s", deviceConfig.Id, protocolConfig.Id)
	return nil
}

func (p *Pebbledb) SaveSystemMetrics(measurement string, tags map[string]string, metrics map[string]any) error {
	if len(metrics) == 0 {
		return nil
	}

	// 创建系统指标数据结构
	systemData := map[string]any{
		"measurement": measurement,
		"tags":        tags,
		"timestamp":   time.Now().Unix(),
		"metrics":     metrics,
	}

	// 序列化数据
	data, err := json.Marshal(systemData)
	if err != nil {
		return gerror.Wrapf(err, "序列化系统指标数据失败, measurement: %s", measurement)
	}

	// 生成键名：system_metrics/{measurement}/{timestamp}
	key := fmt.Sprintf("system_metrics/%s/%d", measurement, time.Now().UnixNano())

	// 写入数据
	err = p.db.Set([]byte(key), data, pebble.Sync)
	if err != nil {
		return gerror.Wrapf(err, "写入系统指标数据失败, measurement: %s", measurement)
	}

	g.Log().Debugf(p.ctx, "系统指标数据写入成功, measurement: %s", measurement)
	return nil
}

func (p *Pebbledb) Close() {
	if p.db != nil {
		err := p.db.Close()
		if err != nil {
			g.Log().Errorf(p.ctx, "关闭PebbleDB失败: %v", err)
		} else {
			g.Log().Infof(p.ctx, "PebbleDB已关闭")
		}
	}
}
