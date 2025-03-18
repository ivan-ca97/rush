package services

import (
	"github.com/google/uuid"
	"github.com/ivan-ca97/rush/backend/dto"
	"github.com/ivan-ca97/rush/backend/models"
)

func (rs *RushServices) GetPersonById(id uuid.UUID) (*dto.PersonGetResponse, error) {
	person, err := rs.Repositories.GetPersonById(id)
	if err != nil {
		return nil, err
	}

	response := &dto.PersonGetResponse{
		Person: *person,
	}

	return response, nil
}

func (rs *RushServices) CreatePerson(body dto.PersonCreateBody) (*dto.PersonCreateResponse, error) {
	person := &models.Person{
		Id:        uuid.New(),
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	person, err := rs.Repositories.CreatePerson(*person)
	if err != nil {
		return nil, err
	}

	response := &dto.PersonCreateResponse{
		Person: *person,
	}

	return response, nil
}

func (rs *RushServices) UpdatePerson(id uuid.UUID, body dto.PersonUpdateBody) (*dto.PersonUpdateResponse, error) {
	person, err := rs.Repositories.GetPersonById(id)
	if err != nil {
		return nil, err
	}

	if body.FirstName != nil {
		person.FirstName = *body.FirstName
	}
	if body.LastName != nil {
		person.LastName = *body.LastName
	}

	person, err = rs.Repositories.UpdatePerson(*person)
	if err != nil {
		return nil, err
	}

	response := &dto.PersonUpdateResponse{
		Person: *person,
	}

	return response, nil
}
