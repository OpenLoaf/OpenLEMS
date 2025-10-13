package remote

import (
	"context"
	"encoding/json"
	"s_db"
	"s_db/s_db_basic"
	s_export_mqtt "s_mqtt"

	v1 "application/api/remote/v1"

	"common/c_log"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// MqttSwitch MQTT服务开关
func (c *ControllerV1) MqttSwitch(ctx context.Context, req *v1.MqttSwitchReq) (res *v1.MqttSwitchRes, err error) {
	c_log.Infof(ctx, "开始设置MQTT服务开关 - 启用状态: %v", req.Enabled)

	// 获取MQTT配置列表设置
	configJson := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingMqttConfigList)
	if configJson == "" {
		c_log.Warning(ctx, "MQTT配置列表为空，使用默认配置")
		configJson = "[]"
	}

	// 解析JSON配置
	var configList []map[string]interface{}
	err = json.Unmarshal([]byte(configJson), &configList)
	if err != nil {
		c_log.Errorf(ctx, "解析MQTT配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "解析MQTT配置失败")
	}

	// 更新所有配置的enabled状态
	for i, config := range configList {
		config["enabled"] = req.Enabled
		configList[i] = config
		c_log.Debugf(ctx, "更新MQTT配置 %d 的启用状态为: %v", i, req.Enabled)
	}

	// 将更新后的配置转换回JSON
	updatedConfigJson, err := json.Marshal(configList)
	if err != nil {
		c_log.Errorf(ctx, "序列化MQTT配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "序列化MQTT配置失败")
	}

	// 保存更新后的配置到数据库
	err = s_db.GetSettingService().SetSettingValueById(ctx, s_db_basic.SystemSettingMqttConfigList.Id, string(updatedConfigJson))
	if err != nil {
		c_log.Errorf(ctx, "保存MQTT配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "保存MQTT配置失败")
	}

	c_log.Infof(ctx, "成功更新MQTT配置 - 启用状态: %v", req.Enabled)

	// 重新加载MQTT服务
	c_log.Info(ctx, "开始重新加载MQTT服务")
	err = s_export_mqtt.ReloadMqtt(ctx)
	if err != nil {
		c_log.Errorf(ctx, "重新加载MQTT服务失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "重新加载MQTT服务失败")
	}

	c_log.Infof(ctx, "MQTT服务开关设置成功 - 启用状态: %v", req.Enabled)
	return &v1.MqttSwitchRes{}, nil
}
