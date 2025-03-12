package rush_db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ivan-ca97/rush/backend/models"
)

func SetupDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection fail:", err)
	}

	db.AutoMigrate(models.User{})
	return db
}
