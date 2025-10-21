package remote

import (
	"common/c_log"
	"context"
	s_export_mqtt "s_mqtt"

	v1 "application/api/remote/v1"
)

// GetMqttStatus 获取MQTT服务状态列表
func (c *ControllerV1) GetMqttStatus(ctx context.Context, req *v1.GetMqttStatusReq) (res *v1.GetMqttStatusRes, err error) {
	// 调用MQTT服务获取状态
	isRunning, clientCount, clientStatusList := s_export_mqtt.GetMqttStatus()

	// 转换客户端状态列表
	var apiClientStatusList []*v1.MqttClientStatus
	for _, clientStatus := range clientStatusList {
		// 转换配置信息
		var apiConfig *v1.MqttConfig
		if clientStatus.Config != nil {
			// 处理密码字段：如果有密码则显示为********，否则为空
			var password string
			if clientStatus.Config.Password != "" {
				password = "********"
			}

			apiConfig = &v1.MqttConfig{
				ServerAddress:      clientStatus.Config.ServerAddress,
				ServerPort:         clientStatus.Config.ServerPort,
				Username:           clientStatus.Config.Username,
				Password:           password,
				UseSSL:             clientStatus.Config.UseSSL,
				InsecureSkipVerify: clientStatus.Config.InsecureSkipVerify,
				ConnectTimeout:     clientStatus.Config.ConnectTimeout,
				ReconnectInterval:  clientStatus.Config.ReconnectInterval,
				KeepAliveTimeout:   clientStatus.Config.KeepAliveTimeout,
				ServiceStandard:    clientStatus.Config.ServiceStandard,
				AllowControl:       clientStatus.Config.AllowControl,
				Enabled:            clientStatus.Config.Enabled,
				DeviceIds:          clientStatus.Config.DeviceIds,
				RewriteChannel:     clientStatus.Config.RewriteChannel,
				PushChannel:        clientStatus.Config.PushChannel,
				SubscribeChannel:   clientStatus.Config.SubscribeChannel,
				UploadPeriod:       clientStatus.Config.UploadPeriod,
			}
		}

		// 转换客户端状态
		apiClientStatus := &v1.MqttClientStatus{
			Config:        apiConfig,
			IsConnected:   clientStatus.IsConnected,
			IsRunning:     clientStatus.IsRunning,
			ClientID:      clientStatus.ClientID,
			SystemNumber:  clientStatus.SystemNumber,
			Topic:         clientStatus.Topic,
			DeviceCount:   clientStatus.DeviceCount,
			UploadPeriod:  clientStatus.UploadPeriod,
			LastPublishAt: clientStatus.LastPublishAt,
			PublishCount:  clientStatus.PublishCount,
			ErrorCount:    clientStatus.ErrorCount,
			StartTime:     clientStatus.StartTime,
		}

		apiClientStatusList = append(apiClientStatusList, apiClientStatus)
	}

	c_log.Infof(ctx, "成功获取MQTT服务状态: 运行状态=%t, 客户端数量=%d", isRunning, clientCount)

	return &v1.GetMqttStatusRes{
		IsRunning:    isRunning,
		ClientCount:  clientCount,
		ClientStatus: apiClientStatusList,
	}, nil
}
