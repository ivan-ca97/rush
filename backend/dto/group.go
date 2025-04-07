package dto

import (
	"github.com/ivan-ca97/rush/backend/models"
)

type CreateGroupBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateGroupBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetAllGroupsResponse struct {
	Groups []models.Group `json:"groups"`
}

type GetGroupByIdResponse struct {
	models.Group
}

type CreateGroupResponse struct {
	models.Group
}

type UpdateGroupResponse struct {
	models.Group
}
