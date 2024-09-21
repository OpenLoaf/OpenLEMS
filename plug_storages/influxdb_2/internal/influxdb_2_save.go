package internal

import (
	"common/c_base"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"strings"
	"time"
)

type Influxdb2 struct {
	client          influxdb2.Client
	ctx             context.Context
	bucketsAPI      api.BucketsAPI
	organization    *domain.Organization
	storageConfig   *c_base.SStorageConfig
	influxdb2Config *sInfluxdb2Config
}

const (
	devicesBucket = "ems/device_metrics"
	//protocolBucket = "ems/protocol_metrics"
	systemBucket = "ems/system_metrics"
)

func NewInfluxdb2(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	if !storageConfig.Enable {
		return nil
	}
	var influxdb2Config = &sInfluxdb2Config{}
	err := gconv.Struct(storageConfig.Params, influxdb2Config)
	if err != nil {
		panic(gerror.Newf("创建Influxdb2实例失败！无法解析params %v", err))
	}

	clt := influxdb2.NewClient(storageConfig.Url, influxdb2Config.Token)
	if ctx == nil {
		ctx = context.Background()
	}

	d := &Influxdb2{
		client:          clt,
		ctx:             ctx,
		bucketsAPI:      clt.BucketsAPI(),
		storageConfig:   storageConfig,
		influxdb2Config: influxdb2Config,
	}

	organization := d.getOrganization()
	d.organization = organization

	// 检查并创建数据库
	err = d.createOrUpdateBucket(devicesBucket, int32(0))
	if err != nil {
		panic(gerror.Newf("创建Bucket:%s 失败！%v", devicesBucket, err))
	}

	//err = d.createOrUpdateBucket(protocolBucket, storageConfig.ProtocolMetricsSurvivalDays)
	//if err != nil {
	//	panic(gerror.Newf("创建Bucket:%s 失败！%v", protocolBucket, err))
	//}
	err = d.createOrUpdateBucket(systemBucket, storageConfig.SystemMetricsSurvivalDays)
	if err != nil {
		panic(gerror.Newf("创建Bucket:%s 失败！%v", systemBucket, err))
	}
	return d
}

func (i *Influxdb2) SaveDevices(deviceId string, deviceType c_base.EDeviceType, fields map[string]any) error {
	return i.write(devicesBucket, string(deviceType), map[string]string{"device_id": deviceId}, fields)
}

func (i *Influxdb2) Close() {
	i.client.Close()
}

func (i *Influxdb2) SaveProtocolMetrics(protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDriverConfig, metrics map[string]any) error {
	tags := map[string]string{
		c_base.ConstDeviceId:        deviceConfig.Id,
		c_base.ConstDeviceName:      deviceConfig.Name,
		c_base.ConstProtocolId:      protocolConfig.Id,
		c_base.ConstProtocolAddress: protocolConfig.Address,
		c_base.ConstProtocolType:    string(protocolConfig.Protocol),
	}
	return i.write(systemBucket, c_base.ConstProtocol, tags, metrics)
}

func (i *Influxdb2) write(bucket, measurement string, tags map[string]string, fields map[string]interface{}, t ...time.Time) error {
	if len(fields) == 0 {
		return nil
	}

	writeAPI := i.client.WriteAPI(i.influxdb2Config.Org, bucket)
	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)

	// Flush writes
	writeAPI.Flush()
	//g.Log().Infof(i.ctx, "设备 %s 写入InfluxDB成功", deviceId)
	return nil
}

func (i *Influxdb2) bucketExists(bucketName string) (bool, error) {
	bucket, err := i.bucketsAPI.FindBucketByName(i.ctx, bucketName)
	if err != nil {
		if strings.HasSuffix(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}
	return bucket != nil, nil
}

func (i *Influxdb2) createOrUpdateBucket(bucketName string, days int32) error {
	if days < 0 {
		g.Log().Noticef(i.ctx, "Bucket \"%s\" 保留期为 %d天，跳过！\n", bucketName, days)
		return nil
	}

	exists, err := i.bucketExists(bucketName)
	if err != nil {
		return err
	}
	if exists {
		g.Log().Debugf(i.ctx, "Bucket \"%s\" 已存在。\n", bucketName)
		// 更新保留策略
		return i.updateBucketRetentionPolicy(bucketName, days)
	}
	return i.createBucket(bucketName, days)
}

func (i *Influxdb2) createBucket(bucketName string, days int32) error {

	tp := domain.RetentionRuleTypeExpire
	retentionRule := domain.RetentionRule{
		Type:         &tp,
		EverySeconds: 60 * 60 * 24 * int64(days),
	}
	_, err := i.bucketsAPI.CreateBucketWithName(i.ctx, i.organization, bucketName, retentionRule)
	if err != nil {
		return err
	}
	g.Log().Infof(i.ctx, "Bucket \"%s\" 创建成功，保留期为 %d天。\n", bucketName, days)
	return nil
}

func (i *Influxdb2) updateBucketRetentionPolicy(bucketName string, days int32) error {
	bucket, err := i.bucketsAPI.FindBucketByName(i.ctx, bucketName)
	if err != nil {
		return err
	}
	if bucket == nil {
		return fmt.Errorf("bucket \"%s\" 不存在!", bucketName)
	}
	if bucket.RetentionRules == nil || len(bucket.RetentionRules) == 0 {
		return fmt.Errorf("bucket \"%s\" 没有保留策略!", bucketName)
	}
	// 更新保留策略
	bucket.RetentionRules[0].EverySeconds = int64(days) * 60 * 60 * 24
	_, err = i.bucketsAPI.UpdateBucket(i.ctx, bucket)
	if err != nil {
		return err
	}
	g.Log().Infof(i.ctx, "Bucket \"%s\" 的保留策略已更新。\n", bucketName)
	return nil
}

func (i *Influxdb2) getOrganization() *domain.Organization {
	organizationsAPI := i.client.OrganizationsAPI()
	organization, err := organizationsAPI.FindOrganizationByName(i.ctx, i.influxdb2Config.Org)
	if err != nil {
		panic(gerror.Newf("Error getting organization info: %v", err))
	}
	return organization
}
