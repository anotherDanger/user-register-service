package main

import (
	"net/http"
	"user_register_service/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctrl *controller.UserControllerImpl) *httprouter.Router {
	r := httprouter.New()
	r.POST("/v1/register", ctrl.Register)
	return r
}

func NewServer(handler http.Handler) *http.Server {

	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
}

func main() {
	server, cleanup, err := InitServer()
	if err != nil {
		panic(err)
	}

	defer cleanup()

	server.ListenAndServe()
}
