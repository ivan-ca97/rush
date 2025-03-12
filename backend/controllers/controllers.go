package controllers

import (
	"net/http"

	"github.com/ivan-ca97/rush/backend/services"
)

type RushControllersInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type RushControllers struct {
	Services services.RushServicesInterface
}

var _ RushControllersInterface = &RushControllers{}
