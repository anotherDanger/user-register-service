package domain

import "github.com/google/uuid"

type Domain struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
