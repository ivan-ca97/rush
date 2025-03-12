package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ivan-ca97/rush/backend/dto"
	"github.com/ivan-ca97/rush/backend/middlewares"
)

func (rc *RushControllers) Login(w http.ResponseWriter, r *http.Request) {
	var loginBody dto.LoginBody

	err := json.NewDecoder(r.Body).Decode(&loginBody)
	if err != nil {
		return
	}

	authentication := middlewares.GetAuthFromContext(r)
	loginResponse, err := rc.Services.Login(loginBody, *authentication)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(loginResponse)
	if err != nil {
		return
	}
}

func (rc *RushControllers) Register(w http.ResponseWriter, r *http.Request) {
	var registerBody dto.RegisterBody

	err := json.NewDecoder(r.Body).Decode(&registerBody)
	if err != nil {
		return
	}

	authentication := middlewares.GetAuthFromContext(r)
	registerResponse, err := rc.Services.Register(registerBody, *authentication)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(registerResponse)
	if err != nil {
		return
	}
}
