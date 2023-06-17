package service

import (
	"github.com/kaikeventura/expense-management/src/dto"
	"github.com/kaikeventura/expense-management/src/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func ConstructUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (service UserService) CreateUser(userDto dto.User) (dto.User, error) {
	createdUser, err := service.repository.SaveUser(userDto)

	if err != nil {
		return dto.User{}, err
	}

	return dto.User{
		Id:       createdUser.Id,
		Username: createdUser.Username,
	}, nil
}

func (service UserService) GetUserByUsername(username string) (dto.User, error) {
	user, err := service.repository.FindUserByUsername(username)

	if err != nil {
		return dto.User{}, err
	}

	return dto.User{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}
