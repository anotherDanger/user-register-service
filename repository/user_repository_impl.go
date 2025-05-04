package repository

import (
	"context"
	"database/sql"
	"log"
	"user_register_service/domain"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Register(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	userId := uuid.New()
	query := "insert into users(id, username, password) values(?, ?, ?)"
	password := []byte(entity.Password)
	hashed, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	_, err = sql.ExecContext(ctx, query, userId, entity.Username, hashed)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	result := domain.Domain{
		Id:       userId,
		Username: entity.Username,
	}

	return &result, nil

}
