package service

import "github.com/kaikeventura/expense-management/src/dto"

type UserService struct {
}

func ConstructUserService() UserService {
	return UserService{}
}

func (service UserService) CreateUser(userDto dto.User) (dto.User, error) {
	return userDto, nil
}
