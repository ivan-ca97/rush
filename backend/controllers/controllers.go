package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ivan-ca97/rush/backend/custom_errors"
	"github.com/ivan-ca97/rush/backend/services"
)

type RushControllersInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)

	GetPersonById(w http.ResponseWriter, r *http.Request)
	CreatePerson(w http.ResponseWriter, r *http.Request)
	UpdatePerson(w http.ResponseWriter, r *http.Request)
}

type RushControllers struct {
	Services services.RushServicesInterface
}

var _ RushControllersInterface = &RushControllers{}

func handleError(w http.ResponseWriter, err error) {
	var expectedErr custom_errors.ExpectedError

	if errors.As(err, &expectedErr) {
		w.WriteHeader(expectedErr.StatusCode())
		json.NewEncoder(w).Encode(expectedErr)
		return
	}

	log.Printf("Unexpected error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error"})
}
