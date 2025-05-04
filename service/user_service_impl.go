package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"user_register_service/domain"
	"user_register_service/helper"
	"user_register_service/repository"
	"user_register_service/web"
)

type UserServiceImpl struct {
	repo *repository.UserRepositoryImpl
	db   *sql.DB
}

func NewUserServiceImpl(repo *repository.UserRepositoryImpl, db *sql.DB) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
		db:   db,
	}
}

func (svc *UserServiceImpl) Register(ctx context.Context, request *web.Request) (*web.Token, error) {

	tx, err := svc.db.Begin()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	result, err := svc.repo.Register(ctx, tx, (*domain.Domain)(request))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	helper.NewTx(tx, &err)

	newReader := strings.NewReader(result.Username)

	response, err := http.Post("http://auth_service:8080/v1/auth", "application/json", newReader)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var token web.Token
	json.Unmarshal(resBody, &token)

	return &token, nil

}
