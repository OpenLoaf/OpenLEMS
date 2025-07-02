package internal

import (
	"common/c_base"
	"context"
	"encoding/json"
	"fmt"
	"sqlite"
	"sqlite/service"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

type PebbledbDatabase struct {
	SystemDb   *pebble.DB
	DeviceDb   *pebble.DB
	ProtocolDb *pebble.DB
}

type Pebbledb struct {
	db                PebbledbDatabase
	ctx               context.Context
	dataRetentionDays int // 数据保留天数，默认30天
}

var configManage service.IConfigManage

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

	// 初始化配置管理
	configManage = sqlite.NewConfigManage(ctx, 1)
	dataRetentionDays := 30 // 默认30天

	// 获取数据保留天数
	dataRetentionDaysValue := configManage.GetSettingValueByName(ctx, "data_retention_days")
	if dataRetentionDaysValue != "" {
		if days, err := strconv.Atoi(dataRetentionDaysValue); err == nil && days > 0 {
			dataRetentionDays = days
		} else {
			g.Log().Warningf(ctx, "数据保留天数配置无效: %s，使用默认值30天", dataRetentionDaysValue)
		}
	}

	// 创建Pebbledb实例
	d := &Pebbledb{
		db:                db,
		ctx:               ctx,
		dataRetentionDays: dataRetentionDays, // 数据保留天数
	}

	// 每天凌晨0点执行一次清理,判断数据超过了保留日期，如果超过了，则删除
	d.startDataCleanupScheduler()

	g.Log().Infof(ctx, "PebbleDB 初始化成功")
	return d
}

// startDataCleanupScheduler 启动数据清理定时任务
func (p *Pebbledb) startDataCleanupScheduler() {
	// 每天凌晨0点执行清理任务
	_, err := gcron.Add(p.ctx, "0 0 * * *", func(ctx context.Context) {
		g.Log().Info(ctx, "开始执行PebbleDB数据清理任务")

		if err := p.cleanupExpiredData(); err != nil {
			g.Log().Errorf(ctx, "PebbleDB数据清理失败: %v", err)
		} else {
			g.Log().Info(ctx, "PebbleDB数据清理完成")
		}
	})

	if err != nil {
		g.Log().Errorf(p.ctx, "启动数据清理定时任务失败: %v", err)
	} else {
		g.Log().Infof(p.ctx, "数据清理定时任务已启动，每天凌晨0点执行，数据保留%d天", p.dataRetentionDays)
	}
}

// cleanupExpiredData 清理过期数据
func (p *Pebbledb) cleanupExpiredData() error {
	// 计算过期时间戳（毫秒）
	expiredTimestamp := time.Now().AddDate(0, 0, -p.dataRetentionDays).UnixMilli()

	// 清理设备数据
	if err := p.cleanupDatabase(p.db.DeviceDb, "device", expiredTimestamp); err != nil {
		return gerror.Wrapf(err, "清理设备数据失败")
	}

	// 清理协议数据
	if err := p.cleanupDatabase(p.db.ProtocolDb, "protocol", expiredTimestamp); err != nil {
		return gerror.Wrapf(err, "清理协议数据失败")
	}

	// 清理系统数据
	if err := p.cleanupDatabase(p.db.SystemDb, "system", expiredTimestamp); err != nil {
		return gerror.Wrapf(err, "清理系统数据失败")
	}

	return nil
}

// cleanupDatabase 清理指定数据库中的过期数据
func (p *Pebbledb) cleanupDatabase(db *pebble.DB, prefix string, expiredTimestamp int64) error {
	iter, err := db.NewIter(nil)
	if err != nil {
		return gerror.Wrapf(err, "创建迭代器失败")
	}
	defer iter.Close()

	var keysToDelete [][]byte
	deletedCount := 0

	// 遍历所有键
	for iter.First(); iter.Valid(); iter.Next() {
		key := string(iter.Key())

		// 检查键是否符合预期格式：prefix/id/timestamp
		if !strings.HasPrefix(key, prefix+"/") {
			continue
		}

		// 从键名中提取时间戳
		parts := strings.Split(key, "/")
		if len(parts) < 3 {
			continue
		}

		timestampStr := parts[len(parts)-1]
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			g.Log().Warningf(p.ctx, "解析时间戳失败，跳过键: %s, 错误: %v", key, err)
			continue
		}

		// 如果数据过期，标记为删除
		if timestamp < expiredTimestamp {
			keysToDelete = append(keysToDelete, iter.Key())
			deletedCount++
		}
	}

	// 批量删除过期数据
	if len(keysToDelete) > 0 {
		batch := db.NewBatch()
		for _, key := range keysToDelete {
			if err := batch.Delete(key, nil); err != nil {
				batch.Close()
				return gerror.Wrapf(err, "添加删除操作到批处理失败")
			}
		}

		if err := batch.Commit(pebble.Sync); err != nil {
			batch.Close()
			return gerror.Wrapf(err, "提交批量删除操作失败")
		}
		batch.Close()

		g.Log().Infof(p.ctx, "成功删除%s数据库中的%d条过期记录", prefix, deletedCount)
	} else {
		g.Log().Infof(p.ctx, "%s数据库中没有需要清理的过期数据", prefix)
	}

	return nil
}

// SetDataRetentionDays 设置数据保留天数
func (p *Pebbledb) SetDataRetentionDays(days int) error {
	if days <= 0 {
		err := gerror.Newf("数据保留天数必须大于0，当前值: %d", days)
		g.Log().Warningf(p.ctx, "%v", err)
		return err
	}

	// 先保存到数据库
	err := configManage.SetSettingValueByName(p.ctx, "data_retention_days", strconv.Itoa(days))
	if err != nil {
		g.Log().Errorf(p.ctx, "设置数据保留天数失败: %v", err)
		return gerror.Wrapf(err, "设置数据保留天数到数据库失败")
	}

	// 只有数据库更新成功后才更新内存中的值
	p.dataRetentionDays = days
	g.Log().Infof(p.ctx, "数据保留天数已设置为: %d天", days)
	return nil
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
