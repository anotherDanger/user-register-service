package service

import (
	"context"
	"user_register_service/web"
)

type UserService interface {
	Register(ctx context.Context, request *web.Request) (*web.Token, error)
}
