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

type PebbledbDatabase struct {
	SystemDb   *pebble.DB
	DeviceDb   *pebble.DB
	ProtocolDb *pebble.DB
}

type Pebbledb struct {
	db  PebbledbDatabase
	ctx context.Context
}

func NewPebbledb(ctx context.Context) c_base.IStorage {

	// 创建PebbleDB实例
	systemDb, err := pebble.Open("./out/pebbledb/system", &pebble.Options{})
	if err != nil {
		panic(gerror.Newf("创建Pebbledb实例失败！%v", err))
	}

	deviceDb, err := pebble.Open("./out/pebbledb/device", &pebble.Options{})
	if err != nil {
		panic(gerror.Newf("创建Pebbledb实例失败！%v", err))
	}

	protocolDb, err := pebble.Open("./out/pebbledb/protocol", &pebble.Options{})
	if err != nil {
		panic(gerror.Newf("创建Pebbledb实例失败！%v", err))
	}
	db := PebbledbDatabase{
		SystemDb:   systemDb,
		DeviceDb:   deviceDb,
		ProtocolDb: protocolDb,
	}

	d := &Pebbledb{
		db:  db,
		ctx: ctx,
	}

	g.Log().Infof(ctx, "PebbleDB 初始化成功")
	return d
}

func (p *Pebbledb) SaveDevices(deviceId string, deviceType c_base.EDeviceType, fields map[string]any) error {
	if len(fields) == 0 {
		return nil
	}

	// 创建设备数据结构
	deviceData := map[string]any{
		"device_id":   deviceId,
		"device_type": deviceType,
		"timestamp":   time.Now().UnixMilli(),
		"metrics":     fields,
	}

	// 序列化数据
	data, err := json.Marshal(deviceData)
	if err != nil {
		return gerror.Wrapf(err, "序列化设备数据失败, deviceId: %s", deviceId)
	}

	// 生成键名：device/{deviceId}/{timestamp}
	key := fmt.Sprintf("device/%s/%d", deviceId, time.Now().UnixMilli())

	// 写入数据
	err = p.db.DeviceDb.Set([]byte(key), data, pebble.Sync)
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
		"timestamp":        time.Now().UnixMilli(),
		"metrics":          metrics,
	}

	// 序列化数据
	data, err := json.Marshal(protocolData)
	if err != nil {
		return gerror.Wrapf(err, "序列化协议指标数据失败, deviceId: %s, protocolId: %s", deviceConfig.Id, protocolConfig.Id)
	}

	// 生成键名：protocol/{protocolId}/{timestamp}
	key := fmt.Sprintf("protocol/%s/%d", protocolConfig.Id, time.Now().UnixMilli())

	// 写入数据
	err = p.db.ProtocolDb.Set([]byte(key), data, pebble.Sync)
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
		"timestamp":   time.Now().UnixMilli(),
		"metrics":     metrics,
	}

	// 序列化数据
	data, err := json.Marshal(systemData)
	if err != nil {
		return gerror.Wrapf(err, "序列化系统指标数据失败, measurement: %s", measurement)
	}

	// 生成键名：system/{measurement}/{timestamp}
	key := fmt.Sprintf("system/%s/%d", measurement, time.Now().UnixMilli())

	// 写入数据
	err = p.db.SystemDb.Set([]byte(key), data, pebble.Sync)
	if err != nil {
		return gerror.Wrapf(err, "写入系统指标数据失败, measurement: %s", measurement)
	}

	g.Log().Debugf(p.ctx, "系统指标数据写入成功, measurement: %s", measurement)
	return nil
}

func (p *Pebbledb) Close() {
	if p.db.SystemDb != nil {
		err := p.db.SystemDb.Close()
		if err != nil {
			g.Log().Errorf(p.ctx, "关闭SystemDb失败: %v", err)
		} else {
			g.Log().Infof(p.ctx, "SystemDb已关闭")
		}
	}
	if p.db.DeviceDb != nil {
		err := p.db.DeviceDb.Close()
		if err != nil {
			g.Log().Errorf(p.ctx, "关闭DeviceDb失败: %v", err)
		} else {
			g.Log().Infof(p.ctx, "DeviceDb已关闭")
		}
	}
	if p.db.ProtocolDb != nil {
		err := p.db.ProtocolDb.Close()
		if err != nil {
			g.Log().Errorf(p.ctx, "关闭ProtocolDb失败: %v", err)
		} else {
			g.Log().Infof(p.ctx, "ProtocolDb已关闭")
		}
	}
}
