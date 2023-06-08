package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/expense-management/src/dto"
	"github.com/kaikeventura/expense-management/src/service"
)

var expenseService service.ExpenseService

func ConstructExpenseController(service service.ExpenseService) {
	expenseService = service
}

func CreateExpense(context *gin.Context) {
	var expense dto.Expense
	err := context.ShouldBindJSON(&expense)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	var createdExpense, expenseError = expenseService.CreateExpense(expense)

	if expenseError != nil {
		context.JSON(400, gin.H{
			"error": expenseError.Error(),
		})

		return
	}

	context.JSON(201, createdExpense)
}

func CreateFixedExpense(context *gin.Context) {
	var fixedExpense dto.FixedExpense
	err := context.ShouldBindJSON(&fixedExpense)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	expenseId := context.Param("expenseId")
	unitExpenseId, err := strconv.ParseUint(expenseId, 10, 16)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Param: " + err.Error(),
		})

		return
	}

	var createdFixedExpense, fixedExpenseError = expenseService.CreateFixedExpense(uint16(unitExpenseId), fixedExpense)

	if fixedExpenseError != nil {
		context.JSON(400, gin.H{
			"error": fixedExpenseError.Error(),
		})

		return
	}

	context.JSON(201, createdFixedExpense)
}
