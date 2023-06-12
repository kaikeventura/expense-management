package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/expense-management/src/controller"
)

func ConfigurationRouter(router *gin.Engine) *gin.Engine {
	main := router.Group("v1/")
	{
		user := main.Group("user")
		{
			user.POST("/", controller.CreateUser)
		}
		expense := main.Group("expense")
		{
			expense.POST("/", controller.CreateExpense)
			expense.POST("/batch", controller.CreateExpenseInBatch)
			expense.POST("/:expenseId/fixed", controller.CreateFixedExpense)
			expense.POST("/:expenseId/purchase", controller.CreatePurchase)
			expense.POST("/:expenseId/credit-card-purchase", controller.CreateCreditCardPurchase)

			expense.GET("/:userId/current", controller.GetCurrentExpense)
		}
	}

	return router
}
