package c_enum

// EUserRole 用户角色枚举
type EUserRole string

const (
	EUserRoleGuest EUserRole = "guest" // 公共访问
	EUserRoleUser  EUserRole = "user"  // 普通用户（只读）
	EUserRoleAdmin EUserRole = "admin" // 管理员（读写）
)
