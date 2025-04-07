package controllers

import (
	"net/http"
)

func (rc *RushControllers) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	GetAllPaginated(rc, w, r, rc.Services.GetAllGroups)
}

func (rc *RushControllers) GetGroupById(w http.ResponseWriter, r *http.Request) {
	GetByUuid(rc, w, r, rc.Services.GetGroupById)
}

func (rc *RushControllers) CreateGroup(w http.ResponseWriter, r *http.Request) {
	Create(rc, w, r, rc.Services.CreateGroup)
}

func (rc *RushControllers) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	UpdateById(rc, w, r, rc.Services.UpdateGroup)
}
