package internal

import (
	"common/c_base"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/influxdata/influxdb/client/v2"
	"time"
)

type Influxdb1 struct {
	ct              client.Client
	ctx             context.Context
	storageConfig   *c_base.SStorageConfig
	influxdb1Config *sInfluxdb1Config
}

func NewInfluxdb1(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	var influxdb1Config = &sInfluxdb1Config{}
	err := gconv.Struct(storageConfig.Params, influxdb1Config)
	if err != nil {
		panic(gerror.Newf("创建Influxdb2实例失败！无法解析params %v", err))
	}
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     storageConfig.Url,
		Username: influxdb1Config.Username,
		Password: influxdb1Config.Password,
	})
	if ctx == nil {
		ctx = context.Background()
	}

	if err != nil {
		panic(gerror.Newf("Error creating InfluxDB Client: %v", err))
	}

	d := &Influxdb1{
		ct:              c,
		ctx:             ctx,
		storageConfig:   storageConfig,
		influxdb1Config: influxdb1Config,
	}

	// 检查并创建数据库
	err = d.createDatabaseIfNotExists()
	if err != nil {
		panic(gerror.Newf("Error creating database: %v", err))
	}

	d.createRetentionIfNotExists("protocol_policy", storageConfig.ProtocolMetricsSurvivalDays)
	d.createRetentionIfNotExists("system_policy", storageConfig.SystemMetricsSurvivalDays)
	return d
}

func (i *Influxdb1) Save(deviceId string, deviceType c_base.EDeviceType, fields map[string]any) error {
	return i.write(string(deviceType), "", map[string]string{"device_id": deviceId}, fields)
}

func (i *Influxdb1) Close() {
	_ = i.ct.Close()
}

func (i *Influxdb1) SaveProtocolMetrics(protocolConfig *c_base.SProtocolConfig, metrics map[string]any) error {
	tags := map[string]string{
		"protocol_id":      protocolConfig.Id,
		"protocol_address": protocolConfig.Address,
		"protocol_type":    string(protocolConfig.Protocol),
	}
	return i.write("protocol_metrics", "protocol_policy", tags, metrics)
}

func (i *Influxdb1) write(name, retentionPolicy string, tags map[string]string, fields map[string]interface{}, t ...time.Time) error {
	if len(fields) == 0 {
		return nil
	}
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision:       "s",
		Database:        i.influxdb1Config.Database,
		RetentionPolicy: retentionPolicy,
	})
	if err != nil {
		g.Log().Errorf(i.ctx, "Error creating InfluxDB BatchPoints: %v", err)
		return err
	}

	point, err := client.NewPoint(name, tags, fields, t...)
	if err != nil {
		g.Log().Errorf(i.ctx, "Error creating InfluxDB Point: %v", err)
		return err
	}
	bp.AddPoint(point)

	err = i.ct.WriteCtx(i.ctx, bp)
	if err != nil {
		g.Log().Errorf(i.ctx, "Error writing to InfluxDB: %v", err)
		return err
	}
	//g.Log().Infof(i.ctx, "设备 %s 写入InfluxDB成功", deviceId)
	return nil
}

func (i *Influxdb1) createRetentionIfNotExists(rpName string, days int32) {
	duration := fmt.Sprintf("%dd", days)
	shardDuration := "1d"
	replication := 1

	// 检查保留策略是否存在
	exists, err := i.retentionPolicyExists(rpName)
	if err != nil {
		panic(gerror.Newf("Error checking retention policy: %v", err))
	}

	if exists {
		g.Log().Infof(i.ctx, "保留策略 \"%s\" 已存在,跳过策略初始化！", rpName)
	} else {
		// 创建保留策略
		err = i.createRetentionPolicy(rpName, duration, shardDuration, replication, false)
		if err != nil {
			panic(gerror.Newf("Error creating retention policy: %v", err))
		}
	}
}

func (i *Influxdb1) retentionPolicyExists(rpName string) (bool, error) {
	q := fmt.Sprintf("SHOW RETENTION POLICIES ON \"%s\"", i.influxdb1Config.Database)
	query := client.Query{
		Command:  q,
		Database: i.influxdb1Config.Database,
	}

	response, err := i.ct.Query(query)
	if err != nil {
		return false, err
	}
	if response.Error() != nil {
		return false, response.Error()
	}

	for _, result := range response.Results {
		for _, series := range result.Series {
			for _, row := range series.Values {
				if len(row) > 0 {
					name, ok := row[0].(string)
					if ok && name == rpName {
						return true, nil
					}
				}
			}
		}
	}

	return false, nil
}

func (i *Influxdb1) createRetentionPolicy(rpName, duration, shardDuration string, replication int, isDefault bool) error {
	q := fmt.Sprintf("CREATE RETENTION POLICY \"%s\" ON \"%s\" DURATION %s REPLICATION %d SHARD DURATION %s",
		rpName, i.influxdb1Config.Database, duration, replication, shardDuration)

	if isDefault {
		q += " DEFAULT"
	}

	query := client.Query{
		Command:  q,
		Database: i.influxdb1Config.Database,
	}

	response, err := i.ct.Query(query)
	if err != nil {
		return err
	}
	if response.Error() != nil {
		return response.Error()
	}

	g.Log().Info(i.ctx, "保留策略创建成功")
	return nil
}

func (i *Influxdb1) createDatabaseIfNotExists() error {
	q := fmt.Sprintf("CREATE DATABASE \"%s\"", i.influxdb1Config.Database)
	query := client.Query{
		Command: q,
	}

	response, err := i.ct.Query(query)
	if err != nil {
		return err
	}
	if response.Error() != nil {
		return response.Error()
	}

	g.Log().Info(i.ctx, "数据库 \"%s\" 已存在或创建成功\n", i.influxdb1Config.Database)
	return nil
}
