package controllers

import (
	"net/http"

	"github.com/ivan-ca97/rush/backend/services"
)

type ControllersBasic interface {
	GetPaginationParameters(request *http.Request) (pageNumber int, pageSize int, err error)
	EncodeResponse(w http.ResponseWriter, response any) error
	HandleError(writer http.ResponseWriter, err error)
}

type RushControllersInterface interface {
	ControllersBasic

	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)

	GetPersonById(w http.ResponseWriter, r *http.Request)
	CreatePerson(w http.ResponseWriter, r *http.Request)
	UpdatePerson(w http.ResponseWriter, r *http.Request)

	GetAllGroups(w http.ResponseWriter, r *http.Request)
	GetGroupById(w http.ResponseWriter, r *http.Request)
	CreateGroup(w http.ResponseWriter, r *http.Request)
	UpdateGroup(w http.ResponseWriter, r *http.Request)
}

type RushControllers struct {
	Services services.RushServicesInterface
}

var _ RushControllersInterface = &RushControllers{}
