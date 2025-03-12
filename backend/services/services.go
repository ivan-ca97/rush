package services

import (
	"github.com/ivan-ca97/rush/backend/dto"
	"github.com/ivan-ca97/rush/backend/middlewares"
	"github.com/ivan-ca97/rush/backend/repositories"
)

type RushServicesInterface interface {
	Login(body dto.LoginBody, authentication middlewares.Authentication) (*dto.LoginResponse, error)
	Register(body dto.RegisterBody, authentication middlewares.Authentication) (*dto.RegisterResponse, error)
}

type RushServices struct {
	Repositories repositories.RushRepositoriesInterface
}

var _ RushServicesInterface = &RushServices{}
