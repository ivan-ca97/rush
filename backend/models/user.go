package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Password string    `json:"password"`
	Username string    `json:"username"`
	EMail    string    `json:"email"`
}
