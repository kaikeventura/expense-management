package database

import (
	"fmt"
	"log"
	"os"

	"github.com/kaikeventura/expense-management/configuration/database/migration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func RunDatabase() {
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	database = db

	migration.ExecuteMigrations(database)
}

func GetDatabase() *gorm.DB {
	return database
}
