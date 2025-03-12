package repositories

import (
	"gorm.io/gorm"

	"github.com/ivan-ca97/rush/backend/middlewares"
	"github.com/ivan-ca97/rush/backend/models"
)

type RushRepositoriesInterface interface {
	Login(user models.User, password string, authentication middlewares.Authentication) (*string, error)
	CreateUser(username string, passwordHash string, eMail string, authentication middlewares.Authentication) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type RushRepositories struct {
	Db *gorm.DB
}

var _ RushRepositoriesInterface = &RushRepositories{}
