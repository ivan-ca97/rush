package repositories

import (
	"errors"

	"github.com/google/uuid"
	ce "github.com/ivan-ca97/rush/backend/custom_errors"
	"github.com/ivan-ca97/rush/backend/models"
	"gorm.io/gorm"
)

func (rr *RushRepositories) GetAllGroups(limit int, offset int) ([]models.Group, error) {
	var groups []models.Group

	err := rr.Db.Limit(limit).Offset(offset).Find(&groups).Error
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (rr *RushRepositories) GetGroupById(id uuid.UUID) (*models.Group, error) {
	var group models.Group

	err := rr.Db.Where("id = ?", id).First(&group).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &ce.RecordNotFoundError{Message: "Group ID not found"}
	}
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (rr *RushRepositories) UpdateGroup(group models.Group) (*models.Group, error) {
	err := rr.Db.Model(&group).Where("id = ?", group.Id).Updates(group).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &ce.RecordNotFoundError{Message: "Group ID not found"}
	}
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (rr *RushRepositories) CreateGroup(group models.Group) (*models.Group, error) {
	err := rr.Db.Create(&group).Error
	if err != nil {
		return nil, err
	}

	return &group, nil
}
