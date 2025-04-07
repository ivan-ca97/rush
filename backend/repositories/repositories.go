package repositories

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/ivan-ca97/rush/backend/middlewares"
	"github.com/ivan-ca97/rush/backend/models"
)

type RushRepositoriesInterface interface {
	Login(user models.User, password string, authentication middlewares.Authentication) (*string, error)
	CreateUser(username string, passwordHash string, eMail string, authentication middlewares.Authentication) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)

	GetPersonById(id uuid.UUID) (*models.Person, error)
	CreatePerson(person models.Person) (*models.Person, error)
	UpdatePerson(person models.Person) (*models.Person, error)

	GetAllGroups(limit int, offset int) ([]models.Group, error)
	GetGroupById(id uuid.UUID) (*models.Group, error)
	CreateGroup(group models.Group) (*models.Group, error)
	UpdateGroup(group models.Group) (*models.Group, error)
}

type RushRepositories struct {
	Db *gorm.DB
}

var _ RushRepositoriesInterface = &RushRepositories{}
