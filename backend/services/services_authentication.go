package services

import (
	"fmt"
	"net/http"

	ce "github.com/ivan-ca97/rush/backend/custom_errors"
	"github.com/ivan-ca97/rush/backend/dto"
	"github.com/ivan-ca97/rush/backend/middlewares"
)

func (rs *RushServices) Login(body dto.LoginBody, authentication middlewares.Authentication) (*dto.LoginResponse, error) {
	user, err := rs.Repositories.GetUserByUsername(body.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, &ce.AuthenticationError{Code: http.StatusUnauthorized, Message: "Username or password incorrect"}
	}

	token, err := rs.Repositories.Login(*user, body.Password, authentication)
	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		Token: token,
	}

	return response, nil
}

func (rs *RushServices) Register(body dto.RegisterBody, authentication middlewares.Authentication) (*dto.RegisterResponse, error) {
	user, err := rs.Repositories.GetUserByUsername(body.Username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		errorMessage := fmt.Sprintf("User %s already exists", body.Username)
		return nil, &ce.AuthenticationError{Code: http.StatusConflict, Message: errorMessage}
	}

	user, err = rs.Repositories.CreateUser(body.Username, body.Password, body.EMail, authentication)
	if err != nil {
		return nil, err
	}

	token, err := rs.Repositories.Login(*user, body.Password, authentication)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		Token: token,
	}
	return response, nil
}
