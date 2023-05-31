package repository

import (
	"log"
	"time"

	"github.com/kaikeventura/expense-management/src/dto"
	"github.com/kaikeventura/expense-management/src/repository/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func ConstructUserRepository(database *gorm.DB) UserRepository {
	return UserRepository{
		database: database,
	}
}

func (repository UserRepository) SaveUser(user dto.User) (entity.User, error) {
	userEntity := buildUserEntity(user)

	err := repository.database.Create(&userEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return entity.User{}, err
	}

	return userEntity, nil
}

func buildUserEntity(user dto.User) entity.User {
	return entity.User{
		Username:  user.Username,
		CreatedAt: time.Now(),
	}
}
