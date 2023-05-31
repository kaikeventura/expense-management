package container

import (
	"github.com/kaikeventura/expense-management/configuration/database"
	"github.com/kaikeventura/expense-management/src/controller"
	"github.com/kaikeventura/expense-management/src/repository"
	"github.com/kaikeventura/expense-management/src/service"
)

func BuildDependencyInjection() {
	userRepository := repository.ConstructUserRepository(database.GetDatabase())
	userService := service.ConstructUserService(userRepository)
	controller.ConstructUserController(userService)
}
