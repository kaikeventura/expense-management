package routes

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/expense-management/src/controller"
)

func ConfigurationRouter(router *gin.Engine) *gin.Engine {
	router.Use(authenticate())
	main := router.Group("v1/")
	{
		user := main.Group("user")
		{
			user.POST("/", controller.CreateUser)

			user.GET("/:username", controller.GetUserByName)
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

func authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Header("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Basic ") {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(authHeader[6:])
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		clientId := credentials[0]
		clientSecret := credentials[1]

		if clientId != os.Getenv("CLIENT_ID") || clientSecret != os.Getenv("CLIENT_SECRET") {
			c.Header("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
