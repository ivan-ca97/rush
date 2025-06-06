package services

import (
	"github.com/google/uuid"
	"github.com/ivan-ca97/rush/backend/dto"
	"github.com/ivan-ca97/rush/backend/middlewares"
	"github.com/ivan-ca97/rush/backend/repositories"
)

type RushServicesInterface interface {
	Login(body dto.LoginBody, authentication middlewares.Authentication) (*dto.LoginResponse, error)
	Register(body dto.RegisterBody, authentication middlewares.Authentication) (*dto.RegisterResponse, error)

	GetPersonById(id uuid.UUID) (*dto.PersonGetResponse, error)
	CreatePerson(body dto.PersonCreateBody) (*dto.PersonCreateResponse, error)
	UpdatePerson(id uuid.UUID, body dto.PersonUpdateBody) (*dto.PersonUpdateResponse, error)

	GetAllGroups(pageNumber int, pageSize int) (*dto.GetAllGroupsResponse, error)
	GetGroupById(id uuid.UUID) (*dto.GetGroupByIdResponse, error)
	CreateGroup(body dto.CreateGroupBody) (*dto.CreateGroupResponse, error)
	UpdateGroup(id uuid.UUID, body dto.UpdateGroupBody) (*dto.UpdateGroupResponse, error)
}

type RushServices struct {
	Repositories repositories.RushRepositoriesInterface
}

var _ RushServicesInterface = &RushServices{}
