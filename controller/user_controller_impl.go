package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"user_register_service/service"
	"user_register_service/web"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	svc *service.UserServiceImpl
}

func NewUserControllerImpl(svc *service.UserServiceImpl) *UserControllerImpl {
	return &UserControllerImpl{
		svc: svc,
	}
}

func (ctrl *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	reqBody := web.Request{}
	json.NewDecoder(r.Body).Decode(&reqBody)

	token, err := ctrl.svc.Register(r.Context(), &reqBody)
	if err != nil {
		w.WriteHeader(400)
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(token)
}
