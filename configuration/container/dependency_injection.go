package container

import (
	"github.com/kaikeventura/expense-management/configuration/database"
	"github.com/kaikeventura/expense-management/src/controller"
	"github.com/kaikeventura/expense-management/src/repository"
	"github.com/kaikeventura/expense-management/src/service"
)

func BuildDependencyInjection() {
	BuildUser()
	BuildExpense()
}

func BuildUser() {
	userRepository := repository.ConstructUserRepository(database.GetDatabase())
	userService := service.ConstructUserService(userRepository)
	controller.ConstructUserController(userService)
}

func BuildExpense() {
	expenseRepository := repository.ConstructExpenseRepository(database.GetDatabase())
	expenseService := service.ConstructExpenseService(expenseRepository)
	controller.ConstructExpenseController(expenseService)
}
