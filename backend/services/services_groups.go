package services

import (
	"github.com/google/uuid"
	"github.com/ivan-ca97/rush/backend/dto"
	"github.com/ivan-ca97/rush/backend/models"
	"github.com/ivan-ca97/rush/backend/utils"
)

func (rs *RushServices) GetAllGroups(pageNumber int, pageSize int) (*dto.GetAllGroupsResponse, error) {
	pageSize, offset := utils.GetLimitAndOffset(pageNumber, pageSize)

	groups, err := rs.Repositories.GetAllGroups(pageSize, offset)
	if err != nil {
		return nil, err
	}

	response := &dto.GetAllGroupsResponse{
		Groups: groups,
	}

	return response, nil
}

func (rs *RushServices) GetGroupById(id uuid.UUID) (*dto.GetGroupByIdResponse, error) {
	group, err := rs.Repositories.GetGroupById(id)
	if err != nil {
		return nil, err
	}

	response := &dto.GetGroupByIdResponse{
		Group: *group,
	}

	return response, nil
}

func (rs *RushServices) CreateGroup(body dto.CreateGroupBody) (*dto.CreateGroupResponse, error) {
	newGroup := models.Group{
		Id:          uuid.New(),
		Name:        body.Name,
		Description: body.Description,
	}

	group, err := rs.Repositories.CreateGroup(newGroup)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateGroupResponse{
		Group: *group,
	}

	return response, nil
}

func (rs *RushServices) UpdateGroup(id uuid.UUID, body dto.UpdateGroupBody) (*dto.UpdateGroupResponse, error) {
	updateGroup := models.Group{
		Id:          id,
		Name:        body.Name,
		Description: body.Description,
	}

	group, err := rs.Repositories.UpdateGroup(updateGroup)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateGroupResponse{
		Group: *group,
	}

	return response, nil
}
