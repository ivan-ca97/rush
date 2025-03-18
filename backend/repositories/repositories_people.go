package repositories

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	ce "github.com/ivan-ca97/rush/backend/custom_errors"
	"github.com/ivan-ca97/rush/backend/models"
	"gorm.io/gorm"
)

func (rr *RushRepositories) GetPersonById(id uuid.UUID) (*models.Person, error) {
	var person models.Person

	idString := id.String()

	err := rr.Db.Model(&person).Where("id = ?", idString).First(&person).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &ce.PersonError{Code: http.StatusNotFound, Message: "Person ID not found"}
	}
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (rr *RushRepositories) CreatePerson(person models.Person) (*models.Person, error) {
	person.Id = uuid.New()

	err := rr.Db.Create(&person).Error
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (rr *RushRepositories) UpdatePerson(person models.Person) (*models.Person, error) {
	err := rr.Db.Model(&person).Where("id = ?", person.Id).Updates(&person).Error
	if err != nil {
		return nil, err
	}

	return &person, nil
}
