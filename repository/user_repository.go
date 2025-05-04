package repository

import (
	"context"
	"database/sql"
	"user_register_service/domain"
)

type UserRepository interface {
	Register(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error)
}
