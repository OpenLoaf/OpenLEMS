package auth

import (
	v1 "application/api/auth/v1"
	"context"
)

type IAuthV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error)
	GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserReq) (res *v1.GetCurrentUserRes, err error)
}
