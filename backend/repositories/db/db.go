package rush_db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ivan-ca97/rush/backend/models"
)

func SetupDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection fail:", err)
	}

	err = db.AutoMigrate(models.User{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(models.Role{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(models.Permission{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(models.Person{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(models.Group{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
