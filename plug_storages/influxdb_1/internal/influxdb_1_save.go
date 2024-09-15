package internal

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/influxdata/influxdb/client/v2"
)

type Influxdb1 struct {
	ct            client.Client
	ctx           context.Context
	storageConfig *c_base.SStorageConfig
}

func NewInfluxdb1(ctx context.Context, storageConfig *c_base.SStorageConfig) c_base.IStorage {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     storageConfig.Url,
		Username: storageConfig.Username,
		Password: storageConfig.Password,
	})
	if ctx == nil {
		ctx = context.Background()
	}

	if err != nil {
		panic(gerror.Newf("Error creating InfluxDB Client: %v", err))
	}
	return &Influxdb1{
		ct:            c,
		ctx:           ctx,
		storageConfig: storageConfig,
	}
}

func (i *Influxdb1) Save(deviceId string, deviceType c_base.EDeviceType, fields map[string]any) error {
	if len(fields) == 0 {
		return nil
	}
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
		Database:  i.storageConfig.Database,
	})
	if err != nil {
		g.Log().Errorf(i.ctx, "Error creating InfluxDB BatchPoints: %v", err)
		return err
	}

	point, err := client.NewPoint(string(deviceType), map[string]string{"device_id": deviceId}, fields)
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

func (i *Influxdb1) Close() {
	_ = i.ct.Close()
}
