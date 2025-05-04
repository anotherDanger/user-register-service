package web

import "github.com/google/uuid"

type Request struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
