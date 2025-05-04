//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"user_register_service/controller"
	"user_register_service/helper"
	"user_register_service/repository"
	"user_register_service/service"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var ServerSet = wire.NewSet(
	helper.NewDb,
	repository.NewUserRepositoryImpl, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewUserServiceImpl, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewUserControllerImpl, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
	NewRouter,
	NewServer, wire.Bind(new(http.Handler), new(*httprouter.Router)),
)

func InitServer() (*http.Server, func(), error) {
	wire.Build(ServerSet)
	return nil, nil, nil
}
