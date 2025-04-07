package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ivan-ca97/rush/backend/dto"
)

func (rc *RushControllers) GetPersonById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	personGetByIdResponse, err := rc.Services.GetPersonById(id)
	if err != nil {
		rc.HandleError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(personGetByIdResponse)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

func (rc *RushControllers) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var personCreateBody dto.PersonCreateBody

	err := json.NewDecoder(r.Body).Decode(&personCreateBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	personCreateResponse, err := rc.Services.CreatePerson(personCreateBody)
	if err != nil {
		rc.HandleError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(personCreateResponse)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

func (rc *RushControllers) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var personUpdateBody dto.PersonUpdateBody

	err := json.NewDecoder(r.Body).Decode(&personUpdateBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	personUpdateResponse, err := rc.Services.UpdatePerson(id, personUpdateBody)
	if err != nil {
		rc.HandleError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(personUpdateResponse)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}
