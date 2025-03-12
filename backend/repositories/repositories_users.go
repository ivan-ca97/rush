package repositories

import (
	"github.com/ivan-ca97/rush/backend/models"
)

func (rr *RushRepositories) GetUserByUsername(username string) (*models.User, error) {
	var user []models.User

	query := rr.Db.
		Table("users").
		Where("id = ?", username).
		Find(&user)

	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected == 0 {
		return nil, nil
	}

	// if query.RowsAffected > 0 {
	// 	return nil, ERROR
	// }

	return &user[0], nil
}
