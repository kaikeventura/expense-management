package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/expense-management/src/dto"
	"github.com/kaikeventura/expense-management/src/service"
)

var userService service.UserService

func ConstructUserController(service service.UserService) {
	userService = service
}

func CreateUser(context *gin.Context) {
	var user dto.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	var createdUser, userError = userService.CreateUser(user)

	if userError != nil {
		context.JSON(404, gin.H{
			"error": userError.Error(),
		})

		return
	}

	context.JSON(201, createdUser)
}
