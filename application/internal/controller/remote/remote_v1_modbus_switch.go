package remote

import (
	"context"
	"encoding/json"
	"s_db"
	"s_db/s_db_basic"
	"s_export_modbus"

	v1 "application/api/remote/v1"

	"common/c_log"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// ModbusSwitch Modbus服务开关
func (c *ControllerV1) ModbusSwitch(ctx context.Context, req *v1.ModbusSwitchReq) (res *v1.ModbusSwitchRes, err error) {
	c_log.Infof(ctx, "开始设置Modbus服务开关 - 启用状态: %v", req.Enabled)

	// 获取Modbus配置设置
	configJson := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingModbusConfig)
	if configJson == "" {
		c_log.Warning(ctx, "Modbus配置为空，使用默认配置")
		configJson = "{}"
	}

	// 解析JSON配置
	var config map[string]interface{}
	err = json.Unmarshal([]byte(configJson), &config)
	if err != nil {
		c_log.Errorf(ctx, "解析Modbus配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "解析Modbus配置失败")
	}

	// 更新配置的enabled状态
	config["enabled"] = req.Enabled
	c_log.Debugf(ctx, "更新Modbus配置的启用状态为: %v", req.Enabled)

	// 将更新后的配置转换回JSON
	updatedConfigJson, err := json.Marshal(config)
	if err != nil {
		c_log.Errorf(ctx, "序列化Modbus配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "序列化Modbus配置失败")
	}

	// 保存更新后的配置到数据库
	err = s_db.GetSettingService().SetSettingValueById(ctx, s_db_basic.SystemSettingModbusConfig.Id, string(updatedConfigJson))
	if err != nil {
		c_log.Errorf(ctx, "保存Modbus配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "保存Modbus配置失败")
	}

	c_log.Infof(ctx, "成功更新Modbus配置 - 启用状态: %v", req.Enabled)

	// 重新加载Modbus服务
	c_log.Info(ctx, "开始重新加载Modbus服务")
	err = s_export_modbus.ReloadModbus(ctx)
	if err != nil {
		c_log.Errorf(ctx, "重新加载Modbus服务失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "重新加载Modbus服务失败")
	}

	c_log.Infof(ctx, "Modbus服务开关设置成功 - 启用状态: %v", req.Enabled)
	return &v1.ModbusSwitchRes{}, nil
}
