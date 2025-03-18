package dto

import (
	"github.com/google/uuid"
	"github.com/ivan-ca97/rush/backend/models"
)

type PersonCreateBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PersonCreateResponse struct {
	models.Person
}

type PersonUpdateBody struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type PersonUpdateResponse struct {
	models.Person
}

type PersonGetByIdBody struct {
	Id uuid.UUID `json:"id"`
}

type PersonGetResponse struct {
	models.Person
}
