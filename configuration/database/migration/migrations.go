package migration

import (
	"log"

	"github.com/kaikeventura/expense-management/src/repository/entity"
	"gorm.io/gorm"
)

func ExecuteMigrations(database *gorm.DB) {
	err := database.AutoMigrate(entity.User{})

	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}
