package c_base

// SReflectPoint 反射遥测点位，用于通过反射调用Get方法获取数据
type SReflectPoint struct {
	*SPoint           // 嵌套基础点位信息
	MethodName string `json:"methodName" v:"required" dc:"反射调用的方法名，如 GetSoc"` // 反射调用的方法名，如 "GetSoc"
}

// GetMethodName 获取反射方法名
func (s *SReflectPoint) GetMethodName() string {
	return s.MethodName
}

// AsProtocolPoint 转换为协议点位，SReflectPoint 不是协议点位，返回 nil
func (s *SReflectPoint) AsProtocolPoint() *SProtocolPoint {
	return nil
}

// 注意：不需要重复实现IPoint接口方法
// 通过结构体嵌套自动继承SPoint的方法实现
// SPoint字段将在启动时验证是否设置
