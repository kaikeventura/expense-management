package container

import (
	"github.com/kaikeventura/expense-management/src/controller"
	"github.com/kaikeventura/expense-management/src/service"
)

func BuildDependencyInjection() {
	userService := service.ConstructUserService()
	controller.ConstructUserController(userService)
}
