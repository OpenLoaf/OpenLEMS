package remote

import (
	"common/c_log"
	"context"
	s_export_mqtt "s_mqtt"

	v1 "application/api/remote/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// ReloadMqtt 重新加载MQTT服务配置
func (c *ControllerV1) ReloadMqtt(ctx context.Context, req *v1.ReloadMqttReq) (res *v1.ReloadMqttRes, err error) {
	// 调用MQTT服务重新加载配置
	err = s_export_mqtt.ReloadMqtt(ctx)
	if err != nil {
		c_log.Errorf(ctx, "重新加载MQTT服务配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "重新加载MQTT服务配置失败")
	}

	c_log.Infof(ctx, "成功重新加载MQTT服务配置")
	return &v1.ReloadMqttRes{}, nil
}
