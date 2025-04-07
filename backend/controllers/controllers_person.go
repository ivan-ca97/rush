package controllers

import (
	"net/http"
)

func (rc *RushControllers) GetPersonById(w http.ResponseWriter, r *http.Request) {
	GetByUuid(rc, w, r, rc.Services.GetPersonById)
}

func (rc *RushControllers) CreatePerson(w http.ResponseWriter, r *http.Request) {
	Create(rc, w, r, rc.Services.CreatePerson)
}

func (rc *RushControllers) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	UpdateById(rc, w, r, rc.Services.UpdatePerson)
}
