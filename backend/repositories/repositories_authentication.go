package repositories

import (
	"net/http"

	"github.com/google/uuid"

	ce "github.com/ivan-ca97/rush/backend/custom_errors"
	"github.com/ivan-ca97/rush/backend/middlewares"
	"github.com/ivan-ca97/rush/backend/models"
)

func (rr *RushRepositories) Login(user models.User, password string, authentication middlewares.Authentication) (*string, error) {
	validated := authentication.ValidatePassword(password, user.Password)
	if !validated {
		return nil, &ce.AuthenticationError{Code: http.StatusUnauthorized, Message: "Username or password incorrect"}
	}

	token, err := authentication.GenerateToken(user.Username, user.Id.String())
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// Creates a new user in the database
func (rr *RushRepositories) CreateUser(username string, password string, eMail string, authentication middlewares.Authentication) (*models.User, error) {
	passwordHash, err := authentication.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Id:       uuid.New(),
		Username: username,
		Password: passwordHash,
		EMail:    eMail,
	}

	err = rr.Db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
