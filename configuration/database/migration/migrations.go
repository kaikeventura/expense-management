package migration

import (
	"log"

	"github.com/kaikeventura/expense-management/src/repository/entity"
	"gorm.io/gorm"
)

func ExecuteMigrations(database *gorm.DB) {
	trucateTables(database)
	err := database.AutoMigrate(entity.User{})

	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}

func trucateTables(database *gorm.DB) {
	database.Exec("truncate table users")
}
