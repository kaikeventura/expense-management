package routes

import (
	"github.com/gin-gonic/gin"
)

func ConfigurationRouter(router *gin.Engine) *gin.Engine {
	main := router.Group("v1/")
	{
		users := main.Group("expense")
		{
			users.GET("/")
		}
	}

	return router
}
