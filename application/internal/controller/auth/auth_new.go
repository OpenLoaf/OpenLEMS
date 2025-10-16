package auth

import (
	"application/api/auth"
)

type Controller struct{}

// NewV1 创建认证控制器
func NewV1() auth.IAuthV1 { return &Controller{} }
